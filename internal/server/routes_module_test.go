package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testServer *httptest.Server
	testClient *http.Client
)

// TestModuleVersionsRoute is responsible for testing the module versions endpoint
func TestModuleVersionsRoute(t *testing.T) {
	testCases := map[string]struct {
		namespace string
		name      string
		system    string
		expected  string
	}{
		"GetModuleVersions": {
			namespace: "hashicorp",
			name:      "consul",
			system:    "aws",
			expected:  "Module Versions",
		},
	}

	// Handle setup and teardown
	func() {
		srv := New(&Options{})
		testServer = httptest.NewServer(srv.handleModuleVersions())
		testClient = testServer.Client()
	}()
	defer func() {
		testServer.Close()
	}()

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			// Build the endpoint to test
			moduleVersionsEndpoint := fmt.Sprintf(
				"%s/%s/%s/%s/versions",
				testServer.URL,
				testCase.namespace,
				testCase.name,
				testCase.system,
			)
			t.Logf("Test Server URL: %s", moduleVersionsEndpoint)

			// Make request to test server
			resp, err := testClient.Post(moduleVersionsEndpoint, "application/json", nil)
			if err != nil {
				t.Errorf("failed making a request to the test server: %s", err)
			}
			defer resp.Body.Close()

			// Read response body so we can test
			respBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("failed reading response body bytes: %s", err)
			}
			t.Logf("Response: %s", string(respBytes))

			// Compare output with a mocked response
			// TODO: return a mocked response for module versions
			if string(respBytes) != testCase.expected {
				t.Errorf(
					"failed getting the expected response,\n  expected: %s\n  recieved: %s\n",
					testCase.expected,
					respBytes,
				)
			}
		})
	}
}
