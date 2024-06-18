package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Amount         float64            `bson:"amount" json:"amount"`
	Method         string             `bson:"method" json:"method"`
	Status         string             `bson:"status" json:"status"`
	CardType       string             `bson:"cardType" json:"cardType"`
	TransactionFee float64            `bson:"transactionFee" json:"transactionFee"`
	CreatedAt      time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt" json:"updatedAt"`
}
