package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentApiRequest struct {
	Method   string `json:"method"`
	CardType string `json:"card_type"`
}

type DataPaymentResponse struct {
	ID             primitive.ObjectID `json:"id"`
	Amount         float64            `json:"amount"`
	Method         string             `json:"method"`
	Status         string             `json:"status"`
	CardType       string             `json:"cardType"`
	TransactionFee float64            `json:"transactionFee"`
}

type PaymentApiResponse struct {
	Status  bool                `json:"status"`
	Message string              `json:"message"`
	Data    DataPaymentResponse `json:"data"`
}

type DataProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"imageUrl"`
	Brand       string  `json:"brand"`
	Rating      float64 `json:"rating"`
}

type ProductApiResponse struct {
	Status  bool                `json:"status"`
	Message string              `json:"message"`
	Data    DataProductResponse `json:"data"`
}
