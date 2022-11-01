package main

import (
	"github.com/gin-gonic/gin"
	"smarthome_GIN/routes"
)

func main() {
	server := gin.Default()
	//open it if need to use function in .tmpl, and see how to use FuncMap
	//routes.LoadFunctionMap(server)
	server.LoadHTMLGlob("resources/**/*")
	routes.WebRouter(server)

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
