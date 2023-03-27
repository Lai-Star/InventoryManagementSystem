package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type RequestIDKey string

const RequestID RequestIDKey = "request_id"

func RequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestIDUuid := uuid.New().String()
		ctx := context.WithValue(req.Context(), RequestID, requestIDUuid)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
