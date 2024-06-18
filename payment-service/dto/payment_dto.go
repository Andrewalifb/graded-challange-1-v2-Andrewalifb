package dto

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Payment Failed Messages
const (
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_CREATE_PAYMENT     = "failed to create payment"
	MESSAGE_FAILED_UPDATE_PAYMENT     = "failed to update payment"
	MESSAGE_FAILED_DELETE_PAYMENT     = "failed to delete payment"
	MESSAGE_FAILED_GET_PAYMENT        = "failed to get payment"
)

// Payment Success Messages
const (
	MESSAGE_SUCCESS_CREATE_PAYMENT = "success create payment"
	MESSAGE_SUCCESS_UPDATE_PAYMENT = "success update payment"
	MESSAGE_SUCCESS_DELETE_PAYMENT = "success delete payment"
	MESSAGE_SUCCESS_GET_PAYMENT    = "success get payment"
)

// Payment Custom Errors
var (
	ErrCreatePayment = errors.New(MESSAGE_FAILED_CREATE_PAYMENT)
	ErrUpdatePayment = errors.New(MESSAGE_FAILED_UPDATE_PAYMENT)
	ErrDeletePayment = errors.New(MESSAGE_FAILED_DELETE_PAYMENT)
	ErrGetPayment    = errors.New(MESSAGE_FAILED_GET_PAYMENT)
)

type PaymentCreateRequest struct {
	Amount   float64 `json:"amount"`
	Method   string  `json:"method"`
	CardType string  `json:"cardType"`
}

type PaymentCreateResponse struct {
	ID             primitive.ObjectID `json:"id"`
	Amount         float64            `json:"amount"`
	Method         string             `json:"method"`
	Status         string             `json:"status"`
	CardType       string             `json:"cardType"`
	TransactionFee float64            `json:"transactionFee"`
}

type PaymentUpdateRequest struct {
	Amount         float64 `json:"amount"`
	Method         string  `json:"method"`
	CardType       string  `json:"cardType"`
	TransactionFee float64 `json:"transactionFee"`
}

type PaymentUpdateResponse struct {
	ID             string  `json:"id"`
	Amount         float64 `json:"amount"`
	Method         string  `json:"method"`
	CardType       string  `json:"cardType"`
	TransactionFee float64 `json:"transactionFee"`
}

type PaymentDeleteResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type PaymentGetResponse struct {
	ID             string  `json:"id"`
	Amount         float64 `json:"amount"`
	Method         string  `json:"method"`
	CardType       string  `json:"cardType"`
	TransactionFee float64 `json:"transactionFee"`
}
