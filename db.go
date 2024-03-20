package main

import (
	"fmt"
	"os/user"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	user, err := user.Current()
	ErrorWithEixt(err)
	dbPath := fmt.Sprintf("%s/%s", user.HomeDir, ".gokeeper.db")
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	ErrorWithEixt(err)
	err = db.AutoMigrate(&Password{})
	ErrorWithEixt(err)
}
