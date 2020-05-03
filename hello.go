package main

import (
	"fmt"

	"github.com/partrickli/zfscraper/magzine"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	c := magzine.Article()
	fmt.Println(c)
}
