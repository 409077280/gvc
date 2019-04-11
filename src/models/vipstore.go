package models

import "time"

type VipStoreModel struct {
	Id         int64     `json:"id"`
	VipId      int64     `json:"vip_id"`     //VIP等级
	CardCrypt  string    `json:"card_crypt"` //卡密
	OwnerId    int64     `json:"owner_id"`   //所有者
	Status     bool      `json:"status"`     //是否使用
	RegisterId   int64     `json:"register_id"`   // VIP注册者
	CreateTime time.Time `json:"create_time"`
	Registerime   time.Time `json:"register_time"` //使用时间
}
