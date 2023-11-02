package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	googleSearchURL = "https://www.google.com/search"
)

func filterQuery(query string) string {
	query = strings.Join(strings.Fields(strings.TrimSpace(query)), " ")
	query = url.QueryEscape(query)
	return query
}

func googleSearch(query, output string) (err error) {
	searchURL := fmt.Sprintf("%s?q=%s", googleSearchURL, filterQuery(query))

	response, err := http.Get(searchURL)
	if err != nil {
		return
	}
	defer response.Body.Close()

	htmlContent, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	outputFile := output + ".html"
	return os.WriteFile(outputFile, htmlContent, os.ModePerm)
}

func main() {
	query := "  Munchkin Cat   is   --    6  @  __  _ cute -  🐈  :) "
	output := "index"
	err := googleSearch(query, output)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output, "file generated successfully!")
}
