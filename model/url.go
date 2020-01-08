package model

import (
	"github.com/jinzhu/gorm"
)

type URL struct {
	gorm.Model
	Address     string `gorm:"type:varchar(250);not null"`
	Threshold   uint
	SuccessCall uint
	FailedCall  uint
	Alert       *Message
	UserID      uint `gorm:"foreignkey:UserID"`
}

func NewURL(address string, threshold uint) *URL {
	return &URL{Address: address, SuccessCall: 0, FailedCall: 0, Threshold: threshold}
}

type Message struct {
	gorm.Model
	Message string
	RefID   uint `gorm:"index:url_id"`
}

func NewAlert(message string, referenceID uint) *Message {
	return &Message{Message: message, RefID: referenceID}
}
