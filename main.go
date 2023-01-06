package main

import (
	"github.com/athun/config"
	"github.com/athun/controllers"
	"github.com/athun/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVeriable()
	config.ConnectDB()
}

func main() {
	r := gin.Default()
	r.POST("/", controllers.CreatUser)
	r.GET("/", controllers.ReadAllUser)
	r.GET("/find/:id", controllers.ReadUser)
	r.Run()
}
