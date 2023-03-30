package models

import (
	"time"
)

type Book struct {
	ID        int       `gorm:"primary_key" json:"id"`
	NameBook  string    `gorm:"type:varchar(100)" json:"name_book" validate:"required"`
	Author    string    `gorm:"type:varchar(100)" json:"author" validate:"required"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}
