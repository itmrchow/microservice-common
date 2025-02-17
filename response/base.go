package response

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type FailedResponse struct {
	BaseResponse
	Error string `json:"error"`
}

type SuccessResponse struct {
	BaseResponse
}
