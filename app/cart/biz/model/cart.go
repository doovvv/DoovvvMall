package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductId uint32 `gorm:"type:int(11);not null;"`
	Qty       uint32 `gorm:"type:int(11);not null;"`
}

func (Cart) TableName() string {
	return "cart"
}
func AddItem(ctx context.Context, db *gorm.DB, item *Cart) error {
	var row Cart
	err := db.WithContext(ctx).
		Model(item).
		Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
		First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if row.ID > 0 {
		return db.WithContext(ctx).Model(&row).UpdateColumn("qty", gorm.Expr("qty+?", item.Qty)).Error
	}
	return db.WithContext(ctx).Create(item).Error
}
func EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	if userId == 0 {
		return errors.New("user_id is required")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userId).Error
}
func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint32) ([]*Cart, error) {
	var rows []*Cart
	if userId == 0 {
		return rows, errors.New("user_id is required")
	}
	err := db.WithContext(ctx).Where(&Cart{UserId: userId}).Find(&rows).Error
	return rows, err
}
