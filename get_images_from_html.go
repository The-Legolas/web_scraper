package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	imgArray := []string{}

	document, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	selc := document.Find("img[src]")

	selc.Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("src")
		if val == "" {
			return
		}
		ref, err := url.Parse(val)
		if err != nil {
			fmt.Println("error;", err)
			return
		}
		r := baseURL.ResolveReference(ref)

		imgArray = append(imgArray, r.String())
	})

	return imgArray, nil
}
