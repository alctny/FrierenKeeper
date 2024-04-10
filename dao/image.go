package dao

import "github.com/alctny/frieren-keeper/model"

type ImageDB struct{}

var _ DB = &ImageDB{}

func (db *ImageDB) Query(any, []model.Password) error {
	return nil
}

func (db *ImageDB) Updata(model.Password) error {
	return nil
}

func (db *ImageDB) Delete(model.Password) error {
	return nil
}

func (db *ImageDB) Insert(model.Password) error {
	return nil
}

func NewImageDB(img string, key string) *ImageDB {
	return nil
}
