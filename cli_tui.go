package main

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
	"github.com/urfave/cli/v2"
)

// var db *gorm.DB
var tui *tview.Application
var page *tview.Pages
var data []Password

func Init() {
	InitHit()
	InitSearch()
	InitForm()
	InitTable()

	tui = tview.NewApplication()
	page = tview.NewPages()

	layout := tview.NewFlex().SetDirection(tview.FlexColumn)
	rightFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(search, 3, 0, false).
		AddItem(form, 0, 7, false).
		AddItem(hit, 0, 3, false)
	leftFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(table, 0, 1, true)

	layout.
		AddItem(leftFlex, 0, 7, false).
		AddItem(rightFlex, 0, 3, false)
	page.AddPage("main", layout, true, true)

	db = NewGormDB()
}

func tuiStart(c *cli.Context) error {
	Init()
	tx := db.Find(&data)
	if tx.Error != nil {
		fmt.Fprint(os.Stderr, "Error: ", tx.Error)
		os.Exit(1)
	}

	SetTabelContent(data)
	tui.SetRoot(page, true).SetFocus(table)
	err := tui.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}

	return nil
}
