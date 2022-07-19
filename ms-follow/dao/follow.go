package dao

import (
	"foodSocialContact/shared/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

//FollowRecord PO
type FollowRecord struct {
	model.BaseModel

	DinerID       int `gorm:"column:diner_id"`
	FollowDinerID int `gorm:"column:follow_diner_id"`
}

type FollowDao struct {
	DB *gorm.DB
}

func (d *FollowDao) SelectFollow(dinerID, followDinerID int) (*FollowRecord, error) {
	var rec FollowRecord
	result := d.DB.Where("follow_diner_id = ? and diner_id = ? ", followDinerID, dinerID).Find(&rec)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "")
	}
	return &rec, nil

}

func (d *FollowDao) Save(dinnerID, followDinerID, followStatus int) (*FollowRecord, error) {
	rec := FollowRecord{
		DinerID:       dinnerID,
		FollowDinerID: followDinerID,
		BaseModel: model.BaseModel{
			Valid: followStatus,
		},
	}
	result := d.DB.Create(&rec)
	if result.Error != nil {
		return nil, result.Error
	}

	return &rec, nil
}

func (d *FollowDao) Update(id, followStatus int) error {

	result := d.DB.Where("id = ? ", id).Update("is_valid", followStatus)

	return result.Error
}
