package routes

import (
	"github.com/gin-gonic/gin"
	"text/template"
)

func LoadFunctionMap(router *gin.Engine) {
	router.SetFuncMap(template.FuncMap{})

}
