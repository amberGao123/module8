package main

import (
	"context"
	"flag"
	"net/http"

	"module8/api"
	"module8/handler"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	defaultConfig = new(api.Config)

	config string
)

func main() {
	flag.StringVar(&config, "config", "./server.toml", "config")

	if _, err := toml.DecodeFile(config, defaultConfig); err != nil {
		logrus.WithError(err).Fatal("decode config")
	}

	r := gin.Default()

	//middleware logger
	r.Use(gin.Logger())
	//middleware recovery
	r.Use(gin.Recovery())

	//register router
	if err := handler.Register(context.Background(), r); err != nil {
		logrus.WithError(err).Errorf("handler Register")
	}

	//sever
	srv := &http.Server{
		Addr:    defaultConfig.Addr,
		Handler: r,
	}
	logrus.Infof("start Server on %s", defaultConfig.Addr)

	//listen
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Errorf("could not listen server: %v", err)
	}

}
