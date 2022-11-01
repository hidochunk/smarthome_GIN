package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smarthome_GIN/database/model"
)

func WebRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		devices := model.GetDeviceByCommunicationType("mqtt")
		c.HTML(http.StatusOK, "index.tmpl", devices)
	})

}
