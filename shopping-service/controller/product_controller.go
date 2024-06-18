package controller

import (
	"net/http"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/dto"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/service"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/utils"
	"github.com/labstack/echo/v4"
)

// var paymentCollection *mongo.Collection = config.GetCollection(config.DB, os.Getenv("COLLECTION_PAYMENT"))

type ProductController interface {
	HandleCreateProductRequest(c echo.Context) error
	HandleGetProductsWithPagingRequest(c echo.Context) error
	HandleGetProductByIDRequest(c echo.Context) error
	HandleUpdateProductRequest(c echo.Context) error
	HandleDeleteProductRequest(c echo.Context) error
}

type productController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

// Create New Product Handler
func (controller *productController) HandleCreateProductRequest(c echo.Context) error {
	var product dto.ProductCreateRequest

	if err := c.Bind(&product); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := controller.productService.ProcessNewProduct(product)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_PRODUCT, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_PRODUCT, result)
	return c.JSON(http.StatusOK, res)
}

// Get All Product With Pagging Handler
func (controller *productController) HandleGetProductsWithPagingRequest(c echo.Context) error {
	var req dto.PaginationRequest
	if err := c.Bind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := controller.productService.FetchProductsWithPaging(req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_PRODUCT, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_PRODUCT, result)
	return c.JSON(http.StatusOK, res)
}

// Get Product By ID Handler
func (controller *productController) HandleGetProductByIDRequest(c echo.Context) error {
	id := c.Param("id")

	result, err := controller.productService.FetchProductByID(id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_PRODUCT, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_PRODUCT, result)
	return c.JSON(http.StatusOK, res)
}

// Update Product By ID Handler
func (controller *productController) HandleUpdateProductRequest(c echo.Context) error {
	id := c.Param("id")
	var product dto.ProductUpdateRequest

	if err := c.Bind(&product); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := controller.productService.UpdateProduct(product, id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_PRODUCT, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_PRODUCT, result)
	return c.JSON(http.StatusOK, res)
}

// Delete Product By ID Handler
func (controller *productController) HandleDeleteProductRequest(c echo.Context) error {
	id := c.Param("id")

	result, err := controller.productService.DeleteProductByID(id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_PRODUCT, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_PRODUCT, result)
	return c.JSON(http.StatusOK, res)
}
