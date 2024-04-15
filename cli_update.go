package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

// update password infomaion
// this method will not attempt to automatically update the encryption status
func update(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return fmt.Errorf("id is empty")
	}
	idStr := ctx.Args().First()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	name := ctx.String("name")
	loginId := ctx.String("loginid")
	password := ctx.String("password")
	bind := ctx.String("bind")
	alias := ctx.String("alias")
	site := ctx.String("site")
	comment := ctx.String("comment")
	isEncrype := 0

	keyFile := ctx.Path("key")
	if keyFile != "" {
		isEncrype = 1
		if password != "" {
			password, err = EncrypeString(password, keyFile)
			if err != nil {
				return err
			}
		}

		if loginId != "" {
			loginId, err = EncrypeString(loginId, keyFile)
			if err != nil {
				return err
			}
		}

		if bind != "" {
			bind, err = EncrypeString(bind, keyFile)
			if err != nil {
				return err
			}
		}
	}

	p := Password{
		Id:       id,
		Name:     name,
		LoginId:  loginId,
		Password: password,
		Bind:     bind,
		Alias:    alias,
		Site:     site,
		Comment:  comment,
		Encrypt:  isEncrype,
	}
	// 此处必须添加 encrypt = isEncrype 条件，保证修改前后加密状态未发生改变
	// 避免出现所有信息都没有加密但 encrypt = 1 或所有信息都有加密但 encrypt = 0 的情况
	tx := db.Model(&Password{}).Where("id = ? AND encrypt = ?", id, isEncrype).Updates(p)
	if tx.Error != nil {
		return tx.Error
	}

	fmt.Println("update password success")
	return nil
}
