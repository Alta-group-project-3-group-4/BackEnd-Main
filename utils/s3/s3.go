package helper

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// CREATE RANDOM STRING
var theSession *session.Session

func GetSession() *session.Session {
	if theSession == nil {
		theSession = initSession()
	}
	return theSession
}
func initSession() *session.Session {
	newSession := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY_IAM"), os.Getenv("SECRET_KEY_IAM"), ""),

		// Region:      aws.String(config.AWS_REGION),
		// Credentials: credentials.NewStaticCredentials(config.ACCESS_KEY_ID, config.ACCESS_KEY_SECRET, ""),
	}))
	return newSession
}

type UploadResult struct {
	Path string `json:"path" xml:"path"`
}

func GetUrlImagesFromAWS(fileData multipart.FileHeader) (string, error) {

	if fileData.Filename != "" && fileData.Size != 0 {
		if fileData.Size > 500000 {
			return "", errors.New("file size max 500kb")
		}
		file, err := fileData.Open()
		if err != nil {
			return "", errors.New("error open fileData")
		}
		tipeNameFile, err := TypeFile(file)
		if err != nil {
			return "", errors.New("file type error only jpg or png file can be upload")
		}
		defer file.Close()
		log.Println("size:", fileData.Filename, file)
		namaFile := GenerateRandomString()
		namaFile = namaFile + tipeNameFile
		fileData.Filename = namaFile
		log.Println(namaFile)
		file2, _ := fileData.Open()
		defer file2.Close()
		uploadURL, err := UploadToS3(fileData.Filename, file2)
		if err != nil {
			return "", errors.New("cannot upload to s3 server error")
		}
		return uploadURL, nil
	}
	return "", nil
}
func UploadToS3(fileName string, src multipart.File) (string, error) {
	// The session the S3 Uploader will use
	sess := GetSession()
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:         aws.String(fileName),
		Body:        src,
		ContentType: aws.String("image/png"),
	})
	// content type penting agar saat link dibuka file tidak langsung auto download
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}
	return result.Location, nil
}

func GenerateRandomString() string {
	rand.Seed(time.Now().Unix())
	str := "AsDfzGhBvCX123456MnBp"
	shuff := []rune(str)

	rand.Shuffle(len(shuff), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})

	return string(shuff)
}
func TypeFile(test multipart.File) (string, error) {
	fileByte, _ := io.ReadAll(test)
	fileType := http.DetectContentType(fileByte)
	TipenamaFile := ""
	if fileType == "image/png" {
		TipenamaFile = ".png"
	} else {
		TipenamaFile = ".jpg"
	}
	if fileType == "image/png" || fileType == "image/jpeg" || fileType == "image/jpg" {
		return TipenamaFile, nil
	}
	return "", errors.New("file type not match")
}

// const charset = "abcdefghijklmnopqrstuvwxyz" +
// 	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// var seededRand *rand.Rand = rand.New(
// 	rand.NewSource(time.Now().UnixNano()))

// func autoGenerate(length int, charset string) string {
// 	b := make([]byte, length)
// 	for i := range b {
// 		b[i] = charset[seededRand.Intn(len(charset))]
// 	}
// 	return string(b)
// }

// func String(length int) string {
// 	return autoGenerate(length, charset)
// }

// // UPLOAD FOTO PROFILE TO AWS S3

// func UploadProfile(c echo.Context) (string, error) {

// 	file, fileheader, err := c.Request().FormFile("image_url")
// 	if err != nil {
// 		log.Print(err)
// 		return "", err
// 	}

// 	randomStr := String(20)

// 	godotenv.Load("local.env")

// 	s3Config := &aws.Config{
// 		Region:      aws.String(os.Getenv("AWS_REGION")),
// 		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY_IAM"), os.Getenv("SECRET_KEY_IAM"), ""),
// 	}
// 	s3Session := session.New(s3Config)

// 	uploader := s3manager.NewUploader(s3Session)

// 	input := &s3manager.UploadInput{
// 		Bucket:      aws.String(os.Getenv("AWS_BUCKET_NAME")),                        // bucket's name
// 		Key:         aws.String("homestay/" + randomStr + "-" + fileheader.Filename), // files destination location
// 		Body:        file,                                                            // content of the file
// 		ContentType: aws.String("image/jpg"),                                         // content type
// 	}
// 	res, err := uploader.UploadWithContext(context.Background(), input)

// 	// RETURN URL LOCATION IN AWS
// 	return res.Location, err
// }

// func UploadProfileUser(c echo.Context) (string, error) {

// 	file, fileheader, err := c.Request().FormFile("file")
// 	if err != nil {
// 		log.Print(err)
// 		return "", err
// 	}

// 	randomStr := String(20)

// 	godotenv.Load(".env")

// 	s3Config := &aws.Config{
// 		Region:      aws.String(os.Getenv("AWS_REGION")),
// 		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY_IAM"), os.Getenv("SECRET_KEY_IAM"), ""),
// 	}
// 	s3Session := session.New(s3Config)

// 	uploader := s3manager.NewUploader(s3Session)

// 	input := &s3manager.UploadInput{
// 		Bucket:      aws.String(os.Getenv("AWS_BUCKET_NAME")),                        // bucket's name
// 		Key:         aws.String("homestay/" + randomStr + "-" + fileheader.Filename), // files destination location
// 		Body:        file,                                                            // content of the file
// 		ContentType: aws.String("image/jpg"),                                         // content type
// 	}
// 	res, err := uploader.UploadWithContext(context.Background(), input)

// 	// RETURN URL LOCATION IN AWS
// 	return res.Location, err
// }
