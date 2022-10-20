package main

import (
	"gohub/config"
	"gohub/features/user/delivery"
	"gohub/features/user/repository"
	"gohub/features/user/services"

	postDel "gohub/features/post/delivery"
	postRepo "gohub/features/post/repository"
	postServ "gohub/features/post/services"

	cData "gohub/features/comments/delivery"
	cRepo "gohub/features/comments/repository"
	cServ "gohub/features/comments/services"

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

	cRepo := cRepo.New(db)
	comS := cServ.New(cRepo)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	cData.New(e, comS)

	postDel.New(e, postS)
	delivery.New(e, uService)

	e.Logger.Fatal(e.Start(":8080"))
}
