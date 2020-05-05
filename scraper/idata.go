package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/partrickli/zfscraper/magzine"
)

// Scrapeidata is searching articles from idata.com
func Scrapeidata(result chan<- int) {
	searchlink := "https://search.cn-ki.net/search?keyword=covid-19&db=CFLS"
	fmt.Println(searchlink)

	c := colly.NewCollector()
	r := make(map[int]magzine.Article, 0)

	c.OnHTML("div.mdui-row.mdui-typo span:nth-child(3)", func(e *colly.HTMLElement) {
		title := e.Text
		i := e.Index

		if atcl, found := r[i]; found {
			atcl.Title = title
			r[i] = atcl
		} else {
			r[i] = magzine.Article{Title: title}
		}
	})

	c.OnHTML("div.mdui-row.mdui-typo h3", func(e *colly.HTMLElement) {
		mz := e.Text
		i := e.Index

		if atcl, found := r[i]; found {
			atcl.Magzine = mz
			r[i] = atcl
		} else {
			r[i] = magzine.Article{Magzine: mz}
		}
	})

	c.OnScraped(func(res *colly.Response) {
		for k, v := range r {
			fmt.Println(k)
			fmt.Printf("%#v\n", v)
		}
		result <- 0
	})

	c.Visit(searchlink)

}
