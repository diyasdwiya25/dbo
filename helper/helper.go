package helper

import (
	"fmt"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func ApiResponse(message string, code int, status string, data interface{}) Response {

	jsonResponse := Response{
		Message: message,
		Code:    code,
		Status:  status,
		Data:    data,
	}

	return jsonResponse
}
func FormatValidationError(err error) []string {
	var errors []string
	fmt.Println(err)
	errors = append(errors, err.Error())

	return errors
}
