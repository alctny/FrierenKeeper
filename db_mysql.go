package main

type MySQLDB struct{}

var _ DB = &MySQLDB{}

func (db *MySQLDB) Query(any, []Password) error {
	return nil
}

func (db *MySQLDB) Updata(Password) error {
	return nil
}

func (db *MySQLDB) Delete(Password) error {
	return nil
}

func (db *MySQLDB) Insert(Password) error {
	return nil
}

func NewMySQLDB(user any) *MySQLDB {
	return nil
}
