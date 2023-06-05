package main

import (
	"gin-crud/controlers"
	"gin-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controlers.PostsCreate)
	r.GET("/posts", controlers.PostsIndex)

	r.Run()
}
