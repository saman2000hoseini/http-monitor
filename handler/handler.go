package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/saman2000hoseini/http-monitor/store"
)

type Handler struct {
	UserStore store.UserStore
	URLStore  store.URLStore
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{store.NewUserStore(db), store.NewURLStore(db)}
}
