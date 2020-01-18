package main

import (
	"github.com/saman2000hoseini/http-monitor/db"
	handler2 "github.com/saman2000hoseini/http-monitor/handler"
	"github.com/saman2000hoseini/http-monitor/monitor"
	"github.com/saman2000hoseini/http-monitor/router"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"time"
)

func main() {
	filename, _ := filepath.Abs("./config.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	r := router.Router()
	db, err := db.FirstSetup()
	if err != nil {
		//TODO
	}
	v1 := r.Group("api")
	handler := handler2.NewHandler(db)
	handler.Setup(v1)
	go monitor.StartMonitoring(time.Duration(config.Monitor.Ticker)*time.Minute, db)
	//go monitor.StartMonitoring(5*time.Second, db)
	r.Logger.Fatal(r.Start(config.Server.Host + config.Server.Port))
}

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Monitor struct {
		Ticker uint `yaml:"ticker"`
	}
}
