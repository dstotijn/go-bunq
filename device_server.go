package bunq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

type deviceServerResponse struct {
	Response []struct {
		ID *struct {
			ID int `json:"id"`
		} `json:"Id,omitempty"`
		DeviceServer *struct {
			ID          int    `json:"id"`
			Created     string `json:"created"`
			Updated     string `json:"updated"`
			Description string `json:"description"`
			IP          string `json:"ip"`
			Status      string `json:"status"`
		} `json:"DeviceServer,omitempty"`
	} `json:"Response"`
}

// A DeviceServer represents a DeviceServe at the bunq API.
type DeviceServer struct {
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Description string
	IP          net.IP
	Status      string
}

// ErrDeviceServerNotFound is returned when a single DeviceServer resource was
// not found.
var ErrDeviceServerNotFound = errors.New("device server not found")

// CreateDeviceServer creates a DeviceServer resource at the bunq API.
func (c *Client) CreateDeviceServer(description string, permittedIPs []net.IP) (*DeviceServer, error) {
	var ips []string
	for i := range permittedIPs {
		ips = append(ips, permittedIPs[i].String())
	}
	body := struct {
		Description  string   `json:"description"`
		Secret       string   `json:"secret"`
		PermittedIPs []string `json:"permitted_ips,omitempty"`
	}{description, c.APIKey, ips}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("bunq: could not encode request body into JSON: %v", err)
	}

	httpMethod := http.MethodPost
	endpoint := apiVersion + "/device-server"
	req, err := http.NewRequest(httpMethod, fmt.Sprintf("%v/%v", c.BaseURL, endpoint), bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, fmt.Errorf("bunq: could not create new request: %v", err)
	}
	setCommonHeaders(req)
	req.Header.Set("X-Bunq-Client-Authentication", c.Token)
	if err = c.addSignature(req, fmt.Sprintf("%v /%v", httpMethod, endpoint), string(bodyJSON)); err != nil {
		return nil, fmt.Errorf("bunq: could not add signature: %v", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("bunq: could not send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bunq: request was unsuccessful: %v", decodeError(resp.Body))
	}

	var dsr deviceServerResponse
	if err = json.NewDecoder(resp.Body).Decode(&dsr); err != nil {
		return nil, fmt.Errorf("bunq: could not decode HTTP response: %v", err)
	}

	deviceServers, err := dsr.DeviceServers()
	if err != nil {
		return nil, fmt.Errorf("bunq: could not parse API response: %v", err)
	}
	if len(deviceServers) == 0 {
		return nil, errors.New("bunq: api response did not contain results")
	}

	return deviceServers[0], nil
}

// GetDeviceServer gets a DeviceServer resource at the bunq API.
func (c *Client) GetDeviceServer(id int) (*DeviceServer, error) {
	httpMethod := http.MethodGet
	endpoint := apiVersion + "/device-server/" + strconv.Itoa(id)
	req, err := http.NewRequest(httpMethod, fmt.Sprintf("%v/%v", c.BaseURL, endpoint), nil)
	if err != nil {
		return nil, fmt.Errorf("bunq: could not create new request: %v", err)
	}
	setCommonHeaders(req)
	req.Header.Set("X-Bunq-Client-Authentication", c.Token)
	if err = c.addSignature(req, fmt.Sprintf("%v /%v", httpMethod, endpoint), ""); err != nil {
		return nil, fmt.Errorf("bunq: could not add signature: %v", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("bunq: could not send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bunq: request was unsuccessful: %v", decodeError(resp.Body))
	}

	var dsr deviceServerResponse
	if err = json.NewDecoder(resp.Body).Decode(&dsr); err != nil {
		return nil, fmt.Errorf("bunq: could not decode HTTP response: %v", err)
	}

	deviceServers, err := dsr.DeviceServers()
	if err != nil {
		return nil, fmt.Errorf("bunq: could not parse API response: %v", err)
	}
	if len(deviceServers) == 0 {
		return nil, ErrDeviceServerNotFound
	}

	return deviceServers[0], nil
}

// ListDeviceServers gets a list of DeviceServer resources at the bunq API.
func (c *Client) ListDeviceServers() ([]*DeviceServer, error) {
	httpMethod := http.MethodGet
	endpoint := apiVersion + "/device-server"
	req, err := http.NewRequest(httpMethod, fmt.Sprintf("%v/%v", c.BaseURL, endpoint), nil)
	if err != nil {
		return nil, fmt.Errorf("bunq: could not create new request: %v", err)
	}
	setCommonHeaders(req)
	req.Header.Set("X-Bunq-Client-Authentication", c.Token)
	if err = c.addSignature(req, fmt.Sprintf("%v /%v", httpMethod, endpoint), ""); err != nil {
		return nil, fmt.Errorf("bunq: could not add signature: %v", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("bunq: could not send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bunq: request was unsuccessful: %v", decodeError(resp.Body))
	}

	var dsr deviceServerResponse
	if err = json.NewDecoder(resp.Body).Decode(&dsr); err != nil {
		return nil, fmt.Errorf("bunq: could not decode HTTP response: %v", err)
	}

	deviceServers, err := dsr.DeviceServers()
	if err != nil {
		return nil, fmt.Errorf("bunq: could not parse API response: %v", err)
	}

	return deviceServers, nil
}

func (dsr *deviceServerResponse) DeviceServers() ([]*DeviceServer, error) {
	var deviceServers []*DeviceServer
	for i := range dsr.Response {
		deviceServer := &DeviceServer{}
		if dsr.Response[i].ID != nil {
			deviceServer.ID = dsr.Response[i].ID.ID
			deviceServers = append(deviceServers, deviceServer)
		}
		if dsr.Response[i].DeviceServer != nil {
			deviceServer.ID = dsr.Response[i].DeviceServer.ID
			deviceServer.Description = dsr.Response[i].DeviceServer.Description
			createdAt, err := parseTimestamp(dsr.Response[i].DeviceServer.Created)
			if err != nil {
				return nil, fmt.Errorf("could not parse created timestamp: %v", err)
			}
			deviceServer.CreatedAt = createdAt
			updatedAt, err := parseTimestamp(dsr.Response[i].DeviceServer.Updated)
			if err != nil {
				return nil, fmt.Errorf("could not parse updated timestamp: %v", err)
			}
			deviceServer.UpdatedAt = updatedAt
			deviceServer.Status = dsr.Response[i].DeviceServer.Status
			deviceServer.IP = net.ParseIP(dsr.Response[i].DeviceServer.IP)
			deviceServers = append(deviceServers, deviceServer)
		}
	}

	return deviceServers, nil
}
