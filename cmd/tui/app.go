package main

import (
	"fmt"
	"os"

	"github.com/alctny/frieren-keeper/dao"
	"github.com/alctny/frieren-keeper/model"
	"github.com/rivo/tview"
	"gorm.io/gorm"
)

var db *gorm.DB
var app *tview.Application
var page *tview.Pages
var data []model.Password

func Init() {
	InitSearch()
	InitForm()
	InitHit()
	InitTable()
	app = tview.NewApplication()
	page = tview.NewPages()
	db = dao.NewGormDB()

	layout := tview.NewFlex().SetDirection(tview.FlexColumn)
	rightFlex := tview.NewFlex().
		AddItem(search, 3, 0, false).
		SetDirection(tview.FlexRow).
		AddItem(form, 0, 7, false).
		AddItem(hit, 0, 3, false)
	layout.
		AddItem(table, 0, 7, false).
		AddItem(rightFlex, 0, 3, false)
	page.AddPage("main", layout, true, true)
}

func run() {
	Init()
	tx := db.Find(&data)
	if tx.Error != nil {
		fmt.Fprint(os.Stderr, "Error: ", tx.Error)
		os.Exit(1)
	}

	SetTabelContent(data)
	app.SetRoot(page, true).SetFocus(table)
	err := app.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}
}

func main() {
	run()
}
