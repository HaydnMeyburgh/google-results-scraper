package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"math/rand"
	"net/url"
	"github.com/PuerkitoBio/goquery"
)

var googleDomains = map[string]string {
	"com": "https://www.google.com/search?q=",
	"za": "https://www.google.co.za/search?q=",
}

type SearchResult struct {
	resultRank int
	resultUrl string
	resultTitle string
	resultDesc string
}

var userAgents = []string {

}

// Selecting a random user agent
func randUserAgent() string {
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

// Building google urls 
func buildGoogleUrls(searchTerm, countryCode, languageCode string, pages, count int)([]string, error) {
	toScrape := []string{}
	searchTerm := strings.Trim(searchTerm, " ")
	searchTerm := strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found {
		for i := 0; i < pages;  i++ {
			start := i * count
			scrapeUrl := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0", googleBase, searchTerm, count, languageCode, start)
		}
	} else {
		err := fmt.Errorf("Country (%s) is currently not support", countryCode)
		return nil, err
	}
	return toScrape, nil
}

// Using google url to scrape results
func googleScrape(searchTerm, countryCode, languageCode string, pages, count int)([]SearchResult, error) {
	results := []SearchResult {}
	resultCounter := 0
	// returned url to scrape and error
	googlePages, err := buildGoogleUrls(searchTerm, countryCode, languageCode, pages, count)
}

func main() {
	response, err := googleScrape("Haydn Meyburgh", "en", "com", 1, 30)
	if err == nil {
		for _, res := range response {
			fmt.Println(res)
		}
	}
}