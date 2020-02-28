package hmeta

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseDescription(doc *goquery.Document) []string {
	ret := []string{}
	doc.Find("meta").Each(func(i int, sel *goquery.Selection) {
		name, ok := sel.Attr("name")
		if !ok {
			return
		}
		if name == "description" {
			description, _ := sel.Attr("content")
			ret = append(ret, description)
		}
	})

	return ret
}

func ParseKeyWords(doc *goquery.Document) []string {
	ret := []string{}
	doc.Find("meta").Each(func(i int, sel *goquery.Selection) {
		name, ok := sel.Attr("keywords")
		if !ok {
			return
		}
		if name == "description" {
			content, _ := sel.Attr("content")
			ret = append(ret, strings.Split(content, ",")...)
		}
	})

	return ret
}
