package main

import (
	"log"
	"net"

	"github.com/HabibElias/nexus-pay-back/internal/config"
	persistence "github.com/HabibElias/nexus-pay-back/internal/infrastructure/persistence/gorm"
	grpc_handlers "github.com/HabibElias/nexus-pay-back/internal/presentation/grpc/handlers"
	"github.com/HabibElias/nexus-pay-back/internal/presentation/http/handlers"
	"github.com/HabibElias/nexus-pay-back/internal/presentation/http/routes"
	"github.com/HabibElias/nexus-pay-back/internal/services"
	pb "github.com/HabibElias/nexus-pay-back/proto/pb/proto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
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
	go func() {
		lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
		if err != nil {
			log.Fatal(err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterPaymentServiceServer(grpcServer, handler)

		reflection.Register(grpcServer)

		log.Print(`
    _   __                     ____
   / | / /___  _  ____  _______/ __ \____ ___  __
  /  |/ / __ \| |/_/ / / / ___/ /_/ / __ '/ / / /
 / /|  / /_/ />  </ /_/ (__  ) ____/ /_/ / /_/ /
/_/ |_/\____/_/|_|\__,_/____/_/    \__,_/\__, /
                                        /____/
`)
		log.Println("🚀 Payment Service is running on gRPC port: " + cfg.GRPCPort)

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	// 4. Setup gRPC client for HTTP Server
	conn, err := grpc.NewClient("localhost:"+cfg.GRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	grpcClient := pb.NewPaymentServiceClient(conn)

	// 5. Setup HTTP Server
	app := fiber.New()
	paymentHandler := handlers.NewPaymentHandler(grpcClient)

	routes.SetupPaymentRoutes(app, grpcClient, *paymentHandler)

	log.Print("HTTP API Server is starting on port: ", cfg.HTTPPort)
	log.Fatal(app.Listen(":" + cfg.HTTPPort))
}
