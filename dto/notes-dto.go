package dto

import "time"

type Create struct {
	Content   string `json:"content" form:"content" binding:"required"`
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Update struct {
	ID        uint   `json:"id" form:"id" binding:"required"`
	Content   string `json:"content" form:"content" binding:"required"`
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
