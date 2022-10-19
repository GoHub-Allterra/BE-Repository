package helper

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func autoGenerate(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return autoGenerate(length, charset)
}


func UploadProfile(c echo.Context) (string, error) {

	file, fileheader, err := c.Request().FormFile("images")
	if err != nil {
		log.Print(err)
		return "", err
	}

	randomStr := String(20)

	s3Config := &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIATMRW76KP55SBRTVX", "RpiEcfYT9wiyTXsXT63F01ldTMcGAagmVRf2Z4ot", ""),
	}
	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)

	input := &s3manager.UploadInput{
		Bucket:      aws.String("gohubalta"),   // bucket's name
		Key:         aws.String("profile/"+randomStr+"-"+fileheader.Filename), // files destination location
		Body:        file,      // content of the file
		ContentType: aws.String("image/jpg"),   // content type
	}
	res, err := uploader.UploadWithContext(context.Background(), input)
	return res.Location, err
}