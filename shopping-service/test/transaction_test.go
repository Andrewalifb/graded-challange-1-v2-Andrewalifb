package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/config"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/controller"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/dto"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/repository"
	"github.com/H8-FTGO-P3/graded-challange-2-v2-Andrewalifb/shopping-service/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_GetAllTransactions_OK(t *testing.T) {

	var DB *mongo.Client = config.ConnectDB()
	e := echo.New()

	var (
		transactionRepository repository.TransactionRepository = repository.NewTranscationRepository(DB)
		transactionService    service.TransactionService       = service.NewTransactionService(transactionRepository)
		transactionController controller.TransactionController = controller.NewTransactionController(transactionService)
	)

	e.GET("/api/v1/all-transactions", transactionController.HandleGetTransactionsRequest)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/all-transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)

	// Check status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check response body
	var response dto.TransactionPaginationResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

}
