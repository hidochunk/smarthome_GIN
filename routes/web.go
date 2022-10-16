package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hidochunk/smarthome_GIN/resources/model"
	"net/http"
)

func WebRouter() *gin.Engine {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		data := new(model.IndexData)
		data.Title = "首頁"
		data.Content = "我的第一個首頁"
		c.HTML(http.StatusOK, "index.html", data)
	})

	return router
}
