package main

import (
	"a-test/controllers/getbody"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

// Article is...
type Article struct {
	Title       string
	URL         string
	Thumbnail   string
	Description string
	Body        string
}

func main() {
	// Declare a new request.
	request := gorequest.New()

	// request get data.
	_, body, errs := request.Get("https://vnexpress.net/kinh-doanh").End()
	if errs != nil {
		panic(errs)
	}
	// get the document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
	//declare a variable rows as a slice of Article. leng=0
	rows := make([]Article, 0)

	doc.Find(".sidebar_1 .list_news").Each(func(i int, sel *goquery.Selection) {
		row := new(Article) // new article
		row.Title = sel.Find("a").Text()
		row.URL, _ = sel.Find("a").Attr("href")
		row.Thumbnail, _ = sel.Find("img").Attr("src")
		row.Description = sel.Find("p.description").Text()
		row.Body = getbody.Getbody(row.URL)
		rows = append(rows, *row)

	})
	// convert to json string.
	bts, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	//

	log.Println(string(bts))
	fmt.Println("Total craw:", len(rows))

}
