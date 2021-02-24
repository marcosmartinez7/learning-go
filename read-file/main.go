package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	fileURI := flag.String("fileURI", "", "location of the file")
	flag.Parse()

	file, err := os.Open(*fileURI)
	if err != nil {
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
