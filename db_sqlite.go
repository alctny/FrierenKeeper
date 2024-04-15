package main

type SQLiteDB struct{}

var _ DB = &SQLiteDB{}

func (db *SQLiteDB) Query(any, []Password) error {
	return nil
}

func (db *SQLiteDB) Updata(Password) error {
	return nil
}

func (db *SQLiteDB) Delete(Password) error {
	return nil
}

func (db *SQLiteDB) Insert(Password) error {
	return nil
}

func NewSQLiteDB(path string) *SQLiteDB {
	return nil
}
