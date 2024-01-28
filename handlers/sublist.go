package handlers

import (
	dto "backend-technical-test/dto/result"
	sublistsdto "backend-technical-test/dto/sublists"
	"backend-technical-test/models"
	"backend-technical-test/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type SublistHandler struct {
	SublistRepository repository.SublistRepository
}

func NewSublistHandler(sublistRepository repository.SublistRepository) *SublistHandler {
	return &SublistHandler{sublistRepository}
}

// @Tags 		Sublists
// @Summary		Get all sublists
// @Router		/sublists [get]
// @Accept		json
// @Produce		json
// @Param		page			query		int				false	"Limit number of sublists by page (default 1)"
// @Param		search			query		string			false	"Search sublist by title or description"
// @Success		200	{object}	dto.SuccessResult{data=[]models.Sublist}
// @Failure		500	{object}	dto.ErrorResult
func (h *SublistHandler) FindSublists(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit := 10
	offset := (page - 1) * limit

	search := c.QueryParam("search")
	if search == "" {
		search = "%"
	}

	sublists, err := h.SublistRepository.FindSublists(offset, limit, search)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: sublists})
}

// @Tags 		Sublists
// @Summary		Get sublist by id
// @Router		/sublists/{id} [get]
// @Accept		json
// @Produce		json
// @Param		id	path	int	true	"Sublist id"
// @Success		200	{object}	dto.SuccessResult{data=models.Sublist}
// @Failure		500	{object}	dto.ErrorResult
func (h *SublistHandler) FindSublistById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	sublist, err := h.SublistRepository.FindSublistById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: sublist})
}

// @Tags 		Sublists
// @Summary		Get sublist by list id
// @Router		/sublists/list/{listId} [get]
// @Accept		json
// @Produce		json
// @Param		listId	path	int	true	"List id"
// @Param		page		query	int	false	"Limit number of sublists by page (default 1)"
// @Param		search		query	string	false	"Search sublist by title or description"
// @Success		200		{object}	dto.SuccessResult{data=[]models.Sublist}
// @Failure		500		{object}	dto.ErrorResult
func (h *SublistHandler) FindSublistByListId(c echo.Context) error {
	listId, _ := strconv.Atoi(c.Param("listId"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit := 10
	offset := (page - 1) * limit

	search := c.QueryParam("search")
	if search == "" {
		search = "%"
	}

	sublists, err := h.SublistRepository.FindSublistByListId(listId, offset, limit, search)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Sublists found", Data: sublists})
}

// @Tags 		Sublists
// @Summary		Create sublist
// @Router		/sublists [post]
// @Accept		json
// @Produce		json
// @Param		request	body	sublistsdto.SublistRequest	true	"Sublist request"
// @Success		200		{object}	dto.SuccessResult{data=models.Sublist}
// @Failure		500		{object}	dto.ErrorResult
func (h *SublistHandler) CreateSublist(c echo.Context) error {
	request := new(sublistsdto.SublistRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	sublist := models.Sublist{
		ListId: request.ListId,
		Title: request.Title,
		Description: request.Description,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data, err := h.SublistRepository.CreateSublist(sublist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			dto.ErrorResult{
				Code: http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Sublist has been created", Data: data})
}

// @Tags 		Sublists
// @Summary		Update sublist
// @Router		/sublists/{id} [put]
// @Accept		json
// @Produce		json
// @Param		id			path		int					true	"Sublist id"
// @Param		sublist		body		sublistsdto.SublistRequest	true	"Update sublist"
// @Success		200			{object}	dto.SuccessResult{data=models.Sublist}
// @Failure		500			{object}	dto.ErrorResult
func (h *SublistHandler) UpdateSublist(c echo.Context) error {
	request := new(sublistsdto.SublistRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	sublistId, _ := strconv.Atoi(c.Param("id"))
	sublist, err := h.SublistRepository.FindSublistById(sublistId)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	if request.Title != "" {
		sublist.Title = request.Title
	}

	if request.Description != "" {
		sublist.Description = request.Description
	}

	sublist.UpdatedAt = time.Now()

	data, err := h.SublistRepository.UpdateSublist(sublist, sublistId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			dto.ErrorResult{
				Code: http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Sublist has been updated", Data: data})
}

// @Tags 		Sublists
// @Summary		Delete sublist
// @Router		/sublists/{id} [delete]
// @Accept		json
// @Produce		json
// @Param		id	path	int	true	"Sublist id"
// @Success		200	{object}	dto.SuccessResult{data=models.Sublist}
// @Failure		500	{object}	dto.ErrorResult
func (h *SublistHandler) DeleteSublist(c echo.Context) error {
	sublistId, _ := strconv.Atoi(c.Param("id"))
	sublist, err := h.SublistRepository.FindSublistById(sublistId)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	data, err := h.SublistRepository.DeleteSublist(sublist)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			dto.ErrorResult{
				Code: http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Sublist has been deleted", Data: data})
}