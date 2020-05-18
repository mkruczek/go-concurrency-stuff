package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	checked := make(chan string)
	urls := []string{"http://google.com", "http://nonExist.io", "http://tour.golang.org", "http://stackoverflo.com"}

	for _, v := range urls {
		go makeRequest(v, checked)
	}

	for s := range checked {
		go func(url string) {
			time.Sleep(5 * time.Second)
			makeRequest(url, checked)
		}(s)
	}
}

func makeRequest(url string, checked chan string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error for %s : %s\n", url, err.Error())
		checked <- url
		return
	}
	sc := resp.StatusCode
	log.Printf("Status Code for %s : %d\n", url, sc)
	checked <- url
}
