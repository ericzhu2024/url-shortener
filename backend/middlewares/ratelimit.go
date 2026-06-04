package middlewares

import (
	"net/http"
	"url-shortener/cache"

	"github.com/gin-gonic/gin"
)

func RateLimit(c *gin.Context) {
	ip := c.ClientIP()

	allowed, err := cache.RateLimit(ip, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		c.Abort()
		return
	}

	if !allowed {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error": "too many requests, max 10 links per minute",
		})
		c.Abort()
		return
	}

	c.Next()
}