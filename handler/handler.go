package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"test/model"
	"test/service"
	"time"
)

// 查询单个订单
func GetOrder(c *gin.Context)  {
    var id model.GetId
    // var idtest = c.Request.FormValue("id")
    // 将request的query中的数据，自动解析到结构体
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
// 更新订单
func UpdateOrder(c *gin.Context){
	var update model.Order
	c.ShouldBind(&update)
	fmt.Printf("这是更新字段：",update)
    err :=service.UpdateOrder(update)
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
// 上传文件
func Upload(c *gin.Context){
	file, err := c.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "请求失败")
		return
	}
	id ,err2:=strconv.Atoi(c.Request.FormValue("id"))
	if err2 != nil {
		c.String(http.StatusBadRequest, "字符转化失败")
		return
	}
	fileName := file.Filename
	fmt.Println("文件名：", fileName)
	// 获取当前路径
	dir, _ := os.Getwd()
	dst :=path.Join(dir,fileName)
	// 保存文件到服务器本地
	// SaveUploadedFile(文件头，保存路径)
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
		return
	}
	err3 :=service.UploadFile(id,dst)
	if err3 != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "url更新失败",
		})
		return
	}
	c.String(http.StatusOK, "上传文件成功")
}
// 下载文件
func Download(c *gin.Context)  {
	var id model.GetId
	c.ShouldBind(&id)
	result,err:=service.GetOrder(id.Id)
	// 数据库中获取url
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "url为空",
		})
		return
	}
	// 创建一个到远程连接的请求
	resp, err := http.Get(result.File_url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data,err2 :=ioutil.ReadAll(resp.Body)
	if err2 != nil{
		panic(err2)
	}
	// 为下载文件设定命名
	fileName :=fmt.Sprintf("download_%d.jpg",time.Now().Unix())
	// 下载文件的存放路径
	filepath :=fmt.Sprintf("./download/%s",fileName)
	// 创建空白文件
	file,err3 :=os.Create(filepath)
	if err3 != nil{
		panic(err3)
	}
	defer file.Close()
	_,err4 :=file.Write(data)
	if err4 != nil{
		panic(err4)
	}
}