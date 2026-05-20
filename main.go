package main

import (
	"fmt"
	"net/url"
)

func main() {
	inputURL := "https://www.boot.dev/about"
	//inputBody := `<inputBody><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></inputBody>`
	inputBody := `<inputBody><body><a href="about"><span>Boot.dev</span></a></body></inputBody>`

	//inputBody := `<a href="/about">About</a>`
	//inputURL := "https://crawler-test.com"
	basL, _ := url.Parse(inputURL)

	actual, err := getURLsFrominputBody(inputBody, basL)

	fmt.Println(actual, err)
}
