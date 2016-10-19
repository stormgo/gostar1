package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func getSelection(doc *goquery.Document, selector string) (s *goquery.Selection) {
	s = doc.Find(selector)
	return s
}

func goq1() {

	//url := "https://godoc.org/github.com/garyburd/redigo/redis?importers"
	url := "https://godoc.org/github.com/attic-labs/testify/suite?importers"

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	selection := getSelection(doc, "table.table.table-condensed")

    node := selection.Get(0)
    if err != nil {
		log.Fatal(err)
	}

    doc = goquery.NewDocumentFromNode(node)
    selection = getSelection(doc,"a")

    github := ""
    selection.Each(func(i int, s *goquery.Selection) {
      github, _ = s.Attr("href")
      fmt.Println(i,github)
    })
}

func main() {
	goq1()
}
