package main

import (
	"os"
	"url-shortener/cache"
	"url-shortener/db"
	"url-shortener/handlers"
	"url-shortener/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.Init()
	cache.Init()

	r := gin.Default()

	// CORS must be registered before routes
    r.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204)
	})
	// apply rate limiting only to the create endpoint
	r.POST("/api/links", middlewares.RateLimit, handlers.CreateLink)

	r.GET("/r/:code", handlers.RedirectLink)
	r.GET("/api/links/:code/stats", handlers.GetStats)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}