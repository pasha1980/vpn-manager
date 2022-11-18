package error

type NotFoundError struct {
	*BaseError
}

func NewNotFoundError(message string, data interface{}) *NotFoundError {
	return &NotFoundError{
		BaseError: NewBaseError(
			404,
			"NOT_FOUND",
			message,
			data,
		),
	}
}
