package models

/* 商品规格表 */
type PropertyModel struct {
	Id           int64     `json:"id"`
	GoodsId      int64     `json:"goods_id"`      //所属商品
	SizeId       int64     `json:"size_id"`       //尺寸或型号
	ColorId      int64     `json:"color_id"`      //颜色
	Img          string    `json:"img"`           //图片
	Stock        int64     `json:"stock"`         //库存
	RealPrice    float64   `json:"real_price"`    //真实价格
	VirtualPrice float64   `json:"virtual_price"` //虚拟价格img
	Status			bool	`json:"status"`
	CreateTime   string `json:"create_time"`
}
