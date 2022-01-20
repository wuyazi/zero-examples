package command

import (
	"bookstore/rpc/book/internal"
	"context"
	"fmt"
	"gitee.com/wuyazi2018/gddd"
	"github.com/tal-tech/go-zero/core/logx"
)

type BookUpdateCommand struct {
	BookId int64
	Price  int64
}

type BookUpdateCommandResult struct {
	BookId int64
	Book   string
	Price  int64
}

func BookUpdateCommandHandle(ctx context.Context, command0 gddd.Command) (result interface{}, err error) {
	cmd, isCommand := command0.(BookUpdateCommand)
	if !isCommand {
		return
	}
	// load account
	book := &BookAggregate{}
	book.Id = cmd.BookId
	var has bool
	has, err = internal.Repository().Load(ctx, book)
	if err != nil {
		return
	}
	if !has {
		return
	}
	logx.Info(book)
	err = book.UpdatePrice(cmd.Price)
	if err != nil {
		return
	}
	ok, saveErr := internal.Repository().Save(ctx, book)
	if saveErr != nil {
		return nil, saveErr
	}
	if !ok {
		err = fmt.Errorf("create account failed")
		return
	}

	result = &BookUpdateCommandResult{
		BookId: book.Id,
		Book:   book.Book,
		Price:  book.Price,
	}

	return book, nil
}
