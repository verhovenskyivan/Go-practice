package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"net/url"
)

func main () {
	args := os.args
	
	if len(args) < 2 {
		fmt.Printf("usage: ./http-get <url>\n")
		os.Exit(1)
	}
	if _ , err := url.ParseRequestURI(args[1]);
	err != nil {
		fmt.Printf("Url invalid: %v" , err)
		os.Exit(1)
	}

	resp, err := http.Get(args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Status: %d: ", resp.StatusCode, string(body))

}