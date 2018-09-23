package common

//AppError construct for application error
type AppError struct {
	Error      error  `json:"error"`
	StatusCode int    `json:"httpStatusCode"`
	Message    string `json:"Message"`
}

//NewAppError returns AppError
func NewAppError(err error, statusCode int, message string) *AppError {
	return &AppError{
		Error:      err,
		StatusCode: statusCode,
		Message:    message,
	}
}
