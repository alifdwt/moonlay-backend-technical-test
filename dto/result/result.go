package dto

import "time"

type SuccessResult struct {
	Code    int         `json:"code" example:"200"`
    Message string      `json:"message" example:"success"`
	Data	interface{} `json:"data"`
}

type ErrorResult struct {
	Code    int         `json:"code" example:"500"`
	Message string      `json:"message" example:"internal server error"`
}

type ListResponse struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
