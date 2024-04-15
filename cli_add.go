package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// add new password
func add(ctx *cli.Context) error {
	keyFile := ctx.String("key")
	name := ctx.String("name")
	loginId := ctx.String("loginid")
	password := ctx.String("password")
	bind := ctx.String("bind")
	alias := ctx.String("alias")
	site := ctx.String("site")
	comment := ctx.String("comment")
	isEncrype := 0

	var err error
	if keyFile != "" {
		isEncrype = 1

		password, err = EncrypeString(password, keyFile)
		if err != nil {
			return err
		}

		loginId, err = EncrypeString(loginId, keyFile)
		if err != nil {
			return err
		}

		if bind != "" {
			bind, err = EncrypeString(bind, keyFile)
			if err != nil {
				return err
			}
		}
	}

	passwd := Password{
		Name:     name,
		LoginId:  loginId,
		Password: password,
		Bind:     bind,
		Alias:    alias,
		Site:     site,
		Comment:  comment,
		Encrypt:  isEncrype,
	}

	tx := db.Create(&passwd)
	if tx.Error != nil {
		return tx.Error
	}
	fmt.Println("add password success")
	return nil
}
