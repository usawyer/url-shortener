package main

import (
	"github.com/usawyer/url-shortener/internal/app"
	"github.com/usawyer/url-shortener/internal/config"
	"github.com/usawyer/url-shortener/internal/logger"
	"github.com/usawyer/url-shortener/internal/util"
)

func main() {
	log := logger.New()

	storeType, err := util.ParseFlag()
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg := config.New(log)
	app := app.New(log, cfg, storeType)
	app.Run()
}
