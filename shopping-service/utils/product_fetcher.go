package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/dto"
)

// Get a Product Details
func FetchProductData(id string) (dto.DataProductResponse, error) {
	productURL := fmt.Sprintf("%s/%s", os.Getenv("FETCH_PRODUCT_URL"), id)
	resp, err := http.Get(productURL)
	if err != nil {
		return dto.DataProductResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.DataProductResponse{}, err
	}

	var response dto.ProductApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return dto.DataProductResponse{}, err
	}

	return response.Data, nil
}
