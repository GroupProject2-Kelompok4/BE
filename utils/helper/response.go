package helper

type DataResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseFormat(code int, message string, data interface{}) DataResponse {
	result := DataResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}

	return result
}
