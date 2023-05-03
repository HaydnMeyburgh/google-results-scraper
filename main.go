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
	rand.Seed(time.Now().Unix())
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
			toScrape = append(toScrape, scrapeUrl)
		}
	} else {
		err := fmt.Errorf("Country (%s) is currently not support", countryCode)
		return nil, err
	}
	return toScrape, nil
}

// Using google url to scrape results
func googleScrape(searchTerm, countryCode, languageCode string, proxyString interface{}, pages, count, backoff int)([]SearchResult, error) {
	results := []SearchResult {}
	resultCounter := 0
	// returns built google url and error
	googlePages, err := buildGoogleUrls(searchTerm, countryCode, languageCode, pages, count)
	if err != nil {
		return nil, err
	}
	// each result from ranging through slice will be passed to 2 functions
	for _, page := range googlePages {
		// will return a response and error
		res, err := scrapeClientRequest(page, proxyString)
		if err != nil {
			return nil, err
		}
		// will return data and error
		data, err := googleParseResult(res, resultCounter)
		if err != nil {
			return nil, err
		}
		resultCounter += len(data)
		for _, result := range data {
			// Append result to results slice
			results = append(results, result)
		}
		time.Sleep(time.Duration(backoff) * time.Second)
	}
	return results, nil
}

func scrapeClientRequest(searchUrl string, proxyString interface{})(*http.Response, error) {
	baseClient := getScrapeClient(proxyString)
	req, _ = http.NewRequest("GET", searchUrl, nil)
	req.Header().Set("User-Agent", randUserAgent())

	res, err := baseClient.Do(req)
	if res.StatusCode != 200 {
		err := fmt.Errorf("Scraper received a non 200 status code suggesting a ban")
		return nil, err	
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func getScrapeClient(proxyString interface{}) *http.Client {
	switch v := proxyString.(type) {
	case string:
		proxyUrl, _ := url.Parse(v)
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	default:
		return &http.Client{}
	}
}

func main() {
	response, err := googleScrape("Haydn Meyburgh", "com", "en", nil, 1, 30, 10)
	if err == nil {
		for _, res := range response {
			fmt.Println(res)
		}
	}
}