package responses

type Responses interface{}

// We can use this internal response to wrap response
type ResponsesImpl struct {
	Code       int         `json:"code"`
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	Pagination interface{} `json:"pagination,omitempty"`
}

func (hrsci *HttpResponseStatusCodesImpl) NewResponses(data any, message string) *ResponsesImpl {
	return &ResponsesImpl{
		Data:    data,
		Code:    hrsci.Code,
		Status:  hrsci.Status,
		Message: message,
	}
}

func (hrsci *HttpResponseStatusCodesImpl) NewResponsesWithPagination(data any, pagination any, message string) *ResponsesImpl {
	return &ResponsesImpl{
		Data:       data,
		Code:       hrsci.Code,
		Status:     hrsci.Status,
		Message:    message,
		Pagination: pagination,
	}
}
