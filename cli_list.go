package main

import (
	"github.com/urfave/cli/v2"
)

// list all passwords
func list(ctx *cli.Context) error {
	passwords := []Password{}
	keyFile := ctx.String("key")

	tx := db.Find(&passwords)
	if tx.Error != nil {
		return tx.Error
	}
	if keyFile != "" {
		var err error
		passwords, err = DecryptPasswords(passwords, keyFile)
		if err != nil {
			return err
		}
	}

	ShowPasswords(passwords)
	return nil
}
