package router

import (
	"github.com/gin-gonic/gin"
	"test/handler"

)

func Init() *gin.Engine {
	router :=gin.Default()
	router.GET("/order/select",handler.Select)
    router.POST("/order/store",handler.Store)
	router.GET("/orders/selectlist",handler.Selectlist)
	router.POST("/order/update",handler.Update)
	return router
}