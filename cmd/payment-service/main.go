package main

import (
	"log"
	"net"

	"github.com/HabibElias/nexus-pay-back/internal/domain/entities"
	persistence "github.com/HabibElias/nexus-pay-back/internal/infrastructure/persistence/gorm"
	grpc_handlers "github.com/HabibElias/nexus-pay-back/internal/presentation/grpc/handlers"
	"github.com/HabibElias/nexus-pay-back/internal/services"
	pb "github.com/HabibElias/nexus-pay-back/proto/pb/proto"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 1. Initialize DB (SQLite for Day 1/2)
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto-migrate the domain entities
	err = db.AutoMigrate(&entities.Payment{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 2. Initialize layers
	repo := persistence.NewPaymentRepositoryImpl(db)
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
