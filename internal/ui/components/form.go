package components

import (
	"github.com/rivo/tview"
)

// CustomForm extends tview.Form with additional functionality
type CustomForm struct {
	*tview.Form
	cancelFunc func()
}

// NewCustomForm creates a new CustomForm
func NewCustomForm() *CustomForm {
	return &CustomForm{
		Form: tview.NewForm(),
	}
}

// AddCancelButton adds a cancel button with a custom function
func (f *CustomForm) AddCancelButton(label string, cancelFunc func()) *CustomForm {
	f.cancelFunc = cancelFunc
	f.AddButton(label, func() {
		if f.cancelFunc != nil {
			f.cancelFunc()
		}
	})
	return f
}

// AddSaveButton adds a save button with a custom function
func (f *CustomForm) AddSaveButton(label string, saveFunc func()) *CustomForm {
	f.AddButton(label, saveFunc)
	return f
}

// SetupForm sets up the form with a title and border
func (f *CustomForm) SetupForm(title string) *CustomForm {
	f.SetBorder(true).
		SetTitle(title).
		SetTitleAlign(tview.AlignLeft)
	return f
}
