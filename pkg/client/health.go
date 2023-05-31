package client

import (
	"encoding/json"
	"net/http"

	"github.com/natrontech/pocketbase-client-go/pkg/models"
)

func (c *Client) HealthCheck() (*models.HealthResponse, error) {
	req, err := http.NewRequest("GET", c.Endpoint+"/api/health", nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	health := models.HealthResponse{}
	err = json.Unmarshal(body, &health)
	if err != nil {
		return nil, err
	}

	return &health, nil
}
