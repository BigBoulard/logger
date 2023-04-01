package main

import (
	"fmt"

	"github.com/BigBoulard/logger/src/conf"
	"github.com/BigBoulard/logger/src/controller"
	"github.com/BigBoulard/logger/src/httpclient"
	"github.com/BigBoulard/logger/src/log"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func main() {
	// load the env variables
	conf.LoadEnv()

	gin.SetMode(conf.Env.GinMode) // mode "debug" from .env

	router = gin.Default()
	err := router.SetTrustedProxies(nil)

	if err != nil {
		log.Fatal("application/main", err)
	}

	controller := controller.NewController(
		httpclient.NewClient(),
	)
	router.GET("/todos", controller.GetTodos)

	err = router.Run(fmt.Sprintf("%s:%s", conf.Env.Host, conf.Env.Port))
	if err != nil {
		log.Fatal("application/main", err)
	}
}
