package controllers

import (
	"net/http"
	"server-test/web-service-gin/middlewares"

	"github.com/gin-gonic/gin"
)

type ArticlesController struct {
}

var articlesPath = "/articles"

func InitializeArticlesController(router *gin.Engine) *ArticlesController {
	ac := &ArticlesController{}
	articlesGroup := router.Group(articlesPath)
	ac.mountMiddleWares(articlesGroup)
	ac.mountRoutes(articlesGroup)
	return ac
}

func (ac *ArticlesController) helloFromArticles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello from articles"})
}

func (ac *ArticlesController) anotherHelloFromArticles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Another hello from articles"})
}

func (ac *ArticlesController) mountRoutes(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/", ac.helloFromArticles)
	routeGroup.GET("/hello", ac.anotherHelloFromArticles)
}

func (ac *ArticlesController) mountMiddleWares(routeGroup *gin.RouterGroup) {
	routeGroup.Use(middlewares.MiddleLogger)
}

// create a new instance of CDSController
