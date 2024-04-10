package main

import "github.com/rivo/tview"

var hit *tview.TextView

func InitHit() {
	hit = tview.NewTextView()
	hit.SetText("这个地方将会添加提示信息")
	hit.
		SetBorder(true).
		SetTitle("Hit").
		SetTitleAlign(tview.AlignLeft)
}
