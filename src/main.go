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

	// GinMode is set to "debug" in the /.env file
	gin.SetMode(conf.Env.GinMode)
	router = gin.Default()

	// creates a controller that uses an httpclient to call https://jsonplaceholder.typicode.com/todos
	controller := controller.NewController(
		httpclient.NewClient(),
	)
	router.GET("/todos", controller.GetTodos)

	// run the router ton host and port specified in the /.env file
	err := router.Run(fmt.Sprintf("%s:%s", conf.Env.Host, conf.Env.Port))
	if err != nil {
		log.Fatal("application/main", err)
	}
}
