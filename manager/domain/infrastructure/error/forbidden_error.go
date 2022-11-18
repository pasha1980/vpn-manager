package error

type AccessDeniedError struct {
	*BaseError
}

func NewForbiddenError(message string) *AccessDeniedError {
	return &AccessDeniedError{
		BaseError: NewBaseError(
			403,
			"ACCESS_DENIED",
			message,
			nil,
		),
	}
}
