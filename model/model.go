package model

type Order struct {
	Id       uint  `json:"id"`
	Order_no  string `json:"order_no"`
	User_name string `json:"user_name"`
	Amount   float64 `json:"amount"`
	Status   string `json:"status"`
	File_url string `json:"file_url"`
}

type  Getid struct {
	ID uint  "form:id"
}

type  GetorderList struct {
	User_name string "form:user_name"
}
