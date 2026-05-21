package main

import (
	"fmt"
	"net/url"
)

type PageData struct {
	URL            string   `json:"url"`
	Heading        string   `json:"heading"`
	FirstParagraph string   `json:"first_paragraph"`
	OutgoingLinks  []string `json:"outgoing_links"`
	ImageURLs      []string `json:"image_urls"`
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
