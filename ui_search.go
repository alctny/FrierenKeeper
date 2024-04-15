package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var search *tview.InputField

func searchEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEnter:
		tui.SetFocus(table)
	}
	return event
}

func searchInType(text string) {
	db.Where("name LIKE ? OR alias LIKE ? OR login_id LIKE ?", "%"+text+"%", "%"+text+"%", "%"+text+"%").Find(&data)
	SetTabelContent(data)
}

func InitSearch() {
	search = tview.NewInputField()
	search.
		SetFieldBackgroundColor(tcell.ColorNone).
		SetChangedFunc(searchInType).
		SetInputCapture(searchEvent)

	search.
		SetBorder(true).
		SetTitle(INPUT_SEARCH).
		SetTitleAlign(tview.AlignLeft)
}
