package svc

import (
	"bookstore/api/internal/config"
	"bookstore/rpc/add/adder"
	"bookstore/rpc/book/bookclient"
	"bookstore/rpc/check/checker"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	Adder      adder.Adder
	BookClient bookclient.Book
	Checker    checker.Checker
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Adder:      adder.NewAdder(zrpc.MustNewClient(c.Add)),
		BookClient: bookclient.NewBook(zrpc.MustNewClient(c.Book)),
		Checker:    checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
