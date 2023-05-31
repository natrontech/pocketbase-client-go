package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/natrontech/pocketbase-client-go/pkg/models"
)

// ListSettings - Returns a list with all available application settings. Secret/password fields are automatically redacted with ****** characters. Only admins can access this action.
func (c *Client) ListSettings(fields *string) (*models.SettingsList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/settings", c.Endpoint), nil)
	if err != nil {
		return nil, err
	}

	if fields != nil {
		q := req.URL.Query()
		q.Add("fields", *fields)
		req.URL.RawQuery = q.Encode()
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	settings := models.SettingsList{}
	err = json.Unmarshal(body, &settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}

// UpdateSettings - Bulk updates application settings and returns the updated settings list. Only admins can access this action.
func (c *Client) UpdateSettings(settings *models.SettingsUpdateRequest) (*models.SettingsList, error) {
	rb, err := json.Marshal(settings)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/settings", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	settingsList := models.SettingsList{}
	err = json.Unmarshal(body, &settingsList)
	if err != nil {
		return nil, err
	}

	return &settingsList, nil
}

// TestS3StorageConnection - Performs a S3 storage connection test. Only admins can access this action.
func (c *Client) TestS3StorageConnection(settings *models.SettingsS3StorageTestRequest) (*models.SettingsResponse, error) {
	rb, err := json.Marshal(settings)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/settings/s3-storage/test", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	testResponse := models.SettingsResponse{}
	err = json.Unmarshal(body, &testResponse)
	if err != nil {
		return nil, err
	}

	return &testResponse, nil
}

// SendTestEmail - Sends a test user email. Only admins can access this action.
func (c *Client) SendTestEmail(settings *models.SettingsSendTestEmailRequest) (*models.SettingsResponse, error) {
	rb, err := json.Marshal(settings)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/settings/send-test-email", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	testResponse := models.SettingsResponse{}
	err = json.Unmarshal(body, &testResponse)
	if err != nil {
		return nil, err
	}

	return &testResponse, nil
}
