package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// get password
func get(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return fmt.Errorf("query is empty")
	}
	query := fmt.Sprintf("%%%s%%", ctx.Args().First())
	passwords := []Password{}
	tx := db.Where("name LIKE ? OR site LIKE ? OR comment LIKE ? OR alias LIKE ?", query, query, query, query).Find(&passwords)
	if tx.Error != nil {
		return tx.Error
	}

	keyFile := ctx.Path("key")
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
