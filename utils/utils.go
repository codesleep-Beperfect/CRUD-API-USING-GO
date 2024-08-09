package utils

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const dsn = "root:Sskdhsu@123@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"

func Connect(){
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal(err.Error())
		panic("Not Connected database")
	}
	fmt.Println("Database Connected Successfully")
	db = d
}

func Migrate() *gorm.DB{
	return db
}