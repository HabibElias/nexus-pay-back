package routes

import (
	"github.com/HabibElias/nexus-pay-back/internal/presentation/http/handlers"
	pb "github.com/HabibElias/nexus-pay-back/proto/pb/proto"
	"github.com/gofiber/fiber/v2"
)

type CreatePaymentReq struct {
	Amount float64 `json:"amount"`
}

func SetupPaymentRoutes(app *fiber.App, grpcClient pb.PaymentServiceClient, payment_handler handlers.PaymentHandler) {
	paymentGroup := app.Group("/api/payments")

	paymentGroup.Post("/", payment_handler.CreatePaymentHandler)
}
