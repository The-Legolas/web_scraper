package main

import (
	"fmt"
	"net/url"
)

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(inputBody, inputURL string) PageData {
	heading := getHeadingFrominputBody(inputBody)
	firstParagraph := getFirstParagraphFrominputBody(inputBody)

	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		fmt.Printf("couldn't parse inputBody: %v", err)
		return PageData{
			URL:            inputURL,
			Heading:        heading,
			FirstParagraph: firstParagraph,
			OutgoingLinks:  nil,
			ImageURLs:      nil,
		}
	}

	imageURLs, err := getImagesFrominputBody(inputBody, parsedURL)
	if err != nil {
		fmt.Printf("couldn't get images: %v", err)
		imageURLs = nil
	}

	outgoingLinks, err := getURLsFrominputBody(inputBody, parsedURL)
	if err != nil {
		fmt.Printf("couldn't get url's: %v", err)
		outgoingLinks = nil
	}

	return PageData{
		URL:            parsedURL.String(),
		Heading:        heading,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      imageURLs,
	}
}
