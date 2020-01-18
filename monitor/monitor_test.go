package monitor

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/saman2000hoseini/http-monitor/db"
	handler2 "github.com/saman2000hoseini/http-monitor/handler"
	"github.com/saman2000hoseini/http-monitor/model"
	"github.com/saman2000hoseini/http-monitor/router"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHTTPCall(t *testing.T) {
	tearDown()
	setup()
	url, _ := h.URLStore.GetByID(1)
	req := HTTPCall(url.Address)
	assert.Equal(t, req, 200)
}

var (
	d *gorm.DB
	h *handler2.Handler
	e *echo.Echo
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	d = db.TestDB()
	db.Migrate(d)
	h = handler2.NewHandler(d)
	e = router.Router()
	loadData()
}

func tearDown() {
	_ = d.Close()
	if err := db.DropTestDB(); err != nil {
		log.Fatal(err)
	}
}

func loadData() error {
	url1 := model.URL{
		Address:     "google.com",
		Threshold:   5,
		SuccessCall: 0,
		FailedCall:  0,
		Alert:       nil,
		UserID:      1,
	}

	url2 := model.URL{
		Address:     "gobyexample.com",
		Threshold:   10,
		SuccessCall: 0,
		FailedCall:  0,
		Alert:       nil,
		UserID:      1,
	}

	url3 := model.URL{
		Address:     "github.com",
		Threshold:   15,
		SuccessCall: 0,
		FailedCall:  0,
		Alert:       nil,
		UserID:      2,
	}

	url4 := model.URL{
		Address:     "www.digikala.com",
		Threshold:   10,
		SuccessCall: 0,
		FailedCall:  0,
		Alert:       nil,
		UserID:      2,
	}

	url5 := model.URL{
		Address:     "https://1995parham.me/about",
		Threshold:   5,
		SuccessCall: 0,
		FailedCall:  0,
		Alert:       nil,
		UserID:      1,
	}

	u1 := model.User{
		Username: "user1",
		URLs:     []model.URL{url1, url2},
	}
	err := u1.HashPassword("secretpass1")
	if err != nil {
		return err
	}
	if err = h.UserStore.Create(&u1); err != nil {
		return err
	}

	u2 := model.User{
		Username: "user2",
		URLs:     []model.URL{url3, url4, url5},
	}
	err = u2.HashPassword("secretPass2")
	if err != nil {
		return err
	}
	if err = h.UserStore.Create(&u2); err != nil {
		return err
	}

	return nil
}
