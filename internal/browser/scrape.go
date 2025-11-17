package browser

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/arrisar/myrientui/internal/myrient"
	"golang.org/x/net/html"
)

var base = "https://myrient.erista.me/files"

type ScrapeResultMsg struct {
	options []myrient.Option
	err     error
}

/**
 * Scrape a page and extract the options
 */
func Scrape(path string) (msg ScrapeResultMsg) {
	uri := fmt.Sprintf("%s%s", base, path)

	doc, err := getDocument(uri)
	if err != nil {
		msg.err = err
		return
	}

	msg.options, msg.err = parseOptions(doc)
	return
}

/**
 * Fetch the document
 */
func getDocument(uri string) (doc *goquery.Document, err error) {
	resp, err := http.Get(uri)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	return
}

/**
 * Parse the document for link table rows
 */
func parseOptions(doc *goquery.Document) (options []myrient.Option, err error) {
	rows := doc.Find("#list tbody tr").Nodes
	for _, row := range rows {
		option := myrient.Option{}
		include := true

		cell := row.FirstChild
		for cell != nil {
			processed := false

			// link
			if !processed && hasClass("link", cell) {
				link := cell.FirstChild
				text := link.FirstChild.Data

				if text == "Parent directory/" || text == "./" || text == "../" {
					include = false
					break
				}

				option.Label = text
				option.Link = getNodeAttr("href", link).Val
				option.IsDir = strings.HasSuffix(option.Link, "/")
				processed = true
			}

			// size
			if !processed && hasClass("size", cell) {
				option.Size = cell.FirstChild.Data
				processed = true
			}

			// date
			if !processed && hasClass("date", cell) {
				option.Date = cell.FirstChild.Data
				processed = true
			}

			// next cell
			cell = cell.NextSibling
		}

		if include {
			options = append(options, option)
		}
	}

	return
}

/**
 * Get a specific attribute from the node
 */
func getNodeAttr(needle string, node *html.Node) (res html.Attribute) {
	for _, attr := range node.Attr {
		if attr.Key == needle {
			res = attr
			return
		}
	}

	return
}

/**
 * Check if a class exists on the node
 */
func hasClass(needle string, node *html.Node) bool {
	attr := getNodeAttr("class", node)
	if attr.Key == "" {
		return false
	}

	classes := strings.Fields(attr.Val)
	for _, class := range classes {
		if class == needle {
			return true
		}
	}

	return false
}
