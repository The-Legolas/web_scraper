package main

import (
	"fmt"
	"net/url"
)

func main() {
	inputURL := "https://www.boot.dev"
	//inputBody := `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`
	inputBody := `<html><body><a href="about"><span>Boot.dev</span></a></body></html>`

	//inputBody := `<a href="/about">About</a>`
	//inputURL := "https://crawler-test.com"
	basL, _ := url.Parse(inputURL)

	actual, err := getURLsFromHTML(inputBody, basL)

	fmt.Println(actual, err)
}
