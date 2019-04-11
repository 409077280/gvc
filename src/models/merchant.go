package models

type MerchantModel struct {
	Id int64 `json:"id"`
	UserName string	`json:"username"` //用户名
	Password	string	`json:"passowrd"`		//密码
	Name string	`json:"name"`			//店铺名
	LegalPerson string	`json:"legal_person"`		//法人名称
	IdCard	string		`json:"id_card"`	//法人身份证号
	CompanyName string	`json:"company_name"`	//公司名
	Description string 	`json:"description"`//店铺简介
	TelePhone string	`json:"telephone"`	//座机
	MobilePhone string	`json:"mobilephone"`	//手机号
	Status bool 	`json:"status"`//使用状态
	LastLoginTime string	`json:"last_login_time"`
	CreateTime string	`json:"create_time"`
	Img	string	`json:"img"`	//商家图标
}
