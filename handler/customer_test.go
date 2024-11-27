package handler

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"project/domain"
	"project/service"
	"project/util"
	"testing"
)

func TestCustomerHandler_Register(t *testing.T) {
	mockCustomerService := service.CustomerServiceMock{}
	mockValidationService := service.ValidationServiceMock{}
	services := service.Service{Customer: &mockCustomerService, Validation: &mockValidationService}
	customerHandler := InitCustomerHandler(util.InitValidator(services, zap.NewNop()))

	tests := []struct {
		name               string
		requestBody        interface{}
		arg1MockSetup      error
		expectedStatusCode int
		expectedBody       domain.Response
	}{
		{
			name: "Customer Successfully Created",
			requestBody: domain.Customer{
				Name:     "Jane Doe",
				Username: "janedoe@mail.com",
				Password: "password",
			},
			arg1MockSetup:      nil,
			expectedStatusCode: http.StatusCreated,
			expectedBody: domain.Response{
				StatusCode: http.StatusCreated,
				Message:    util.StrPtr("Customer registered successfully."),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err, "Failed to marshal request body")

			req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(requestBody))

			w := httptest.NewRecorder()

			customer := domain.Customer{}
			err = json.Unmarshal(requestBody, &customer)
			assert.NoError(t, err, "Failed to unmarshal request body")
			mockValidationService.On("IsUnique", "users", "username", customer.Username).Once().Return(true, nil)
			mockCustomerService.On("Register", &customer).Once().Return(tt.arg1MockSetup)

			customerHandler.Registration(w, req)

			res := w.Result()
			assert.Equal(t, tt.expectedStatusCode, res.StatusCode)

			var responseBody domain.Response
			err = json.NewDecoder(w.Body).Decode(&responseBody)
			assert.NoError(t, err, "Failed to decode response body")
			assert.Equal(t, tt.expectedBody.StatusCode, responseBody.StatusCode)
			assert.Equal(t, *tt.expectedBody.Message, *responseBody.Message)
		})
	}
}
