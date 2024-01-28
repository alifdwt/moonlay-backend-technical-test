package routes

import (
	"backend-technical-test/handlers"
	"backend-technical-test/pkg/postgres"
	"backend-technical-test/repository"

	"github.com/labstack/echo/v4"
)

func ListRoutes(e *echo.Group) {
	listRepository := repository.NewListRepository(postgres.DB)
	listHandler := handlers.NewListHandler(listRepository)

	e.GET("/lists", listHandler.FindLists)
	e.GET("/lists/:id", listHandler.FindListById)
	e.POST("/lists", listHandler.CreateList)
	e.PUT("/lists/:id", listHandler.UpdateList)
	e.DELETE("/lists/:id", listHandler.DeleteList)
}