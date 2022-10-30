package responses

// Mapping Successful Responses https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#successful_responses

func (hrsc *HttpResponse) Ok(status string) *HttpResponse {
	return &HttpResponse{
		Code:   200,
		Status: SetStatus(status, "OK"),
	}
}

func (hrsc *HttpResponseStatusCodesImpl) Created(status string) *HttpResponse {
	return &HttpResponse{
		Code:   201,
		Status: SetStatus(status, "CREATED"),
	}
}
