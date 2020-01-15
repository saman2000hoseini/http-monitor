package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/saman2000hoseini/http-monitor/model"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./myDB.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

//create tables from models
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{}, &model.URL{}, &model.Message{}).Error
	return err
}

//make database and create tables for the first time
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

func TestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../myDB_test.db")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	if err := os.Remove("./../myDB_test.db"); err != nil {
		return err
	}
	return nil
}
