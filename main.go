package main

import (
	"learn_gin/controllers"
	"learn_gin/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
  initializers.ConnectWithDb()
}

func main() {
  router := gin.Default() //creates a router
  // alternative : 	port := os.Getenv("PORT") // Reads env ( import os )
  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "hello",
    })
  })

  router.POST("/createPosts", controllers.PostsCreate)
  router.GET("/getPosts", controllers.PostIndex)
  router.GET("/posts/:id", controllers.PostById)

  router.Run() // listens on 0.0.0.0:8080 by default or we can pass any port as router.Run(":" + port)
}