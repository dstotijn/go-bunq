package bunq

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type userResponse struct {
	Response []struct {
		UserCompany *UserCompany `json:"UserCompany"`
	} `json:"Response"`
}

// ErrUserNotFound is returned when a single User resource was not found.
var ErrUserNotFound = errors.New("user not found")

// GetUser gets a User resource at the bunq API.
func (c *Client) GetUser(id int) (interface{}, error) {
	httpMethod := http.MethodGet
	endpoint := apiVersion + "/user/" + strconv.Itoa(id)
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

	var userResp userResponse
	if err = json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		return nil, fmt.Errorf("bunq: could not decode HTTP response: %v", err)
	}

	users, err := userResp.users()
	if err != nil {
		return nil, fmt.Errorf("bunq: could not parse API response: %v", err)
	}
	if len(users) == 0 {
		return nil, ErrUserNotFound
	}

	return users[0], nil
}

// ListUsers gets a list of User resources at the bunq API.
func (c *Client) ListUsers() ([]interface{}, error) {
	httpMethod := http.MethodGet
	endpoint := apiVersion + "/user"
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

	var userResp userResponse
	if err = json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		return nil, fmt.Errorf("bunq: could not decode HTTP response: %v", err)
	}

	users, err := userResp.users()
	if err != nil {
		return nil, fmt.Errorf("bunq: could not parse API response: %v", err)
	}

	return users, nil
}

func (userResp *userResponse) users() ([]interface{}, error) {
	users := make([]interface{}, len(userResp.Response))
	for i := range userResp.Response {
		if userResp.Response[i].UserCompany != nil {
			users[i] = *userResp.Response[i].UserCompany
		}
	}

	return users, nil
}
