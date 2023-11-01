package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func filterQuery(query string) string {
	query = strings.Join(strings.Fields(strings.TrimSpace(query)), " ")
	query = strings.ReplaceAll(query, " ", "+")
	query = html.EscapeString(query)
	return query
}

func googleSearch(query, output string) (err error) {
	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", filterQuery(query))

	response, err := http.Get(searchURL)
	if err != nil {
		return
	}
	defer response.Body.Close()

	htmlContent, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	htmlFile, err := os.Create(output + ".html")
	if err != nil {
		return
	}
	defer htmlFile.Close()

	_, err = htmlFile.Write(htmlContent)
	if err != nil {
		return
	}

	return
}

func main() {
	query := "Munchkin Cat"
	output := "index"
	err := googleSearch(query, output)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output, "file generated successfully!")
}
