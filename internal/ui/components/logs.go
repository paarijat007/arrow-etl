package components

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

// LogViewer represents a log viewer component
type LogViewer struct {
	*tview.TextView
}

// NewLogViewer creates a new LogViewer
func NewLogViewer() *LogViewer {
	lv := &LogViewer{
		TextView: tview.NewTextView(),
	}
	lv.SetDynamicColors(true).
		SetScrollable(true).
		SetWordWrap(true).
		SetChangedFunc(func() {
			lv.ScrollToEnd()
		})
	lv.SetBorder(true).SetTitle("Logs")
	return lv
}

// AddLog adds a new log entry
func (lv *LogViewer) AddLog(level, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	var coloredLevel string
	switch level {
	case "INFO":
		coloredLevel = "[blue]INFO[white]"
	case "WARNING":
		coloredLevel = "[yellow]WARNING[white]"
	case "ERROR":
		coloredLevel = "[red]ERROR[white]"
	default:
		coloredLevel = level
	}
	logEntry := fmt.Sprintf("%s %s: %s\n", timestamp, coloredLevel, message)
	fmt.Fprint(lv, logEntry)
}

// Clear clears all log entries
func (lv *LogViewer) Clear() {
	lv.SetText("")
}
