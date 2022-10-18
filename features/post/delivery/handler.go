package delivery

import (
	"gohub/config"
	"gohub/features/post/domain"
	"gohub/middlewares"
	"net/http"

	// "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/labstack/echo"
)

type postHandler struct {
	PostUsecase domain.PostUsecase
}

func New(e *echo.Echo, ps domain.PostUsecase) {
	handler := &postHandler{
		PostUsecase: ps,
	}

	e.POST("/myposts", handler.AddPosting(), middleware.JWT([]byte(config.JWT_SECRET)))
	e.GET("/posts", handler.SelectAll())

}

func (ph *postHandler) SelectAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ph.PostUsecase.GetAllPosts()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "failed get all data",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all data",
			"data":    res,
		})

	}
}

func (ph *postHandler) AddPosting() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input Request
		id := middlewares.ExtractToken(c)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "failed bind data",
			})
		}

		cnv := ToDomain(input)
		_, err := ph.PostUsecase.AddPost(cnv, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "failed",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
		})
	}
}
