package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/saman2000hoseini/http-monitor/model"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{}, &model.URL{}, &model.Message{}).Error
	return err
}

func FirstSetup() (*gorm.DB, error) {
	db, err := NewDB()
	if err != nil {
		return nil, errors.New("Error on creating db")
	}
	err = Migrate(db)
	if err != nil {
		return nil, errors.New("Error on creating tables")
	}
	return db, nil
}
