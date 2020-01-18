package monitor

import (
	"fmt"
	"github.com/carlescere/scheduler"
	"github.com/jinzhu/gorm"
	handler2 "github.com/saman2000hoseini/http-monitor/handler"
	"github.com/saman2000hoseini/http-monitor/model"
	"net/http"
	"strings"
	"sync"
	"time"
)

var handler *handler2.Handler
var wg *sync.WaitGroup

func worker(urls <-chan model.URL) {
	for url := range urls {
		monitorURL(&url)
	}
}

//foreach user dedicate goroutine to monitor added urls
func StartMonitoring(d time.Duration, db *gorm.DB) {
	wg = new(sync.WaitGroup)
	scheduler.Every().Day().At("00:00").Run(resetURLs)
	handler = handler2.NewHandler(db)
	ticker := time.NewTicker(d)
	for {
		wg.Wait()
		var urls []model.URL
		db.Find(&urls)
		urlsChan := make(chan model.URL, len(urls))
		go func() {
			for i := 0; i < 30; i++ {
				go worker(urlsChan)
			}
		}()
		for _, url := range urls {
			wg.Add(1)
			urlsChan <- url
		}
		close(urlsChan)
		<-ticker.C
	}
}

func monitorURL(url *model.URL) {
	defer wg.Done()
	if HTTPCall(url.Address)/100 == 2 {
		handler.URLStore.SuccessCall(url)
	} else {
		err := handler.URLStore.FailedCall(url)
		if err != nil {
			fmt.Println(err)
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

//extract url address into standard format
func reFormat(a string) string {
	strings.Replace(a, "", "www.", 1)
	if !strings.Contains(a, "http://") && !strings.Contains(a, "https://") {
		a = "http://" + a
	}
	return a
}

func resetURLs() {
	wg.Add(1)
	defer wg.Done()
	handler.URLStore.Reset()
}
