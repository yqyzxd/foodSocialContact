package dao

import (
	dinerpb "foodSocialContact/ms-diner/proto/gen"
	"foodSocialContact/shared/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DinerRecord struct {
	model.BaseModel
	*dinerpb.Diner
}

type DinerDao struct {
	DB *gorm.DB
}

func (d *DinerDao) FindUserByUsername(username string) (*DinerRecord, error) {
	var rec DinerRecord
	result := d.DB.Where("username = ?", username).First(&rec)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "")
	}
	return &rec, nil
}
