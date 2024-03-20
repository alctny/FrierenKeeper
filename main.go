package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var VERSION = "0.3.1"
var app = cli.NewApp()
var db *gorm.DB

func main() {
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
