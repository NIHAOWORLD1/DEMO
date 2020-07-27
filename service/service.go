package service

import (
    "test/db"
    "test/model"
)
// 查询订单
func GetOrder(id uint) ( model.Order, error) {
     result,err:=db.GetOrder(id)
     var err2 error
    if err != nil {
        err2 =err
        return result,err2
    }
    return result,err2
}
// 创建订单
func CreateOrder(order model.Order) (id uint,err error) {
    result := db.CreateOrder(order)
    id =order.Id
    if result != nil {
        err = result
        return
    }
    return
}
// 查询订单列表
func SearchOrderList(condition model.GetCondition)(result []model.Order,err2 error){
    name :="%"+condition.User_name+"%"
    result,err := db.SearchOrderList(name)
    if err != nil {
         err2 =err
        return
    }
    return
}
// 更新订单
func UpdateOrder(update model.Order) (err error){
    result := db.UpdateOrder(update)
    if result != nil {
        err = result
        return
    }
    return
}
// 上传文件
func UploadFile(id int, dst string) error{
    rid :=uint(id)
    var upload model.Order
    upload.File_url=dst
    upload.Id=rid
    result :=db.UpdateOrder(upload)
    return result
}
