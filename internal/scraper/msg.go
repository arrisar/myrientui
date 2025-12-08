package scraper

type StartedMsg struct{}

type ResultsMsg struct {
	Index Index
	Err   error
}

type IndexMsg struct {
	Found   int
	Indexed int
}
