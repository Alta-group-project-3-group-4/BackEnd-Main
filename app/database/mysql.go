package database

import (
	"airbnb/app/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(app config.AppConfig) *gorm.DB {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", app.DBUSERNAME, app.DBPASS, app.DBHOST, app.DBPORT, app.DBNAME)

	DB, err := gorm.Open(mysql.Open(URL), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return DB
}
