package domain_models

import "time"

type BaseModel struct {
	CreateDate *time.Time `json:"createDate"`
	CreateBy   string     `json:"createBy"`
	UpdateDate *time.Time `json:"updateDate"`
	UpdateBy   string     `json:"updateBy"`
	IsDeleted  bool       `json:"isDeleted"`
}
