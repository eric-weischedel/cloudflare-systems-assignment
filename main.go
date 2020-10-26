package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/url"
)

const defaultURL string = "https://cloudflare-assignment.eric-weischedel.workers.dev/links"

func main() {
	flag.Usage = func() {
		fmt.Println("This CLI is used to make HTTP requests to a given endpoint and profile the server's performance.")
		flag.PrintDefaults()
	}
	var urlFlag = flag.String("url", defaultURL, "The URL to send the HTTP request to")
	var profileFlag = flag.Int("profile", 10, "The number of requests to make")
	flag.Parse()

	fmt.Println(*profileFlag)

	url, parseError := url.Parse(*urlFlag)
	if parseError != nil {
		log.Fatal(parseError)
	}

	connection, dialError := net.Dial("tcp", url.Host+":80")
	if dialError != nil {
		log.Fatal(dialError)
	}

	rt := fmt.Sprintf("GET %v HTTP/1.1\r\n", url.Path)
	rt += fmt.Sprintf("Host: %v\r\n", url.Host)
	rt += fmt.Sprintf("Connection: close\r\n")
	rt += fmt.Sprintf("\r\n")

	_, writeError := connection.Write([]byte(rt))
	if writeError != nil {
		log.Fatal(writeError)
	}

	response, readError := ioutil.ReadAll(connection)
	if readError != nil {
		log.Fatal(readError)
	}
	fmt.Println(string(response))

	connection.Close()
}
