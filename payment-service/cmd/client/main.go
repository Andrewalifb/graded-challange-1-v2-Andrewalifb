package main

import (
	"flag"
	"log"
	"os"

	pb "github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/proto"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/controller"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "server:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPaymentServiceClient(conn)

	var paymentController controller.PaymentController = controller.NewPaymentController(client)

	e := echo.New()

	var log = logrus.New()

	log.Out = os.Stdout

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logrus.StandardLogger().Out,
	}))

	routes.PaymentRoute(e, paymentController)

	e.Logger.Fatal(e.Start(":8080"))
}
