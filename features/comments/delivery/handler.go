package delivery

import (
	"errors"
	"gohub/features/comments/domain"
	"gohub/middlewares"
	"log"
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
		var comment CommentFormat
		id_post := c.Param("id")


		if err := c.Bind(&comment); err != nil {
			c.JSON(http.StatusBadRequest, errors.New("cannot bind data"))
		}

		idUser := uint(middlewares.ExtractToken(c))
		idCnv, _ := strconv.Atoi(id_post)
		idPost := uint(idCnv)
		comment.IdPost = idPost
		comment.IdUser = idUser
		data := ToDomain(comment)
		log.Print(data)
		_, err1 := ch.PostUsecase.Insert(data)
		if err1 != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("error from server"))
		}
		return c.JSON(http.StatusCreated, map[string]string{"msg": "insert comment success"})
	}
}
