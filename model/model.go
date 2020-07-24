package model

type Order struct {
	Id       uint  `json:"id" form:"id"`
	Order_no  string `json:"order_no" form:"order_no"`
	User_name string `json:"user_name" form:"user_name"`
	Amount   float64 `json:"amount" form:"amount"`
	Status   string `json:"status" form:"status"`
	File_url string `json:"file_url" form:"file_url"`
}

type  Getid struct {
	Id uint  `form:"id"`
}

type  GetorderList struct {
	User_name string `form:"user_name"`
}
