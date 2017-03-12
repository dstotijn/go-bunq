package bunq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/%v", c.BaseURL, endpoint), bytes.NewReader(bodyJSON))
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

	var apiResp installationResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("bunq: could not decode HTTP response: %v", err)
	}

	if len(apiResp.Response) == 0 {
		return nil, errors.New("bunq: api response did not contain results")
	}

	installation := &Installation{}
	for i := range apiResp.Response {
		if apiResp.Response[i].ID != nil {
			installation.ID = apiResp.Response[i].ID.ID
			continue
		}
		if apiResp.Response[i].Token != nil {
			installation.Token.ID = apiResp.Response[i].Token.ID
			installation.Token.Token = apiResp.Response[i].Token.Token
			createdAt, err := parseTimestamp(apiResp.Response[i].Token.Created)
			if err != nil {
				return nil, fmt.Errorf("bunq: could not parse created timestamp: %v", err)
			}
			installation.Token.CreatedAt = createdAt
			updatedAt, err := parseTimestamp(apiResp.Response[i].Token.Updated)
			if err != nil {
				return nil, fmt.Errorf("bunq: could not parse updated timestamp: %v", err)
			}
			installation.Token.UpdatedAt = updatedAt
			continue
		}
		if apiResp.Response[i].ServerPublicKey != nil {
			installation.ServerPublicKey = apiResp.Response[i].ServerPublicKey.ServerPublicKey
		}
	}

	return installation, nil
}
