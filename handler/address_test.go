package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"project/domain"
	"project/service"
	"project/util"
	"testing"
)

func TestAddressHandler_All(t *testing.T) {
	mockService := service.AddressServiceMock{}
	services := service.Service{Address: &mockService}
	addressHandler := InitAddressHandler(util.InitValidator(services, zap.NewNop()))

	req := httptest.NewRequest(http.MethodPost, "/api/addresses", bytes.NewBuffer(nil))
	w := httptest.NewRecorder()
	mockService.On("Index", "").Once().Return(nil, nil)
	addressHandler.All(w, req)

	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var responseBody domain.Response
	err := json.NewDecoder(w.Body).Decode(&responseBody)
	assert.NoError(t, err, "Failed to decode response body")
	assert.Equal(t, http.StatusOK, responseBody.StatusCode)
}

func TestAddressHandler_Create(t *testing.T) {
	mockService := service.AddressServiceMock{}
	services := service.Service{Address: &mockService}
	addressHandler := InitAddressHandler(util.InitValidator(services, zap.NewNop()))

	tests := []struct {
		name               string
		requestBody        interface{}
		arg1MockSetup      error
		expectedStatusCode int
		expectedBody       domain.Response
	}{
		{
			name: "Address Created",
			requestBody: domain.Address{
				FullName: "Jane Doe",
				Email:    "janedoe@mail.com",
				Detail:   "Address Detail",
			},
			arg1MockSetup:      nil,
			expectedStatusCode: http.StatusCreated,
			expectedBody: domain.Response{
				StatusCode: http.StatusCreated,
				Message:    util.StrPtr("address successfully created"),
			},
		},
		{
			name: "Create Address Failed Missing Email",
			requestBody: domain.Address{
				FullName: "Jane Doe",
				Email:    "",
				Detail:   "Address Detail",
			},
			arg1MockSetup:      errors.New("invalid input"),
			expectedStatusCode: http.StatusUnprocessableEntity,
			expectedBody: domain.Response{
				StatusCode: http.StatusUnprocessableEntity,
				Message:    util.StrPtr("invalid input"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err, "Failed to marshal request body")

			req := httptest.NewRequest(http.MethodPost, "/api/addresses", bytes.NewBuffer(requestBody))

			w := httptest.NewRecorder()
			var customerAddress domain.Address
			err = json.Unmarshal(requestBody, &customerAddress)
			assert.NoError(t, err, "Failed to unmarshal request body")
			mockService.On("Create", &customerAddress, "").Once().Return(tt.arg1MockSetup)

			addressHandler.Create(w, req)

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
