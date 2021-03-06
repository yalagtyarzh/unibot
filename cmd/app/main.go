package main

import (
	"flag"
	"github.com/yalagtyarzh/unibot/internal"
	"github.com/yalagtyarzh/unibot/internal/config"
	"github.com/yalagtyarzh/unibot/pkg/logging"
	"log"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "config", "configs/dev.yml", "config file path")
}

func main() {
	flag.Parse()

	log.Print("config initializing")
	cfg := config.GetConfig(cfgPath)

	log.Print("logger initializing")
	logging.Init(cfg.AppConfig.LogLevel)
	logger := logging.GetLogger()

	logger.Println("Creating Application")
	app, err := internal.NewApp(logger, cfg)
	if err != nil {
		logger.Fatal(err)
	}

	app.Run()
}
