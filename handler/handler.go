package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/model"
	"test/service"
)

func Select(c *gin.Context)  {
    var id model.Getid
    //// 将request的query中的数据，自动解析到结构体
	if err := c.ShouldBind(&id); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.String(http.StatusOK, "Success")
	service.Select(id.ID)

}
func Store(c *gin.Context){
	var order model.Order
	c.ShouldBind(&order)
	id,err:= service.Store(order)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "添加成功",
		"data":    id,
	})

}