package models

type GoodsModel struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`		//营销标题
	Description string    `json:"description"`  //详情页概述
	MerchantId  int64     `json:"merchant_id"`  //商家ID
	ExpressId   int64     `json:"express_id"`   //运费选择
	Discount	float64	  `json:"discount"`		//活动折扣
	ProductCode string    `json:"product_code"` //货号
	BarCode     string    `json:"bar_code"`     //条形码
	CategoryId  int64     `json:"category_id"`  //所属分类 ID
	SalesVolume int64     `json:"sales_volume"` //销售量
	Detail 		string    `json:"detail"` 		//商品详情富文本
	Display		bool      `json:"display"`		//是否上架
	Status		bool	  `json:"status"`		//是否删除
	CreateTime  string 	  `json:"create_time"`
}
