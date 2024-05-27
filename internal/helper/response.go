package helper

import "github.com/labstack/echo/v4"

type BaseResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseData(code int, message string, data interface{}) *BaseResponse {
	var response BaseResponse

	response.Code = code
	response.Message = message

	if data != nil {
		response.Data = data
	}

	return &response
}

func ErrorHandler(c echo.Context, statusCode int, errorMessage string) error {
	response := ResponseData(statusCode, errorMessage, nil)
	return c.JSON(statusCode, response)
}

type ErrorResponseJson struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type SuccessResponseJson struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type SuccessResponseJsonWithPagenation struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}
type SuccessResponseJsonWithPagenationAndCount struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Count      int         `json:"count_data,omitempty"`
}

type SuccessResponseJsonWithPagenationAndCountAll struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Count      interface{} `json:"count,omitempty"`
}

type SuccessResponseJsonWithPaginationAndCount struct {
	Status      bool        `json:"status"`
	Message     string      `json:"message"`
	DataMessage string      `json:"data_message,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Pagination  interface{} `json:"pagination,omitempty"`
	Count       interface{} `json:"count,omitempty"`
}

func ErrorResponse(message string) ErrorResponseJson {
	return ErrorResponseJson{
		Status:  false,
		Message: message,
	}
}

func SuccessResponse(message string) SuccessResponseJson {
	return SuccessResponseJson{
		Status:  true,
		Message: message,
	}
}
