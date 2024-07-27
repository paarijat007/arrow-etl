package ui

import (
	"etl-ui/internal/ui/pages"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TerminalUI struct {
	app   *tview.Application
	pages *tview.Pages
}

func NewTerminalUI() *TerminalUI {
	return &TerminalUI{
		app:   tview.NewApplication(),
		pages: tview.NewPages(),
	}
}

func (t *TerminalUI) Run() error {
	t.setupPages()
	t.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			t.pages.SwitchToPage("main")
		}
		return event
	})
	return t.app.SetRoot(t.pages, true).EnableMouse(true).Run()
}

func (t *TerminalUI) setupPages() {
	mainMenu := pages.NewMainMenu(t.pages)
	t.pages.AddPage("main", mainMenu, true, true)
	t.pages.AddPage("create_job", pages.NewCreateJob(t.pages), true, false)
	t.pages.AddPage("manage_jobs", pages.NewManageJobs(t.pages), true, false)
	t.pages.AddPage("run_job", pages.NewRunJob(t.pages), true, false)
	t.pages.AddPage("settings", pages.NewSettings(t.pages), true, false)

	// Handle exit action
	mainMenu.(*tview.Flex).GetItem(1).(*tview.Flex).GetItem(2).(*tview.List).SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		if index == 4 { // Exit option
			t.app.Stop()
		}
	})
}
