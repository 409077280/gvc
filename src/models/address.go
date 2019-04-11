package models

type AddressModel struct {
	Id int64		`json:"id"`
	UserId int64	`json:"user_id"`
	Country string	`json:"country"`
	Province string	`json:"province"` //省
	City string		`json:"city"`	//市
	County string	`json:"county"`	//区或者县
	Detail string	`json:"detail"` //详细地址
	CreateTime	string 	`json:"create_time"`
}

