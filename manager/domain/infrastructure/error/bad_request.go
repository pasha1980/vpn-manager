package error

type BadRequestError struct {
	*BaseError
}

func NewBadRequestError(message string, data interface{}) *BadRequestError {
	return &BadRequestError{
		BaseError: NewBaseError(
			400,
			"BAD_REQUEST",
			message,
			data,
		),
	}
}
