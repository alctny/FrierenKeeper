package dao

import "github.com/alctny/frieren-keeper/model"

type MySQLDB struct{}

var _ DB = &MySQLDB{}

func (db *MySQLDB) Query(any, []model.Password) error {
	return nil
}

func (db *MySQLDB) Updata(model.Password) error {
	return nil
}

func (db *MySQLDB) Delete(model.Password) error {
	return nil
}

func (db *MySQLDB) Insert(model.Password) error {
	return nil
}

func NewMySQLDB(user any) *MySQLDB {
	return nil
}
