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
	return us.db.Save(u).Error
}

func (us *URLStore) FailedCall(u *model.URL) error {
	u.FailedCall++
	var err error
	u.Alert, err = us.GetAlert(u.ID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if gorm.IsRecordNotFoundError(err) {
		u.Alert = nil
	}
	if u.FailedCall >= u.Threshold && u.Alert == nil {
		u.Alert = model.NewAlert("Critical threshold violated", u.ID)
	}
	return us.db.Save(u).Error
}

func (us *URLStore) PublishAlert(u *model.URL) (string, error) {
	m := u.Alert.Message
	u.Alert = nil
	return m, us.db.Save(u).Error
}

func (us *URLStore) GetByID(id uint) (*model.URL, error) {
	var url model.URL
	err := us.db.Where(&model.URL{Model: gorm.Model{ID: id}}).First(&url).Error
	return &url, err
}

func (us *URLStore) GetByUser(id uint) ([]model.URL, error) {
	var urls []model.URL
	err := us.db.Where(&model.URL{UserID: id}).Find(&urls).Error
	return urls, err
}

func (us *URLStore) GetAlert(id uint) (*model.Message, error) {
	alert := &model.Message{}
	err := us.db.Where(&model.Message{RefID: id}).First(alert).Error
	return alert, err
}

func (us *URLStore) Create(u *model.URL) error {
	return us.db.Create(u).Error
}

func (us *URLStore) Update(u *model.URL) error {
	return us.db.Save(u).Error
	//return us.db.Model(u).Update(u).Error
}
