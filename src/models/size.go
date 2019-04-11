package models

/* 商品尺码表 */
type Size struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	MerchantId	int64	`json:"merchant_id"` //所属商家
	Status		bool	`json:"status"`			//是否软删除
	UpdateTime string 	`json:"update_time"`
}
