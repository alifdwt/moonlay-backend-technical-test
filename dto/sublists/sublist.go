package sublistsdto

type SublistRequest struct {
	Title string `json:"title" form:"title" validate:"required" example:"Get milk"`
	Description string `json:"description" form:"description" example:"Ultramilk 1 L"`
	ListId int `json:"list_id" form:"list_id" validate:"required" example:"1"`
}

type SublistResponse struct {
	Id int `json:"id" example:"1"`
	Title string `json:"title" example:"Get milk"`
	Description string `json:"description" example:"Ultramilk 1 L"`
	ListId int `json:"list_id" example:"1"`
}