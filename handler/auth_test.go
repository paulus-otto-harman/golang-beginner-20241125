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
	"time"
)

func TestAuthHandler_Login(t *testing.T) {
	mockService := service.AuthServiceMock{}
	services := service.Service{Auth: &mockService}
	authHandler := InitAuthHandler(util.InitValidator(services, zap.NewNop()))

	tests := []struct {
		name               string
		requestBody        interface{}
		arg1MockSetup      *domain.Session
		arg2MockSetup      error
		expectedStatusCode int
		expectedBody       domain.Response
	}{
		{
			name: "Success Login",
			requestBody: domain.User{
				Username: "test@example.com",
				Password: "password123",
			},
			arg1MockSetup: &domain.Session{
				Token:     "blablabla",
				ExpiredAt: time.Now(),
			},
			arg2MockSetup:      nil,
			expectedStatusCode: http.StatusOK,
			expectedBody: domain.Response{
				StatusCode: http.StatusOK,
				Message:    util.StrPtr("user authenticated"),
			},
		},
		{
			name:               "Invalid Request Body",
			requestBody:        "invalid-json",
			arg1MockSetup:      nil,
			arg2MockSetup:      nil,
			expectedStatusCode: http.StatusBadRequest,
			expectedBody: domain.Response{
				StatusCode: http.StatusBadRequest,
				Message:    util.StrPtr("invalid request body"),
			},
		},
		{
			name: "Missing Username or Password",
			requestBody: domain.User{
				Username: "",
				Password: "",
			},
			arg1MockSetup:      nil,
			arg2MockSetup:      nil,
			expectedStatusCode: http.StatusUnprocessableEntity,
			expectedBody: domain.Response{
				StatusCode: http.StatusUnprocessableEntity,
				Message:    util.StrPtr("invalid input"),
			},
		},
		{
			name: "Authentication Failed Username",
			requestBody: domain.User{
				Username: "test@example.com",
				Password: "wrongpassword",
			},
			arg1MockSetup:      nil,
			arg2MockSetup:      errors.New("invalid username"),
			expectedStatusCode: http.StatusUnauthorized,
			expectedBody: domain.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    util.StrPtr("invalid username and/or password."),
			},
		},
		{
			name: "Authentication Failed Password",
			requestBody: domain.User{
				Username: "test@example.com",
				Password: "wrongpassword",
			},
			arg1MockSetup:      nil,
			arg2MockSetup:      errors.New("invalid password"),
			expectedStatusCode: http.StatusUnauthorized,
			expectedBody: domain.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    util.StrPtr("invalid username and/or password."),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err, "Failed to marshal request body")

			req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(requestBody))

			w := httptest.NewRecorder()

			mockService.On("Login", tt.requestBody).Once().Return(tt.arg1MockSetup, tt.arg2MockSetup)

			authHandler.Login(w, req)

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
