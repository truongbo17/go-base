package response

type BaseResponse struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"statusCode"`
	RequestId  string      `json:"requestId,omitempty"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}
