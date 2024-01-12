package controllers

import (
	"server-test/web-service-gin/initializers"
	"server-test/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

// return all posts
func Index(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}


func PostShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

// Create a new post
func PostCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Body: body.Body, Title: body.Title}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post

	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	initializers.DB.Delete(&post)

	c.JSON(200, gin.H{
		"post": post,
	})
}
