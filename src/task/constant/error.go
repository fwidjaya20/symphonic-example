package constant

import "net/http"

const (
	ErrInvalidPayload = "T-001"
	ErrCreateRecord   = "T-002"
)

func ToHttpStatusCode(code string) int {
	switch code {
	case ErrCreateRecord:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
