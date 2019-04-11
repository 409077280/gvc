package models

type VipModel struct {
	Id                int64     `json:"id"`
	Name              string    `json:"name"`              //VIP名称
	ShoppingDiscounts float64   `json:"shopping_discount"` //购物折扣率
	Brokerage         float64   `json:"brokerage"`         //佣金率
	Drawings          float64   `json:"drawings"`          //提现手续费折扣
	Price             float64   `json:"price"`             //单价
	Img               string    `json:"img"`               //VIP标识
	UpdateTime        string `json:"update_time"`
}
