package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var table *tview.Table

func tableEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlQ:
		tui.Stop()
	case tcell.KeyEnter: // 编辑

		current := CurrentSelectedData()
		SetFormContent(current)
		form.SetTitle("Update").SetTitleAlign(tview.AlignLeft)
		tui.SetFocus(form)

	case tcell.KeyDelete: // 删除
		DeleteSelected()

	case tcell.KeyCtrlN: // 新增
		tui.SetFocus(form)
		// form.SetTitle("New").SetTitleAlign(tview.AlignLeft)
		// f := app.GetFocus()

	case tcell.Key('d'): // 解密

	case tcell.Key('e'): // 加密

	case tcell.Key('i'): // 导入

	case tcell.Key('p'): // 导出

	case tcell.Key('?'): // show healper

	case tcell.KeyCtrlF: // search
		tui.SetFocus(search)

	default:
	}
	return event
}

func InitTable() {
	table = tview.NewTable()
	table.SetInputCapture(tableEvent)
	table.SetTitle(" GOKEEPER ").SetTitleAlign(tview.AlignLeft)
	table.SetSelectable(true, false).SetBorder(true).SetTitle("Frieren Keeper")
	table.Focus(func(p tview.Primitive) {
		hit.SetText("table get focus")
	})
}

// SetTabelContent 使用 []data.Password 填充 table
func SetTabelContent(data []Password) {
	newCell := func(text string) *tview.TableCell { return tview.NewTableCell(text) }
	table.Clear()
	for i, d := range data {
		table.
			SetCell(i, 0, newCell(fmt.Sprintf("%d", d.Id))).
			SetCell(i, 1, newCell(d.Name)).
			SetCell(i, 2, newCell(d.Alias)).
			SetCell(i, 3, newCell(d.LoginId)).
			SetCell(i, 4, newCell(d.Password)).
			SetCell(i, 5, newCell(d.Bind)).
			SetCell(i, 6, newCell(d.Site)).
			SetCell(i, 7, newCell(d.Comment))
	}
}

// CurrentSelectedData 获取当前焦点数据
func CurrentSelectedData() *Password {
	if data == nil {
		return nil
	}
	r, _ := table.GetSelection()
	password := data[r]
	return &password
}

// DeleteSelected 删除数据
func DeleteSelected() {
	p := CurrentSelectedData()
	if p == nil {
		hit.SetText("no data delete")
		return
	}

	// TODO 添加确认窗口
	tx := db.Delete(p)
	if tx.Error != nil {
		hit.SetText(fmt.Sprint("delete error: ", tx.Error))
		tui.SetFocus(table)
		return
	}
	tx = db.Find(&data)
	if tx.Error != nil {
		hit.SetText(fmt.Sprint("flush error: ", tx.Error))
		tui.SetFocus(table)
		return
	}
	SetTabelContent(data)
	hit.SetText("delete success")
}
