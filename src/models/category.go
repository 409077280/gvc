package models


type CategoryModel struct {
	Id         int64     `json:"id"`
	ParentId   int64     `json:"parent_id"`
	Name       string    `json:"name"`
	Ico        string    `json:"ico"`    //图标
	Status     bool      `json:"status"` //是否禁用
	UpdateTime string `json:"update_time"`
}
