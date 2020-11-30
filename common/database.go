package common

import (
	"fmt"
	"gindemo/model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"net/url"
)

var DB * gorm.DB


func InitDB() *gorm.DB{
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))

	db,err := gorm.Open(driverName,args)

	if err != nil{
		panic("failed to connect databases! err is : " + err.Error() )

	}

	DB = db

	db.AutoMigrate(&model.User{})
	return db
}

func GetDB() *gorm.DB  {
	return DB
}