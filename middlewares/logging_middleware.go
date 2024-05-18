// middlewares/logging_middleware.go
package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)

		log.Printf("[%s] %s %s %d %s",
			c.Request.Method,
			c.Request.RequestURI,
			c.ClientIP(),
			c.Writer.Status(),
			duration,
		)
	}
}
