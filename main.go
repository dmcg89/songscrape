package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// queryPtr := flag.String("artist", "jennings-waylon", " Help text.")
	// dirPtr := flag.String("dir", ".", "Help Text")
	// flag.Parse()
	// query := *queryPtr
	// dir := *dirPtr

	// } else {
	// 	// TODO: append current file
	// 	fmt.Println("file already exists")
	// }
	// scrape()
	c := colly.NewCollector(
	// colly.AllowedDomains("cowboylyrics.com"),
	)
	query := "sample"
	fName := query + ".txt"
	file, err := os.Create(fName)

	selector := "body > table:nth-child(4) > tbody > tr:nth-child(1) > td:nth-child(2) > table:nth-child(2) > tbody > tr > td:nth-child(2) > table > tbody > tr > td:nth-child(2) > table> tbody > tr > td > table > tbody > tr > td > ol > li > a"

	// On every a element which has href attribute call callback
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
		// fmt.Println(collectLyrics(r.URL.String()))
		// if checkFiles(dir, query) == false {
		// 	// create new txt file if file if file doesnt exist
		// 	fName := query + ".txt"
		// 	file, err := os.Create(fName)
		// 	defer file.Close()
		// 	check(err)
		// }
		cnts := strings.Join(collectLyrics(r.URL.String())[:], ",")
		b := []byte(cnts)
		check(err)
		// fmt.Println(cnts)
		_, err1 := file.Write(b)
		check(err1)
	})

	c.Visit("https://www.cowboylyrics.com/lyrics/jennings-waylon.html")
	// c.Visit("https://www.cowboylyrics.com/lyrics/" + query + ".html")
}

func collectLyrics(url string) []string {
	c := colly.NewCollector()
	knownUrls := []string{}
	xpath := "/html/body/table[2]/tbody/tr[1]/td[2]/table[2]/tbody/tr/td[2]/table/tbody/tr[1]/td[1]/text()"
	c.OnXML(xpath, func(e *colly.XMLElement) {
		knownUrls = append(knownUrls, e.Text)
	})
	c.Visit(url)
	return knownUrls
}

func checkFiles(dir, query string) bool {
	// checks to see if file already exists in directory
	files, err := ioutil.ReadDir(dir)
	check(err)
	for _, file := range files {
		if file.Name() == query+".txt" {
			return true
		}
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
