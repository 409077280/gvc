package models

/* 商品收藏 */
type CollectModel struct {
	Id int64	`json:"id"`
	UserId int64	`json:"user_id"`
	GoodsId int64	`json:"goods_id"`
	CreateTime string	`json:"create_time"`
}
