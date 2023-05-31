package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/natrontech/pocketbase-client-go/pkg/models"
)

// AuthWithPassword - Authenticate an admin with their email and password.
func (c *Client) AuthWithPassword() (*AuthResponse, error) {
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

// AuthRefresh - Returns a new auth response (token and admin data) for already authenticated admin.
func (c *Client) AuthRefresh() (*AuthResponse, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/admins/auth-refresh", c.Endpoint), nil)
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

// RequestPasswordReset - Sends a password reset email to an admin.
func (c *Client) RequestPasswordReset(email string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/admins/request-password-reset", c.Endpoint), strings.NewReader(fmt.Sprintf(`{"email":"%s"}`, email)))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}

// ConfirmPasswordReset - Confirms a password reset request and sets a new admin password.
func (c *Client) ConfirmPasswordReset(token, password, passwordConfirm string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/admins/confirm-password-reset", c.Endpoint), strings.NewReader(fmt.Sprintf(`{"token":"%s","password":"%s","password_confirm":"%s"}`, token, password, passwordConfirm)))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}

// ListAdmins - Returns a paginated admins list. Only admins can access this action.
func (c *Client) ListAdmins() (*models.AdminList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/admins", c.Endpoint), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	admins := models.AdminList{}
	err = json.Unmarshal(body, &admins)
	if err != nil {
		return nil, err
	}

	return &admins, nil
}

// ViewAdmin - Return a single admin by its ID. Only admins can access this action.
func (c *Client) ViewAdmin(id string) (*models.Admin, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/admins/%s", c.Endpoint, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	admin := models.Admin{}
	err = json.Unmarshal(body, &admin)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

// CreateAdmin - Creates a new admin. Only admins can access this action.
func (c *Client) CreateAdmin(admin *models.AdminCreateRequest) (*models.Admin, error) {
	rb, err := json.Marshal(admin)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/admins", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	a := models.Admin{}
	err = json.Unmarshal(body, &a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// UpdateAdmin - Update a single admin model by its ID. Only admins can access this action.
func (c *Client) UpdateAdmin(id string, admin *models.AdminUpdateRequest) (*models.Admin, error) {
	rb, err := json.Marshal(admin)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/admins/%s", c.Endpoint, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	a := models.Admin{}
	err = json.Unmarshal(body, &a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// DeleteAdmin - Deletes a single admin by its id. Only admins can access this action.
func (c *Client) DeleteAdmin(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/admins/%s", c.Endpoint, id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}
