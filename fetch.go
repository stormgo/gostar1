package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
    "golang.org/x/net/html"
)

func main() {

	//url := "https://godoc.org/github.com/garyburd/redigo/redis?importers"
	url := "https://godoc.org/github.com/attic-labs/testify/suite?importers"

	resp, _ := http.Get(url)
	//bytes, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println("HTML:\n\n", string(bytes))

	//resp.Body.Close()

    z := html.NewTokenizer(resp.Body)

    for {
        tt := z.Next()
        //fmt.Println(tt)

        switch {
        case tt == html.ErrorToken:
        	fmt.Println("End of the document, we're done")
            return
        case tt == html.StartTagToken:
            t := z.Token()

            isAnchor := t.Data == "a"
            if isAnchor {
                fmt.Println(t)
            }
        }

    }

}
