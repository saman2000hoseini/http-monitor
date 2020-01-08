package main

import (
	"github.com/saman2000hoseini/http-monitor/db"
	handler2 "github.com/saman2000hoseini/http-monitor/handler"
	"github.com/saman2000hoseini/http-monitor/router"
	"github.com/saman2000hoseini/http-monitor/utils"
	"time"
)

func main() {
	r := router.Router()
	db, err := db.FirstSetup()
	if err != nil {
		//TODO
	}
	v1 := r.Group("api")
	handler := handler2.NewHandler(db)
	handler.Setup(v1)
	go utils.StartMonitoring(15*time.Minute, db)
	r.Logger.Fatal(r.Start("localhost:54321"))
}
