package responses

type IHttpSuccess interface {
	GetStatusCode() int
	GetMessage() interface{}
	GetData() interface{}
	GetType() string
}

type HttpSuccess struct {
	StatusCode int
	Data       interface{}
	Type       string
}

const (
	SUCCESS = "SUCCESS"
)

func (he *HttpSuccess) GetStatusCode() int {
	return he.StatusCode
}

func (he *HttpSuccess) GetData() interface{} {
	return he.Data
}

func (he *HttpSuccess) GetType() string {
	return he.Type
}

func (he *HttpSuccess) Error() string {
	return he.Type
}

// Status Code 200
func Success(data interface{}) *HttpSuccess {
	return &HttpSuccess{
		StatusCode: 200,
		Data:       data,
		Type:       SUCCESS,
	}
}

// Status Code 201
func Created(data interface{}) *HttpSuccess {
	return &HttpSuccess{
		StatusCode: 201,
		Data:       data,
		Type:       SUCCESS,
	}
}
