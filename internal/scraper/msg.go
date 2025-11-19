package scraper

type StartedMsg struct{}

type ResultsMsg struct {
	Results []Result
	Err     error
}
