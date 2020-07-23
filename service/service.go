package service

import (
    "test/db"
    "test/model"
)

func Select(id uint)  {
    db.Select(id)
}
func Store(order model.Order) (id uint,err error) {
    result := db.Store(order)
    id =order.ID
    if result != nil {
        err = result
        return
    }
    return
}
func Selectlist(condition model.GetorderList)(result []model.Order,err2 error){
    name :="%"+condition.User_name+"%"
    result,err := db.Selectlist(name)
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