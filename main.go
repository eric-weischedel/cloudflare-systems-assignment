package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"sort"
	"time"
)

const defaultURL string = "https://cloudflare-assignment.eric-weischedel.workers.dev/links"

func makeRequest(url url.URL) time.Duration {
	beginningTime := time.Now()

	connection, dialError := net.Dial("tcp", url.Host+":80")
	if dialError != nil {
		log.Fatal(dialError)
	}

	httpString := fmt.Sprintf("GET %v HTTP/1.1\r\n", url.Path)
	httpString += fmt.Sprintf("Host: %v\r\n", url.Host)
	httpString += fmt.Sprintf("Connection: close\r\n")
	httpString += fmt.Sprintf("\r\n")

	_, writeError := connection.Write([]byte(httpString))
	if writeError != nil {
		log.Fatal(writeError)
	}

	connection.Close()

	endingTime := time.Now()

	return endingTime.Sub(beginningTime)
}

func main() {
	// set up command line arguments
	flag.Usage = func() {
		fmt.Println("This CLI is used to make HTTP requests to a given endpoint and profile the server's performance.")
		flag.PrintDefaults()
	}
	urlFlag := flag.String("url", defaultURL, "The URL to send the HTTP request to")
	profileFlag := flag.Int("profile", 10, "The number of requests to make")
	flag.Parse()

	// parse url
	url, parseError := url.Parse(*urlFlag)
	if parseError != nil {
		log.Fatal(parseError)
	}

	// make requests
	durations := make([]time.Duration, *profileFlag)
	var total int64 = 0
	for i := 0; i < *profileFlag; i++ {
		duration := makeRequest(*url)
		durations[i] = duration
		total += int64(duration)
	}

	sort.Slice(durations, func(i, j int) bool {
		return int64(durations[i]) < int64(durations[j])
	})

	// generate output
	fmt.Println("Number of requests:", *profileFlag)
	fmt.Println("Fastest time:", durations[0])
	fmt.Println("Slowest time:", durations[*profileFlag-1])
	fmt.Println("Mean time:", time.Duration(total/int64(*profileFlag)))
	fmt.Println("Median time:", durations[*profileFlag/2])
	// fmt.Println("Percentage succeeded:", *profileFlag)
	// fmt.Println("Error codes:", *profileFlag)
	// fmt.Println("Smallest response (in bytes):", *profileFlag)
	// fmt.Println("Largest response (in bytes):", *profileFlag)

}
