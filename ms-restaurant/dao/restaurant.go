package dao

import (
	"foodSocialContact/shared/model"
	"gorm.io/gorm"
)

type RestaurantRecord struct {
	model.BaseModel
}

type RestaurantDao struct {
	DB *gorm.DB
}

func (dao *RestaurantDao) getById(id int) (*RestaurantRecord, error) {

	return nil, nil
}
