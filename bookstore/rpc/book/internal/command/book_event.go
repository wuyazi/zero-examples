package command

import "time"

type BookCreated struct {
	Book       string    `json:"book"`
	Price      int64     `json:"price"`
	CreateTime time.Time `json:"create_time"`
}

type BookChangedPrice struct {
	Price int64 `json:"price"`
}
