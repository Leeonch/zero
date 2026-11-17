package httpapi

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
)

const timestampLayout = "2006-01-02T15:04:05.000Z"

// Response is the JSON envelope for IAM HTTP responses.
type Response[T any] struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      T      `json:"data"`
	Timestamp string `json:"timestamp"`
	Path      string `json:"path,omitempty"`
}

// WriteSuccess writes a successful response with the given HTTP status.
func WriteSuccess[T any](w http.ResponseWriter, r *http.Request, status int, data T) {
	httpx.WriteJsonCtx(r.Context(), w, status, Response[T]{
		Code:      0,
		Message:   "success",
		Data:      data,
		Timestamp: time.Now().UTC().Format(timestampLayout),
	})
}

// WriteError writes an error response with a message safe to return to the caller.
func WriteError(w http.ResponseWriter, r *http.Request, status int, message string) {
	httpx.WriteJsonCtx(r.Context(), w, status, Response[any]{
		Code:      -1,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().UTC().Format(timestampLayout),
		Path:      r.URL.RequestURI(),
	})
}
