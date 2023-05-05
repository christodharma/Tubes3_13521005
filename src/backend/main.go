package main

import (
	"backend/controllers/chatcontroller"
	"backend/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main()  {
	r := gin.Default()
	model.ConnectDatabase()
	// r.GET("/", _) should bring to landing page/chat page
	r.GET("/chats", chatcontroller.Index)
	r.GET("/chat/:id", chatcontroller.Show)
	r.GET("/chat", chatcontroller.TextInput)
	r.POST("/chat", chatcontroller.Create)
	r.PUT("/chat/:id", chatcontroller.Update)
	r.DELETE("/chat", chatcontroller.Delete)
	r.Use(cors.New(cors.Config{ 
		AllowOrigins: []string{"http://localhost:3000"}, 
		AllowMethods: []string{"GET", "POST", "DELETE"}, 
		AllowHeaders: []string{"Origin", "Content-Type"}, 
}))
	r.Run()
}