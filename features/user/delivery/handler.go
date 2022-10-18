package delivery

import (
	"bytes"
	"encoding/base64"
	"gohub/config"
	"gohub/features/user/domain"
	"gohub/middlewares"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
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

func upload(c echo.Context) error {

	img, err := c.FormFile("images")

	aws_access_key_id := "AKIATMRW76KP65TYH5H2" 
	aws_secret_access_key := "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" 
	token := "" 
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)   
	_, err =  creds.Get() 
	if err != nil { 
		return err
	}   
	cfg := aws.NewConfig().WithRegion("ap-southeast-1").WithCredentials(creds)   
	svc := s3.New(session.New(), cfg) 
	
	file, err := os.Open(img.Filename)   
	if err != nil { 
		return err
	}   
	defer file.Close()   
	fileInfo, _ := file.Stat()   
	size := fileInfo.Size()   
	buffer := make([]byte, size) // read file content to buffer 
	file.Read(buffer)   
	fileBytes := bytes.NewReader(buffer) 
	fileType := http.DetectContentType(buffer) 
	path := "/media/" + file.Name()   

	params := &s3.PutObjectInput{ 
		Bucket: aws.String("gohubalta"), 
		Key: aws.String(path), 
		Body: fileBytes, 
		ContentLength: aws.Int64(size), 
		ContentType: aws.String(fileType),
	}   
	resp, err := svc.PutObject(params) 
	if err != nil { 
		return err
	}   
	log.Printf("response %s", awsutil.StringValue(resp)) 
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
			// Get name
			name := c.FormValue("name")
			// Get avatar
			avatar, err := c.FormFile("images")
			if err != nil {
				return err
			}
		
			// Source
			src, err := avatar.Open()
			if err != nil {
				return err
			}
			defer src.Close()
		

			position := strings.Index(avatar.Filename, ",")
			if position == -1 {
				// image format doesn't match with 'data:image/png;base64,iVBO...'
			}
			// decode the base64 string, removing 'data:image/png;base64,' from it
			reader := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(avatar.Filename[position+1:]))

			data, err := ioutil.ReadAll(reader)
			if err != nil {
				// error handler
			}
			// you write the file in wherever you want
			ioutil.WriteFile("public/image.png", data, 0644)

			return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
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

		return c.JSON(http.StatusCreated, SuccessResponseNoData("berhasil registtrasi"))
	}

}
