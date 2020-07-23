package model

type Order struct {
	ID       uint
	Order_no  string
	User_name string
	Amount   float64
	Status   string
	File_url string
}
type  Getid struct {
	ID uint  "form:id"
}