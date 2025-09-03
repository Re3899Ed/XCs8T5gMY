// 代码生成时间: 2025-09-03 14:08:00
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
    "strings"
    "golang.org/x/net/html"
)

// ScrapeWebContent defines a function to scrape content from a webpage.
// It takes a URL as an argument and returns the HTML content of the webpage or an error.
func ScrapeWebContent(url string) (string, error) {
    // Send an HTTP GET request to the specified URL.
    response, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer response.Body.Close()

    // Check if the response status code is 200 OK.
    if response.StatusCode != http.StatusOK {
        return "", fmt.Errorf("received non-200 status code: %d", response.StatusCode)
    }

    // Read the body of the response into a string.
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return "", err
    }

    // Convert the body into a string and return it.
    return string(body), nil
}

// main is the entry point of the program.
func main() {
    // URL to scrape.
    url := "http://example.com"

    // Call the ScrapeWebContent function and handle the result.
    content, err := ScrapeWebContent(url)
    if err != nil {
        log.Println("Error scraping content:", err)
        return
    }

    // Print the scraped content to the console.
    fmt.Println("Scraped content:")
    fmt.Println(string(content))
}
