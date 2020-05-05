package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "idata.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"标题", "作者"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("li.aca_algo", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("h2"),
			e.ChildText("div.caption_author"),
		})
	})

	c.OnHTML("a.sb_pagN", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// start scaping the page under the link found
		e.Request.Visit(link)
	})

	searchlink := "https://cn.bing.com/academic/search?q=covid-19"
	c.Visit(searchlink)

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
