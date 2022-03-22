package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
)

func TestMw() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Test Middleware "+ time.Now().String())
		c.Next()
	}
}