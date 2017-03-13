package bunq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type installationResponse struct {
	Response []struct {
		ID *struct {
			ID int `json:"id"`
		} `json:"Id,omitempty"`
		Token *struct {
			ID      int    `json:"id"`
			Created string `json:"created"`
			Updated string `json:"updated"`
			Token   string `json:"token"`
		} `json:"Token,omitempty"`
		ServerPublicKey *struct {
			ServerPublicKey string `json:"server_public_key"`
		} `json:"ServerPublicKey,omitempty"`
	} `json:"Response"`
}

// A Token is used for authenticating requests to the bunq API.
type Token struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Token     string
}

// An Installation represents an installation resource at the bunq API.
type Installation struct {
	ID              int
	Token           Token
	ServerPublicKey string
}

// ErrInstallationNotFound is returned when a single Installation resource was
// not found.
var ErrInstallationNotFound = errors.New("installation not found")

// CreateInstallation creates an installation resource at the bunq API.
func (c *Client) CreateInstallation() (*Installation, error) {
	publicKey, err := c.publicKey()
	if err != nil {
		return nil, fmt.Errorf("bunq: could not get public key: %v", err)
	}

	body := struct {
		string `json:"client_public_key"`
	}{string(publicKey)}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("bunq: could not encode request body into JSON: %v", err)
	}

	endpoint := apiVersion + "/installation"
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%v/%v", c.BaseURL, endpoint), bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, fmt.Errorf("bunq: could not create new request: %v", err)
	}
	setCommonHeaders(req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("bunq: could not send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bunq: request was unsuccessful: %v", decodeError(resp.Body))
	}

	var insResp installationResponse
	if err = json.NewDecoder(resp.Body).Decode(&insResp); err != nil {
		return nil, fmt.Errorf("bunq: could not decode HTTP response: %v", err)
	}

	if len(insResp.Response) == 0 {
		return nil, errors.New("bunq: api response did not contain results")
	}

	installation, err := insResp.Installation()
	if err != nil {
		return nil, fmt.Errorf("bunq: could not parse API response: %v", err)
	}

	return installation, nil
}

// GetInstallation gets an Installation resource at the bunq API.
func (c *Client) GetInstallation(id int) (*Installation, error) {
	httpMethod := http.MethodGet
	endpoint := apiVersion + "/installation/" + strconv.Itoa(id)
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

	var insResp installationResponse
	if err = json.NewDecoder(resp.Body).Decode(&insResp); err != nil {
		return nil, fmt.Errorf("bunq: could not decode HTTP response: %v", err)
	}

	if len(insResp.Response) == 0 {
		return nil, ErrInstallationNotFound
	}

	installation, err := insResp.Installation()
	if err != nil {
		return nil, fmt.Errorf("bunq: could not parse API response: %v", err)
	}

	return installation, nil
}

// GetInstallationID gets the installation ID of the Installation used for the
// current session.
func (c *Client) GetInstallationID() (int, error) {
	httpMethod := http.MethodGet
	endpoint := apiVersion + "/installation"
	req, err := http.NewRequest(httpMethod, fmt.Sprintf("%v/%v", c.BaseURL, endpoint), nil)
	if err != nil {
		return 0, fmt.Errorf("bunq: could not create new request: %v", err)
	}
	setCommonHeaders(req)
	req.Header.Set("X-Bunq-Client-Authentication", c.Token)
	if err = c.addSignature(req, fmt.Sprintf("%v /%v", httpMethod, endpoint), ""); err != nil {
		return 0, fmt.Errorf("bunq: could not add signature: %v", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("bunq: could not send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bunq: request was unsuccessful: %v", decodeError(resp.Body))
	}

	var insResp installationResponse
	if err = json.NewDecoder(resp.Body).Decode(&insResp); err != nil {
		return 0, fmt.Errorf("bunq: could not decode HTTP response: %v", err)
	}

	if len(insResp.Response) == 0 {
		return 0, ErrInstallationNotFound
	}

	installation, err := insResp.Installation()
	if err != nil {
		return 0, fmt.Errorf("bunq: could not parse API response: %v", err)
	}

	return installation.ID, nil
}

func (insResp *installationResponse) Installation() (*Installation, error) {
	installation := &Installation{}
	for i := range insResp.Response {
		if insResp.Response[i].ID != nil {
			installation.ID = insResp.Response[i].ID.ID
			continue
		}
		if insResp.Response[i].Token != nil {
			installation.Token.ID = insResp.Response[i].Token.ID
			installation.Token.Token = insResp.Response[i].Token.Token
			createdAt, err := parseTimestamp(insResp.Response[i].Token.Created)
			if err != nil {
				return nil, fmt.Errorf("could not parse created timestamp: %v", err)
			}
			installation.Token.CreatedAt = createdAt
			updatedAt, err := parseTimestamp(insResp.Response[i].Token.Updated)
			if err != nil {
				return nil, fmt.Errorf("could not parse updated timestamp: %v", err)
			}
			installation.Token.UpdatedAt = updatedAt
			continue
		}
		if insResp.Response[i].ServerPublicKey != nil {
			installation.ServerPublicKey = insResp.Response[i].ServerPublicKey.ServerPublicKey
		}
	}

	return installation, nil
}
