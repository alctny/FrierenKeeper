package main

import (
	"os"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

var VERSION = "0.3.2"
var app *cli.App
var db *gorm.DB

func init() {
	app = cli.NewApp()
	app.Commands = []*cli.Command{
		{
			Name:   "init",
			Usage:  "init password manager",
			Action: initDB,
			Flags:  []cli.Flag{},
		},
		{
			Name:   "add",
			Usage:  "add new password",
			Action: add,
			Flags: []cli.Flag{
				// required
				&cli.StringFlag{
					Name:     "name",
					Required: true,
					Usage:    "name for password",
				},
				&cli.StringFlag{
					Name:     "loginid",
					Required: true,
					Usage:    "login id",
				},
				&cli.StringFlag{
					Name:     "password",
					Required: true,
					Usage:    "password",
				},
				// optional
				&cli.PathFlag{
					Name:  "key",
					Usage: "path to key file",
				},
				&cli.StringFlag{
					Name:  "bind",
					Usage: "bind email/tel/third party account",
				},
				&cli.StringFlag{
					Name:  "comment",
					Usage: "notes for password",
				},
				&cli.StringFlag{
					Name:  "alias",
					Usage: "alias for name",
				},
				&cli.StringFlag{
					Name:  "site",
					Usage: "witch site to use this password",
				},
			},
		},
		{
			Name:   "get",
			Usage:  "get password",
			Action: get,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:  "key",
					Usage: "path to key file",
				},
			},
		},
		{
			Name:   "list",
			Usage:  "list all password",
			Action: list,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:  "key",
					Usage: "path to key file",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "update password",
			Action: update,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "name for password",
				},
				&cli.StringFlag{
					Name:  "loginid",
					Usage: "login id",
				},
				&cli.StringFlag{
					Name:  "password",
					Usage: "password",
				},
				&cli.PathFlag{
					Name:  "key",
					Usage: "path to key file",
				},
				&cli.StringFlag{
					Name:  "bind",
					Usage: "bind email/tel/third party account",
				},
				&cli.StringFlag{
					Name:  "comment",
					Usage: "notes for password",
				},
				&cli.StringFlag{
					Name:  "alias",
					Usage: "alias for name",
				},
				&cli.StringFlag{
					Name:  "site",
					Usage: "witch site to use this password",
				},
			},
		},
		{
			Name:   "remove",
			Usage:  "remove password",
			Action: remove,
		},
		{
			Name:   "import",
			Usage:  "import passwords from csv, the csv file format: name,loginId,password,bind,alias,site,comment",
			Action: importFromCsv,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:  "key",
					Usage: "path to key file",
				},
			},
		},
		{
			Name:   "export",
			Usage:  "export to csv",
			Action: export2Csv,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:  "key",
					Usage: "path to key file",
				},
				&cli.PathFlag{
					Name:  "output",
					Usage: "output file",
				},
			},
		},
		{
			Name:   "decrypt",
			Usage:  "decrypt all",
			Action: decryptAll,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:     "key",
					Usage:    "path to key file",
					Required: true,
				},
			},
		},
		{
			Name:   "encrypt",
			Usage:  "encrypt all",
			Action: encryptAll,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:     "key",
					Usage:    "path to key file",
					Required: true,
				},
			},
		},
		{
			Name:   "cover",
			Usage:  "encrypt local file",
			Action: lsbCover,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:     "key",
					Usage:    "path to key file",
					Required: true,
				},
			},
		},
		{
			Name:   "decover",
			Usage:  "decrypt local file",
			Action: deLsbCover,
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:     "key",
					Usage:    "path to key file",
					Required: true,
				},
			},
		},
		{
			Name:   "gen",
			Usage:  "generate password",
			Action: generate,
		},
		{
			Name:   "tui",
			Usage:  "start with tui",
			Action: tuiStart,
		},
	}

	app.Version = VERSION
	app.Usage = "password manager"
	app.EnableBashCompletion = true
	app.Name = "gokeeper"
	app.Authors = []*cli.Author{{Name: "Alctny", Email: "ltozvxe@gmail.com"}}
	db = NewGormDB()
}

func main() {
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
