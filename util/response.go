package util

import (
	"encoding/json"
	"net/http"
	"project/domain"
)

func Response(w http.ResponseWriter) ResponseWriter {
	return ResponseWriter{Writer: w}
}

type ResponseWriter struct {
	Writer http.ResponseWriter
}

func (response ResponseWriter) Json(statusCode int, message string, data ...interface{}) {
	r := domain.Response{
		StatusCode: statusCode,
		Data:       nil,
	}

	if message != "" {
		r.Message = &message
	}

	if len(data) > 0 {
		r.Data = data[0]
	}

	response.Writer.WriteHeader(statusCode)

	json.NewEncoder(response.Writer).Encode(r)
}

func (response ResponseWriter) ValidationFail(data interface{}) {
	message := "invalid input"
	response.Json(http.StatusUnprocessableEntity, message, data)
}
