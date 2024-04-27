package responses

type Response struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   error  `json:"error,omitempty"`
}

func NewResponse(message string, data any, err error) *Response {
	return &Response{Message: message, Data: data, Error: err}
}
