package routes

import (
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/controller"
	"github.com/labstack/echo/v4"
)

func TransactionRoute(e *echo.Echo, transactionController controller.TransactionController) {
	routes := e.Group("/api")

	routesV1 := routes.Group("/v1")
	// Create New Transaction
	routesV1.POST("/transactions", transactionController.HandleCreateTransactionRequest)
	// Get All Transaction With Pagging
	routesV1.GET("/all-transactions", transactionController.HandleGetTransactionsRequest)
	// Get Transaction By ID
	routesV1.GET("/transactions/:id", transactionController.HandleGetTransactionByIDRequest)
	// Update Transaction By ID
	routesV1.PUT("/transactions/:id", transactionController.HandleUpdateTransactionRequest)
	// Delete Transaction By ID
	routesV1.DELETE("/transactions/:id", transactionController.HandleDeleteTransactionRequest)
	// Completed Payment For a Transaction ID, auto create payment on payment-service
	routesV1.PATCH("/transactions/:id/payment", transactionController.HandleUpdateTransactionAfterPayment)

}
