package grpc_handlers

import (
	"context"

	"github.com/HabibElias/nexus-pay-back/internal/services"
	pb "github.com/HabibElias/nexus-pay-back/proto/pb/proto"
)

type Handler struct {
	pb.UnimplementedPaymentServiceServer
	service *services.PaymentService
}

func NewHandler(s *services.PaymentService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreatePayment(
	ctx context.Context,
	req *pb.CreatePaymentRequest,
) (*pb.Payment, error) {

	payment, err := h.service.CreatePayment(ctx, req.Amount)
	if err != nil {
		return nil, err
	}

	return &pb.Payment{
		Id:     payment.ID,
		Amount: payment.Amount,
	}, nil
}
