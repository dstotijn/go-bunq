package bunq

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

type deviceResponse struct {
	Response []struct {
		DevicePhone *struct {
			ID          int    `json:"id"`
			Created     Time   `json:"created"`
			Updated     Time   `json:"updated"`
			Description string `json:"description"`
			PhoneNumber string `json:"phone_number"`
			OS          string `json:"os"`
			Status      string `json:"status"`
		} `json:"DevicePhone,omitempty"`
		DeviceServer *struct {
			ID          int    `json:"id"`
			Created     Time   `json:"created"`
			Updated     Time   `json:"updated"`
			Description string `json:"description"`
			IP          string `json:"ip"`
			Status      string `json:"status"`
		} `json:"DeviceServer,omitempty"`
	} `json:"Response"`
}

// A DevicePhone represents a Device at the bunq API.
type DevicePhone struct {
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Description string
	PhoneNumber string
	OS          string
	Status      string
}

// ErrDeviceNotFound is returned when a single Device resource was
// not found.
var ErrDeviceNotFound = errors.New("device not found")

// GetDevice gets a Device resource at the bunq API.
func (c *Client) GetDevice(id int) (interface{}, error) {
	httpMethod := http.MethodGet
	endpoint := apiVersion + "/device/" + strconv.Itoa(id)
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

	var devResp deviceResponse
	if err = json.NewDecoder(resp.Body).Decode(&devResp); err != nil {
		return nil, fmt.Errorf("bunq: could not decode HTTP response: %v", err)
	}

	devices, err := devResp.devices()
	if err != nil {
		return nil, fmt.Errorf("bunq: could not parse API response: %v", err)
	}
	if len(devices) == 0 {
		return nil, ErrDeviceNotFound
	}

	return devices[0], nil
}

// ListDevices gets a list of Device resources at the bunq API.
func (c *Client) ListDevices() ([]interface{}, error) {
	httpMethod := http.MethodGet
	endpoint := apiVersion + "/device"
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

	var devResp deviceResponse
	if err = json.NewDecoder(resp.Body).Decode(&devResp); err != nil {
		return nil, fmt.Errorf("bunq: could not decode HTTP response: %v", err)
	}

	devices, err := devResp.devices()
	if err != nil {
		return nil, fmt.Errorf("bunq: could not parse API response: %v", err)
	}

	return devices, nil
}

func (devResp *deviceResponse) devices() ([]interface{}, error) {
	devices := make([]interface{}, len(devResp.Response))
	for i := range devResp.Response {
		if devResp.Response[i].DevicePhone != nil {
			device := DevicePhone{}
			device.ID = devResp.Response[i].DevicePhone.ID
			device.CreatedAt = time.Time(devResp.Response[i].DevicePhone.Created)
			device.UpdatedAt = time.Time(devResp.Response[i].DevicePhone.Updated)
			device.Description = devResp.Response[i].DevicePhone.Description
			device.PhoneNumber = devResp.Response[i].DevicePhone.PhoneNumber
			device.OS = devResp.Response[i].DevicePhone.OS
			device.Status = devResp.Response[i].DevicePhone.Status
			devices[i] = device
		}
		if devResp.Response[i].DeviceServer != nil {
			device := DeviceServer{}
			device.ID = devResp.Response[i].DeviceServer.ID
			device.Description = devResp.Response[i].DeviceServer.Description
			device.CreatedAt = time.Time(devResp.Response[i].DeviceServer.Created)
			device.UpdatedAt = time.Time(devResp.Response[i].DeviceServer.Updated)
			device.Status = devResp.Response[i].DeviceServer.Status
			device.IP = net.ParseIP(devResp.Response[i].DeviceServer.IP)
			devices[i] = device
		}
	}

	return devices, nil
}
