package service
import(
	"fmt"
	"test/db"
	"test/model"
	"testing"
)

func TestSelect(t *testing.T) {
	db.Init()
	id :=uint(1)
	req,err:=GetOrder(id)
    if err != nil{
    	t.Error()
	}
	if req.Id == id {
		t.Log()
		fmt.Print(req)
	}
}

func TestCreateOrder(t *testing.T) {
	 db.Init()
     testdata :=model.Order{1,"1","niexiaojun",1000,"seccess",""}
     id,err :=CreateOrder(testdata)
	if err != nil {
       t.Error()
	}
	if id == testdata.Id{
		t.Log()
	}
}
func TestUpdateOrder(t *testing.T) {
	db.Init()
	testdata :=model.Order{1,"1","niexiaojun",1000,"seccess",""}
	err :=db.UpdateOrder(testdata)
	if err != nil {
		t.Error()
	}
	t.Log()
}

func TestSearchOrderList(t *testing.T) {
	db.Init()
	testdata :=model.Order{1,"2","huangxiaohua",2000,"seccess",""}
	err :=db.UpdateOrder(testdata)
	if err != nil {
		t.Error()
	}
	t.Log()
}