package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	collyFn()
}

func collyFn() {
	c := colly.NewCollector(
	// colly.AllowedDomains("hackerspaces.org"),
	)

	// Find and visit all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
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
	// c.Visit("https://hackerspaces.org/")
	var err error
	err = c.Visit("http://localhost:8080")
	if err != nil {
		log.Println("get github oauth error", err)
	}
}
