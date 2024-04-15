package main

type ImageDB struct{}

var _ DB = &ImageDB{}

func (db *ImageDB) Query(any, []Password) error {
	return nil
}

func (db *ImageDB) Updata(Password) error {
	return nil
}

func (db *ImageDB) Delete(Password) error {
	return nil
}

func (db *ImageDB) Insert(Password) error {
	return nil
}

func NewImageDB(img string, key string) *ImageDB {
	return nil
}
