package controller

import (
	"net/http"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/dto"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/service"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/utils"
	"github.com/labstack/echo/v4"
)

// var paymentCollection *mongo.Collection = config.GetCollection(config.DB, os.Getenv("COLLECTION_PAYMENT"))

type TransactionController interface {
	HandleCreateTransactionRequest(c echo.Context) error
	HandleGetTransactionsRequest(c echo.Context) error
	HandleGetTransactionByIDRequest(c echo.Context) error
	HandleUpdateTransactionRequest(c echo.Context) error
	HandleDeleteTransactionRequest(c echo.Context) error
	HandleUpdateTransactionAfterPayment(c echo.Context) error
}

type transactionController struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &transactionController{
		transactionService: transactionService,
	}
}

// Create New Transaction Handler
func (controller *transactionController) HandleCreateTransactionRequest(c echo.Context) error {
	var transaction dto.TransactionCreateRequest

	if err := c.Bind(&transaction); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := controller.transactionService.ProcessNewTransaction(transaction)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_TRANSACTION, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_TRANSACTION, result)
	return c.JSON(http.StatusOK, res)
}

// Get All Transaction With PAging Handler
func (controller *transactionController) HandleGetTransactionsRequest(c echo.Context) error {
	var req dto.PaginationRequest
	if err := c.Bind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := controller.transactionService.FetchTransactionsWithPaging(req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TRANSACTION, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_TRANSACTION, result)
	return c.JSON(http.StatusOK, res)
}

// Get Transaction By ID Handler
func (controller *transactionController) HandleGetTransactionByIDRequest(c echo.Context) error {
	id := c.Param("id")

	result, err := controller.transactionService.FetchTransactionByID(id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TRANSACTION, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_TRANSACTION, result)
	return c.JSON(http.StatusOK, res)
}

// Update Transaction By ID Handler
func (controller *transactionController) HandleUpdateTransactionRequest(c echo.Context) error {
	id := c.Param("id")
	var transaction dto.TransactionUpdateRequest

	if err := c.Bind(&transaction); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := controller.transactionService.UpdateTransaction(transaction, id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_TRANSACTION, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_TRANSACTION, result)
	return c.JSON(http.StatusOK, res)
}

// Delete Transaction By ID Handler
func (controller *transactionController) HandleDeleteTransactionRequest(c echo.Context) error {
	id := c.Param("id")

	result, err := controller.transactionService.DeleteTransactionByID(id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_TRANSACTION, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_TRANSACTION, result)
	return c.JSON(http.StatusOK, res)
}

// Update Transaction Status And Create Payment from Payment-Service Handler
func (controller *transactionController) HandleUpdateTransactionAfterPayment(c echo.Context) error {
	id := c.Param("id")
	var transaction dto.PatchTransactionByIdRequest

	if err := c.Bind(&transaction); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := controller.transactionService.UpdateTransactionAfterPayment(transaction, id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_TRANSACTION, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_TRANSACTION, result)
	return c.JSON(http.StatusOK, res)
}
