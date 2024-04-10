package dao

import (
	"fmt"
	"os"
	"os/user"

	"github.com/alctny/frieren-keeper/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DB interface {
	Query(any, []model.Password) error
	Updata(model.Password) error
	Delete(model.Password) error
	Insert(model.Password) error
}

func NewGormDB() *gorm.DB {

	var dbPath string
	_, err := os.Stat("gokeeper.db")
	if err == nil {
		dbPath = "gokeeper.db"
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		dbPath = fmt.Sprintf("%s/%s", user.HomeDir, ".gokeeper.db")
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Password{})
	if err != nil {
		panic(err)
	}

	return db
}
