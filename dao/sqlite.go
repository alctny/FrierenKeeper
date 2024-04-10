package dao

import (
	"github.com/alctny/frieren-keeper/model"
)

type SQLiteDB struct{}

var _ DB = &SQLiteDB{}

func (db *SQLiteDB) Query(any, []model.Password) error {
	return nil
}

func (db *SQLiteDB) Updata(model.Password) error {
	return nil
}

func (db *SQLiteDB) Delete(model.Password) error {
	return nil
}

func (db *SQLiteDB) Insert(model.Password) error {
	return nil
}

func NewSQLiteDB(path string) *SQLiteDB {
	return nil
}
