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
	db.HasTable(&model.Order{})
	db.SingularTable(true)
	db.AutoMigrate(&model.Order{})
}
func Select(id uint) (model.Order,error) {
	//不设置数据库表会被加s而找不到
	var order model.Order
	result :=db.Where("id = ?",id).Find(&order)
	//相当于select * from order where id = id
	return order,result.Error
}

func Store(order model.Order)  error{
	tx :=db.Begin()
	tx.SingularTable(true)
	//不设置数据库表会被加s而找不到
	result :=tx.Create(&order)
	tx.Commit()
	return result.Error
}
func Selectlist(con string)(orders []model.Order,err error){
	tx :=db.Begin()
	tx.SingularTable(true)
	//不设置数据库表会被加s而找不到
	if len(con) != 0 {
		result := tx.Where("user_name LIKE ?", con).Find(&orders)
		err = result.Error
		tx.Commit()
	}else {
		result := tx.Find(&orders)
		err = result.Error
		tx.Commit()
	}
	return
}

func UpdateOrder(renew model.Order) error{
	fmt.Print(renew)
	tx :=db.Begin()
	tx.SingularTable(true)
	//不设置数据库表会被加s而找不到
	result :=tx.Save(&renew)
	tx.Commit()
	return result.Error
}