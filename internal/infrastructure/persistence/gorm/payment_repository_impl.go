package persistence

import (
	"context"

	"github.com/HabibElias/nexus-pay-back/internal/domain/entities"
	"github.com/HabibElias/nexus-pay-back/internal/domain/repositories"
	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepositoryImpl(db *gorm.DB) repositories.PaymentRepository {
	return &PaymentRepositoryImpl{db: db}
}

func (r *PaymentRepositoryImpl) Create(ctx context.Context, payment *entities.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}
