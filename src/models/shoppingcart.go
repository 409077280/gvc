package models

/*购物车*/
type ShoppingCartModel struct {
	Id int64	`json:"id"`
	UserId int64	`json:"user_id"`
	GoodsId int64	`json:"goods_id"`
	PropertyId int64	`json:"property_id"`	//商品属性
	CreateTime string	`json:"create_time"`
}

