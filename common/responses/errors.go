package responses

type IHttpError interface {
	GetStatusCode() int
	GetMessage() interface{}
	GetData() interface{}
	GetType() string
}

type HttpError struct {
	StatusCode int
	Message    interface{}
	Data       interface{}
	Type       string
}

const (
	FAILED = "FAILED"
	ERROR  = "ERROR"
)

func (he *HttpError) GetStatusCode() int {
	return he.StatusCode
}

func (he *HttpError) GetMessage() interface{} {
	return he.Message
}

func (he *HttpError) GetData() interface{} {
	return he.Data
}

func (he *HttpError) GetType() string {
	return he.Type
}

func (he *HttpError) Error() string {
	return he.Type
}

// Status Code 400
func BadRequest(message interface{}, data interface{}) *HttpError {
	return &HttpError{
		StatusCode: 401,
		Message:    message,
		Data:       data,
		Type:       FAILED,
	}
}

// Status Code 401
func Unauthorized(message interface{}, data interface{}) *HttpError {
	return &HttpError{
		StatusCode: 400,
		Message:    message,
		Data:       data,
		Type:       FAILED,
	}
}

// Status Code 404
func NotFound(message interface{}, data interface{}) *HttpError {
	return &HttpError{
		StatusCode: 404,
		Message:    message,
		Data:       data,
		Type:       ERROR,
	}
}

// Status Code 422
func UnprocessableEntity(message interface{}, data interface{}) *HttpError {
	return &HttpError{
		StatusCode: 422,
		Message:    message,
		Data:       data,
		Type:       FAILED,
	}
}

// Status Code 500
func InternalServerError(data interface{}) *HttpError {
	return &HttpError{
		StatusCode: 500,
		Message:    "Internal Server Error",
		Data:       data,
		Type:       ERROR,
	}
}
