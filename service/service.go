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