package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"test/model"
)
var db *gorm.DB
var err error
func init(){
      db,err :=gorm.Open("mysql","root:123456@/demo_order?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.HasTable("demo_order")
      defer db.Close()
}
func Select(id uint)[]model.Order{
	tx :=db.Begin()
	tx.SingularTable(true)
	//不设置数据库表会被加s而找不到
	var order []model.Order
	tx.Where("id = ?",id).Find(&order)
	//相当于select * from order where id = name
	tx.Commit()
	return order
}

func Store(order model.Order)  error{
	tx :=db.Begin()
	tx.SingularTable(true)
	//不设置数据库表会被加s而找不到
	result :=tx.Create(&order)
	tx.Commit()
	return result.Error
}
func Selectlist(con string)([]model.Order,error){
	tx :=db.Begin()
	tx.SingularTable(true)
	//不设置数据库表会被加s而找不到
	var orders []model.Order
	result :=tx.Where("name LIKE ?", con).Find(&orders)
	tx.Commit()
	return orders,result.Error
}

func UpdateOrder(renew model.Order) error{
	tx :=db.Begin()
	tx.SingularTable(true)
	//不设置数据库表会被加s而找不到
	result :=tx.Save(&renew)
	tx.Commit()
	return result.Error
}