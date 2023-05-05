package model

import(
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase(){
	var database, err = gorm.Open(mysql.Open("root:admin@tcp(localhost:3306)/ChatTMBOK"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&Chat{})
	DB = database
}