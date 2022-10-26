package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"smarthome_GIN/database/model"
	"smarthome_GIN/routes"
)

func main() {
	test := model.GetDeviceByDeviceType("lamp")
	fmt.Println(test[0])

	server := gin.Default()
	server.LoadHTMLGlob("resources/views/*")
	routes.WebRouter(server)

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
