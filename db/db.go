package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"test/model"
)
var db *gorm.DB
func Init(){
	var err error
	db,err =gorm.Open("mysql","root:123456@/demo_order?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	err = db.Error
	if err != nil {
		panic(err)
	}
	flag := db.HasTable("order")
	if !flag {
		println("err")
	}

}
func Select(id uint) (model.Order,error) {
	db.SingularTable(true)
	//不设置数据库表会被加s而找不到
	var order model.Order
	result :=db.Where("id = ?",id).Find(&order)
	//相当于select * from order where id = name
	return order,result.Error
}

func Store(order model.Order)  error{
	fmt.Println(db)
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