package entity

import (
	"time"
)

type Note struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Content   string    `json:"content"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `gorm:"type:datetime;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;autoCreateTime" json:"updated_at"`
}
