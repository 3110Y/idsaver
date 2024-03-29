package entity

import (
	"time"
)

type IDEntity struct {
	Id       string
	Uid      string
	Len      uint64
	Allow    bool
	CreateAt time.Time `db:"create_at"`
	UpdateAt time.Time `db:"update_at"`
}
