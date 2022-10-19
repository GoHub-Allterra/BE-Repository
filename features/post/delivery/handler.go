package delivery

import (
	"gohub/config"
	"gohub/features/post/domain"
	"gohub/middlewares"
	"gohub/utils/helper"
	"log"
	"net/http"
	"strconv"

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
	e.GET("/posts/comments/:id", handler.SelectId())
	e.GET("/myposts", handler.GetAllMyPosts(), middleware.JWT([]byte(config.JWT_SECRET)))
	e.PUT("/posts/:id", handler.PutId(), middleware.JWT([]byte(config.JWT_SECRET)))
	e.DELETE("/posts/:id", handler.DeletePosts(), middleware.JWT([]byte(config.JWT_SECRET)))

}

func (ph *postHandler) DeletePosts() echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		idProd, _ := strconv.Atoi(id)
		idFromToken := middlewares.ExtractToken(c)
		row, errDel := ph.PostUsecase.DeletedPost(idProd, idFromToken)
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

func (ph *postHandler) PutId() echo.HandlerFunc {
	return func(c echo.Context) error {
		idToken := middlewares.ExtractToken(c)
		id := c.Param("id")
		idConv, _ := strconv.Atoi(id)
		if idConv < 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "param must be number",
			})
		}

		var update Request
		err := c.Bind(&update)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "failed bind data",
			})
		}

		file, err := c.FormFile("images")
		if file != nil {
			res, err := helper.UploadPosts(c)
			if err != nil {
				return err
			}
			log.Print(res)
			update.Images = res
		}
		if err != nil {
			return err
		}

		var insert domain.Post
		if update.Caption != "" {
			insert.Caption = update.Caption
		}
		if update.Images != "" {
			insert.Images = update.Images
		}
		insert.ID = uint(idConv)
		row, _ := ph.PostUsecase.UpdatePost(idConv, idToken, insert)
		if row == 1 {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success update data",
			})
		} else {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "failed update data",
			})
		}
	}
}

func (ph *postHandler) SelectId() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idConv, _ := strconv.Atoi(id)
		if idConv < 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "param must be number",
			})
		}
		res, resErr := ph.PostUsecase.SelectById(idConv)
		if resErr != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "failed get by id",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "succes get by id",
			"data":    res,
		})
	}
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
			"message": "success get all posts",
			"data":    res,
		})

	}
}

func (ph *postHandler) GetAllMyPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := middlewares.ExtractToken(c)
		respost, err := ph.PostUsecase.GetMyPosts(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "error server",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all my posts",
			"posts":   respost,
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

		file, err := c.FormFile("images")
		if file != nil {
			res, err := helper.UploadPosts(c)
			if err != nil {
				return err
			}
			log.Print(res)
			input.Images = res
		}
		if err != nil {
			return err
		}

		cnv := ToDomain(input)
		_, errposts := ph.PostUsecase.AddPost(cnv, id)
		if errposts != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "failed",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success created posts",
		})
	}
}
