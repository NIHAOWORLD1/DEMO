package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"test/model"
)
var db *gorm.DB
func Init(){
	var err error
	db,err =gorm.Open("mysql","root:123456@/demo_orde?charset=utf8&parseTime=True&loc=Local")
	//检查数据库连接错误
	if err != nil {
		Createdatabese("demo_orde")
	}
	//是否已经存在表
	flag := db.HasTable("order")
	if !flag {
		db.HasTable(&model.Order{})
		//不设置数据库表会被加s而找不到
		db.SingularTable(true)
		db.AutoMigrate(&model.Order{})
	}
}
func Createdatabese(name string) {
	//连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	//defer db.Close()
    //exec函数为执行sql语句
	_, err = db.Exec("CREATE DATABASE "+name)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
}

func Select(id uint) (model.Order,error) {
	var order model.Order
	result :=db.Where("id = ?",id).Find(&order)
	//相当于select * from order where id = id
	return order,result.Error
}

func Store(order model.Order)  error{
	tx :=db.Begin()
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