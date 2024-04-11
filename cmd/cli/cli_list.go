package main

import (
	"github.com/alctny/frieren-keeper/model"
	"github.com/urfave/cli/v2"
)

// list all passwords
func list(ctx *cli.Context) error {
	passwords := []model.Password{}
	keyFile := ctx.String("key")

	tx := db.Find(&passwords)
	if tx.Error != nil {
		return tx.Error
	}
	if keyFile != "" {
		var err error
		passwords, err = model.DecryptPasswords(passwords, keyFile)
		if err != nil {
			return err
		}
	}

	model.ShowPasswords(passwords)
	return nil
}
