package pages

import (
	"github.com/rivo/tview"
)

func NewMainMenu(pages *tview.Pages) tview.Primitive {
	list := tview.NewList().
		AddItem("Create New Job", "Set up a new ETL job", 'c', func() {
			pages.SwitchToPage("create_job")
		}).
		AddItem("Manage Existing Jobs", "View and edit saved jobs", 'm', func() {
			pages.SwitchToPage("manage_jobs")
		}).
		AddItem("Run Job", "Execute a configured ETL job", 'r', func() {
			pages.SwitchToPage("run_job")
		}).
		AddItem("Settings", "Modify application settings", 's', func() {
			pages.SwitchToPage("settings")
		}).
		AddItem("Exit", "Exit the application", 'q', func() {
			// We'll handle this in the main application loop
		})

	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("ETL Platform"), 3, 1, false).
			AddItem(list, 0, 1, true).
			AddItem(nil, 0, 1, false), 40, 1, true).
		AddItem(nil, 0, 1, false)
}
