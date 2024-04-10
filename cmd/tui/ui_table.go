package main

import (
	"fmt"

	"github.com/alctny/frieren-keeper/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var table *tview.Table

func tableEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlQ:
		app.Stop()
	case tcell.KeyEnter: // 编辑

		current := CurrentSelectedData()
		SetFormContent(current)
		form.SetTitle("Update").SetTitleAlign(tview.AlignLeft)
		app.SetFocus(form)

	case tcell.KeyDelete: // 删除
		DeleteSelected()

	case tcell.KeyCtrlN: // 新增
		app.SetFocus(form)
		// form.SetTitle("New").SetTitleAlign(tview.AlignLeft)
		// f := app.GetFocus()

	case tcell.Key('d'): // 解密

	case tcell.Key('e'): // 加密

	case tcell.Key('i'): // 导入

	case tcell.Key('p'): // 导出

	case tcell.Key('?'): // show healper

	case tcell.KeyCtrlF: // search
		app.SetFocus(search)

	default:
	}
	return event
}

func InitTable() {
	table = tview.NewTable()
	table.SetInputCapture(tableEvent)
	table.SetTitle(" GOKEEPER ").SetTitleAlign(tview.AlignLeft)
	table.SetSelectable(true, false).SetBorder(true).SetTitle(" GOKEEPER ")
}

// SetTabelContent 使用 []data.Password 填充 table
func SetTabelContent(data []model.Password) {
	newCell := func(text string) *tview.TableCell { return tview.NewTableCell(text) }
	newTitle := func(text string) *tview.TableCell {
		return tview.
			NewTableCell(text).
			SetBackgroundColor(tcell.ColorWhite).
			SetTextColor(tcell.ColorBlack)
	}

	table.Clear()

	table.
		SetCell(0, 0, newTitle(LEABLE_ID)).
		SetCell(0, 1, newTitle(LEABLE_NAME)).
		SetCell(0, 2, newTitle(LEABLE_ALIAS)).
		SetCell(0, 3, newTitle(LEABLE_LOGIN_ID)).
		SetCell(0, 4, newTitle(LEABLE_PASSWORD)).
		SetCell(0, 5, newTitle(LEABLE_BIND)).
		SetCell(0, 6, newTitle(LEABLE_SITE)).
		SetCell(0, 7, newTitle(LEABLE_COMMENT))

	for i, d := range data {
		table.
			SetCell(i+1, 0, newCell(fmt.Sprintf("%d", d.Id))).
			SetCell(i+1, 1, newCell(d.Name)).
			SetCell(i+1, 2, newCell(d.Alias)).
			SetCell(i+1, 3, newCell(d.LoginId)).
			SetCell(i+1, 4, newCell(d.Password)).
			SetCell(i+1, 5, newCell(d.Bind)).
			SetCell(i+1, 6, newCell(d.Site)).
			SetCell(i+1, 7, newCell(d.Comment))
	}
}

// CurrentSelectedData 获取当前焦点数据
func CurrentSelectedData() *model.Password {
	r, _ := table.GetSelection()
	if r <= 1 {
		return nil
	}
	hit.SetText(fmt.Sprintf("selected row: %d", r))
	password := data[r-1]
	return &password
}

// DeleteSelected 删除数据
func DeleteSelected() {
	p := CurrentSelectedData()
	if p == nil {
		hit.SetText("no data delete")
		return
	}

	tx := db.Delete(p)
	if tx.Error != nil {
		hit.SetText(fmt.Sprint("delete error: ", tx.Error))
		app.SetFocus(table)
		return
	}
	tx = db.Find(&data)
	if tx.Error != nil {
		hit.SetText(fmt.Sprint("flush error: ", tx.Error))
		app.SetFocus(table)
		return
	}
	SetTabelContent(data)
	hit.SetText("delete success")
}
