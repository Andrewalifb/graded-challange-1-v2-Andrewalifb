package dto

import (
	"errors"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/entity"
)

// Product Failed Messages
const (
	MESSAGE_FAILED_CREATE_PRODUCT = "failed to create product"
	MESSAGE_FAILED_UPDATE_PRODUCT = "failed to update product"
	MESSAGE_FAILED_DELETE_PRODUCT = "failed to delete product"
	MESSAGE_FAILED_GET_PRODUCT    = "failed to get product"
)

// Product Success Messages
const (
	MESSAGE_SUCCESS_CREATE_PRODUCT = "success create product"
	MESSAGE_SUCCESS_UPDATE_PRODUCT = "success update product"
	MESSAGE_SUCCESS_DELETE_PRODUCT = "success delete product"
	MESSAGE_SUCCESS_GET_PRODUCT    = "success get product"
)

// Product Custom Errors
var (
	ErrCreateProduct = errors.New(MESSAGE_FAILED_CREATE_PRODUCT)
	ErrUpdateProduct = errors.New(MESSAGE_FAILED_UPDATE_PRODUCT)
	ErrDeleteProduct = errors.New(MESSAGE_FAILED_DELETE_PRODUCT)
	ErrGetProduct    = errors.New(MESSAGE_FAILED_GET_PRODUCT)
)

type ProductCreateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"imageUrl"`
	Brand       string  `json:"brand"`
	Rating      float64 `json:"rating"`
}

type ProductCreateResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"imageUrl"`
	Brand       string  `json:"brand"`
	Rating      float64 `json:"rating"`
}

type ProductUpdateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"imageUrl"`
	Brand       string  `json:"brand"`
	Rating      float64 `json:"rating"`
}

type ProductUpdateResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"imageUrl"`
	Brand       string  `json:"brand"`
	Rating      float64 `json:"rating"`
}

type ProductDeleteResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"` // e.g., "deleted"
}

type ProductGetResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"imageUrl"`
	Brand       string  `json:"brand"`
	Rating      float64 `json:"rating"`
}

type ProductPaginationResponse struct {
	Data []ProductGetResponse `json:"data"`
	PaginationResponse
}

type GetAllProductRepositoryResponse struct {
	Data []entity.Product `json:"data"`
	PaginationResponse
}
