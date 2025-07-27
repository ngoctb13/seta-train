package transaction

import (
	"context"

	"gorm.io/gorm"
)

type TxnManager interface {
	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
}
