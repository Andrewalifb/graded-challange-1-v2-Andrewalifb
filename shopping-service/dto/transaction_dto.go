package dto

import (
	"errors"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/entity"
)

// Transaction Failed Messages
const (
	MESSAGE_FAILED_CREATE_TRANSACTION = "failed to create transaction"
	MESSAGE_FAILED_UPDATE_TRANSACTION = "failed to update transaction"
	MESSAGE_FAILED_DELETE_TRANSACTION = "failed to delete transaction"
	MESSAGE_FAILED_GET_TRANSACTION    = "failed to get transaction"
)

// Transaction Success Messages
const (
	MESSAGE_SUCCESS_CREATE_TRANSACTION = "success create transaction"
	MESSAGE_SUCCESS_UPDATE_TRANSACTION = "success update transaction"
	MESSAGE_SUCCESS_DELETE_TRANSACTION = "success delete transaction"
	MESSAGE_SUCCESS_GET_TRANSACTION    = "success get transaction"
)

// Transaction Custom Errors
var (
	ErrCreateTransaction = errors.New(MESSAGE_FAILED_CREATE_TRANSACTION)
	ErrUpdateTransaction = errors.New(MESSAGE_FAILED_UPDATE_TRANSACTION)
	ErrDeleteTransaction = errors.New(MESSAGE_FAILED_DELETE_TRANSACTION)
	ErrGetTransaction    = errors.New(MESSAGE_FAILED_GET_TRANSACTION)
)

type ProductDetailRequest struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type TransactionCreateRequest struct {
	Products []ProductDetailRequest `json:"products"`
	UserID   string                 `json:"userId"`
	BranchID string                 `json:"branchId"`
}

type ProductDetailResponse struct {
	ProductID   string  `json:"productId"`
	ProductName string  `json:"productName"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	SubTotal    float64 `json:"subTotal"`
}

type TransactionCreateResponse struct {
	ID              string                  `json:"id"`
	Products        []ProductDetailResponse `json:"products"`
	PaymentID       string                  `json:"paymentId"`
	UserID          string                  `json:"userId"`
	BranchID        string                  `json:"branchId"`
	Status          string                  `json:"status"`
	TotalAmount     float64                 `json:"totalAmount"`
	DeliveryDetails string                  `json:"deliveryDetails"`
}

type TransactionUpdateRequest struct {
	Products        []ProductDetailRequest `json:"products"`
	PaymentID       string                 `json:"paymentId"`
	UserID          string                 `json:"userId"`
	BranchID        string                 `json:"branchId"`
	Status          string                 `json:"status"`
	TotalAmount     float64                `json:"totalAmount"`
	DeliveryDetails string                 `json:"deliveryDetails"`
}

type TransactionUpdateResponse struct {
	ID              string                  `json:"id"`
	Products        []ProductDetailResponse `json:"products"`
	PaymentID       string                  `json:"paymentId"`
	UserID          string                  `json:"userId"`
	BranchID        string                  `json:"branchId"`
	Status          string                  `json:"status"`
	TotalAmount     float64                 `json:"totalAmount"`
	DeliveryDetails string                  `json:"deliveryDetails"`
}

type TransactionDeleteResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type TransactionGetResponse struct {
	ID              string                  `json:"id"`
	Products        []ProductDetailResponse `json:"products"`
	PaymentID       string                  `json:"paymentId"`
	UserID          string                  `json:"userId"`
	Status          string                  `json:"status"`
	TotalAmount     float64                 `json:"totalAmount"`
	DeliveryDetails string                  `json:"deliveryDetails"`
}

type TransactionPaginationResponse struct {
	Data []TransactionGetResponse `json:"data"`
	PaginationResponse
}

type GetAllTransactionRepositoryResponse struct {
	Data []entity.Transaction `json:"data"`
	PaginationResponse
}

type PatchTransactionByIdRequest struct {
	PaymentData PaymentApiRequest `json:"paymentData"`
}

type PatchTransactionByIdResponse struct {
	PaymentData     DataPaymentResponse    `json:"paymenyData"`
	TransactionData TransactionGetResponse `json:"transactionData"`
}
