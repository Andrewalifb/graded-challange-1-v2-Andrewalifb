package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductDetails struct {
	ProductID primitive.ObjectID `bson:"productId" json:"productId"`
	Quantity  int                `bson:"quantity" json:"quantity"`
	Price     float64            `bson:"price" json:"price"`
	SubTotal  float64            `bson:"subTotal" json:"subTotal"`
}

type Transaction struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Products        []ProductDetails   `bson:"products" json:"products"`
	PaymentID       string             `bson:"paymentId" json:"paymentId"`
	UserID          string             `bson:"userId" json:"userId"`
	BranchID        string             `bson:"branchId" json:"branchId"`
	Status          string             `bson:"status" json:"status"`
	TotalAmount     float64            `bson:"totalAmount" json:"totalAmount"`
	DeliveryDetails string             `bson:"deliveryDetails" json:"deliveryDetails"`
	CreatedAt       time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type PaymentApi struct {
	Amount   float64 `bson:"amount" json:"amount"`
	Method   string  `bson:"method" json:"method"`
	CardType string  `bson:"cardType" json:"cardType"`
}
