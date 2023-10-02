package constant

import "net/http"

const (
	ErrInvalidPayload       = "T-001"
	ErrCreateRecord         = "T-002"
	ErrPublishCreatedRecord = "T-003"
)

func ToHttpStatusCode(code string) int {
	switch code {
	case ErrCreateRecord, ErrPublishCreatedRecord:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
