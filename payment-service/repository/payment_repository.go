package repository

import (
	"context"
	"time"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/config"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/entity"
	pb "github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PaymentRepository interface {
	CreatePayment(payment *entity.Payment) error
	GetPayment(id string) (*pb.Payment, error)
	UpdatePayment(payment *entity.Payment) error
	DeletePayment(id string) error
	GetAllPayments() ([]*pb.Payment, error)
}

type paymentRepository struct {
	paymentCollection *mongo.Collection
}

func NewPaymentRepository(DB *mongo.Client) PaymentRepository {
	return &paymentRepository{
		paymentCollection: config.GetCollection(DB, "payment"),
	}
}

func (r *paymentRepository) CreatePayment(payment *entity.Payment) error {
	payment.ID = primitive.NewObjectID()
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()

	_, err := r.paymentCollection.InsertOne(context.Background(), payment)
	if err != nil {
		return err
	}
	return nil
}

func (r *paymentRepository) GetPayment(id string) (*pb.Payment, error) {
	var payment entity.Payment
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.paymentCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&payment)
	if err != nil {
		return nil, err
	}
	return &pb.Payment{
		Id:             payment.ID.Hex(),
		Amount:         payment.Amount,
		Method:         payment.Method,
		Status:         payment.Status,
		CardType:       payment.CardType,
		TransactionFee: payment.TransactionFee,
	}, nil
}

func (r *paymentRepository) UpdatePayment(payment *entity.Payment) error {
	update := bson.M{
		"$set": bson.M{
			"amount":          payment.Amount,
			"method":          payment.Method,
			"card_type":       payment.CardType,
			"transaction_fee": payment.TransactionFee,
			"updated_at":      time.Now(),
		},
	}
	_, err := r.paymentCollection.UpdateOne(context.Background(), bson.M{"_id": payment.ID}, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *paymentRepository) DeletePayment(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.paymentCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}

func (r *paymentRepository) GetAllPayments() ([]*pb.Payment, error) {
	var payments []*pb.Payment

	cur, err := r.paymentCollection.Find(context.Background(), bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var payment entity.Payment
		err := cur.Decode(&payment)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &pb.Payment{
			Id:             payment.ID.Hex(),
			Amount:         payment.Amount,
			Method:         payment.Method,
			Status:         payment.Status,
			CardType:       payment.CardType,
			TransactionFee: payment.TransactionFee,
		})
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return payments, nil
}
