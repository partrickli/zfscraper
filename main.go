package main

import (
	"fmt"

	"github.com/partrickli/zfscraper/scraper"
)

func main() {
	rc := make(chan int)
	scraper.Scrapeidata(rc)

	result := <-rc
	fmt.Println(result)
}
