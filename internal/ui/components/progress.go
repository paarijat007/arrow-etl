package components

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

// ProgressBar represents a progress bar component
type ProgressBar struct {
	*tview.TextView
	maxValue int
	current  int
}

// NewProgressBar creates a new ProgressBar
func NewProgressBar(maxValue int) *ProgressBar {
	pb := &ProgressBar{
		TextView: tview.NewTextView(),
		maxValue: maxValue,
		current:  0,
	}
	pb.SetDynamicColors(true)
	pb.SetBorder(true)
	pb.update()
	return pb
}

// SetProgress sets the current progress
func (pb *ProgressBar) SetProgress(value int) {
	if value > pb.maxValue {
		value = pb.maxValue
	}
	pb.current = value
	pb.update()
}

// GetProgress returns the current progress
func (pb *ProgressBar) GetProgress() int {
	return pb.current
}

// update updates the progress bar display
func (pb *ProgressBar) update() {
	width := 50
	filled := int(float64(pb.current) / float64(pb.maxValue) * float64(width))
	bar := fmt.Sprintf("[green]%s[white]%s",
		strings.Repeat("█", filled),       // Full block character
		strings.Repeat("░", width-filled), // Light shade character
	)
	percentage := int(float64(pb.current) / float64(pb.maxValue) * 100)
	pb.SetText(fmt.Sprintf("%s %3d%%", bar, percentage))
}
