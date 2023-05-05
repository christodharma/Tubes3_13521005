package model

import(
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase(){
	var database, err = gorm.Open(postgres.Open("postgresql://QZYpViscGxsMENNoRvYLuySFehjPQxEu:vkersMizkROgEAIFgcboXvzoLVrCnEnu@db.thin.dev/f212dad5-bb33-4d21-ab00-ca1a5b3999d1"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&Chat{})
	DB = database
}