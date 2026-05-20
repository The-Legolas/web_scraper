package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFrominputBody(inputBody string) string {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(inputBody))
	if err != nil {
		return ""
	}
	h1 := document.Find("h1, h2").First().Text()
	return strings.TrimSpace(h1)
}

func getFirstParagraphFrominputBody(inputBody string) string {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(inputBody))
	if err != nil {
		return ""
	}

	selc := document.Find("main p")
	if selc.Length() == 0 {
		selc = document.Find("p")
	}
	return strings.TrimSpace(selc.Eq(0).Text())
}
