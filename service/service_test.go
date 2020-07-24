package service
import(
	"fmt"
	"test/db"
	"testing"
)

func TestSelect(t *testing.T) {
	db.Init()
	id :=uint(1)
	req,err:=Select(id)
    if err != nil{
    	t.Error()
	}
	if req.Id == id {
		t.Log()
		fmt.Print(req)
	}
}