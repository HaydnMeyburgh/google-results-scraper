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

func googleScrape()([]SearchResult, err) {
	results := []SearchResult {}
	resultCounter := 0
	buildGoogleUrls()
}

func main() {
	response, err := googleScrape("Haydn Meyburgh")
	if err == nil {
		for _, res := range response {
			fmt.Println(res)
		}
	}
}