package delivery

import (
	"gohub/features/comments/domain"
	"gohub/middlewares"
	"net/http"
	"strconv"

	// _ "github.com/labstack/echo"
	// _ "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

type postHandler struct {
	PostUsecase domain.ServiceInterface
}

func New(e *echo.Echo, ps domain.ServiceInterface) {
	handler := &postHandler{
		PostUsecase: ps,
	}

	e.POST("/comments/:id", handler.AddComment(), middlewares.JWTMiddleware())
	e.DELETE("/comments/:id", handler.DeleteCom(), middlewares.JWTMiddleware())

}

func (ch *postHandler) DeleteCom() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idProd, _ := strconv.Atoi(id)
		idFromToken := middlewares.ExtractToken(c)
		row, errDel := ch.PostUsecase.DeleteId(idProd, idFromToken)
		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "failed server error",
			})
		}
		if row != 1 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "failed delete",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete posts",
		})
	}
}
func (ch *postHandler) AddComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		var reqComment Request
		if err := c.Bind(&reqComment); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "failed bind data",
			})
		}
		id, err := strconv.Atoi(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "error get id")
		}

		var commentData domain.Comments = domain.Comments{
			Post_ID: uint(id),
			User_ID: uint(middlewares.ExtractToken(c)),
			Comment: reqComment.Comment,
		}

		_, err = ch.PostUsecase.Insert(commentData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success add comment",
		})

	}
}
