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
	"time"
)

func TestAuthHandler_Login(t *testing.T) {
	mockService := service.AuthServiceMock{}
	logger := zap.NewNop()

	services := service.Service{Auth: &mockService}

	authHandler := InitAuthHandler(util.InitValidator(services, logger))

	t.Run("Success Login", func(t *testing.T) {
		requestBody, err := json.Marshal(domain.User{
			Username: "test@example.com",
			Password: "password123",
		})

		assert.NoError(t, err, "Failed to marshal request body")

		// request
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))

		// Prepare response recorder
		w := httptest.NewRecorder()

		// setup mock
		mockService.On("Login", domain.User{
			Username: "test@example.com",
			Password: "password123",
		}).Once().Return(&domain.Session{
			Token:     "blablabla",
			ExpiredAt: time.Now(),
		}, nil)

		// Call the handler
		authHandler.Login(w, req)

		// Assert status code
		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		// Periksa body response
		var responseBody domain.Response
		err = json.NewDecoder(w.Body).Decode(&responseBody)
		assert.NoError(t, err, "Failed to decode response body")
		assert.Equal(t, "user authenticated", *responseBody.Message)
	})
}
