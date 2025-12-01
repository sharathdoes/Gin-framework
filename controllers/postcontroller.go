package controllers

import (
	"learn_gin/initializers"
	"learn_gin/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Title string
		Body string
		Tags pq.StringArray
	}
	c.Bind(&body)
	newPost :=models.Post{Title:body.Title, Body:body.Body, Tags: body.Tags}
	result:=initializers.DB.Create(&newPost)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post":newPost,
	})
}

func PostIndex(c *gin.Context){
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"post":posts,
	})
}

func PostById(c *gin.Context){
	var post models.Post
	id:=c.Param("id")
	initializers.DB.Find(&post, id)

	c.JSON(200, gin.H{
		"post":post,
	})
}

