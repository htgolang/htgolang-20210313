package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	// id
	// class class="x y"
	// tagName
	document, _ := goquery.NewDocument("http://localhost:9999/static/index.html")

	document.Find("table").Each(func(index int, section *goquery.Selection) {
		fmt.Println(index)
		fmt.Println(section.Attr("border"))
		// fmt.Println(section.Text())
		// fmt.Println(section.Html())
	})

	document.Find("a").Each(func(index int, section *goquery.Selection) {
		fmt.Println(section.Attr("href"))
	})
}
