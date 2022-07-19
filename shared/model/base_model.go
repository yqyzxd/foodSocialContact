package model

import (
	"time"
)

//BaseModel 数据库公共字段
type BaseModel struct {
	//id create update delete valid
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:create_date"`
	UpdatedAt time.Time `gorm:"column:update_date"`
	Valid     int       `gorm:"column:is_valid"`
}
