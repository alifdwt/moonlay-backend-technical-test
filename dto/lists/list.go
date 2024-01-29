package listsdto

import (
	"backend-technical-test/models"
	"mime/multipart"
)

type ListRequest struct {
	Id int `json:"id,omitempty" form:"id" example:"1"`
	Title string `json:"title" form:"title" validate:"required" example:"Shopping List"`
	Description string `json:"description" form:"description" example:"My shopping list"`
	File *multipart.FileHeader `json:"file,omitempty" form:"file" example:"file.pdf"`
}

type ListResponse struct {
	Id int `json:"id" example:"1"`
	Title string `json:"title" example:"Shopping List"`
	Description string `json:"description" example:"My shopping list"`
	File string `json:"file" example:"uploads/file.pdf"`
	Sublists []models.Sublist `json:"sublists"`
}