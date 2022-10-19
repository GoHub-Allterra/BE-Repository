package main

import (
	"gohub/config"
	"gohub/features/user/delivery"
	"gohub/features/user/repository"
	"gohub/features/user/services"

	postDel "gohub/features/post/delivery"
	postRepo "gohub/features/post/repository"
	postServ "gohub/features/post/services"

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

	postRepo := postRepo.New(db)
	postS := postServ.New(postRepo)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	postDel.New(e, postS)
	delivery.New(e, uService)

	e.Logger.Fatal(e.Start(":80"))
}
