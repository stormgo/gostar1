package main

import (
  "fmt"
  "log"

  "github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {

    //url := "https://godoc.org/github.com/garyburd/redigo/redis?importers"
	url := "https://godoc.org/github.com/attic-labs/testify/suite?importers"


  doc, err := goquery.NewDocument(url)
  if err != nil {
    log.Fatal(err)
  }

  // Find the review items
  doc.Find("table").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    band := s.Find("a").Text()
    //title := s.Find("i").Text()
    //fmt.Printf("Review %d: %s - %s\n", i, band, title)
    fmt.Println(band)
  })
}

func main() {
  ExampleScrape()
}
