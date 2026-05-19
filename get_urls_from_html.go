package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	urlArray := []string{}

	document, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	selc := document.Find("a[href]")
	selc.Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("href")
		if val == "" {
			return
		}
		ref, err := url.Parse(val)
		if err != nil {
			fmt.Println("error;", err)
			return
		}
		r := baseURL.ResolveReference(ref)

		urlArray = append(urlArray, r.String())
	})

	return urlArray, nil
}
