package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/model"
	"test/service"
)
// 查询单个订单
func GetOrder(c *gin.Context)  {
    var id model.GetId
    //var idtest = c.Request.FormValue("id")
    //// 将request的query中的数据，自动解析到结构体
	c.ShouldBind(&id)
	fmt.Printf("结果",id)
	result,err2 :=service.GetOrder(id.Id)
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
// 添加订单
func CreateOrder(c *gin.Context){
	var order model.Order
	c.ShouldBind(&order)
	id,err:= service.CreateOrder(order)
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

// 查询订单列表
func SearchOrderList(c *gin.Context){
	var condition model.GetCondition
	c.ShouldBind(&condition)
	result, err :=service.SearchOrderList(condition)
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
//更新订单
func UpdateOrder(c *gin.Context){
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