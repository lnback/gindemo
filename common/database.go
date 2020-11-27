package common

import (
	"fmt"
	"gindemo/model"
	"github.com/jinzhu/gorm"
)

var DB * gorm.DB


func InitDB() *gorm.DB{
	driverName := "mysql"
	host := "192.168.33.30"
	port := "3306"
	database := "test"
	username := "root"
	password := "123456"
	charset := "utf8"

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

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