package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	// Selectors for in state and out of state tuition
	inStateSelector = "#Body > main > div > div > div > div.large-9.columns > table > tbody:nth-child(2) > tr:nth-child(1) > td:nth-child(2)"
	outStateSelector = "#Body > main > div > div > div > div.large-9.columns > table > tbody:nth-child(2) > tr:nth-child(1) > td:nth-child(3)"
)

func main() {
	inState, outState := "", ""
	fmt.Println("This is a simple program to web scrape tuition from Weber State University")

	res, err := http.Get("https://www.weber.edu/admissions/shared/costs.html")
	if err != nil {
		log.Fatal("Error when getting the page: ", err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error when creating document: ", err)
	}

	// Scrape tuition for in state and out of state
	inState = document.Find(inStateSelector).First().Text()
	outState = document.Find(outStateSelector).First().Text()

	fmt.Println("In state tuition: ", inState)
	fmt.Println("Out of state tuition: ", outState)
}
