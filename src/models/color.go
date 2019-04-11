package models

type ColorModel struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	MerchantId	int64	`json:"merchant_id"`	//所属商家
	Status		bool	`json:"status"`			//软删除
	UpdateTime string 	`json:"update_time"`
}
