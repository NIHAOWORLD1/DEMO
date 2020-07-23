package router

import (
	"github.com/gin-gonic/gin"
	"test/handler"

)

func Init() *gin.Engine {
	router :=gin.Default()
	router.GET("/order/menbee/select",handler.Select)

	return router
}