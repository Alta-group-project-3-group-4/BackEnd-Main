package helper

import (
	"airbnb/app/config"
	"errors"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Uploader interface {
	UploadImageS3(file multipart.FileHeader, homestayId string) (string, error)
}

var ObjectURL string = "https://try123ok.s3.ap-southeast-1.amazonaws.com/"

func UploadImageS3(file multipart.FileHeader, homestayId string) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	// ext := filepath.Ext(file.Filename)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("try123ok"),
		Key:    aws.String("files/homestay/" + homestayId + "/" + file.Filename),
		Body:   src,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", errors.New("problem with upload image")
	}
	path := ObjectURL + "files/homestay/" + homestayId + "/" + file.Filename
	return path, nil
}
