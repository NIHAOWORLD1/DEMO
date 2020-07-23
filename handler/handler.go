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
func Update(c *gin.Context){

}