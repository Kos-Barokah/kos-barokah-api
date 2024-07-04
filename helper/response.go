package helper

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ResponseValidation struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
}

func FormatResponse(status bool, message string, data any) any {
	var response = new(Response)
	response.Status = status
	response.Message = message
	response.Data = data
	return *response
}

func FormatResponseValidation(status bool, message string, msgErr any) any {
	var response = new(ResponseValidation)
	response.Status = status
	response.Message = message
	response.Error = msgErr
	return response
}

type ApiResponse[T any] struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}
