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
	result,err2 :=service.Select(id.ID)
	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "没有这个用户",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "查询成功",
		"data":    result,
	})

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
func Selectlist(c *gin.Context){
	var condition model.GetorderList
	c.ShouldBind(&condition)
	result, err :=service.Selectlist(condition)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data":   result,
	})
}

func Update(c *gin.Context){
	var renew model.Order
	c.ShouldBind(&renew)
    err :=service.UpdateOrder(renew)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": "修改成功",
	})
}