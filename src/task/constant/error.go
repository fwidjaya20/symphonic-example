package constant

import "net/http"

const (
	ErrInvalidPayload       = "T-001"
	ErrCreateRecord         = "T-002"
	ErrPublishCreatedRecord = "T-003"
	ErrRetrieveRecord       = "T-004"
)

func ToHttpStatusCode(code string) int {
	switch code {
	case ErrInvalidPayload:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
