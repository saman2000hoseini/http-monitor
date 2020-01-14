package monitor

import (
	"fmt"
	"github.com/jinzhu/gorm"
	handler2 "github.com/saman2000hoseini/http-monitor/handler"
	"github.com/saman2000hoseini/http-monitor/model"
	"net/http"
	"strings"
	"sync"
	"time"
)

var handler *handler2.Handler

func StartMonitoring(d time.Duration, db *gorm.DB) {
	var wg sync.WaitGroup
	handler = handler2.NewHandler(db)
	ticker := time.NewTicker(d)
	for {
		var users []model.User
		db.Find(&users)
		for _, user := range users {
			wg.Add(1)
			go MonitorURLs(&user, &wg)
		}
		wg.Wait()
		<-ticker.C
	}

}

func MonitorURLs(u *model.User, wg *sync.WaitGroup) {
	defer wg.Done()
	user, err := handler.UserStore.GetByID(u.ID)
	if err != nil {
		return
	}
	user.URLs, _ = handler.URLStore.GetByUser(user.ID)
	for _, url := range user.URLs {
		if HTTPCall(url.Address)/100 == 2 {
			handler.URLStore.SuccessCall(&url)
		} else {
			err := handler.URLStore.FailedCall(&url)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func HTTPCall(a string) int {
	url := reFormat(a)
	resp, err := http.Get(url)
	if err != nil {
		//TODO
		return 500
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
