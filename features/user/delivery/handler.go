package delivery

import (
	"bytes"
	"gohub/config"
	"gohub/features/user/domain"
	"gohub/middlewares"
	"net/http"
	"os"

	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	srv domain.Service
}



func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}
	e.Static("/users/update", "./public")
	e.POST("/login", handler.Login())                                                       // LOGIN USER
	e.POST("/register", handler.AddUser())                                                  // REGISTER USER
	e.GET("/users/:id", handler.GetUser(), middleware.JWT([]byte(config.JWT_SECRET)))       // GET USER BY ID
	e.DELETE("/users", handler.DeleteUser(), middleware.JWT([]byte(config.JWT_SECRET)))     // DELETE USER BY ID
	e.PUT("/users/update", handler.UpdateUser(), middleware.JWT([]byte(config.JWT_SECRET))) // UPDATE USER BY ID                                  // ADD PROFILE PHOTOS
}

const (
    AWS_S3_REGION = "ap-southeast-1"
    AWS_S3_BUCKET = "gohubalta"
)

func uploadFile(session *session.Session, uploadFileDir string) error {
	var c echo.Context
	file, err := c.FormFile("images")
    upFile, err := os.Open(file.Filename)
    if err != nil {
        return err
    }
    defer upFile.Close()
    
    upFileInfo, _ := upFile.Stat()
    var fileSize int64 = upFileInfo.Size()
    fileBuffer := make([]byte, fileSize)
    upFile.Read(fileBuffer)
    
    _, err = s3.New(session).PutObject(&s3.PutObjectInput{
        Bucket:               aws.String(AWS_S3_BUCKET),
        Key:                  aws.String(uploadFileDir),
        ACL:                  aws.String("private"),
        Body:                 bytes.NewReader(fileBuffer),
        ContentLength:        aws.Int64(fileSize),
        ContentType:          aws.String(http.DetectContentType(fileBuffer)),
        ContentDisposition:   aws.String("attachment"),
        ServerSideEncryption: aws.String("AES256"),
    })
    return err
}

func (us *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind data"))
		}

		cnv := ToDomain(input)
		res, token, err := us.srv.Login(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("login failed"))
		}
		res.Token = token
		return c.JSON(http.StatusOK, SuccessResponseWithData("login successful", ToResponse(res, "login")))
	}
}


func (us *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		file, err := c.FormFile("images")
		if file != nil {

			// session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
			// if err != nil {
			// 	log.Fatal(err)
			// }
		
			// // Upload Files
			// err = uploadFile(session, file.Filename)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			input.Images = file.Filename
		}

		id := middlewares.ExtractToken(c)
		input.ID = uint(id)
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
		id := middlewares.ExtractToken(c)
		toUint := uint(id)
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

		return c.JSON(http.StatusCreated, SuccessResponseNoData("berhasil registrasi"))
	}

}
