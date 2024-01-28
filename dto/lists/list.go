package listsdto

import "backend-technical-test/models"

type ListRequest struct {
	Title string `json:"title" form:"title" validate:"required" example:"Shopping List"`
	Description string `json:"description" form:"description" example:"My shopping list"`
}

type ListResponse struct {
	Id int `json:"id" example:"1"`
	Title string `json:"title" example:"Shopping List"`
	Description string `json:"description" example:"My shopping list"`
	Sublists []models.Sublist `json:"sublists"`
}