package bunq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
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

	endpoint := apiVersion + "/device-server"
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/%v", c.BaseURL, endpoint), bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, fmt.Errorf("bunq: could not create new request: %v", err)
	}
	setCommonHeaders(req)
	req.Header.Set("X-Bunq-Client-Authentication", c.Token)
	if err = c.addSignature(req, "POST /"+endpoint, string(bodyJSON)); err != nil {
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

	var apiResp deviceServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("bunq: could not decode HTTP response: %v", err)
	}

	if len(apiResp.Response) == 0 {
		return nil, errors.New("bunq: api response did not contain results")
	}

	deviceServer := &DeviceServer{}
	for i := range apiResp.Response {
		if apiResp.Response[i].ID != nil {
			deviceServer.ID = apiResp.Response[i].ID.ID
			continue
		}
		if apiResp.Response[i].DeviceServer != nil {
			deviceServer.ID = apiResp.Response[i].ID.ID
			deviceServer.Description = apiResp.Response[i].DeviceServer.Description
			createdAt, err := parseTimestamp(apiResp.Response[i].DeviceServer.Created)
			if err != nil {
				return nil, fmt.Errorf("bunq: could not parse created timestamp: %v", err)
			}
			deviceServer.CreatedAt = createdAt
			updatedAt, err := parseTimestamp(apiResp.Response[i].DeviceServer.Updated)
			if err != nil {
				return nil, fmt.Errorf("bunq: could not parse updated timestamp: %v", err)
			}
			deviceServer.UpdatedAt = updatedAt
			continue
		}
	}

	return deviceServer, nil
}
