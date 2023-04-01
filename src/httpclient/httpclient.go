package httpcli

import (
	"errors"
	"fmt"

	"github.com/BigBoulard/logger/src/conf"
	"github.com/go-resty/resty/v2"
)

type client struct {
	Resty *resty.Client
}

type httpClient interface {
	GetTodos()
}

func NewClient() httpClient {
	conf.LoadEnv() // load env vars
	r := resty.New()
	r.SetDebug(conf.Env.AppMode)
	r.SetBaseURL("http://jsonplaceholder.typicode.com")

	return &client{
		Resty: r,
	}
	return &client{}
}

func (c *client) GetTodos() error {
	resp, err := c.Resty.
		R().
		SetHeader("Accept", "application/json").
		Get("/todo/1")

	if err != nil {
		return err
	}

	if resp.StatusCode() > 399 {
		return errors.New(fmt.Sprintf("%s - an error occurred", resp.StatusCode()))
	}

	return nil
}
