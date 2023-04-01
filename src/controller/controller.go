package controller

import (
	"net/http"

	"github.com/BigBoulard/logger/src/httpclient"
	"github.com/BigBoulard/logger/src/log"
	"github.com/gin-gonic/gin"
)

func NewController(httpClient httpclient.HttpClient) ctrl {
	return &controller{
		httpClient: httpClient,
	}
}

type ctrl interface {
	GetTodos(c *gin.Context)
}

type controller struct {
	httpClient httpclient.HttpClient
}

func (ctrl *controller) GetTodos(c *gin.Context) {
	todos, err := ctrl.httpClient.GetTodos()
	if err != nil {
		log.Error("controller/GetTodos", err)
		c.JSON(500, err) // we return 500 just for simplicity
		return
	}
	c.JSON(http.StatusOK, todos)
}
