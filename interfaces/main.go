package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

type botWritter struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("error ", err)
		os.Exit(1)
	}

	io.Copy(botWritter{}, resp.Body)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (botWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
