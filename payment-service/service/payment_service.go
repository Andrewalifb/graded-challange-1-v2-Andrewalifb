package service

import (
	"context"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/entity"
	pb "github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/payment-service/repository"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, req *pb.PaymentCreateRequest) (*pb.PaymentCreateResponse, error)
	GetPayment(ctx context.Context, req *pb.PaymentGetRequest) (*pb.PaymentGetResponse, error)
	UpdatePayment(ctx context.Context, req *pb.PaymentUpdateRequest) (*pb.PaymentUpdateResponse, error)
	DeletePayment(ctx context.Context, req *pb.PaymentDeleteRequest) (*pb.PaymentDeleteResponse, error)
	GetAllPayments(ctx context.Context, req *pb.PaymentGetAllRequest) (*pb.PaymentGetAllResponse, error)
}

type paymentService struct {
	pb.UnimplementedPaymentServiceServer
	repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) pb.PaymentServiceServer {
	return &paymentService{repo: repo}
}

func (s *paymentService) CreatePayment(ctx context.Context, req *pb.PaymentCreateRequest) (*pb.PaymentCreateResponse, error) {
	payment := &entity.Payment{
		Amount:         req.Amount,
		Method:         req.Method,
		CardType:       req.CardType,
		Status:         "completed",
		TransactionFee: req.Amount * 0.01,
	}

	err := s.repo.CreatePayment(payment)
	if err != nil {
		return nil, err
	}

	return &pb.PaymentCreateResponse{
		Payment: &pb.Payment{
			Id:             payment.ID.Hex(),
			Amount:         payment.Amount,
			Method:         payment.Method,
			Status:         payment.Status,
			CardType:       payment.CardType,
			TransactionFee: payment.TransactionFee,
		},
	}, nil
}

func (s *paymentService) GetPayment(ctx context.Context, req *pb.PaymentGetRequest) (*pb.PaymentGetResponse, error) {
	payment, err := s.repo.GetPayment(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.PaymentGetResponse{
		Payment: payment,
	}, nil
}

func (s *paymentService) UpdatePayment(ctx context.Context, req *pb.PaymentUpdateRequest) (*pb.PaymentUpdateResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	payment := &entity.Payment{
		ID:             objectID,
		Amount:         req.Amount,
		Method:         req.Method,
		CardType:       req.CardType,
		TransactionFee: req.TransactionFee,
	}

	err = s.repo.UpdatePayment(payment)
	if err != nil {
		return nil, err
	}

	return &pb.PaymentUpdateResponse{
		Payment: &pb.Payment{
			Id:             payment.ID.Hex(),
			Amount:         payment.Amount,
			Method:         payment.Method,
			CardType:       payment.CardType,
			TransactionFee: payment.TransactionFee,
		},
	}, nil
}

func (s *paymentService) DeletePayment(ctx context.Context, req *pb.PaymentDeleteRequest) (*pb.PaymentDeleteResponse, error) {
	err := s.repo.DeletePayment(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.PaymentDeleteResponse{
		Id:     req.Id,
		Status: "success",
	}, nil
}

func (s *paymentService) GetAllPayments(ctx context.Context, req *pb.PaymentGetAllRequest) (*pb.PaymentGetAllResponse, error) { // Updated method
	payments, err := s.repo.GetAllPayments()
	if err != nil {
		return nil, err
	}
	return &pb.PaymentGetAllResponse{
		Payments: payments,
	}, nil
}
