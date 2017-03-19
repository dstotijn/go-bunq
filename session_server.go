package bunq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type sessionServerResponse struct {
	Response []struct {
		ID *struct {
			ID int `json:"id"`
		} `json:"Id,omitempty"`
		Token *struct {
			ID    int    `json:"id"`
			Token string `json:"token"`
		} `json:"Token,omitempty"`
		UserCompany *UserCompany `json:"UserCompany,omitempty"`
	} `json:"Response"`
}

// A SessionToken is used for authenticating requests to the bunq API.
type SessionToken struct {
	ID    int
	Token string
}

// A Session represents a SessionServer at the bunq API.
type Session struct {
	ID          int
	Token       SessionToken
	UserCompany UserCompany
}

// CreateSession creates a new session for a DeviceServer at the bunq API.
func (c *Client) CreateSession() (*Session, error) {
	body := struct {
		Secret string `json:"secret"`
	}{c.APIKey}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("bunq: could not encode request body into JSON: %v", err)
	}

	httpMethod := http.MethodPost
	endpoint := apiVersion + "/session-server"
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

	var ssr sessionServerResponse
	if err = json.NewDecoder(resp.Body).Decode(&ssr); err != nil {
		return nil, fmt.Errorf("bunq: could not decode HTTP response: %v", err)
	}

	if len(ssr.Response) == 0 {
		return nil, errors.New("bunq: api response did not contain results")
	}

	session, err := ssr.session()
	if err != nil {
		return nil, fmt.Errorf("bunq: could not parse API response: %v", err)
	}

	return session, nil
}

func (ssr *sessionServerResponse) session() (*Session, error) {
	session := &Session{}
	for i := range ssr.Response {
		if ssr.Response[i].ID != nil {
			session.ID = ssr.Response[i].ID.ID
			continue
		}
		if ssr.Response[i].Token != nil {
			session.Token.ID = ssr.Response[i].Token.ID
			session.Token.Token = ssr.Response[i].Token.Token
			continue
		}
		if ssr.Response[i].UserCompany != nil {
			session.UserCompany = *ssr.Response[i].UserCompany
		}
	}

	return session, nil
}
