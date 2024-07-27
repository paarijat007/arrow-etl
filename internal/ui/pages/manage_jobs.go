package pages

import (
	"github.com/rivo/tview"
)

func NewManageJobs(pages *tview.Pages) tview.Primitive {
	list := tview.NewList().
		AddItem("Job 1", "MySQL to CSV", 'a', nil).
		AddItem("Job 2", "PostgreSQL to ClickHouse", 'b', nil).
		AddItem("Back to Main Menu", "Return to the main menu", 'q', func() {
			pages.SwitchToPage("main")
		})

	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("Manage Existing Jobs"), 3, 1, false).
			AddItem(list, 0, 1, true).
			AddItem(nil, 0, 1, false), 40, 1, true).
		AddItem(nil, 0, 1, false)
}
