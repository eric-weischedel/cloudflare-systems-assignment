package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	url, parseError := url.Parse("https://cloudflare-assignment.eric-weischedel.workers.dev/links")
	if parseError != nil {
		log.Fatal(parseError)
	}
	response, fetchError := http.Get(url.String())
	if fetchError != nil {
		fmt.Println(fetchError)
	}

	defer response.Body.Close()

	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		log.Fatal(readError)
	}

	fmt.Println(string(body))
}
