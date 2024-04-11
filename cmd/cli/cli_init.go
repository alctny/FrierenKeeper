package main

import (
	"github.com/alctny/frieren-keeper/model"
	"github.com/urfave/cli/v2"
)

// create sqlite file and create table
func initDB(ctx *cli.Context) error {
	db.AutoMigrate(&model.Password{})
	return db.Error
}
