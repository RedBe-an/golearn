package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func main() {
	var results = make(map[string]string)
	urls := []string{
		"https://www.google.com",
		"https://www.stackoverflow.com",
		"https://www.discord.com",
		"https://www.reddit.com",
		"https://www.instagram.com",
		"https://www.go.dev",
		"https://www.github.com",
	}
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking :", url)
	resp, err := http.Get(url)
	if resp == nil {
		fmt.Println(err)
		return errRequestFailed
	}
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
