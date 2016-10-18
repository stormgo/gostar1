package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://godoc.org/github.com/garyburd/redigo/redis?importers"

	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("HTML:\n\n", string(bytes))

	resp.Body.Close()
}
