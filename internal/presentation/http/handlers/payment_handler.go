package handlers

import (
	"context"
	"log"
	"time"

	pb "github.com/HabibElias/nexus-pay-back/proto/pb/proto"
	"github.com/gofiber/fiber/v2"
)

type CreatePaymentReq struct {
	Amount float64 `json:"amount"`
}

type PaymentHandler struct {
	grpcClient pb.PaymentServiceClient
}

func NewPaymentHandler(grpcClient pb.PaymentServiceClient) *PaymentHandler {
	return &PaymentHandler{grpcClient: grpcClient}
}

func (h *PaymentHandler) CreatePaymentHandler(c *fiber.Ctx) error {
	var req CreatePaymentReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := h.grpcClient.CreatePayment(ctx, &pb.CreatePaymentRequest{
		Amount: req.Amount,
	})
	if err != nil {
		log.Printf("Error calling gRPC CreatePayment: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":     res.Id,
		"amount": res.Amount,
	})
}
