package response

type BaseResponse struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	RequestId  string      `json:"request_id,omitempty"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}
