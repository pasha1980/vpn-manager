package error

type BaseErrorInterface interface {
	GetErrorData() *BaseError
}

type BaseError struct {
	Code      int
	ErrorType string
	Message   string
	Data      interface{}
}

func (e *BaseError) Error() string {
	return "Internal"
}

func (e *BaseError) GetErrorData() *BaseError {
	return e
}

func NewBaseError(code int, errorType string, errMessage string, data interface{}) *BaseError {
	if data == nil {
		data = make(map[string]any)
	}
	return &BaseError{
		Code:      code,
		ErrorType: errorType,
		Message:   errMessage,
		Data:      data,
	}
}
