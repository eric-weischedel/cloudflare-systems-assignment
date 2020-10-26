package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const defaultURL string = "https://cloudflare-assignment.eric-weischedel.workers.dev/links"

/*
	Makes a request to the specified URL and returns
	1) the duration it took in milliseconds,
	2) the length of the response in bytes,
	3) and the HTTP status code
*/
func makeRequest(url url.URL) (time.Duration, int, int) {
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

	response, err := ioutil.ReadAll(connection)
	if err != nil {
		log.Fatal(err)
	}

	// get status code
	i := strings.Index(string(response), " ")
	codeString := string(response)[i+1 : i+4]
	code, convertError := strconv.Atoi(codeString)
	if convertError != nil {
		log.Fatal(convertError)
	}

	connection.Close()

	return time.Now().Sub(beginningTime), len(response), code
}

func main() {
	// parse arguments
	flag.Usage = func() {
		fmt.Println("This CLI is used to make HTTP requests to a given endpoint and profile the server's performance.")
		flag.PrintDefaults()
	}
	urlFlag := flag.String("url", defaultURL, "The URL to send the HTTP request to")
	profileFlag := flag.Int("profile", 10, "The number of requests to make")
	flag.Parse()

	url, parseError := url.Parse(*urlFlag)
	if parseError != nil {
		log.Fatal(parseError)
	}

	// make http requests
	durations := make([]time.Duration, *profileFlag)
	lengths := make([]int, *profileFlag)
	statusCodes := make([]int, *profileFlag)
	total := int64(0)
	for i := 0; i < *profileFlag; i++ {
		duration, length, statusCode := makeRequest(*url)
		durations[i] = duration
		lengths[i] = length
		statusCodes[i] = statusCode
		total += int64(duration)
	}

	// sort durations
	sort.Slice(durations, func(i, j int) bool {
		return int64(durations[i]) < int64(durations[j])
	})

	// sort lengths
	sort.Slice(lengths, func(i, j int) bool {
		return lengths[i] < lengths[j]
	})

	// parse status codes
	errorCodes := make([]int, 0)
	errorCount := 0
	for i := 0; i < len(statusCodes); i++ {
		if statusCodes[i] > 399 {
			errorCodes = append(errorCodes, statusCodes[i])
			errorCount++
		}
	}

	// generate output
	fmt.Println("Number of requests:", *profileFlag)
	fmt.Println("Fastest time:", durations[0])
	fmt.Println("Slowest time:", durations[*profileFlag-1])
	fmt.Println("Mean time:", time.Duration(total/int64(*profileFlag)))
	fmt.Println("Median time:", durations[*profileFlag/2])
	fmt.Println("Percentage succeeded:", (float32(*profileFlag-errorCount)/float32(*profileFlag))*100)
	fmt.Println("Error codes:", errorCodes)
	fmt.Println("Smallest response (in bytes):", lengths[0])
	fmt.Println("Largest response (in bytes):", lengths[*profileFlag-1])

}
