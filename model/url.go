package model

import (
	"time"
)

const alert = "Critical threshold violated!"

type URL struct {
	ID         uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Address    string `gorm:"type:varchar(250);unique_index;not null"`
	Threshold  uint
	ErrorCount uint
	CreatedAt  time.Time
	UserID     uint
}

func NewURL(address string, threshold uint) *URL {
	return &URL{Address: address, ErrorCount: 0, CreatedAt: time.Now(), Threshold: threshold}
}
