package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/dto"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/entity"
)

// Function to Call Payment-Service REST API
func CreatePaymentForTransaction(request entity.PaymentApi) (dto.DataPaymentResponse, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return dto.DataPaymentResponse{}, err
	}

	resp, err := http.Post(os.Getenv("PAYMENT_API_URL"), "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return dto.DataPaymentResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dto.DataPaymentResponse{}, err
	}

	var response dto.PaymentApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return dto.DataPaymentResponse{}, err
	}

	return response.Data, nil
}
