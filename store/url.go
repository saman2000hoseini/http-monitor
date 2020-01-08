package store

import (
	"github.com/jinzhu/gorm"
	"github.com/saman2000hoseini/http-monitor/model"
)

type URLStore struct {
	db *gorm.DB
}

func NewURLStore(db *gorm.DB) *URLStore {
	return &URLStore{db}
}

func (us *URLStore) SuccessCall(u *model.URL) error {
	u.SuccessCall++
	return us.db.Model(u).Update(u).Error
}

func (us *URLStore) FailedCall(u *model.URL) error {
	u.FailedCall++
	return us.db.Model(u).Update(u).Error
}
