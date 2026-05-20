package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFrominputBodyAbsolute(t *testing.T) {
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
			inputBody: `<inputBody><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></inputBody>`,
			expected:  []string{"https://crawler-test.com"},
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
			inputBody: `<inputBody><body><a href="/about"><span>Boot.dev</span></a></body></inputBody>`,
			expected:  []string{"https://crawler-test.com/about"},
		},
		{
			name:      "test bad sub link",
			inputURL:  "https://crawler-test.com",
			inputBody: `<inputBody><body><a href="bad"><span>Boot.dev</span></a></body></inputBody>`,
			expected:  []string{"https://crawler-test.com/bad"},
		},
		{
			name:      "test two links",
			inputURL:  "https://crawler-test.com",
			inputBody: `<inputBody><body><a href="https://crawler-test1.com"><span>Boot.dev</span></a><a href="https://crawler-test2.com"><span>Boot.dev</span></a></body></inputBody>`,
			expected:  []string{"https://crawler-test1.com", "https://crawler-test2.com"},
		},
		{
			name:      "test two different types of links",
			inputURL:  "https://crawler-test.com",
			inputBody: `<inputBody><body><a href="blog"><span>Boot.dev</span></a><a href="https://crawler-test2.com"><span>Boot.dev</span></a></body></inputBody>`,
			expected:  []string{"https://crawler-test.com/blog", "https://crawler-test2.com"},
		},
		{
			name:      "test mix",
			inputURL:  "https://crawler-test.com/blog",
			inputBody: `<inputBody><body><a href=""></a><a href="/ok"></a><a href="https://external.com"></a><a href="relative"></a></body></inputBody>`,
			expected:  []string{"https://crawler-test.com/ok", "https://external.com", "https://crawler-test.com/relative"},
		},
		{
			name:      "invalid href URL",
			inputURL:  "https://crawler-test.com",
			inputBody: `<inputBody>body><a href=":\\invalidURL"><span>Boot.dev</span></a></body></inputBody>`,
			expected:  nil,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("couldn't parse input URL: %v", err)
				return
			}

			actual, err := getURLsFrominputBody(tc.inputBody, baseURL)

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
