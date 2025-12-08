package scraper

import (
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/net/html"
)

var base = "https://myrient.erista.me/files"

func New(cacheDir string) Scraper {
	s := Scraper{}

	s.Config = Config{}
	s.Config.CacheDir = cacheDir
	s.Config.Workers = 32

	return s
}

func (s Scraper) MsgScrape(path string) tea.Cmd {
	return tea.Batch(
		func() tea.Msg { return StartedMsg{} },
		func() tea.Msg {
			index, err := s.DoScrape(path)
			return ResultsMsg{index, err}
		})
}

func (s Scraper) DoScrape(path string) (index Index, err error) {
	if index, err := s.CachedIndexRead(path); err == nil {
		return index, nil
	}

	if index, err = s.ScrapePath(path); err != nil {
		return
	}

	s.CachedIndexWrite(index)
	return
}

/**
 * Scrape a page and extract the results
 */
func (s Scraper) ScrapePath(path string) (index Index, err error) {
	index.Path = path

	var doc *goquery.Document
	doc, err = s.getDocument(base + path)
	if err != nil {
		return
	}

	err = s.parseResults(doc, &index)
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
func (s Scraper) parseResults(doc *goquery.Document, index *Index) error {
	rows := doc.Find("#list tbody tr").Nodes
	for _, row := range rows {
		result := Link{}
		include := true
		currDir := false

		cell := row.FirstChild
		for cell != nil {
			processed := false

			// link
			if !processed && include && s.hasClass("link", cell) {
				processed = true
				link := cell.FirstChild
				text := link.FirstChild.Data

				if text == "./" {
					currDir = true
				}

				if text == "Parent directory/" || text == "./" || text == "../" {
					include = false
				} else {
					result.Label = text
					result.Link = s.getNodeAttr("href", link).Val
					result.IsDir = strings.HasSuffix(result.Link, "/")
				}
			}

			// size
			if !processed && include && s.hasClass("size", cell) {
				processed = true
				if cell.FirstChild.Data != "-" {
					result.Size = cell.FirstChild.Data
				}
			}

			// date
			if !processed && s.hasClass("date", cell) {
				processed = true
				if date, err := time.Parse("02-Jan-2006 15:04", cell.FirstChild.Data); err == nil {
					result.UpdatedAt = date
					if currDir {
						index.UpdatedAt = date
					}
				}

			}

			// next cell
			cell = cell.NextSibling
		}

		if include {
			result.IndexedAt = time.Now().Truncate(time.Second).UTC()
			index.Links = append(index.Links, result)
		}
	}

	return nil
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

	for class := range strings.FieldsSeq(attr.Val) {
		if class == needle {
			return true
		}
	}

	return false
}
