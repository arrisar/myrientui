package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/net/html"
)

var base = "https://myrient.erista.me/files"

type Scraper struct {
	results chan ResultsMsg
}

func New() Scraper {
	s := Scraper{}
	s.results = make(chan ResultsMsg)
	return s
}

func (s Scraper) StartScrape(path string) tea.Cmd {
	return tea.Batch(
		func() tea.Msg { return StartedMsg{} },
		func() tea.Msg {
			results, err := s.DoScrape(path)
			return ResultsMsg{results, err}
		})
}

/**
 * Scrape a page and extract the results
 */
func (s Scraper) DoScrape(path string) (res []Result, err error) {
	uri := fmt.Sprintf("%s%s", base, path)

	var doc *goquery.Document
	doc, err = s.getDocument(uri)
	if err != nil {
		return
	}

	res, err = s.parseResults(doc)
	return
}

/**
 * Fetch the document
 */
func (s Scraper) getDocument(uri string) (doc *goquery.Document, err error) {
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
func (s Scraper) parseResults(doc *goquery.Document) (results []Result, err error) {
	rows := doc.Find("#list tbody tr").Nodes
	for _, row := range rows {
		result := Result{}
		include := true

		cell := row.FirstChild
		for cell != nil {
			processed := false

			// link
			if !processed && s.hasClass("link", cell) {
				link := cell.FirstChild
				text := link.FirstChild.Data

				if text == "Parent directory/" || text == "./" || text == "../" {
					include = false
					break
				}

				result.Label = text
				result.Link = s.getNodeAttr("href", link).Val
				result.IsDir = strings.HasSuffix(result.Link, "/")
				processed = true
			}

			// size
			if !processed && s.hasClass("size", cell) {
				result.Size = cell.FirstChild.Data
				processed = true
			}

			// date
			if !processed && s.hasClass("date", cell) {
				result.Date = cell.FirstChild.Data
				processed = true
			}

			// next cell
			cell = cell.NextSibling
		}

		if include {
			results = append(results, result)
		}
	}

	return
}

/**
 * Get a specific attribute from the node
 */
func (s Scraper) getNodeAttr(needle string, node *html.Node) (res html.Attribute) {
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
func (s Scraper) hasClass(needle string, node *html.Node) bool {
	attr := s.getNodeAttr("class", node)
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
