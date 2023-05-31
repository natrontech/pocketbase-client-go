package client_test

import (
	"testing"

	"github.com/natrontech/pocketbase-client-go/pkg/client"
	"github.com/stretchr/testify/require"
)

func TestAuthWithPassword(t *testing.T) {
	testCases := []struct {
		name          string
		identity      string
		password      string
		expectedError string
	}{
		{
			name:          "Should fail when identity is empty",
			identity:      "",
			password:      "password",
			expectedError: "identity and password must be set",
		},
		{
			name:          "Should fail when password is empty",
			identity:      "identity",
			password:      "",
			expectedError: "identity and password must be set",
		},
		// You can add more test cases here.
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := client.Client{
				Auth: client.AuthStruct{
					Identity: tc.identity,
					Password: tc.password,
				},
				// Add more initial fields if required.
			}
			_, err := c.AuthWithPassword()

			if tc.expectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAuthRefresh(t *testing.T) {
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
			_, err := c.AuthRefresh()

			if tc.expectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
