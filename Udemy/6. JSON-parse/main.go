package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"net/url"
    "io"
)

type Words struct {
    page string 'json:"page"'
    words []string 'json:"words"'
}

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
	if resp.StatusCode != 200 {
        fmt.Printf("Status: %d: ", resp.StatusCode, string(body))
        os.Exit(1)
    }

    var words Words
    err = json.Unmarshal(body, &words)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("JSON page %v", words.page, words.Words)
}
