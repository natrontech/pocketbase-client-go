package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	assert := assert.New(t)
	endpoint := "http://localhost:8090"
	identity := "admin@natron.io"
	password := "0123456789"

	c, err := NewClient(&endpoint, &identity, &password)
	assert.NoError(err)
	assert.Equal(endpoint, c.Endpoint)
}
