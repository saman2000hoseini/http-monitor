package store

import (
	"github.com/jinzhu/gorm"
	"github.com/saman2000hoseini/http-monitor/model"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStore {
	return UserStore{db}
}

func (us *UserStore) Create(u *model.User) error {
	return us.db.Create(u).Error
}

func (us *UserStore) Update(u *model.User) error {
	return us.db.Save(u).Error
	//return us.db.Model(u).Update(u).Error
}

func (us *UserStore) GetURLs(u *model.User) ([]model.URL, error) {
	var urls []model.URL
	err := us.db.Preload("URls").Preload("Alert").Where(u).Find(u).Error
	urls = u.URLs
	return urls, err
}

/*
func (us *UserStore) AddURL(user *model.User,u model.URL ) error{

	return us.db.
}*/

func (us *UserStore) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := us.db.Where(&model.User{Username: username}).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return &user, nil
	}
	return &user, err
}

func (us *UserStore) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := us.db.Where(&model.User{Model: gorm.Model{ID: id}}).Preload("URLs").First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return &user, nil
	}
	return &user, err
}
