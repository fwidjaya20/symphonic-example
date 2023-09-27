package exception

type Exception struct {
	Code     string
	Err      error
	Message  string
	Metadata *any
}

func New(err error, code string, message string, metadata *any) Exception {
	return Exception{
		Code:     code,
		Err:      err,
		Message:  message,
		Metadata: metadata,
	}
}

func (e Exception) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return ""
}
