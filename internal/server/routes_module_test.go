package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	srv        *Server          // Instance of application server to test
	testServer *httptest.Server // HTTP test server to mock the response
	testClient *http.Client     // HTTP client created from the test server
)

func setUp(endpoint http.HandlerFunc) {
	srv = New(&Options{})
	testServer = httptest.NewServer(endpoint)
	testClient = testServer.Client()
}

func tearDown() {
	testServer.Close()
}

// TestModuleVersionsRoute is responsible for testing the module versions endpoint
func TestModuleVersionsRoute(t *testing.T) {
	testCases := map[string]struct {
		namespace string
		name      string
		system    string
		expected  []byte
	}{
		"GetModuleVersions": {
			namespace: "hashicorp",
			name:      "consul",
			system:    "aws",
			expected:  []byte("Module Versions"),
		},
	}

	// Handle setup and teardown
	setUp(srv.handleModuleVersions())
	defer tearDown()

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
			t.Logf("Module Versions Endpoint: %s", moduleVersionsEndpoint)

			// Make request to test server
			resp, err := testClient.Get(moduleVersionsEndpoint)
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
			if i := bytes.Compare(respBytes, testCase.expected); i != 0 {
				t.Errorf(
					"failed getting the expected response,\n  expected: %s\n  recieved: %s\n",
					string(testCase.expected),
					string(respBytes),
				)
			}
		})
	}
}

// TestModuleDownloadRoute is responsible for testing the module download endpoint
func TestModuleDownloadRoute(t *testing.T) {
	testCases := map[string]struct {
		namespace string
		name      string
		system    string
		version   string
		expected  []byte
	}{
		"GetModuleDownload": {
			namespace: "hashicorp",
			name:      "consul",
			system:    "aws",
			version:   "0.1.0",
			expected:  []byte("Module Download"),
		},
	}

	// Handle setup and teardown
	setUp(srv.handleModuleDownload())
	defer tearDown()

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			// Build the endpoint to test
			moduleDownloadEndpoint := fmt.Sprintf(
				"%s/%s/%s/%s/%s/download",
				testServer.URL,
				testCase.namespace,
				testCase.name,
				testCase.system,
				testCase.version,
			)
			t.Logf("Module Download Endpoint: %s", moduleDownloadEndpoint)

			// Make request to test server
			resp, err := testClient.Get(moduleDownloadEndpoint)
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
			if i := bytes.Compare(respBytes, testCase.expected); i != 0 {
				t.Errorf(
					"failed getting the expected response,\n  expected: %s\n  recieved: %s\n",
					string(testCase.expected),
					string(respBytes),
				)
			}
		})
	}
}
