package app

import (
	"fmt"

	"github.com/arrisar/myrientui/internal/browser"
	"github.com/arrisar/myrientui/internal/config"
	"github.com/arrisar/myrientui/internal/scraper"
	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	Browser tea.Model
	Config  config.Config
	Scraper scraper.Scraper
}

func New() App {
	a := App{}
	a.Config = config.New()
	a.Browser = browser.New()
	a.Scraper = scraper.New(a.Config.Data.Storage.CacheDir)
	return a
}

/**
 * MODEL
 */

func (a App) Init() tea.Cmd {
	// go a.Scraper.IndexAll()
	return a.Browser.Init()
}

func (a App) View() string {
	return a.Browser.View()
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var c, cmd tea.Cmd
	var forward bool

	a, cmd, forward = a.handleUpdateMsg(msg)
	c = tea.Batch(c, cmd)

	if forward {
		a.Browser, cmd = a.Browser.Update(msg)
		c = tea.Batch(c, cmd)
	}

	return a, c
}

/**
 * HANDLERS
 */

func (a App) handleUpdateMsg(msg tea.Msg) (app App, cmd tea.Cmd, forward bool) {
	app = a

	switch msg := msg.(type) {
	case browser.OptionSelectedMsg:
		app, cmd = app.handleOptionSelectedMsg(msg)
		forward = true
	case scraper.StartedMsg:
		app, cmd = app.handleStartedMsg(msg)
		forward = false
	case scraper.ResultsMsg:
		app, cmd = app.handleResultsMsg(msg)
		forward = false
	default:
		forward = true
	}

	return
}

func (a App) handleOptionSelectedMsg(msg browser.OptionSelectedMsg) (App, tea.Cmd) {
	uri := "/"
	for _, v := range msg.Path {
		uri += v.String()
	}

	return a, a.Scraper.MsgScrape(uri)
}

func (a App) handleStartedMsg(scraper.StartedMsg) (App, tea.Cmd) {
	return a, func() tea.Msg {
		return browser.OptionsLoadingMsg{}
	}
}

func (a App) handleResultsMsg(msg scraper.ResultsMsg) (App, tea.Cmd) {
	if msg.Err != nil {
		fmt.Println(msg.Err)
		return a, nil
	}

	a.Scraper.IndexPreload(msg.Index)
	return a, func() tea.Msg {
		o := browser.OptionsChangedMsg{}

		for _, r := range msg.Index.Links {
			o.Options = append(o.Options, browser.Option(r.Label))
		}

		return o
	}
}
