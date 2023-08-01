package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
	}

	for _, link := range links {
		pingUrl(link)
	}
}

func pingUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
		return
	}

	fmt.Println(url, "is up")
}
