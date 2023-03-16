package config

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/viper"
)

var (
	JWT_KEY           string = ""
	KEYID             string = ""
	ACCESSKEY         string = ""
	MIDTRANSSERVERKEY string = ""
)

func InitEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("SECRETKEY"); found {
		app.JWTKEY = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBUSERNAME"); found {
		app.DBUSERNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("KEYID"); found {
		app.keyid = val
		isRead = false
		KEYID = val
	}
	if val, found := os.LookupEnv("ACCESSKEY"); found {
		app.accesskey = val
		isRead = false
		ACCESSKEY = val
	}

	if val, found := os.LookupEnv("MIDTRANSSERVERKEY"); found {
		app.MIDTRANSSERVERKEY = val
		isRead = false
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPASS = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHOST = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		app.DBPORT = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBNAME = val
		isRead = false
	}

	err2 := viper.Unmarshal(&app)
	if err2 != nil {
		log.Println("error parse config : ", err2.Error())
		return nil
	}

	if isRead {
		viper.AddConfigPath("./")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}

		app.JWTKEY = viper.Get("SECRETKEY").(string)
		app.DBUSERNAME = viper.Get("DBUSERNAME").(string)
		app.DBPASS = viper.Get("DBPASS").(string)
		app.DBHOST = viper.Get("DBHOST").(string)
		app.DBPORT, _ = viper.Get("DBPORT").(string)
		app.DBNAME = viper.Get("DBNAME").(string)

		app.keyid = os.Getenv("KEYID")
		app.accesskey = os.Getenv("ACCESSKEY")
	}

	return &app
}

func S3Config() *session.Session {
	s3Config := &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(KEYID, ACCESSKEY, ""),
	}
	s3Session, _ := session.NewSession(s3Config)
	return s3Session
}
