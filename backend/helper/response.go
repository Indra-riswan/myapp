package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type Emptyobj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

func BuildErrors(message, errors string, data interface{}) Response {
	SplitErr := strings.Split(errors, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  SplitErr,
		Data:    data,
	}
	return res
}
