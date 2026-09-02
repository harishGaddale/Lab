package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://go.dev/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(lnk string, c chan string) {
	_, err := http.Get(lnk)
	if err != nil {
		fmt.Println(lnk, " is down")
		c <- lnk
		return
	}
	fmt.Println(lnk, " is up")
	c <- lnk
}
