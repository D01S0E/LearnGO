package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	urls := "https://httpbin.org"

	resp, err := http.Get(urls)

	if err != nil {
		log.Fatal(err)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Printf("%s", body)
	}
	resp.Body.Close()
}
