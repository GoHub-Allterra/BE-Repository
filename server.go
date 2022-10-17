package main

import (
	"gohub/config"
	"gohub/features/user/delivery"
	"gohub/features/user/repository"
	"gohub/features/user/services"

	"gohub/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)
	uRepo := repository.New(db)
	uService := services.New(uRepo)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	delivery.New(e, uService)
	
	e.Logger.Fatal(e.Start(":8001"))
}
