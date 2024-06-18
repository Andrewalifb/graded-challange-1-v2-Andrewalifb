package service

import (
	"context"
	"time"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/dto"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/entity"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService interface {
	ProcessNewProduct(req dto.ProductCreateRequest) (dto.ProductCreateResponse, error)
	FetchProductsWithPaging(req dto.PaginationRequest) (dto.ProductPaginationResponse, error)
	FetchProductByID(id string) (dto.ProductGetResponse, error)
	UpdateProduct(req dto.ProductUpdateRequest, id string) (dto.ProductUpdateResponse, error)
	DeleteProductByID(id string) (dto.ProductDeleteResponse, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

// Create New Product Service
func (service *productService) ProcessNewProduct(req dto.ProductCreateRequest) (dto.ProductCreateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newProduct := entity.Product{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Stock:       req.Stock,
		ImageURL:    req.ImageURL,
		Brand:       req.Brand,
		Rating:      req.Rating,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	productCreated, err := service.productRepo.InsertNewProduct(ctx, newProduct)
	if err != nil {
		return dto.ProductCreateResponse{}, err
	}

	return dto.ProductCreateResponse{
		ID:          productCreated.ID.Hex(),
		Name:        productCreated.Name,
		Description: productCreated.Description,
		Price:       productCreated.Price,
		Category:    productCreated.Category,
		Stock:       productCreated.Stock,
		ImageURL:    productCreated.ImageURL,
		Brand:       productCreated.Brand,
		Rating:      productCreated.Rating,
	}, nil
}

// Get All Product With Pagging Service
func (service *productService) FetchProductsWithPaging(req dto.PaginationRequest) (dto.ProductPaginationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dataWithPaginate, err := service.productRepo.SelectProductsWithPaging(ctx, req)
	if err != nil {
		return dto.ProductPaginationResponse{}, err
	}
	responseData := make([]dto.ProductGetResponse, len(dataWithPaginate.Data))
	for i, product := range dataWithPaginate.Data {
		responseData[i] = dto.ProductGetResponse{
			ID:          product.ID.Hex(),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Category:    product.Category,
			Stock:       product.Stock,
			ImageURL:    product.ImageURL,
			Brand:       product.Brand,
			Rating:      product.Rating,
		}
	}

	return dto.ProductPaginationResponse{
		Data:               responseData,
		PaginationResponse: dataWithPaginate.PaginationResponse,
	}, nil
}

// Get Product By ID Service
func (service *productService) FetchProductByID(id string) (dto.ProductGetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.ProductGetResponse{}, err
	}

	product, err := service.productRepo.SelectProductByID(ctx, objectID)
	if err != nil {
		return dto.ProductGetResponse{}, err
	}

	return dto.ProductGetResponse{
		ID:          product.ID.Hex(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		Stock:       product.Stock,
		ImageURL:    product.ImageURL,
		Brand:       product.Brand,
		Rating:      product.Rating,
	}, nil
}

// Update Product By ID Service
func (service *productService) UpdateProduct(req dto.ProductUpdateRequest, id string) (dto.ProductUpdateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.ProductUpdateResponse{}, err
	}

	updatedProduct := entity.Product{
		ID:          objectID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Stock:       req.Stock,
		ImageURL:    req.ImageURL,
		Brand:       req.Brand,
		Rating:      req.Rating,
		UpdatedAt:   time.Now(),
	}

	productUpdated, err := service.productRepo.UpdateProductByID(ctx, objectID, updatedProduct)
	if err != nil {
		return dto.ProductUpdateResponse{}, err
	}

	return dto.ProductUpdateResponse{
		ID:          productUpdated.ID.Hex(),
		Name:        productUpdated.Name,
		Description: productUpdated.Description,
		Price:       productUpdated.Price,
		Category:    productUpdated.Category,
		Stock:       productUpdated.Stock,
		ImageURL:    productUpdated.ImageURL,
		Brand:       productUpdated.Brand,
		Rating:      productUpdated.Rating,
	}, nil
}

// Delete Product By ID Service
func (service *productService) DeleteProductByID(id string) (dto.ProductDeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.ProductDeleteResponse{}, err
	}

	err = service.productRepo.DeleteProductByID(ctx, objectID)
	if err != nil {
		return dto.ProductDeleteResponse{}, err
	}

	return dto.ProductDeleteResponse{
		ID:     id,
		Status: "deleted",
	}, nil
}
