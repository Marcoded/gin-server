package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	MountRoutes(routeGroup *gin.RouterGroup)
}

