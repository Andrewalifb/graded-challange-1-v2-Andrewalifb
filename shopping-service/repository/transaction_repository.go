package repository

import (
	"context"
	"math"
	"os"
	"time"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/config"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/dto"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository interface {
	InsertNewTransaction(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error)
	SelectTransactionsWithPaging(ctx context.Context, req dto.PaginationRequest) (dto.GetAllTransactionRepositoryResponse, error)
	SelectTransactionByID(ctx context.Context, id primitive.ObjectID) (entity.Transaction, error)
	UpdateTransactionByID(ctx context.Context, id primitive.ObjectID, transaction entity.Transaction) (entity.Transaction, error)
	DeleteTransactionByID(ctx context.Context, id primitive.ObjectID) error
	FindTransactions(ctx context.Context, filter bson.M) (*mongo.Cursor, error)
	PatchTransactionByID(ctx context.Context, id primitive.ObjectID, transaction entity.Transaction) (entity.Transaction, error)
}

type transactionRepository struct {
	transactionCollection *mongo.Collection
}

func NewTranscationRepository(DB *mongo.Client) TransactionRepository {
	return &transactionRepository{
		transactionCollection: config.GetCollection(DB, os.Getenv("COLLECTION_TRANSACTION")),
	}
}

// Create New Transaction Repository
func (repository *transactionRepository) InsertNewTransaction(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	_, err := repository.transactionCollection.InsertOne(ctx, transaction)
	if err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}

// Get All Transaction With Pagging Repository
func (repository *transactionRepository) SelectTransactionsWithPaging(ctx context.Context, req dto.PaginationRequest) (dto.GetAllTransactionRepositoryResponse, error) {
	result := make([]entity.Transaction, 0)

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	totalCount, err := repository.CountTransactionCollectionData(ctx)
	if err != nil {
		return dto.GetAllTransactionRepositoryResponse{}, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(req.Limit)))

	curr, err := repository.transactionCollection.Find(ctx, bson.D{{}}, dto.NewMongoPaginate(req.Limit, req.Page).GetPaginatedOpts())
	if err != nil {
		return dto.GetAllTransactionRepositoryResponse{}, err
	}

	for curr.Next(ctx) {
		var transaction entity.Transaction
		if err := curr.Decode(&transaction); err != nil {
			return dto.GetAllTransactionRepositoryResponse{}, err
		}
		result = append(result, transaction)
	}
	return dto.GetAllTransactionRepositoryResponse{
		Data: result,
		PaginationResponse: dto.PaginationResponse{
			Limit:   req.Limit,
			Page:    req.Page,
			MaxPage: int64(totalPages),
			Count:   int64(totalCount),
		},
	}, nil
}

// Count All Transaction Data Inside Transaction Collection
func (repository *transactionRepository) CountTransactionCollectionData(ctx context.Context) (int, error) {
	totalCount, err := repository.transactionCollection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return 0, err
	}
	return int(totalCount), nil
}

// Get Transaction By ID Repository
func (repository *transactionRepository) SelectTransactionByID(ctx context.Context, id primitive.ObjectID) (entity.Transaction, error) {
	var transaction entity.Transaction
	err := repository.transactionCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&transaction)
	if err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}

// Update Transaction By ID Repository
func (repository *transactionRepository) UpdateTransactionByID(ctx context.Context, id primitive.ObjectID, transaction entity.Transaction) (entity.Transaction, error) {
	_, err := repository.transactionCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": transaction})
	if err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}

// Delete Transaction By ID Repository
func (repository *transactionRepository) DeleteTransactionByID(ctx context.Context, id primitive.ObjectID) error {
	_, err := repository.transactionCollection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// Get Transaction Data By ID For Cron Job
func (repository *transactionRepository) FindTransactions(ctx context.Context, filter bson.M) (*mongo.Cursor, error) {
	return repository.transactionCollection.Find(ctx, filter)
}

// Update Transaction and Create Payment Repository
func (repository *transactionRepository) PatchTransactionByID(ctx context.Context, id primitive.ObjectID, transaction entity.Transaction) (entity.Transaction, error) {
	update := bson.M{
		"paymentId":       transaction.PaymentID,
		"status":          transaction.Status,
		"deliveryDetails": transaction.DeliveryDetails,
		"updatedAt":       time.Now(),
	}

	_, err := repository.transactionCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	if err != nil {
		return entity.Transaction{}, err
	}
	return transaction, nil
}
