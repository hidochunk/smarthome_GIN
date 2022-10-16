package routes

import "github.com/gin-gonic/gin"

func WebRouter() *gin.Engine {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	return router
}
