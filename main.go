package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	inState, outState := "", ""
	fmt.Println("This is a simple program to web scrape tuition from Weber State University")

	res, err := http.Get("https://www.weber.edu/admissions/shared/costs.html")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Getting document failed: ")
	}

	inState = document.Find("#Body > main > div > div > div > div.large-9.columns > table > tbody:nth-child(2) > tr:nth-child(1) > td:nth-child(2)").First().Text()

	outState = document.Find("#Body > main > div > div > div > div.large-9.columns > table > tbody:nth-child(2) > tr:nth-child(1) > td:nth-child(3)").First().Text()

	fmt.Println("In state tuition: ", inState)
	fmt.Println("Out of state tuition: ", outState)
}
