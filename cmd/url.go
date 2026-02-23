package main

import (
	"fmt"
	"github.com/praffq/go-url"
	"log"
)

func main() {
	uri, err := url.Parse("http://example.com/sample")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Scheme:", uri.Scheme)
	fmt.Println("Host:", uri.Host)
	fmt.Println("Path:", uri.Path)
	fmt.Println("String:", uri.String())
}
