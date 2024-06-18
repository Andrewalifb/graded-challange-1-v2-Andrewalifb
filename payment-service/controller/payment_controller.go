package controller

import (
	"net/http"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/dto"
	pb "github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/proto"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/utils"
	"github.com/labstack/echo/v4"
)

type PaymentController interface {
	HandleCreatePaymentRequest(ctx echo.Context) error
	HandleUpdatePaymentRequest(ctx echo.Context) error
	HandleDeletePaymentRequest(ctx echo.Context) error
	HandleGetPaymentRequest(ctx echo.Context) error
	HandleGetAllPaymentsRequest(ctx echo.Context) error
}

type paymentController struct {
	service pb.PaymentServiceClient
}

func NewPaymentController(service pb.PaymentServiceClient) PaymentController {
	return &paymentController{service: service}
}

func (c *paymentController) HandleCreatePaymentRequest(ctx echo.Context) error {
	var payment pb.PaymentCreateRequest

	if err := ctx.Bind(&payment); err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	res, err := c.service.CreatePayment(ctx.Request().Context(), &payment)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_PAYMENT, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_PAYMENT, res)
	return ctx.JSON(http.StatusOK, successResponse)
}

func (c *paymentController) HandleUpdatePaymentRequest(ctx echo.Context) error {
	var payment pb.PaymentUpdateRequest

	if err := ctx.Bind(&payment); err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	res, err := c.service.UpdatePayment(ctx.Request().Context(), &payment)
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_PAYMENT, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_PAYMENT, res)
	return ctx.JSON(http.StatusOK, successResponse)
}

func (c *paymentController) HandleDeletePaymentRequest(ctx echo.Context) error {
	id := ctx.Param("id")

	res, err := c.service.DeletePayment(ctx.Request().Context(), &pb.PaymentDeleteRequest{Id: id})
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_PAYMENT, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_PAYMENT, res)
	return ctx.JSON(http.StatusOK, successResponse)
}

func (c *paymentController) HandleGetPaymentRequest(ctx echo.Context) error {
	id := ctx.Param("id")

	res, err := c.service.GetPayment(ctx.Request().Context(), &pb.PaymentGetRequest{Id: id})
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_PAYMENT, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_PAYMENT, res)
	return ctx.JSON(http.StatusOK, successResponse)
}

func (c *paymentController) HandleGetAllPaymentsRequest(ctx echo.Context) error {
	res, err := c.service.GetAllPayments(ctx.Request().Context(), &pb.PaymentGetAllRequest{})
	if err != nil {
		errorResponse := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_PAYMENT, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	successResponse := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_PAYMENT, res)
	return ctx.JSON(http.StatusOK, successResponse)
}
