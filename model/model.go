package model

type Order struct {
	ID       uint
	Order_no  string
	User_name string
	Amount   float64
	Status   string
	File_url string
}
var Orders Order

type  Getid struct {
	ID uint  "form:id"
}

type  GetorderList struct {
	User_name string "form:id"
}
