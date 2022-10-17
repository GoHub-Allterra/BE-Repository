package delivery

import (
	"net/http"
	"gohub/features/user/domain"
	"strconv"

	// "strconv"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}
	e.GET("/user", handler.ShowAllUser()) // GET ALL USER
	e.POST("/user", handler.AddUser())    // ADD USER
	e.GET("/user/:id", handler.GetUser()) // GET USER BY ID
	e.DELETE("/user", handler.DeleteUser()) // DELETE USER BY ID
	e.PUT("/user", handler.UpdateUser()) // UPDATE USER BY ID
}

func (us *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		id := c.QueryParam("id")
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		cnvID, err:= strconv.Atoi(id)
		fixID := uint(cnvID)
		cnv := ToDomain(input)
		res, err := us.srv.UpdateUser(fixID, cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("update user failed"))
		}
		return c.JSON(http.StatusOK, SuccessResponse("update user successful", res))
	}
}

func (us *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.QueryParam("id")
		cnvId, _ := strconv.Atoi(id)
		toUint := uint(cnvId)
		_, err := us.srv.DeleteUser(toUint)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("delete user failed"))
		}
		return c.JSON(http.StatusOK, FailResponse("delete user successful"))
	}
}

func (us *userHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnvId, err := strconv.Atoi(id)
		toUint := uint(cnvId)
		res, err := us.srv.Get(toUint)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("user not found"))
		}
		return c.JSON(http.StatusOK, SuccessResponse("get data berhasil", ToResponse(res, "reg")))
	}
}

func (us *userHandler) AddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := us.srv.AddUser(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("berhasil register", ToResponse(res, "reg")))
	}

}

func (us *userHandler) ShowAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := us.srv.ShowAllUser()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get all user", ToResponse(res, "all")))
	}
}
