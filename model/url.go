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
	Alert       *Message `gorm:"null"`
	UserID      uint
}

func NewURL(address string, threshold uint, userID uint) *URL {
	return &URL{Address: address, SuccessCall: 0, FailedCall: 0, Threshold: threshold, UserID: userID}
}

type Message struct {
	gorm.Model
	Message    string
	FailedCall uint
	RefID      uint `gorm:"not null"`
}

func NewAlert(message string, failedCall uint, referenceID uint) *Message {
	return &Message{Message: message, FailedCall: failedCall, RefID: referenceID}
}
