package main

import (
	"os"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/config"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/controller"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/repository"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/routes"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/service"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var DB *mongo.Client = config.ConnectDB()

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

	var (
		transactionRepository repository.TransactionRepository = repository.NewTranscationRepository(DB)
		transactionService    service.TransactionService       = service.NewTransactionService(transactionRepository)
		transactionController controller.TransactionController = controller.NewTransactionController(transactionService)
	)

	var (
		productRepository repository.ProductRepository = repository.NewProductRepository(DB)
		productService    service.ProductService       = service.NewProductService(productRepository)
		productController controller.ProductController = controller.NewProductController(productService)
	)

	routes.TransactionRoute(e, transactionController)
	routes.ProductRoute(e, productController)

	utils.StartCronJob(transactionRepository)

	e.Logger.Fatal(e.Start(":8081"))
}
