package browser

type BrowserResizeMsg struct {
	BrowserHeight int
	BrowserWidth  int
	HeaderHeight  int
	HeaderWidth   int
	ListHeight    int
	ListWidth     int
	FooterHeight  int
	FooterWidth   int
}

type FilterStateMsg struct {
	Enabled bool
	Active  bool
	Applied bool
	Value   string
}

type OptionSelectedMsg struct {
	Option Option
	Path   []Option
}

type OptionsChangedMsg struct {
	Options []Option
}

type OptionsLoadingMsg struct{}
