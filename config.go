package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func (cfg *config) setPageData(normalizedURL string, data PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.pages[normalizedURL] = data
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, ok := cfg.pages[normalizedURL]
	if ok {
		return
	}
	isFirst = true
	cfg.pages[normalizedURL] = PageData{URL: normalizedURL}

	return
}

func configure(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	parsedURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		pages:              make(map[string]PageData),
		baseURL:            parsedURL,
		mu:                 new(sync.Mutex),
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 new(sync.WaitGroup),
		maxPages:           maxPages,
	}, nil
}
