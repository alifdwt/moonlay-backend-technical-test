package main

import (
	"backend-technical-test/database"
	"backend-technical-test/pkg/postgres"
	"backend-technical-test/routes"
	"fmt"

	_ "backend-technical-test/docs"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title 				Moonlay Academy - Backend Test (GOLANG)
// @version 			1.0
// @description 		To do list backend handlers

// @contact.name 		Alif Dewantara
// @contact.url 		https://github.com/alifdwt
// @contact.email 		aputradewantara@gmail.com

// @license.name 		MIT
// @license.url 		https://opensource.org/licenses/MIT

// @host 				localhost:5000
// @BasePath 			/api/v1

func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	postgres.DatabaseInit()
	database.RunMigration()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowedHeaders: []string{"*"},
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	routes.RouteInit(e.Group("/api/v1"))
	fmt.Println("server running localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}