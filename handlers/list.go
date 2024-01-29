package handlers

import (
	listsdto "backend-technical-test/dto/lists"
	dto "backend-technical-test/dto/result"
	"backend-technical-test/models"
	"backend-technical-test/repository"
	"backend-technical-test/util"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ListHandler struct {
	ListRepository repository.ListRepository
}

func NewListHandler(listRepository repository.ListRepository) *ListHandler {
	return &ListHandler{
		ListRepository: listRepository,
	}
}

// @Tags 		Lists
// @Summary		Get all lists
// @Description	Get all lists from database. Use query parameter to limit and offset
// @Router		/lists [get]
// @Accept		json
// @Produce		json
// @Param		page			query		int				false	"Limit number of lists by page (default 1)"
// @Param		withSublists	query		bool			false	"Include sublists in response"
// @Param		search			query		string			false	"Search list by title or description"
// @Success		200				{object}	dto.SuccessResult{data=[]models.List}
// @Failure		500				{object}	dto.ErrorResult
func (h *ListHandler) FindLists(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit := 10
	offset := (page - 1) * limit

	withSublists := c.QueryParam("withSublists")
	if withSublists == "" {
		withSublists = "false"
	}

	search := c.QueryParam("search")
	if search == "" {
		search = "%"
	}

	if withSublists == "true" {
		lists, err := h.ListRepository.GetLists(offset, limit, true, search)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				dto.ErrorResult{
					Code: http.StatusInternalServerError,
					Message: err.Error(),
				})
		}

		return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: lists})
	}

	lists, err := h.ListRepository.GetLists(offset, limit, false, search)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			dto.ErrorResult{
				Code: http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Lists found", Data: lists})
}

// @Tags 		Lists
// @Summary		Get list by id
// @Router		/lists/{id} [get]
// @Accept		json
// @Produce		json
// @Param		id	path	int	true	"List id"
// @Success		200	{object}	dto.SuccessResult{data=models.List}
// @Failure		500	{object}	dto.ErrorResult
func (h *ListHandler) FindListById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	list, err := h.ListRepository.FindListById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "List found", Data: list})
}

// @Tags 		Lists
// @Summary		Create list
// @Router		/lists [post]
// @Accept		json
// @Produce		json
// @Param		list	body		listsdto.ListRequest	true	"Create list"
// @Success		200		{object}	dto.SuccessResult{data=models.List}
// @Failure		500		{object}	dto.ErrorResult
func (h *ListHandler) CreateList(c echo.Context) error {
	request := new(listsdto.ListRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	list := models.List{
		Title:       request.Title,
		Description: request.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	file, err := c.FormFile("file")
	if err == nil {
		filePath, err := util.SaveFile(file)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		list.File = filePath
	}

	createdList, err := h.ListRepository.CreateList(list)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "List has been created",
		Data:    createdList,
	})
}

// @Tags 		Lists
// @Summary		Update list
// @Router		/lists/{id} [put]
// @Accept		json
// @Produce		json
// @Param		id			path		int					true	"List id"
// @Param		list		body		listsdto.ListRequest	true	"Update list"
// @Success		200			{object}	dto.SuccessResult{data=models.List}
// @Failure		500			{object}	dto.ErrorResult
func (h *ListHandler) UpdateList(c echo.Context) error {
	request := new(listsdto.ListRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	listId, _ := strconv.Atoi(c.Param("id"))
	list, err := h.ListRepository.FindListById(listId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if request.Title != "" {
		list.Title = request.Title
	}

	if request.Description != "" {
		list.Description = request.Description
	}

	list.UpdatedAt = time.Now()

	file, err := c.FormFile("file")
	if err == nil {
		filePath, err := util.SaveFile(file)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		if list.File != "" {
			if err := os.Remove(list.File); err != nil {
				return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
					Code:    http.StatusInternalServerError,
					Message: "Failed to delete old file",
				})
			}
		}
		list.File = filePath
	}

	updatedList, err := h.ListRepository.UpdateList(list, listId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// Respon dengan hasil pembaruan
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code:    http.StatusOK,
		Message: "List has been updated",
		Data:    updatedList,
	})
}


// @Tags 		Lists
// @Summary		Delete list
// @Router		/lists/{id} [delete]
// @Accept		json
// @Produce		json
// @Param		id	path	int	true	"List id"
// @Success		200	{object}	dto.SuccessResult{data=models.List}
// @Failure		500	{object}	dto.ErrorResult
func (h *ListHandler) DeleteList(c echo.Context) error {
	listId, _ := strconv.Atoi(c.Param("id"))
	list, err := h.ListRepository.FindListById(listId)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	data, err := h.ListRepository.DeleteList(list)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			dto.ErrorResult{
				Code: http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "List has been deleted", Data: data})
}