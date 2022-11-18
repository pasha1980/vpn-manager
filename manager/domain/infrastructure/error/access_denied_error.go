package error

type AccessDeniedError struct {
	*BaseError
}

func NewAccessDeniedError(message string) *AccessDeniedError {
	return &AccessDeniedError{
		BaseError: NewBaseError(
			403,
			"ACCESS_DENIED",
			message,
			nil,
		),
	}
}
