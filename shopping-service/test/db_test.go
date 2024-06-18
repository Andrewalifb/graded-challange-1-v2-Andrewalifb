package tests

import (
	"context"
	"testing"
	"time"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/config"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpDatabaseConnection() *mongo.Client {
	var DB *mongo.Client = config.ConnectDB()
	return DB
}

func Test_DBConnection(t *testing.T) {
	db := SetUpDatabaseConnection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Test ping to database
	err := db.Ping(ctx, nil)
	assert.NoError(t, err, "Should successfully connect to the database")

	// Disconnect from the database
	err = db.Disconnect(ctx)
	assert.NoError(t, err, "Should successfully disconnect from the database")
}
