package main

import (
	"github.com/urfave/cli/v2"
)

// create sqlite file and create table
func initDB(ctx *cli.Context) error {
	db.AutoMigrate(&Password{})
	return db.Error
}
