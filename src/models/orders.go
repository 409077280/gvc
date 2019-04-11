package models

type OrdersModel struct {
	Id             string  `json:"id"`
	OrderCode      string  `json:"order_code"`      // 订单编号
	GoodsTitle     string  `json:"goods_title"`		// 商品名称
	PropertySize   string  `json:"property_size"`   // 属性尺寸或型号
	PropertyColor  string  `json:"property_color"`  // 属性颜色
	PropertyImg    string  `json:"property_img"`    // 属性图片
	PropertyPrice  float64 `json:"property_price"`  // 真实价格
	GoodsNumber    bool    `json:"goods_number"`    // 购买数量
	GoodsDiscount  int64   `json:"goods_discount"`  // 当前商品折扣
	VipDiscount    float64 `json:"vip_discount"`    // VIP购物折扣
	ExpressStatus  bool    `json:"express_status"`  // 是否使用邮寄
	ExpressAddress string  `json:"express_address"` // 当前订单邮寄地址
	ExpressCompany string  `json:"express_company"` // 当前订单邮寄地址
	ExpressCode    string  `json:"express_code"`    // 物流单号
	ExpressFee 	   float64 `json:"express_fee"` 	// 当前邮寄费用
	CurrentPrice   float64 `json:"current_price"`   // 当前商品价格
	TotalPrice     float64 `json:"total_price"`   	// 当前总价格
	PayStatus      bool    `json:"pay_status"`      // 支付状态
	PaymentCode    string  `json:"payment_code"`    // 支付流水号
	InvoiceStatus  bool    `json:"invoice_status"`  // 是否开发票
	InvoiceCode    string  `json:"invoice_code"`    // 发票编号
	GoodsId        int64   `json:"goods_id"`        // 商品编号
	PropertyId     int64   `json:"property_id"`     // 规格ID
	UserId         int64   `json:"user_id"`         // 购买者ID
	MerchantId     int64   `json:"merchant_id"`  	// 商家ID
	ExpressId	   int64   `json:"express_id"`		// 物流类型ID
	Status         bool    `json:"status"`			// 是否取消订单
	CreateTime     string  `json:"create_time"`
}
