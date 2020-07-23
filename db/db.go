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
	db1 :=db
	defer db1.Close()
	db1.SingularTable(true)
	//不设置数据库表会被加s而找不到
	var order []model.Order
	db1.Where("id = ?",id).Find(&order)
	//相当于select * from order where id = name
	return order
}

func Store(order model.Order)  error{
	db1 :=db
	defer db1.Close()
	db1.SingularTable(true)
	//不设置数据库表会被加s而找不到
	tx :=db1.Create(&order)
	return tx.Error
}