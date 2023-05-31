package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/natrontech/pocketbase-client-go/pkg/models"
)

// ListCollections - Returns a paginated Collections list. Only admins can access this action.
func (c *Client) ListCollections() (*models.CollectionList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/collections", c.Endpoint), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	collections := models.CollectionList{}
	err = json.Unmarshal(body, &collections)
	if err != nil {
		return nil, err
	}

	return &collections, nil
}

// ViewCollection - Returns a single Collection by its ID or name. Only admins can access this action.
func (c *Client) ViewCollection(id string) (*models.Collection, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/collections/%s", c.Endpoint, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	collection := models.Collection{}
	err = json.Unmarshal(body, &collection)
	if err != nil {
		return nil, err
	}

	return &collection, nil
}

// CreateCollection - Creates a new Collection. Only admins can access this action.
func (c *Client) CreateCollection(collection *models.CollectionCreateRequest) (*models.Collection, error) {
	rb, err := json.Marshal(collection)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/collections", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	col := models.Collection{}
	err = json.Unmarshal(body, &col)
	if err != nil {
		return nil, err
	}

	return &col, nil
}

// UpdateCollection - Updates a single Collection by its ID or name. Only admins can access this action.
func (c *Client) UpdateCollection(id string, collection *models.CollectionUpdateRequest) (*models.Collection, error) {
	rb, err := json.Marshal(collection)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/collections/%s", c.Endpoint, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	col := models.Collection{}
	err = json.Unmarshal(body, &col)
	if err != nil {
		return nil, err
	}

	return &col, nil
}

// DeleteCollection - Deletes a single Collection by its ID or name. Only admins can access this action.
func (c *Client) DeleteCollection(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/collections/%s", c.Endpoint, id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}

// ImportCollections - Bulk imports the provided Collections configuration. Only admins can access this action.
func (c *Client) ImportCollections(collections *models.CollectionImportRequest) error {
	rb, err := json.Marshal(collections)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/collections/import", c.Endpoint), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}
