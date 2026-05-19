package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	selc := document.Find("h1")
	if selc.Size() == 0 {
		selc = document.Find("h2")
	}
	return selc.First().Text()
}

func getFirstParagraphFromHTML(html string) string {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	selc := document.Find("main p")
	if selc.Length() == 0 {
		selc = document.Find("p")
	}
	return selc.Eq(0).Text()
}
