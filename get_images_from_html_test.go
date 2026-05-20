package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestGetImagesFrominputBodyRelative(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		{
			name:      "test function",
			inputURL:  "https://crawler-test.com",
			inputBody: `<inputBody><body><img src="https://crawler-test.com/logo.png" alt="Logo"></body></inputBody>`,
			expected:  []string{"https://crawler-test.com/logo.png"},
		},
		{
			name:      "test no link",
			inputURL:  "https://crawler-test.com",
			inputBody: `<inputBody><body><a href=""><span>Boot.dev</span></a></body></inputBody>`,
			expected:  nil,
		},
		{
			name:      "test sub link",
			inputURL:  "https://crawler-test.com",
			inputBody: `<inputBody><body><img src="/logo.png" alt="Logo"></body></inputBody>`,
			expected:  []string{"https://crawler-test.com/logo.png"},
		},
		{
			name:      "test bad sub link",
			inputURL:  "https://crawler-test.com",
			inputBody: `<inputBody><body><img src="bad.png" alt="Logo"></body></inputBody>`,
			expected:  []string{"https://crawler-test.com/bad.png"},
		},
		{
			name:      "test two links",
			inputURL:  "https://crawler-test.com",
			inputBody: `<inputBody><body><img src="/logo1.png" alt="Logo"><img src="/logo2.png" alt="Logo"></body></inputBody>`,
			expected:  []string{"https://crawler-test.com/logo1.png", "https://crawler-test.com/logo2.png"},
		},
		{
			name:      "test two different types of links",
			inputURL:  "https://crawler-test.com",
			inputBody: `<inputBody><body><img src="https://crawler-test.com/blog/logo.png" alt="Logo"><img src="/logo.png" alt="Logo"></body></inputBody>`,
			expected:  []string{"https://crawler-test.com/blog/logo.png", "https://crawler-test.com/logo.png"},
		},
		{
			name:      "test mix",
			inputURL:  "https://crawler-test.com/blog",
			inputBody: `<inputBody><body><img src="" alt="Logo"><img src="/ok" alt="Logo"><img src="https://external.com/logo.png" alt="Logo"><img src="/logo.png" alt="Logo"></body></inputBody>`,
			expected:  []string{"https://crawler-test.com/ok", "https://external.com/logo.png", "https://crawler-test.com/logo.png"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("couldn't parse input URL: %v", err)
				return
			}

			actual, err := getImagesFrominputBody(tc.inputBody, baseURL)

			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected %q, got %q", i, tc.name, tc.expected, actual)
			}
		})
	}
}
