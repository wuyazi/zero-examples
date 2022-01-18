package command

import "time"

type BookCreated struct {
	Book       string    `json:"book"`
	Price      int64     `json:"price"`
	CreateTime time.Time `json:"create_time"`
}

func (t *BookCreated) EventName() (name string) {
	name = "BookCreated"
	return
}

type BookChangedPrice struct {
	Price      int64     `json:"price"`
	CreateTime time.Time `json:"create_time"`
}

func (t *BookChangedPrice) EventName() (name string) {
	name = "BookUpdatedPrice"
	return
}
