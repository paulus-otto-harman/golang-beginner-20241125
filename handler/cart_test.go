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

func TestCartHandler_Get(t *testing.T) {
	mockCartService := service.CartServiceMock{}
	services := service.Service{Cart: &mockCartService}
	cartHandler := InitCartHandler(services, zap.NewNop())

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/cart", bytes.NewBuffer(nil))
	req.Header.Set("token", "")

	mockCartService.On("Get", req.Header.Get("token")).Once().Return(&domain.Cart{}, nil)
	cartHandler.Get(w, req)

	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var responseBody domain.Response
	err := json.NewDecoder(w.Body).Decode(&responseBody)
	assert.NoError(t, err, "Failed to decode response body")
}

func TestCartHandler_Create(t *testing.T) {
	mockCartService := service.CartServiceMock{}
	services := service.Service{Cart: &mockCartService}
	cartHandler := InitCartHandler(services, zap.NewNop())

	tests := []struct {
		name               string
		requestBody        interface{}
		arg1MockSetup      error
		expectedStatusCode int
		expectedBody       domain.Response
	}{
		{
			name: "Add To Cart Success",
			requestBody: domain.CartItem{
				ProductId: 1,
				Quantity:  1,
			},
			arg1MockSetup:      nil,
			expectedStatusCode: http.StatusOK,
			expectedBody: domain.Response{
				StatusCode: http.StatusOK,
				Message:    util.StrPtr("Product successfully added to cart"),
			},
		},
		{
			name: "Add To Cart Failed Missing Quantity",
			requestBody: domain.CartItem{
				ProductId: 1,
			},
			arg1MockSetup:      errors.New("failed to add product to cart"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody: domain.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    util.StrPtr("failed to add product to cart"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err, "Failed to marshal request body")

			req := httptest.NewRequest(http.MethodPost, "/api/cart", bytes.NewBuffer(requestBody))
			req.Header.Set("token", "")

			w := httptest.NewRecorder()

			cartItem := domain.CartItem{}
			err = json.Unmarshal(requestBody, &cartItem)
			assert.NoError(t, err, "Failed to unmarshal request body")
			mockCartService.On("Store", cartItem, req.Header.Get("token")).Once().Return(tt.arg1MockSetup)

			cartHandler.Create(w, req)

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
