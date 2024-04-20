package utils

import (
	"my-app/constants"

	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var uploader *s3manager.Uploader

func init() {
	awsSession, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(constants.Region),
			Credentials: credentials.NewStaticCredentials(
				constants.AccessKey,
				constants.SecretKey,
				"",
			),
		},
	})

	if err != nil {
		panic(err)
	}

	uploader = s3manager.NewUploader(awsSession)
}

func saveFile(fileReader io.Reader, fileHeader *multipart.FileHeader) (string, error) {
	// Upload the file to S3 using the fileReader
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(constants.BucketName),
		Key:    aws.String(fileHeader.Filename),
		Body:   fileReader,
	})
	if err != nil {
		return "", err
	}

	// Get the URL of the uploaded file
	url := fmt.Sprintf("https://%s.s3.us-east-2.amazonaws.com/%s", constants.BucketName, fileHeader.Filename)

	return url, nil
}

func UploadFile(c *gin.Context) string {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err.Error()
	}

	var errors []string
	var uploadedURLs string

	files := form.File["files"]

	for _, file := range files {
		fileHeader := file

		f, err := fileHeader.Open()
		if err != nil {
			errors = append(errors, fmt.Sprintf("Error opening file %s: %s", fileHeader.Filename, err.Error()))
			continue
		}
		defer f.Close()

		uploadedURL, err := saveFile(f, fileHeader)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Error saving file %s: %s", fileHeader.Filename, err.Error()))
		} else {
			return uploadedURL
		}
	}
	if len(errors) > 0 {
		return ""
	}
	return uploadedURLs
}

func DeleteFile(fileName string) error {
	svc := s3.New(session.New(&aws.Config{
		Region: aws.String(constants.Region),
		Credentials: credentials.NewStaticCredentials(
			constants.AccessKey,
			constants.SecretKey,
			"",
		),
	}))

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(constants.BucketName),
		Key:    aws.String(fileName),
	})

	return err
}

func ChangeFileUpload(c *gin.Context, nameFileDelete string) string {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err.Error()
	}

	var errors []string
	var uploadedURLs string

	files := form.File["files"]

	for _, file := range files {
		fileHeader := file

		f, err := fileHeader.Open()
		if err != nil {
			errors = append(errors, fmt.Sprintf("Error opening file %s: %s", fileHeader.Filename, err.Error()))
			continue
		}
		defer f.Close()

		// Delete the existing file
		err = DeleteFile(path.Base(nameFileDelete))
		if err != nil {
			errors = append(errors, fmt.Sprintf("Error deleting file %s: %s", path.Base(nameFileDelete), err.Error()))
		}

		uploadedURL, err := saveFile(f, fileHeader)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Error saving file %s: %s", fileHeader.Filename, err.Error()))
		} else {
			return uploadedURL
		}
	}
	if len(errors) > 0 {
		return ""
	}
	return uploadedURLs
}
