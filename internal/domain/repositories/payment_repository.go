package repositories

import (
	"context"

	"github.com/HabibElias/nexus-pay-back/internal/domain/entities"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *entities.Payment) error
}
