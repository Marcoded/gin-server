package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MiddleLogger(c *gin.Context) {
	fmt.Println("I run only on groups!")
	c.Next()
}
