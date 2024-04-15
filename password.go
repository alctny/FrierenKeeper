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

func DecryptPasswords(ps []Password, keyFile string) ([]Password, error) {
	result := []Password{}
	var err error

	for _, p := range ps {
		if p.Encrypt == 0 {
			result = append(result, p)
			continue
		}

		p.Encrypt = 0
		p.Password, err = DecryptString(p.Password, keyFile)
		if err != nil {
			return nil, err
		}

		p.LoginId, err = DecryptString(p.LoginId, keyFile)
		if err != nil {
			return nil, err
		}

		if p.Bind != "" {
			p.Bind, err = DecryptString(p.Bind, keyFile)
			if err != nil {
				return nil, err
			}
		}

		result = append(result, p)
	}

	return result, nil
}

func EncryptPasswords(ps []Password, keyFile string) ([]Password, error) {
	result := []Password{}
	var err error

	for _, p := range ps {
		if p.Encrypt == 0 {
			result = append(result, p)
			continue
		}

		p.Encrypt = 0
		p.Password, err = EncrypeString(p.Password, keyFile)
		if err != nil {
			return nil, err
		}

		p.LoginId, err = EncrypeString(p.LoginId, keyFile)
		if err != nil {
			return nil, err
		}

		if p.Bind != "" {
			p.Bind, err = EncrypeString(p.Bind, keyFile)
			if err != nil {
				return nil, err
			}
		}

		result = append(result, p)
	}

	return result, nil
}
