package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

// remove password
func remove(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return fmt.Errorf("id is empty")
	}
	idStr := ctx.Args().First()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	tx := db.Delete(&Password{Id: id})
	if tx.Error != nil {
		return tx.Error
	}
	fmt.Println("remove password success")
	return nil
}
