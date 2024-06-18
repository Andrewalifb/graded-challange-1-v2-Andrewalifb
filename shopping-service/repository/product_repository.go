package repository

import (
	"context"
	"math"
	"os"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/config"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/dto"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	InsertNewProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	SelectProductsWithPaging(ctx context.Context, req dto.PaginationRequest) (dto.GetAllProductRepositoryResponse, error)
	SelectProductByID(ctx context.Context, id primitive.ObjectID) (entity.Product, error)
	UpdateProductByID(ctx context.Context, id primitive.ObjectID, product entity.Product) (entity.Product, error)
	DeleteProductByID(ctx context.Context, id primitive.ObjectID) error
}

type productRepository struct {
	productCollection *mongo.Collection
}

func NewProductRepository(DB *mongo.Client) ProductRepository {
	return &productRepository{
		productCollection: config.GetCollection(DB, os.Getenv("COLLECTION_PRODUCT")),
	}
}

// Create New Product Repository
func (repository *productRepository) InsertNewProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	_, err := repository.productCollection.InsertOne(ctx, product)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

// Get All Product With Pagging Repository
func (repository *productRepository) SelectProductsWithPaging(ctx context.Context, req dto.PaginationRequest) (dto.GetAllProductRepositoryResponse, error) {
	result := make([]entity.Product, 0)

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	totalCount, err := repository.CountProductCollectionData(ctx)
	if err != nil {
		return dto.GetAllProductRepositoryResponse{}, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(req.Limit)))

	curr, err := repository.productCollection.Find(ctx, bson.D{{}}, dto.NewMongoPaginate(req.Limit, req.Page).GetPaginatedOpts())
	if err != nil {
		return dto.GetAllProductRepositoryResponse{}, err
	}

	for curr.Next(ctx) {
		var product entity.Product
		if err := curr.Decode(&product); err != nil {
			return dto.GetAllProductRepositoryResponse{}, err
		}
		result = append(result, product)
	}

	return dto.GetAllProductRepositoryResponse{
		Data: result,
		PaginationResponse: dto.PaginationResponse{
			Limit:   req.Limit,
			Page:    req.Page,
			MaxPage: int64(totalPages),
			Count:   int64(totalCount),
		},
	}, nil
}

// Get Count All Product Data On Collection Product
func (repository *productRepository) CountProductCollectionData(ctx context.Context) (int, error) {
	totalCount, err := repository.productCollection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return 0, err
	}
	return int(totalCount), nil
}

// Get Product By ID Repository
func (repository *productRepository) SelectProductByID(ctx context.Context, id primitive.ObjectID) (entity.Product, error) {
	var product entity.Product
	err := repository.productCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

// Update Product Data By ID Repository
func (repository *productRepository) UpdateProductByID(ctx context.Context, id primitive.ObjectID, product entity.Product) (entity.Product, error) {
	_, err := repository.productCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": product})
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

// Delete Product By ID Repository
func (repository *productRepository) DeleteProductByID(ctx context.Context, id primitive.ObjectID) error {
	_, err := repository.productCollection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
