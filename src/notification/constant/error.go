package constant

import "net/http"

const (
	ErrClientNotFound = "N-001"
	ErrSendMessage    = "N-002"
)

func ToHttpStatusCode(code string) int {
	switch code {
	default:
		return http.StatusInternalServerError
	}
}
