package pages

import (
	"etl-ui/internal/ui/components"
	"etl-ui/pkg/config"
	"strconv"

	"github.com/rivo/tview"
)

func NewSettings(pages *tview.Pages) tview.Primitive {
	cfg := config.GetConfig()

	form := components.NewCustomForm()

	form.AddDropDown("Log Level", []string{"INFO", "DEBUG", "ERROR"}, getIndex(cfg.LogLevel, []string{"INFO", "DEBUG", "ERROR"}), nil)
	form.AddInputField("Max Concurrent Jobs", strconv.Itoa(cfg.MaxConcurrentJobs), 20, nil, nil)
	form.AddCheckbox("Enable Parallel Processing", cfg.EnableParallel, nil)
	form.AddInputField("Chunk Size", strconv.Itoa(cfg.ChunkSize), 20, nil, nil)
	form.AddDropDown("Default Source Type", []string{"MySQL", "PostgreSQL", "CSV", "ClickHouse"}, getIndex(cfg.DefaultSourceType, []string{"MySQL", "PostgreSQL", "CSV", "ClickHouse"}), nil)
	form.AddDropDown("Default Destination Type", []string{"MySQL", "PostgreSQL", "CSV", "ClickHouse"}, getIndex(cfg.DefaultDestinationType, []string{"MySQL", "PostgreSQL", "CSV", "ClickHouse"}), nil)

	form.AddButton("Save", func() {
		// Save settings
		_, cfg.LogLevel = form.GetFormItem(0).(*tview.DropDown).GetCurrentOption()
		maxJobs, _ := strconv.Atoi(form.GetFormItem(1).(*tview.InputField).GetText())
		cfg.MaxConcurrentJobs = maxJobs
		cfg.EnableParallel = form.GetFormItem(2).(*tview.Checkbox).IsChecked()
		chunkSize, _ := strconv.Atoi(form.GetFormItem(3).(*tview.InputField).GetText())
		cfg.ChunkSize = chunkSize
		_, cfg.DefaultSourceType = form.GetFormItem(4).(*tview.DropDown).GetCurrentOption()
		_, cfg.DefaultDestinationType = form.GetFormItem(5).(*tview.DropDown).GetCurrentOption()

		config.SaveConfig("config.json")
		pages.SwitchToPage("main")
	})

	form.AddButton("Cancel", func() {
		pages.SwitchToPage("main")
	})

	form.SetupForm("Settings")

	return form
}

func getIndex(value string, options []string) int {
	for i, v := range options {
		if v == value {
			return i
		}
	}
	return 0
}
