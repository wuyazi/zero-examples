package command

import (
	"bookstore/rpc/book/internal"
	"context"
	"fmt"
	"gitee.com/wuyazi2018/gddd"
)

type BookCreateCommand struct {
	Book  string
	Price int64
}

type BookCreateCommandResult struct {
	BookId int64
	Book   string
	Price  int64
}

func BookCreateCommandHandle(ctx context.Context, command0 gddd.Command) (result interface{}, err error) {
	cmd, isCommand := command0.(BookCreateCommand)
	if !isCommand {
		return
	}
	// create account
	book := &BookAggregate{}
	err = book.InitId()
	if err != nil {
		return
	}
	err = book.Create(cmd.Book, cmd.Price)
	if err != nil {
		return
	}
	ok, saveErr := internal.Repository().Save(ctx, book)
	if saveErr != nil {
		return
	}
	if !ok {
		err = fmt.Errorf("create account failed")
		return
	}

	result = &BookCreateCommandResult{
		BookId: book.Id,
		Book:   book.Book,
		Price:  book.Price,
	}

	return book, nil
}
