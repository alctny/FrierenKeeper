package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/alctny/frieren-keeper/model"
	"github.com/alctny/frieren-keeper/util"
	"github.com/urfave/cli/v2"
)

// import passwords from csv, csv format(no header)):
// name,loginId,password,bind,alias,site,comment
func importFromCsv(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return fmt.Errorf("csv file is empty")
	}
	csvFile := ctx.Args().First()
	f, err := os.Open(csvFile)
	if err != nil {
		return err
	}
	defer f.Close()
	csvData, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}
	if len(csvData) < 1 {
		return nil
	}

	keyFile := ctx.Path("key")

	if len(csvData[0]) < 7 {
		return fmt.Errorf("csv format error, should be 'name,loginId,password,bind,alias,site,comment', and no header")
	}
	passwords := []model.Password{}
	for _, line := range csvData[1:] {
		p := model.Password{
			Name:     line[0],
			LoginId:  line[1],
			Password: line[2],
			Bind:     line[3],
			Alias:    line[4],
			Site:     line[5],
			Comment:  line[6],
		}
		if keyFile != "" {
			p.Encrypt = 1
			p.Password, err = util.EncrypeString(p.Password, keyFile)
			if err != nil {
				return err
			}
			p.LoginId, err = util.EncrypeString(p.LoginId, keyFile)
			if err != nil {
				return err
			}
			p.Bind, err = util.EncrypeString(p.Bind, keyFile)
			if err != nil {
				return err
			}
		}
		passwords = append(passwords, p)
	}
	tx := db.Create(&passwords)
	if err != nil {
		return err
	}
	fmt.Println("import passwords success: ", tx.RowsAffected)
	return nil
}
