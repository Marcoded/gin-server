package main

import (
	"net/http"
	"server-test/web-service-gin/controllers"
	"server-test/web-service-gin/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
}
func main() {

	router := gin.Default()
	router.GET("/", gretting)
	router.POST("/posts", controllers.PostCreate)
	router.GET("/posts", controllers.Index)
	router.PUT("/posts/:id", controllers.PostUpdate)
	router.GET("/posts/:id", controllers.PostShow)
	router.DELETE("/posts/:id", controllers.PostDelete)
	router.Run()

}
func gretting(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "The cake is a lie"})

}
