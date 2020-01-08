package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type URL struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Address     string `gorm:"type:varchar(250);unique_index;not null"`
	Threshold   uint
	SuccessCall uint
	FailedCall  uint
	CreatedAt   time.Time
	Alert       *Message
	UserID      uint `gorm:"foreignkey:UserID"`
}

func NewURL(address string, threshold uint) *URL {
	return &URL{Address: address, SuccessCall: 0, FailedCall: 0, CreatedAt: time.Now(), Threshold: threshold}
}

type Message struct {
	gorm.Model
	Message     string
	ReferenceID uint `gorm:"foreignkey:ID"`
}
