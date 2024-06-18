package utils

import (
	"context"
	"os"
	"time"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Decrease Product Quantity After Transaction Created
func DecreaseProductQuantity(productID string, quantity int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collectionName := os.Getenv("COLLECTION_PRODUCT")
	collection := config.GetCollection(config.Db, collectionName)

	id, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return err
	}

	update := bson.M{
		"$inc": bson.M{
			"stock": -quantity, // Change "quantity" to "stock"
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}
