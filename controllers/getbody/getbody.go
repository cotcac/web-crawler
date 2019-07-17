package getbody

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

// Getbody is a function...
func Getbody(url string) string {
	fmt.Println("get body")
	request := gorequest.New()

	// request get data.
	_, body, errs := request.Get(url).End()
	if errs != nil {
		panic(errs)
	}
	// get the document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
	// get body
	b, _ := doc.Find("article.content_detail").Html()
	return b
}
