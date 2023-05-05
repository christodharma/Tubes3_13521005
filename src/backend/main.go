package main

import (
	"backend/controllers/chatcontroller"
	"backend/model"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	model.ConnectDatabase()
	
	r.GET("/chats", chatcontroller.Index)
	r.GET("/chat/:id", chatcontroller.Show)
	r.POST("/chat", chatcontroller.Create)
	r.PUT("/chat/:id", chatcontroller.Update)
	r.DELETE("/chat", chatcontroller.Delete)
	r.Run()
}