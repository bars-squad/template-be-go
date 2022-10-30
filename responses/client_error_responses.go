package responses

// Maaping Client Error Responses https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#client_error_responses

func (hrsc *HttpResponseStatusCodesImpl) BadRequest(status string) *HttpResponse {
	return &HttpResponse{
		Code:   400,
		Status: SetStatus(status, "BAD_REQUEST"),
	}
}

func (hrsc *HttpResponseStatusCodesImpl) Unathorized(status string) *HttpResponse {
	return &HttpResponse{
		Code:   401,
		Status: SetStatus(status, "UNAUTHORIZED"),
	}
}

func (hrsc *HttpResponseStatusCodesImpl) Forbidden(status string) *HttpResponse {
	return &HttpResponse{
		Code:   403,
		Status: SetStatus(status, "FORBIDDEN"),
	}
}

func (hrsc *HttpResponseStatusCodesImpl) NotFound(status string) *HttpResponse {
	return &HttpResponse{
		Code:   404,
		Status: SetStatus(status, "NOT_FOUND"),
	}
}

func (hrsc *HttpResponseStatusCodesImpl) Conflict(status string) *HttpResponse {
	return &HttpResponse{
		Code:   409,
		Status: SetStatus(status, "CONFLICT"),
	}
}

func (hrsc *HttpResponseStatusCodesImpl) UnprocessableEntity(status string) *HttpResponse {
	return &HttpResponse{
		Code:   422,
		Status: SetStatus(status, "UNPROCESSABLE_ENTITY"),
	}
}
