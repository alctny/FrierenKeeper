package main

type PostgreSQLDB struct{}

var _ DB = &PostgreSQLDB{}

func (p *PostgreSQLDB) Query(any, []Password) error {
	return nil
}

func (p *PostgreSQLDB) Updata(Password) error {
	return nil
}

func (p *PostgreSQLDB) Delete(Password) error {
	return nil
}

func (p *PostgreSQLDB) Insert(Password) error {
	return nil
}
