package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// encrypt all passwors, loginId, bind
func encryptAll(ctx *cli.Context) error {
	keyFile := ctx.Path("key")
	passwords := []Password{}
	tx := db.Where("encrypt = 0").Find(&passwords)
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.Begin()
	var err error
	for _, p := range passwords {
		p.Password, err = EncrypeString(p.Password, keyFile)
		if err != nil {
			return err
		}

		p.LoginId, err = EncrypeString(p.LoginId, keyFile)
		if err != nil {
			return err
		}

		p.Bind, err = EncrypeString(p.Bind, keyFile)
		if err != nil {
			return err
		}
		p.Encrypt = 1
		result := tx.Save(p)
		if result.Error != nil {
			return result.Error
		}
	}
	tx = tx.Commit()
	if tx.Error != nil {
		return tx.Error
	}
	fmt.Println("encrypt success")

	return nil
}
