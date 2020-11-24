package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://amazon.com",
		"https://naver.com",
		"https://golang.org",
		"https://stackoverflow.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		return
	}

	fmt.Println(link, " is up!")
}
