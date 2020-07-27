package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"test/model"
)
var db *gorm.DB
func Init(){
	var err error
	db,err =gorm.Open("mysql","root:123456@/demo_orde?charset=utf8&parseTime=True&loc=Local")
	// 检查数据库连接错误
	if err != nil {
		Createdatabese("demo_orde")
	}
	//是否已经存在表
	flag := db.HasTable("order")
	if !flag {
		db.HasTable(&model.Order{})
		//不设置数据库表会被加s而找不到
		db.AutoMigrate(&model.Order{})
	}
}
func Createdatabese(name string) {
	// 连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
    // exec函数为执行sql语句
	_, err = db.Exec("CREATE DATABASE "+name)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}
}
// 查询订单
func GetOrder(id uint) (model.Order,error) {
	db.SingularTable(true)
	var order model.Order
	result :=db.Where("id = ?",id).Find(&order)
	// 相当于select * from order where id = id
	return order,result.Error
}
// 创建订单
func CreateOrder(order model.Order) error{
	tx :=db.Begin()
	tx.SingularTable(true)
	if err :=tx.Create(&order).Error;err != nil{
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
// 查询订单
func SearchOrderList(con string)(orders []model.Order,err error){
	tx :=db.Begin()
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
// 更新订单
func UpdateOrder(renew model.Order) error{
	tx :=db.Begin()
    //tx.SingularTable(true)
	if err :=tx.Model(&renew).Updates(renew).Error;err != nil{
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}