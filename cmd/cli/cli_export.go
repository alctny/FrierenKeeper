package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/alctny/frieren-keeper/model"
	"github.com/urfave/cli/v2"
)

// export all passwords to csv
// TODO: export to csv
func export2Csv(ctx *cli.Context) error {
	keyFile := ctx.Path("key")
	passwords := []model.Password{}
	if keyFile != "" {
		tx := db.Find(&passwords)
		if tx.Error != nil {
			return tx.Error
		}
		var err error
		passwords, err = model.DecryptPasswords(passwords, keyFile)
		if err != nil {
			return err
		}
	} else {
		tx := db.Where("encrypt = 0").Find(&passwords)
		if tx.Error != nil {
			return tx.Error
		}
	}

	outputFile := ctx.Path("output")
	if outputFile == "" {
		outputFile = "output.csv"
	}
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()
	csvWriter := csv.NewWriter(f)
	defer csvWriter.Flush()
	for _, p := range passwords {
		csvWriter.Write([]string{p.Name, p.LoginId, p.Password, p.Bind, p.Alias, p.Site, p.Comment})
	}
	fmt.Println("export passwords success to: ", outputFile)
	return nil
}
