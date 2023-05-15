package httpgin

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomMW(c *gin.Context) {
	t := time.Now().UTC()

	c.Next()

	latency := time.Since(t)
	status := c.Writer.Status()

	log.Println("latency", latency, "method", c.Request.Method, "path", c.Request.URL.Path, "status", status)
}
