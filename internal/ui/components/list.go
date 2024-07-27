package components

import (
	"github.com/rivo/tview"
)

// CustomList extends tview.List with additional functionality
type CustomList struct {
	*tview.List
	items []string
}

// NewCustomList creates a new CustomList
func NewCustomList() *CustomList {
	return &CustomList{
		List:  tview.NewList(),
		items: []string{},
	}
}

// AddItems adds multiple items to the list
func (l *CustomList) AddItems(items ...string) *CustomList {
	for _, item := range items {
		l.items = append(l.items, item)
		l.AddItem(item, "", 0, nil)
	}
	return l
}

// SetSelectedFunc sets a function to be called when an item is selected
func (l *CustomList) SetSelectedFunc(handler func(index int, item string)) *CustomList {
	l.List.SetSelectedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		handler(index, l.items[index])
	})
	return l
}

// SetupList sets up the list with a title and border
func (l *CustomList) SetupList(title string) *CustomList {
	l.SetBorder(true).
		SetTitle(title).
		SetTitleAlign(tview.AlignLeft)
	return l
}
