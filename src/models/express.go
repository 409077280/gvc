package models

type ExpressModel struct {
	Id       int64   `json:"id"`
	MerchantId	int64	`json:"merchant_id"`
	Company  string  `json:"company"`   //公司名称
	Item string  `json:"item"` //项目名称
	Price    float64 `json:"price"`     //项目价格
	Create_time string `json:"create_time"`
}
