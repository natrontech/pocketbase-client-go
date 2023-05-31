package client_test

import (
	"testing"

	"github.com/natrontech/pocketbase-client-go/pkg/client"
	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError string
	}{
		// You can add more test cases here.
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := client.Client{
				// Add more initial fields if required.
			}
			_, err := c.HealthCheck()

			if tc.expectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
