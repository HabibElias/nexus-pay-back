package services

import (
	"context"

	"github.com/HabibElias/nexus-pay-back/internal/domain/entities"
	repositories "github.com/HabibElias/nexus-pay-back/internal/domain/repositories"
	"github.com/google/uuid"
)

type PaymentService struct {
	repo repositories.PaymentRepository
}

func NewPaymentService(repo repositories.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) CreatePayment(ctx context.Context, amount float64) (*entities.Payment, error) {
	payment := &entities.Payment{
		ID:     uuid.New().String(),
		Amount: amount,
	}

	if err := s.repo.Create(ctx, payment); err != nil {
		return nil, err
	}

	return payment, nil
}
