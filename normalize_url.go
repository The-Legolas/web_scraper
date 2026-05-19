package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	noSpace := strings.TrimSpace(rawURL)
	outputURL, err := url.Parse(noSpace)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	path := outputURL.Host + outputURL.Path
	normalized := strings.TrimSuffix(path, "/")
	lowerCase := strings.ToLower(normalized)

	return lowerCase, nil
}
