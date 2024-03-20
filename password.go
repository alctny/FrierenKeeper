package main

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

type Password struct {
	Id       int
	Name     string
	LoginId  string
	Password string
	Bind     string
	Alias    string
	Site     string
	Comment  string
	Encrypt  int
}

// 展示记录的密码，自动调整列宽，且自适应中文字符
func ShowPasswords(ps []Password) {
	aliasWidth := runewidth.StringWidth("Alias")
	nameWidth := runewidth.StringWidth("Name")
	loginIdWidth := runewidth.StringWidth("LoginID")
	passwordWidth := runewidth.StringWidth("Password")
	bindWidth := runewidth.StringWidth("Bind")
	commentWidth := runewidth.StringWidth("Comment")

	for _, p := range ps {
		if runewidth.StringWidth(p.Name) > nameWidth {
			nameWidth = runewidth.StringWidth(p.Name)
		}
		if runewidth.StringWidth(p.LoginId) > loginIdWidth {
			loginIdWidth = runewidth.StringWidth(p.LoginId)
		}
		if runewidth.StringWidth(p.Password) > passwordWidth {
			passwordWidth = runewidth.StringWidth(p.Password)
		}
		if runewidth.StringWidth(p.Bind) > bindWidth {
			bindWidth = runewidth.StringWidth(p.Bind)
		}
		if runewidth.StringWidth(p.Comment) > commentWidth {
			commentWidth = runewidth.StringWidth(p.Comment)
		}
		if runewidth.StringWidth(p.Alias) > aliasWidth {
			aliasWidth = runewidth.StringWidth(p.Alias)
		}
	}

	text := strings.Builder{}
	fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		runewidth.FillRight("ID", 2),
		runewidth.FillRight("Name", nameWidth),
		runewidth.FillRight("LoginID", loginIdWidth),
		runewidth.FillRight("Alias", aliasWidth),
		runewidth.FillRight("Password", passwordWidth),
		runewidth.FillRight("Bind", bindWidth),
		runewidth.FillRight("Comment", commentWidth),
	)
	for _, ps := range ps {
		text.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			fmt.Sprintf("%02d", ps.Id),
			runewidth.FillRight(ps.Name, nameWidth),
			runewidth.FillRight(ps.LoginId, loginIdWidth),
			runewidth.FillRight(ps.Alias, aliasWidth),
			runewidth.FillRight(ps.Password, passwordWidth),
			runewidth.FillRight(ps.Bind, bindWidth),
			runewidth.FillRight(ps.Comment, commentWidth),
		))
	}

	fmt.Print(text.String())
}

func DecryptPasswords(ps []Password, keyFile string) []Password {
	result := []Password{}
	keyByte := FileHash256(keyFile)
	for _, p := range ps {
		if p.Encrypt == 1 {
			p.Encrypt = 0
			p.Password = Decrypt(p.Password, keyByte)
			p.LoginId = Decrypt(p.LoginId, keyByte)
			if p.Bind != "" {
				p.Bind = Decrypt(p.Bind, keyByte)
			}
		}
		result = append(result, p)
	}

	return result
}
