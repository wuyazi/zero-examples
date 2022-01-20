package command

import (
	"fmt"
	"gitee.com/wuyazi2018/gddd"
	"strings"
	"time"
)

type BookAggregate struct {
	gddd.AbstractAggregate
	Book       string    `json:"book"`
	Price      int64     `json:"price"`
	CreateTime time.Time `json:"create_time"`
}

func (b *BookAggregate) Create(book string, price int64) (err error) {
	book = strings.TrimSpace(book)
	if book == "" {
		err = fmt.Errorf("book is empty")
		return
	}
	if price == 0 {
		err = fmt.Errorf("price is 0")
		return
	}

	utcNow := time.Now().UTC()
	b.Apply(&BookCreated{
		Book:       book,
		Price:      price,
		CreateTime: utcNow,
	})
	return
}

func (b *BookAggregate) OnBookCreated(event *BookCreated) (err error) {
	b.Book = event.Book
	b.Price = event.Price
	b.CreateTime = event.CreateTime
	return
}

func (b *BookAggregate) UpdatePrice(price int64) (err error) {
	if price == 0 {
		err = fmt.Errorf("price is 0")
		return
	}

	b.Apply(&BookChangedPrice{
		Price: price,
	})
	return
}

func (b *BookAggregate) OnBookChangedPrice(event *BookChangedPrice) (err error) {
	b.Price = event.Price
	return
}
