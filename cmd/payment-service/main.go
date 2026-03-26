package main

import (
	"log"
	"net"

	"github.com/HabibElias/nexus-pay-back/internal/config"
	persistence "github.com/HabibElias/nexus-pay-back/internal/infrastructure/persistence/gorm"
	grpc_handlers "github.com/HabibElias/nexus-pay-back/internal/presentation/grpc/handlers"
	"github.com/HabibElias/nexus-pay-back/internal/services"
	pb "github.com/HabibElias/nexus-pay-back/proto/pb/proto"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()

	// setup database
	config.SetupDatabase(*cfg)

	// 2. Initialize layers
	repo := persistence.NewPaymentRepositoryImpl(config.DB)
	service := services.NewPaymentService(repo)
	handler := grpc_handlers.NewHandler(service)

	// 3. Start gRPC Server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, handler)

	log.Println("Server running on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
