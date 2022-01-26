package errorx

var (
	ApiNotFound       = NewErrorX(1000001, "api not found")
	ServiceNotFound   = NewErrorX(1000002, "service not found")
	GetServiceError   = NewErrorX(1000003, "get service error")
	ParseServiceError = NewErrorX(1000004, "parse service error")
	InvalidParamError = NewErrorX(1000004, "invalid param error")
)

type ErrorX struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func NewErrorX(code int32, message string) *ErrorX {

	return &ErrorX{
		Code:    code,
		Message: message,
	}
}

func (x ErrorX) Error() string {
	return x.Message
}
