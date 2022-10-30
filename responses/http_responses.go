package responses

// You can say this file is a main of responses folder

// HTTP response status codes indicate whether a specific HTTP request has been successfully completed. Responses are grouped in five classes. Please read these article https://developer.mozilla.org/en-US/docs/Web/HTTP/Status

type HttpResponse struct {
	Code   int
	Status string
}

type HttpResponseStatusCodes interface {
	Ok(status string) HttpResponse
	Created(status string) HttpResponse
	BadRequest(status string) HttpResponse
	Unauthorized(status string) HttpResponse
	Forbidden(status string) HttpResponse
	NotFound(status string) HttpResponse
	Conflict(status string) HttpResponse
	UnprocessableEntity(status string) HttpResponse
}

type HttpResponseStatusCodesImpl struct{}

func SetStatus(status string, defaultStatus string) string {
	if status != "" {
		return status
	}
	return defaultStatus
}
