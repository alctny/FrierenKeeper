package main

import (
	"fmt"
	"strconv"

	"github.com/alctny/frieren-keeper/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var form *tview.Form

func fromEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlQ:
		SetFormContent(&model.Password{})
		app.SetFocus(table)
	}
	return event
}

func ConfirmSelected() {
	idStr := form.GetFormItemByLabel(LEABLE_ID).(*tview.TextView).GetText(false)
	name := form.GetFormItemByLabel(LEABLE_NAME).(*tview.InputField).GetText()
	loginId := form.GetFormItemByLabel(LEABLE_LOGIN_ID).(*tview.InputField).GetText()
	alias := form.GetFormItemByLabel(LEABLE_ALIAS).(*tview.InputField).GetText()
	password := form.GetFormItemByLabel(LEABLE_PASSWORD).(*tview.InputField).GetText()
	bind := form.GetFormItemByLabel(LEABLE_BIND).(*tview.InputField).GetText()
	site := form.GetFormItemByLabel(LEABLE_SITE).(*tview.InputField).GetText()
	comment := form.GetFormItemByLabel(LEABLE_COMMENT).(*tview.TextArea).GetText()
	encrypt, _ := form.GetFormItemByLabel(LEABLE_ENCRYPT).(*tview.DropDown).GetCurrentOption()
	id, _ := strconv.ParseInt(idStr, 10, 64)
	p := model.Password{
		Id:       int(id),
		Name:     name,
		LoginId:  loginId,
		Password: password,
		Bind:     bind,
		Alias:    alias,
		Site:     site,
		Comment:  comment,
		Encrypt:  encrypt,
	}

	tx := db.Save(&p)
	if tx.Error != nil {
		hit.SetText(fmt.Sprint("update error: ", tx.Error))
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
	hit.SetText("update success")
	SetFormContent(&model.Password{})
	// 重新设置焦点，否则再次获得焦点的时候会让按扭获得焦点，交互体验不好
	form.SetFocus(1)
	app.SetFocus(table)
}

func SetFormContent(current *model.Password) {
	form.SetFocus(1)
	if current == nil {
		return
	}
	form.GetFormItemByLabel(LEABLE_ID).(*tview.TextView).SetText(fmt.Sprint(current.Id))
	form.GetFormItemByLabel(LEABLE_NAME).(*tview.InputField).SetText(current.Name)
	form.GetFormItemByLabel(LEABLE_LOGIN_ID).(*tview.InputField).SetText(current.LoginId)
	form.GetFormItemByLabel(LEABLE_ALIAS).(*tview.InputField).SetText(current.Alias)
	form.GetFormItemByLabel(LEABLE_PASSWORD).(*tview.InputField).SetText(current.Password)
	form.GetFormItemByLabel(LEABLE_BIND).(*tview.InputField).SetText(current.Bind)
	form.GetFormItemByLabel(LEABLE_SITE).(*tview.InputField).SetText(current.Site)
	form.GetFormItemByLabel(LEABLE_COMMENT).(*tview.TextArea).SetText(current.Comment, true)
	form.GetFormItemByLabel(LEABLE_ENCRYPT).(*tview.DropDown).SetCurrentOption(current.Encrypt)
}

// 初始化表单
func InitForm() {
	form = tview.NewForm()
	form.
		AddTextView(LEABLE_ID, "", 0, 1, true, false).
		AddInputField(LEABLE_NAME, "", 0, nil, nil).
		AddInputField(LEABLE_LOGIN_ID, "", 0, nil, nil).
		AddInputField(LEABLE_ALIAS, "", 0, nil, nil).
		AddInputField(LEABLE_PASSWORD, "", 0, nil, nil).
		AddInputField(LEABLE_BIND, "", 0, nil, nil).
		AddInputField(LEABLE_SITE, "", 0, nil, nil).
		AddTextArea(LEABLE_COMMENT, "", 0, 10, 0, nil).
		AddDropDown(LEABLE_ENCRYPT, []string{"Image", "None"}, 1, nil).
		AddButton(BUTTON_CONFIRM, ConfirmSelected).
		AddButton(BUTTON_CANCLE, func() { SetFormContent(&model.Password{}); app.SetFocus(table) })

	form.SetInputCapture(fromEvent).SetBorder(true)
}
