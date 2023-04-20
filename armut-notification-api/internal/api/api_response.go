package api

type ApiResponse struct {
	Data    *interface{} `json:"Data"`
	Message string       `json:"Message"`
}

func Ok(data interface{}) *ApiResponse {
	apiResponse := ApiResponse{
		Data:    &data,
		Message: "Success",
	}
	return &apiResponse
}

func Error(message string) *ApiResponse {
	apiResponse := ApiResponse{
		Message: message,
	}

	return &apiResponse
}
