package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]

	maxConcurrency := 3
	var err error
	if len(os.Args) >= 3 {
		maxConcurrency, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error converting max concurrency to number")
			return
		}
	}

	maxPages := 10
	if len(os.Args) == 4 {
		maxPages, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Error converting max concurrency to number")
			return
		}
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error  - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %v...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL := range cfg.pages {
		fmt.Printf("found: %s\n", normalizedURL)
	}
}
