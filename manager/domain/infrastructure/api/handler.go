package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	apiError "vpn-manager/domain/infrastructure/error"
)

func NewErrorHandler() func(err error, ctx echo.Context) {
	return func(err error, c echo.Context) {
		var statusCode int
		var errorType string
		var message string
		var data interface{}

		if err.Error() == "Internal" {
			baseError := err.(apiError.BaseErrorInterface)
			errorData := baseError.GetErrorData()

			statusCode = errorData.Code
			errorType = errorData.ErrorType
			message = errorData.Message
			data = errorData.Data
		} else {
			switch err.(type) {
			case *echo.HTTPError:
				httpError := err.(*echo.HTTPError)
				statusCode = httpError.Code
				errorType = "UNDEFINED_ERROR"
				message = fmt.Sprintf("%s", httpError.Message)
			default:
				statusCode = http.StatusInternalServerError
				errorType = "INTERNAL_ERROR"
				message = err.Error()
			}
		}

		response := map[string]interface{}{
			"type":    errorType,
			"message": message,
			"data":    data,
		}

		c.JSON(statusCode, response)
	}
}
