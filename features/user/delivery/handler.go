package delivery

import (
	"bytes"
	"gohub/config"
	"gohub/features/user/domain"
	"gohub/middlewares"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}
	e.POST("/login", handler.Login())                                                       // LOGIN USER
	e.POST("/register", handler.AddUser())                                                  // REGISTER USER
	e.GET("/users/:id", handler.GetUser(), middleware.JWT([]byte(config.JWT_SECRET)))       // GET USER BY ID
	e.DELETE("/users", handler.DeleteUser(), middleware.JWT([]byte(config.JWT_SECRET)))     // DELETE USER BY ID
	e.PUT("/users/update", handler.UpdateUser(), middleware.JWT([]byte(config.JWT_SECRET))) // UPDATE USER BY ID                                  // ADD PROFILE PHOTOS
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

const (
	AWS_S3_REGION = "ap-southeast-1"
	AWS_S3_BUCKET = "gohub-alterra"
)

const (
	S3_REGION = "ap-southeast-1"
	S3_BUCKET = "gohub-alterra"
)

func uploadFileToS3(fileName string) error {

	s, err := session.NewSession(&aws.Config{
		Region: aws.String(S3_REGION),
		Credentials: credentials.NewStaticCredentials(
			"AKIATMRW76KP65TYH5H2",
			"zskfWYUM6IwoOU6+2L9Z9Iv3SgUDo7YTKIboFaCK",
			""),
	})

	if err != nil {
		log.Fatal(err, "ERROR SESSION")
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err, "ERROR OPEN FILE")
		return err
	}
	defer file.Close()

	// get the file size and read
	// the file content into a buffer
	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, s3err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(S3_BUCKET),
		Key:                  aws.String(fileName),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})

	return s3err
}

func (us *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		id := middlewares.ExtractToken(c)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		image, err := c.FormFile("images")
		if image != nil {
			upload := uploadFileToS3(image.Filename)
			if upload != nil {
				log.Fatal("ERROR IMAGE")
			}
		}

		fixID := uint(id)
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

		return c.JSON(http.StatusCreated, SuccessResponseNoData("berhasil registtrasi"))
	}

}
