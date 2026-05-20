package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFrominputBody(inputBodyBody string, baseURL *url.URL) ([]string, error) {
	var imgArray []string

	document, err := goquery.NewDocumentFromReader(strings.NewReader(inputBodyBody))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse inputBody: %w", err)
	}

	selc := document.Find("img[src]")

	selc.Each(func(i int, s *goquery.Selection) {
		val, ok := s.Attr("src")
		if !ok || strings.TrimSpace(val) == "" {
			return
		}
		ref, err := url.Parse(val)
		if err != nil {
			fmt.Printf("couldn't parse src %q: %v\n", val, err)
			return
		}
		r := baseURL.ResolveReference(ref)

		imgArray = append(imgArray, r.String())
	})

	return imgArray, nil
}
