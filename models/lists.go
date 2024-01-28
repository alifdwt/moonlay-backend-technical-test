package models

import "time"

// Todo lists model
type List struct {
	Id          int    		`json:"id" gorm:"primary_key" example:"1"`
	Title       string 		`json:"title" gorm:"type:varchar(255)" example:"Weekly Shopping List"`
	Description string		`json:"description" gorm:"type:varchar(255)" example:"My shopping list"`
	Sublists    []Sublist 	`json:"sublists" gorm:"foreignKey:ListId;references:Id"`
	CreatedAt   time.Time 	`json:"created_at" example:"2022-01-01T00:00:00Z"`
	UpdatedAt   time.Time 	`json:"updated_at" example:"2022-01-01T00:00:00Z"`
	// Priority    int    `json:"priority"`
	// IsDone      bool   `json:"is_done"`
	// UserId      int    `json:"user_id"`
	// User        User   `json:"user"`
}