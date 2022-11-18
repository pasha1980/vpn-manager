package api

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	apiError "vpn-manager/domain/infrastructure/error"
)

func NewErrorHandler() func(err error, ctx echo.Context) {
	return func(err error, c echo.Context) {
		var statusCode int
		var errorType string
		var message string
		var data interface{}

		switch err.(type) {
		case *apiError.BaseError:
			baseError := err.(apiError.BaseErrorInterface)
			errorData := baseError.GetErrorData()

			statusCode = errorData.Code
			errorType = errorData.ErrorType
			message = errorData.Message
			data = errorData.Data

		case *echo.HTTPError:
			httpError := err.(*echo.HTTPError)
			statusCode = httpError.Code
			errorType = "UNDEFINED_ERROR"
			message = fmt.Sprintf("%s", httpError.Message)
		default:
			statusCode = http.StatusInternalServerError
			errorType = "INTERNAL_ERROR"
			message = "Internal server error"
		}

		response := map[string]interface{}{
			"type":    errorType,
			"message": message,
			"data":    data,
		}
		jsonData, _ := json.Marshal(response)

		if statusCode < 500 {
			log.Debug(string(jsonData))
		} else {
			log.Error(string(jsonData))
		}

		c.JSON(statusCode, jsonData)

	}
}
