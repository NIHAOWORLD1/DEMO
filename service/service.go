package service

import (
    "test/db"
    "test/model"
)

func GetOrder(id uint) ( model.Order, error) {
     result,err:=db.GetOrder(id)
     var err2 error
    if err != nil {
        err2 =err
        return result,err2
    }
    return result,err2
}
func CreateOrder(order model.Order) (id uint,err error) {
    result := db.CreateOrder(order)
    id =order.Id
    if result != nil {
        err = result
        return
    }
    return
}
func SearchOrderList(condition model.GetCondition)(result []model.Order,err2 error){
    name :="%"+condition.User_name+"%"
    result,err := db.SearchOrderList(name)
    if err != nil {
         err2 =err
        return
    }
    return
}

func UpdateOrder(renew model.Order) (err error){
    result := db.UpdateOrder(renew)
    if result != nil {
        err = result
        return
    }
    return
}