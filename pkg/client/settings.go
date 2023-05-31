package client

import (
	"encoding/json"
	"fmt"
	"net/http"

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
