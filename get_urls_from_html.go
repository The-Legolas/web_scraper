package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFrominputBody(inputBodyBody string, baseURL *url.URL) ([]string, error) {
	var urlArray []string

	document, err := goquery.NewDocumentFromReader(strings.NewReader(inputBodyBody))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse inputBody: %w", err)
	}

	selc := document.Find("a[href]")
	selc.Each(func(i int, s *goquery.Selection) {
		val, ok := s.Attr("href")
		if !ok || strings.TrimSpace(val) == "" {
			return
		}
		ref, err := url.Parse(val)
		if err != nil {
			fmt.Printf("couldn't parse hrf %q: %v\n", val, err)
			return
		}
		r := baseURL.ResolveReference(ref)

		urlArray = append(urlArray, r.String())
	})

	return urlArray, nil
}
