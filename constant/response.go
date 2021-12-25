package constant

const (
	ResponseInternalServerErr = "server unavailable"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type WrappedResponse struct {
	Base *Response   `json:"base"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
	}
}

func NewWrapSuccessResponse(v interface{}) *WrappedResponse {
	return NewWrapResponse(0, "success", v)
}

func NewWrapFailResponse(v interface{}) *WrappedResponse {
	return NewWrapResponse(-1, "fail", v)
}

func NewWrapResponse(code int, msg string, v interface{}) *WrappedResponse {
	return &WrappedResponse{
		Base: &Response{
			Code: code,
			Msg:  msg,
		},
		Data: v,
	}
}
