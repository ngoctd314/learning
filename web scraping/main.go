package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
	// colly.AllowedDomains("hackerspaces.org"),
	)

	var err error
	err = c.Visit("https://github.com/login/oauth/authorize?client_id=09f284058a21a54ac468&redirect_uri=http://localhost:8080/oauth/redirect")
	if err != nil {
		log.Println("get github oauth error", err)
	}

	// Find and visit all links
	c.OnHTML("p", func(e *colly.HTMLElement) {
		// Print link
		// fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		// err := c.Visit(e.Request.AbsoluteURL(link))
		// if err != nil {
		// 	log.Println(err)
		// }
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")
}
