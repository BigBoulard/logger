package application

import (
	"fmt"

	"github.com/BigBoulard/logger/src/conf"
	"github.com/BigBoulard/logger/src/controller"
	"github.com/BigBoulard/logger/src/log"
	"github.com/BigBoulard/logger/src/repository"
	"github.com/BigBoulard/logger/src/service"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func StartApplication() {
	// load the env variables
	conf.LoadEnv()

	gin.SetMode(conf.Env.GinMode) // mode "debug" from .env

	router = gin.Default()
	err := router.SetTrustedProxies(nil)

	router.Use(log.Middleware())

	if err != nil {
		log.Fatal("application/main", err)
	}

	controller := controller.NewController(
		service.NewService(
			repository.NewRepo(),
		),
	)
	router.GET("/products/:product-id", controller.GetProduct)

	err = router.Run(fmt.Sprintf("%s:%s", conf.Env.Host, conf.Env.Port))
	if err != nil {
		log.Fatal("application/main", err)
	}
}
