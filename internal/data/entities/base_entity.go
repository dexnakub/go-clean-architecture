package data_entities

import "time"

type BaseEntity struct {
	CreateDate *time.Time `gorm:"type:datetime"`
	CreateBy   string     `gorm:"type:nvarchar(100)"`
	UpdateDate *time.Time `gorm:"type:datetime"`
	UpdateBy   string     `gorm:"type:nvarchar(100)"`
	IsDeleted  bool       `gorm:"default:false"`
}
