package store

import (
	"github.com/jinzhu/gorm"
	"github.com/saman2000hoseini/http-monitor/model"
)

type URLStore struct {
	db *gorm.DB
}

func NewURLStore(db *gorm.DB) URLStore {
	return URLStore{db}
}

func (us *URLStore) SuccessCall(u *model.URL) error {
	u.SuccessCall++
	return us.db.Model(u).Update(u).Error
}

func (us *URLStore) FailedCall(u *model.URL) error {
	u.FailedCall++
	if u.FailedCall >= u.Threshold && u.Alert == nil {
		u.Alert = model.NewAlert("Critical threshold violated", u.ID)
	}
	return us.db.Model(u).Update(u).Error
}

func (us *URLStore) PublishAlert(u *model.URL) (string, error) {
	m := u.Alert.Message
	u.Alert = nil
	return m, us.db.Model(u).Update(u).Error
}

func (us *URLStore) GetByID(id uint) (*model.URL, error) {
	var url model.URL
	err := us.db.Where(&model.URL{Model: gorm.Model{ID: id}}).Preload("Alert").First(&url).Error
	return &url, err
}

func (us *URLStore) Create(u *model.URL) error {
	return us.db.Create(u).Error
}

func (us *URLStore) Update(u *model.URL) error {
	return us.db.Save(u).Error
	//return us.db.Model(u).Update(u).Error
}
