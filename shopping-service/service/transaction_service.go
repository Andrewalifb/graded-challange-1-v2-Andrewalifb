package service

import (
	"context"
	"time"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/dto"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/entity"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/repository"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionService interface {
	ProcessNewTransaction(req dto.TransactionCreateRequest) (dto.TransactionCreateResponse, error)
	FetchTransactionsWithPaging(req dto.PaginationRequest) (dto.TransactionPaginationResponse, error)
	FetchTransactionByID(id string) (dto.TransactionGetResponse, error)
	UpdateTransaction(req dto.TransactionUpdateRequest, id string) (dto.TransactionUpdateResponse, error)
	DeleteTransactionByID(id string) (dto.TransactionDeleteResponse, error)
	UpdateTransactionAfterPayment(req dto.PatchTransactionByIdRequest, id string) (dto.PatchTransactionByIdResponse, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{transactionRepo: transactionRepo}
}

// Create New Transaction Service
func (service *transactionService) ProcessNewTransaction(req dto.TransactionCreateRequest) (dto.TransactionCreateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Total Amount On Each Transaction
	var countTotalAmount float64

	productDetails := make([]entity.ProductDetails, len(req.Products))
	for i, product := range req.Products {
		// Get Product Price to Count Sub Total and Total Amount
		productData, err := utils.FetchProductData(product.ProductID)
		if err != nil {
			// handle error
			return dto.TransactionCreateResponse{}, err
		}
		productID, err := primitive.ObjectIDFromHex(product.ProductID)
		if err != nil {
			// handle error
			return dto.TransactionCreateResponse{}, err
		}
		productDetails[i] = entity.ProductDetails{
			ProductID: productID,
			Quantity:  product.Quantity,
			Price:     productData.Price,
			SubTotal:  float64(product.Quantity) * productData.Price,
		}
		countTotalAmount += float64(product.Quantity) * productData.Price
	}

	newTransaction := entity.Transaction{
		ID:              primitive.NewObjectID(),
		Products:        productDetails,
		PaymentID:       "",
		UserID:          req.UserID,
		BranchID:        req.BranchID,
		Status:          "pending",
		TotalAmount:     countTotalAmount,
		DeliveryDetails: "waiting for payment completion",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	transactionCreated, err := service.transactionRepo.InsertNewTransaction(ctx, newTransaction)
	if err != nil {
		return dto.TransactionCreateResponse{}, err
	}

	// Decrease the product quantity in the product collection after transaction is successfully created
	for _, product := range req.Products {
		err = utils.DecreaseProductQuantity(product.ProductID, product.Quantity)
		if err != nil {
			return dto.TransactionCreateResponse{}, err
		}
	}

	productDetailResponses := make([]dto.ProductDetailResponse, len(transactionCreated.Products))
	for i, product := range transactionCreated.Products {
		// Get Each Product Details for Response
		productData, err := utils.FetchProductData(product.ProductID.Hex())
		if err != nil {
			return dto.TransactionCreateResponse{}, err
		}
		productDetailResponses[i] = dto.ProductDetailResponse{
			ProductID:   product.ProductID.Hex(),
			ProductName: productData.Name,
			Quantity:    product.Quantity,
			Price:       product.Price,
			SubTotal:    product.SubTotal,
		}
	}

	return dto.TransactionCreateResponse{
		ID:              transactionCreated.ID.Hex(),
		Products:        productDetailResponses,
		PaymentID:       transactionCreated.PaymentID,
		UserID:          transactionCreated.UserID,
		BranchID:        transactionCreated.BranchID,
		Status:          transactionCreated.Status,
		TotalAmount:     transactionCreated.TotalAmount,
		DeliveryDetails: transactionCreated.DeliveryDetails,
	}, nil
}

// Get All Transaction With Pagging Service
func (service *transactionService) FetchTransactionsWithPaging(req dto.PaginationRequest) (dto.TransactionPaginationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dataWithPaginate, err := service.transactionRepo.SelectTransactionsWithPaging(ctx, req)
	if err != nil {
		return dto.TransactionPaginationResponse{}, err
	}
	responseData := make([]dto.TransactionGetResponse, len(dataWithPaginate.Data))
	for i, transaction := range dataWithPaginate.Data {
		products := make([]dto.ProductDetailResponse, len(transaction.Products))
		for j, product := range transaction.Products {
			productData, err := utils.FetchProductData(product.ProductID.Hex())
			if err != nil {
				return dto.TransactionPaginationResponse{}, err
			}
			products[j] = dto.ProductDetailResponse{
				ProductID:   product.ProductID.Hex(),
				ProductName: productData.Name,
				Quantity:    product.Quantity,
				Price:       productData.Price,
				SubTotal:    float64(product.Quantity) * productData.Price,
			}
		}

		responseData[i] = dto.TransactionGetResponse{
			ID:              transaction.ID.Hex(),
			Products:        products,
			PaymentID:       transaction.PaymentID,
			UserID:          transaction.UserID,
			Status:          transaction.Status,
			TotalAmount:     transaction.TotalAmount,
			DeliveryDetails: transaction.DeliveryDetails,
		}
	}

	return dto.TransactionPaginationResponse{
		Data:               responseData,
		PaginationResponse: dataWithPaginate.PaginationResponse,
	}, nil
}

// Get Transacton By ID Service
func (service *transactionService) FetchTransactionByID(id string) (dto.TransactionGetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.TransactionGetResponse{}, err
	}

	transaction, err := service.transactionRepo.SelectTransactionByID(ctx, objectID)
	if err != nil {
		return dto.TransactionGetResponse{}, err
	}

	products := make([]dto.ProductDetailResponse, len(transaction.Products))
	for j, product := range transaction.Products {
		productData, err := utils.FetchProductData(product.ProductID.Hex())
		if err != nil {
			// handle error
			return dto.TransactionGetResponse{}, err
		}
		products[j] = dto.ProductDetailResponse{
			ProductID:   product.ProductID.Hex(),
			ProductName: productData.Name,
			Quantity:    product.Quantity,
			Price:       productData.Price,
			SubTotal:    float64(product.Quantity) * productData.Price,
		}
	}

	return dto.TransactionGetResponse{
		ID:              transaction.ID.Hex(),
		Products:        products,
		PaymentID:       transaction.PaymentID,
		UserID:          transaction.UserID,
		Status:          transaction.Status,
		TotalAmount:     transaction.TotalAmount,
		DeliveryDetails: transaction.DeliveryDetails,
	}, nil
}

// Update Transaction By ID Service
func (service *transactionService) UpdateTransaction(req dto.TransactionUpdateRequest, id string) (dto.TransactionUpdateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.TransactionUpdateResponse{}, err
	}

	productDetails := make([]entity.ProductDetails, len(req.Products))
	for i, product := range req.Products {
		productData, err := utils.FetchProductData(product.ProductID)
		if err != nil {
			return dto.TransactionUpdateResponse{}, err
		}
		var countTotalAmount float64
		productID, err := primitive.ObjectIDFromHex(product.ProductID)
		if err != nil {
			return dto.TransactionUpdateResponse{}, err
		}
		productDetails[i] = entity.ProductDetails{
			ProductID: productID,
			Quantity:  product.Quantity,
			Price:     productData.Price,
			SubTotal:  float64(product.Quantity) * productData.Price,
		}
		countTotalAmount += float64(product.Quantity) * productData.Price
	}

	updatedTransaction := entity.Transaction{
		ID:              objectID,
		Products:        productDetails,
		PaymentID:       req.PaymentID,
		UserID:          req.UserID,
		BranchID:        req.BranchID,
		Status:          req.Status,
		TotalAmount:     req.TotalAmount,
		DeliveryDetails: req.DeliveryDetails,
		UpdatedAt:       time.Now(),
	}

	transactionUpdated, err := service.transactionRepo.UpdateTransactionByID(ctx, objectID, updatedTransaction)
	if err != nil {
		return dto.TransactionUpdateResponse{}, err
	}

	productDetailResponses := make([]dto.ProductDetailResponse, len(transactionUpdated.Products))
	for i, product := range transactionUpdated.Products {
		productData, err := utils.FetchProductData(product.ProductID.Hex())
		if err != nil {
			// handle error
			return dto.TransactionUpdateResponse{}, err
		}
		productDetailResponses[i] = dto.ProductDetailResponse{
			ProductID:   product.ProductID.Hex(),
			ProductName: productData.Name,
			Quantity:    product.Quantity,
			Price:       product.Price,
			SubTotal:    product.SubTotal,
		}
	}

	return dto.TransactionUpdateResponse{
		ID:              transactionUpdated.ID.Hex(),
		Products:        productDetailResponses,
		PaymentID:       transactionUpdated.PaymentID,
		UserID:          transactionUpdated.UserID,
		BranchID:        transactionUpdated.BranchID,
		Status:          transactionUpdated.Status,
		TotalAmount:     transactionUpdated.TotalAmount,
		DeliveryDetails: transactionUpdated.DeliveryDetails,
	}, nil
}

// Delete Transaction By ID Service
func (service *transactionService) DeleteTransactionByID(id string) (dto.TransactionDeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.TransactionDeleteResponse{}, err
	}

	err = service.transactionRepo.DeleteTransactionByID(ctx, objectID)
	if err != nil {
		return dto.TransactionDeleteResponse{}, err
	}

	return dto.TransactionDeleteResponse{
		ID:     id,
		Status: "deleted",
	}, nil
}

// Update Transaction Status and Create Payment Service
func (service *transactionService) UpdateTransactionAfterPayment(req dto.PatchTransactionByIdRequest, id string) (dto.PatchTransactionByIdResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.PatchTransactionByIdResponse{}, err
	}

	// Get Transaction Detail By ID
	transactionData, err := service.transactionRepo.SelectTransactionByID(ctx, objectID)
	if err != nil {
		return dto.PatchTransactionByIdResponse{}, err
	}

	// Prepare Data to Insert On Payment-Service
	payment := entity.PaymentApi{
		Amount:   transactionData.TotalAmount,
		Method:   req.PaymentData.Method,
		CardType: req.PaymentData.CardType,
	}
	// Call Payment- Service API to Create New Payment For This Transaction
	paymentData, err := utils.CreatePaymentForTransaction(payment)
	if err != nil {
		return dto.PatchTransactionByIdResponse{}, err
	}

	// Update Transaction Data
	updatedTransaction := entity.Transaction{
		ID:              objectID,
		Products:        transactionData.Products,
		PaymentID:       paymentData.ID.Hex(), // Get Payment ID
		UserID:          transactionData.UserID,
		BranchID:        transactionData.BranchID,
		Status:          paymentData.Status, // Get Payment Status
		TotalAmount:     transactionData.TotalAmount,
		DeliveryDetails: "shipped", // Update Delivaery Status To Shipped
		UpdatedAt:       time.Now(),
	}

	// Exceute Update Transaction Data
	transactionUpdated, err := service.transactionRepo.PatchTransactionByID(ctx, objectID, updatedTransaction)
	if err != nil {
		return dto.PatchTransactionByIdResponse{}, err
	}

	productDetailResponses := make([]dto.ProductDetailResponse, len(transactionUpdated.Products))
	for i, product := range transactionUpdated.Products {
		// Get Transaction Data Details
		productData, err := utils.FetchProductData(product.ProductID.Hex())
		if err != nil {
			return dto.PatchTransactionByIdResponse{}, err
		}
		productDetailResponses[i] = dto.ProductDetailResponse{
			ProductID:   product.ProductID.Hex(),
			ProductName: productData.Name,
			Quantity:    product.Quantity,
			Price:       productData.Price,
			SubTotal:    float64(product.Quantity) * productData.Price,
		}
	}

	return dto.PatchTransactionByIdResponse{
		PaymentData: paymentData,
		TransactionData: dto.TransactionGetResponse{
			ID:              transactionData.ID.Hex(),
			Products:        productDetailResponses,
			PaymentID:       updatedTransaction.PaymentID,
			UserID:          transactionData.UserID,
			Status:          updatedTransaction.Status,
			TotalAmount:     transactionData.TotalAmount,
			DeliveryDetails: updatedTransaction.DeliveryDetails,
		},
	}, nil
}
