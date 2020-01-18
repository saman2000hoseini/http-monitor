package store

import (
	"github.com/jinzhu/gorm"
	"github.com/saman2000hoseini/http-monitor/model"
	"time"
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

//check if failed calls passed the threshold to throw alert
func (us *URLStore) FailedCall(u *model.URL) error {
	u.FailedCall++
	var err error
	err = us.db.Save(u).Error
	if u.FailedCall >= u.Threshold {
		u.Alert, err = us.GetAlert(u.ID)
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return err
		}
		if gorm.IsRecordNotFoundError(err) {
			u.Alert = model.NewAlert("Critical threshold violated", u.FailedCall, u.ID)
		} else {
			u.Alert.FailedCall = u.FailedCall
		}
		return us.db.Save(u.Alert).Error
	}
	return err
}

//thrown alert has been seen so it should reset
func (us *URLStore) PublishAlert(u *model.URL) error {
	u.FailedCall = 0
	us.db.Save(u)
	return us.db.Delete(u.Alert).Error
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
	err := us.db.Table("messages").Order("created_at desc").Where("ref_id = ?", id).First(alert).Error
	//	err := us.db.Where(&model.Message{RefID: id}).First(alert).Error
	return alert, err
}

func (us *URLStore) Create(u *model.URL) error {
	return us.db.Create(u).Error
}

func (us *URLStore) Update(u *model.URL) error {
	return us.db.Save(u).Error
	//return us.db.Model(u).Update(u).Error
}

func (us *URLStore) GetTotalFailure(u *model.URL) (uint, error) {
	var alerts []model.Message
	var sum uint
	t := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	err := us.db.Unscoped().Where("created_at >= ? AND ref_id = ?", t, u.ID).Find(&alerts).Error
	for _, alert := range alerts {
		sum += alert.FailedCall
	}
	if u.FailedCall < u.Threshold {
		sum += u.FailedCall
	}
	return sum, err
}

func (us *URLStore) Reset() {
	var urls []model.URL
	us.db.Find(&urls)
	for _, url := range urls {
		if url.FailedCall >= url.Threshold {
			url.Alert, _ = us.GetAlert(url.ID)
			us.PublishAlert(&url)
		}
		url.FailedCall = 0
		url.SuccessCall = 0
		us.db.Save(url)
	}
}
