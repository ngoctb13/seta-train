package transaction

import (
	"context"

	"gorm.io/gorm"
)

type GormTxnManager struct {
	db *gorm.DB
}

func NewGormTxnManager(db *gorm.DB) *GormTxnManager {
	return &GormTxnManager{
		db: db,
	}
}

func (t *GormTxnManager) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
