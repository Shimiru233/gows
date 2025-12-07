package routes

import (
	"firstgo/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 静态文件
	r.Static("/static", "./web/static")

	// 模板
	r.LoadHTMLGlob("web/templates/*")

	// 首页
	r.GET("/", handler.Home)

	return r
}
