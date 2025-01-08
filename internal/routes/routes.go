package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	r.GET("/shorten", getUrls)
	r.GET("/shorten/:shortCode/stats", getStatURL)
	r.POST("/shorten", createUrl)
	r.PUT("/shorten/:shortCode", updateURL)
	r.DELETE("/shorten/:shortCode", deleteUrl)

	r.GET("/:shortCode", redirect)
}
