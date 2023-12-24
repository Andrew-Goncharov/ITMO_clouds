package schemes

const (
	statusError = "ERROR"
)

type SimpleWebError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"msg"`
}

func NewSimpleWebError(statusCode int, message string) SimpleWebError {
	return SimpleWebError{StatusCode: statusCode, Message: message}
}
