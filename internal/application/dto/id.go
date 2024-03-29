package dto

import "time"

type IdDTO struct {
	Id       *string
	Uid      string
	Len      uint64
	Allow    bool
	CreateAt *time.Time
	UpdateAt *time.Time
}
