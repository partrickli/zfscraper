package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	articles := [][]string{
		{"title", "auther", "magzine"},
		{"covid19", "zf", "medicine"},
		{"golang for scraper", "lgy", "golang"},
	}

	f, err := os.Create("./output.csv")
	if err != nil {
		log.Fatalln("error creating file", err)
	}

	w := csv.NewWriter(f)
	w.WriteAll(articles)
	if err := w.Error(); err != nil {
		log.Fatalln("error writing file", err)
	}
}
