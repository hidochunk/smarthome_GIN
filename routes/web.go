package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smarthome_GIN/form"
)

func WebRouter(router *gin.Engine) *gin.Engine {
	router.GET("/", func(c *gin.Context) {
		data := new(form.IndexData)
		data.Title = "首頁"
		data.Content = "我的第一個首頁"
		c.HTML(http.StatusOK, "index.html", data)
	})

	return router
}
