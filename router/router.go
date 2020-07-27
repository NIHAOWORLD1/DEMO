package router

import (
	"github.com/gin-gonic/gin"
	"test/handler"
)

func Init() *gin.Engine {
	router :=gin.Default()
	router.GET("/order/getorder",handler.GetOrder)
    router.POST("/order/Createorder",handler.CreateOrder)
	router.GET("/orders/Searchorderlist",handler.SearchOrderList)
	router.POST("/order/updateorder",handler.UpdateOrder)
	router.POST("/file/upload",handler.Upload)
	router.POST("/file/download",handler.Download)
	return router
}