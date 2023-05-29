package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// SignIn - Get a new token
func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.Identity == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("identity and password must be set")
	}

	rb, err := json.Marshal(c.Auth)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/admins/auth-with-password", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
