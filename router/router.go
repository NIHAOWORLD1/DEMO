package router

import (
	"github.com/gin-gonic/gin"
	"test/handler"

)

func Init() *gin.Engine {
	router :=gin.Default()
	router.GET("/order/select",handler.Select)
    router.POST("/orders/select/",handler.Update)
	return router
}