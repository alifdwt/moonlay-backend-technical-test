package routes

import (
	"backend-technical-test/handlers"
	"backend-technical-test/pkg/postgres"
	"backend-technical-test/repository"

	"github.com/labstack/echo/v4"
)

func SublistRoutes(e *echo.Group) {
	sublistRepository := repository.NewSublistRepository(postgres.DB)
	sublistHandler := handlers.NewSublistHandler(sublistRepository)

	e.GET("/sublists", sublistHandler.FindSublists)
	e.GET("/sublists/:id", sublistHandler.FindSublistById)
	e.GET("/sublists/list/:listId", sublistHandler.FindSublistByListId)
	e.POST("/sublists", sublistHandler.CreateSublist)
	e.PUT("/sublists/:id", sublistHandler.UpdateSublist)
	e.DELETE("/sublists/:id", sublistHandler.DeleteSublist)
}