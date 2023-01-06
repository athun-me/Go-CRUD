package controllers

import (
	"github.com/athun/config"
	"github.com/athun/models"
	"github.com/gin-gonic/gin"
)

func CreatUser(c *gin.Context) {
	var Data struct {
		Name  string
		Place string
		Age   int64
	}

	c.Bind(&Data)

	user := models.User{
		Name:  Data.Name,
		Place: Data.Place,
		Age:   Data.Age,
	}
	db := config.ConnectDB()
	result := db.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func ReadAllUser(c *gin.Context) {
	var user []models.User
	db := config.ConnectDB()
	result := db.Find(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func ReadUser(c *gin.Context) {
	//Get id from URL
	id := c.Param("id")

	var user []models.User
	db := config.ConnectDB()
	result := db.First(&user, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
