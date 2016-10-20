package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func getRow(selection *goquery.Selection) {

	node := selection.Get(0)

	doc := goquery.NewDocumentFromNode(node)
    selection1 := getSelection(doc,"td")

	path := ""
	synopsis := ""
	selection1.Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			selection2 := getSelection(doc,"a")
			path, _ = selection2.Attr("href")
      		fmt.Println("column ",i, " ",path)
		}
		if i == 1 {
			synopsis = s.Text()
			if synopsis != "" {
			 	fmt.Println("column ",i, " ",synopsis)
			} else {
				fmt.Println("column ",i)
			}
		}
    })
}

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
    selection = getSelection(doc,"tr")

	// So now you want to create your map and pass the map into
	// the function which builds it up with each td
	// namely the path and the synopsis

	// pass in the selection and get back the map

    //path := ""
	//synopsis := ""

    selection.Each(func(i int, s *goquery.Selection) {
      	fmt.Println("row ",i)
		getRow(s)
    })
}

func main() {
	goq1()
}
