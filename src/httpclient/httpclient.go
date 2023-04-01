package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/BigBoulard/logger/src/conf"
	"github.com/BigBoulard/logger/src/domain"
	"github.com/go-resty/resty/v2"
)

type client struct {
	Resty *resty.Client
}

type HttpClient interface {
	GetTodos() ([]domain.Todo, error)
}

func NewClient() HttpClient {
	conf.LoadEnv() // load env vars
	r := resty.New()
	if conf.Env.AppMode == "debug" { // Debug mode is taken from the env vars
		r.SetDebug(true)
	}
	r.SetBaseURL("http://jsonplaceholder.typicode.com")

	return &client{
		Resty: r,
	}
}

func (c *client) GetTodos() ([]domain.Todo, error) {
	resp, err := c.Resty.
		R().
		SetHeader("Accept", "application/json").
		Get("/todos")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() > 399 {
		return nil, errors.New(fmt.Sprintf("%d - an error occurred", resp.StatusCode()))
	}

	var todos []domain.Todo
	err = json.Unmarshal(resp.Body(), &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
