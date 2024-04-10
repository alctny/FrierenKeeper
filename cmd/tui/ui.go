package main

import "github.com/gdamore/tcell/v2"

//

// form lable and table title
const (
	LEABLE_ID       = "ID"
	LEABLE_NAME     = "Name"
	LEABLE_LOGIN_ID = "Login ID"
	LEABLE_ALIAS    = "Alias"
	LEABLE_PASSWORD = "Password"
	LEABLE_BIND     = "Bind"
	LEABLE_SITE     = "Site"
	LEABLE_COMMENT  = "Comment"
	LEABLE_ENCRYPT  = "Encrypt"
	BUTTON_CONFIRM  = "Comfirm"
	BUTTON_CANCLE   = "Cancle"
	INPUT_SEARCH    = " "
)

// 主题配置
const (
	COLOR_SELECT = tcell.ColorRed
	COLOR_NORMAL = tcell.ColorWhite
)

// 主要组件
const (
	TABLE  = "table"
	FORM   = "form"
	SEARCH = "search"
	HIT    = "hit"
)
