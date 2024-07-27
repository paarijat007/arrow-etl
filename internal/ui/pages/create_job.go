package pages

import (
	"github.com/rivo/tview"
)

func NewCreateJob(pages *tview.Pages) tview.Primitive {
	form := tview.NewForm().
		AddDropDown("Source Type", []string{"MySQL", "PostgreSQL", "CSV", "ClickHouse"}, 0, nil).
		AddInputField("Host", "localhost", 30, nil, nil).
		AddInputField("Port", "3306", 10, nil, nil).
		AddInputField("Database", "", 30, nil, nil).
		AddInputField("Username", "", 30, nil, nil).
		AddPasswordField("Password", "", 30, '*', nil).
		AddDropDown("Transformation Type", []string{"Basic", "Custom SQL", "Apache Arrow"}, 0, nil).
		AddTextArea("SQL Query", "", 40, 10, 0, nil).
		AddDropDown("Destination Type", []string{"MySQL", "PostgreSQL", "CSV", "ClickHouse"}, 0, nil).
		AddButton("Save Job", func() {
			// TODO: Implement job saving logic
			pages.SwitchToPage("main")
		}).
		AddButton("Back to Main Menu", func() {
			pages.SwitchToPage("main")
		})

	form.SetBorder(true).SetTitle("Create New Job").SetTitleAlign(tview.AlignLeft)

	return form
}
