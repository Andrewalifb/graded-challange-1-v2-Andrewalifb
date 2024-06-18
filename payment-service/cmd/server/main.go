package main

import (
	"log"
	"net"

	pb "github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/proto"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/config"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/repository"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func main() {
	var DB *mongo.Client = config.ConnectDB()

	var paymentRepository repository.PaymentRepository = repository.NewPaymentRepository(DB)
	var paymentService pb.PaymentServiceServer = service.NewPaymentService(paymentRepository)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, paymentService)

	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
