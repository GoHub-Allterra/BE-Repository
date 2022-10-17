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
	e.POST("/user", handler.AddUser())    // REGISTER USER
	e.GET("/user/:id", handler.GetUser()) // GET USER BY ID
	e.DELETE("/user", handler.DeleteUser()) // DELETE USER BY ID
	e.PUT("/user/update/:id", handler.UpdateUser()) // UPDATE USER BY ID
	// e.POST("/login", handler.Login()) // LOGIN USER
}

func (us *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		id := c.Param("id")
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		cnvID, err:= strconv.Atoi(id)
		fixID := uint(cnvID)
		input.ID = fixID
		cnv := ToDomain(input)
		_, err = us.srv.UpdateUser(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("update user failed"))
		}
		return c.JSON(http.StatusOK, SuccessResponseNoData("update user successful"))
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
		return c.JSON(http.StatusOK, SuccessResponseNoData("delete user successful"))
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
		return c.JSON(http.StatusOK, SuccessResponseWithData("get data berhasil", ToResponse(res, "get")))
	}
}

func (us *userHandler) AddUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		_, err := us.srv.AddUser(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponseNoData("berhasil registtrasi"))
	}

}
