package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	links := []string{
		"http://bitcoin.org",
		"http://google.com",
		"http://stackoverflow.com",
		"http://facebook.com",
	}

	c := make(chan string) // create channel

	for _, link := range links {
		go printStatus(link, c) // create subroutine
	}

	for l := range c { // blocking call
		go func(l string, c chan string) {
			time.Sleep(time.Second * 5)
			printStatus(l, c)
		}(l, c)
	}

}

func isLinkOnline(link string) bool {
	_, error := http.Get(link)
	if error != nil {
		return false
	}
	return true
}

func printStatus(link string, c chan string) {
	fmt.Printf("Is %s responding? %s \n", link, strconv.FormatBool(isLinkOnline(link)))
	c <- link
}
