package utils

import (
	"github.com/jinzhu/gorm"
	"github.com/saman2000hoseini/http-monitor/model"
	"github.com/saman2000hoseini/http-monitor/store"
	"net/http"
	"strings"
	"time"
)

var userS *store.UserStore
var urlS *store.URLStore

func StartMonitoring(d time.Duration, db *gorm.DB) {
	userS = store.NewUserStore(db)
	urlS = store.NewURLStore(db)
	ticker := time.NewTicker(d)
	for {
		var users []model.User
		db.Find(&users)
		for _, user := range users {
			go MonitorURLs(&user)
		}
		<-ticker.C
	}

}

func MonitorURLs(u *model.User) {
	urls, err := userS.GetURLs(u)
	if err != nil {
		return
	}
	for _, url := range urls {
		if HTTPCall(url.Address)/100 == 2 {
			urlS.SuccessCall(&url)
		} else {
			urlS.FailedCall(&url)
		}
	}
}

func HTTPCall(a string) int {
	url := reFormat(a)
	resp, err := http.Get(url)
	if err != nil {
		//TODO
	}
	return resp.StatusCode
}

func reFormat(a string) string {
	strings.Replace(a, "", "www.", 1)
	if !strings.Contains(a, "http://") {
		a = "http://" + a
	}
	return a
}
