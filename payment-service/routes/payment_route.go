package routes

import (
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/controller"
	"github.com/labstack/echo/v4"
)

func PaymentRoute(e *echo.Echo, paymentController controller.PaymentController) {
	routes := e.Group("/api")

	routesV1 := routes.Group("/v1")
	// Create New Payment
	routesV1.POST("/payments", paymentController.HandleCreatePaymentRequest)
	// Update Existing Payment
	routesV1.PUT("/payments/:id", paymentController.HandleUpdatePaymentRequest)
	// Delete Payment
	routesV1.DELETE("/payments/:id", paymentController.HandleDeletePaymentRequest)
	// Get Payment by ID
	routesV1.GET("/payments/:id", paymentController.HandleGetPaymentRequest)
	// Get All Payments
	routesV1.GET("/payments", paymentController.HandleGetAllPaymentsRequest)
}
