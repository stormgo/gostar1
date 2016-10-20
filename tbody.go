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
    selection = getSelection(doc,"tbody")

	node = selection.Get(0)
    if err != nil {
		log.Fatal(err)
	}

	doc = goquery.NewDocumentFromNode(node)
    selection = getSelection(doc,"td")

    //path := ""
	//synopsis := ""

    selection.Each(func(i int, s *goquery.Selection) {
      fmt.Println(i)
    })
}

func main() {
	goq1()
}
