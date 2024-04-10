package dao

import "github.com/alctny/frieren-keeper/model"

type PostgreSQLDB struct{}

var _ DB = &PostgreSQLDB{}

func (p *PostgreSQLDB) Query(any, []model.Password) error {
	return nil
}

func (p *PostgreSQLDB) Updata(model.Password) error {
	return nil
}

func (p *PostgreSQLDB) Delete(model.Password) error {
	return nil
}

func (p *PostgreSQLDB) Insert(model.Password) error {
	return nil
}
