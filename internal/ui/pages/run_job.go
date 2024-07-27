package pages

import (
	"etl-ui/internal/ui/components"

	"github.com/rivo/tview"
)

func NewRunJob(pages *tview.Pages) tview.Primitive {
	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	jobInfo := tview.NewTextView().
		SetText("Job: MyETLJob1\nSource: MySQL - etl_source\nTransform: Custom SQL Query\nDestination: CSV - /path/to/output.csv").
		SetBorder(true).
		SetTitle("Job Information")

	progressBar := components.NewProgressBar(100)
	progressBar.SetProgress(50)
	progressBar.SetTitle("Progress")

	logViewer := components.NewLogViewer()
	logViewer.AddLog("INFO", "Job started")
	logViewer.AddLog("INFO", "Connected to MySQL source")
	logViewer.AddLog("INFO", "Extracted 10000 rows")
	logViewer.AddLog("WARNING", "Transformation in progress...")

	buttons := tview.NewFlex().
		AddItem(tview.NewButton("Back").SetSelectedFunc(func() {
			pages.SwitchToPage("main")
		}), 0, 1, false).
		AddItem(tview.NewButton("Pause"), 0, 1, false).
		AddItem(tview.NewButton("Stop"), 0, 1, false)

	flex.AddItem(jobInfo, 0, 1, false).
		AddItem(progressBar, 3, 1, false).
		AddItem(logViewer, 0, 3, false).
		AddItem(buttons, 3, 1, true)

	return flex
}
