package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Endpoint represents a single endpoint
const Endpoint = "http://localhost:8090"

// Client -
type Client struct {
	Endpoint   string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	Token string `json:"token"`
	Admin struct {
		Id      string `json:"id"`
		Created string `json:"created"`
		Updated string `json:"updated"`
		Email   string `json:"email"`
		Avatar  int    `json:"avatar"`
	} `json:"admin"`
}

// NewClient -
func NewClient(endpoint, identity, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Endpoint:   Endpoint,
	}

	if endpoint != nil {
		c.Endpoint = *endpoint
	}

	if identity == nil && password == nil {
		return &c, nil
	}

	c.Auth = AuthStruct{
		Identity: *identity,
		Password: *password,
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	if token == "" {
		req.Header.Set("Authorization", token)
	}

	// set content type to json charset utf-8
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
