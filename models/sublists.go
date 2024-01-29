package models

import "time"

// Todo sublists model
type Sublist struct {
	Id          int    `json:"id" gorm:"primary_key" example:"1"`
	Title       string `json:"title" gorm:"type:varchar(255)" example:"Get milk"`
	Description string `json:"description" gorm:"type:varchar(255)" example:"Ultramilk 1 L"`
	ListId      int    `json:"list_id" example:"1"`
	List        List   `json:"list" gorm:"foreignKey:ListId"`
	CreatedAt   time.Time `json:"created_at" example:"2022-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2022-01-01T00:00:00Z"`
	File		string	`json:"file" example:"uploads/file.pdf"`
}