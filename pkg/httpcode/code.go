package httpcode

type ErrCode struct {
	Code int
	Msg  string
}

// NewErrCode 实例化错误码
func NewErrCode(code int, msg string) ErrCode {
	return ErrCode{code, msg}
}

// UpMsg 自定义错误返回.
func (slf ErrCode) UpMsg(msg string) ErrCode {
	return NewErrCode(slf.Code, msg)
}

var (
	Code200       = NewErrCode(200, "Success")
	ParamErr      = NewErrCode(400, "param err")
	TokenNotValid = NewErrCode(401, "token invalid")
	TokenNotFound = NewErrCode(403, "token not found")
	TooManyReq    = NewErrCode(429, "too many request")
	ServiceErr    = NewErrCode(500, "service err")
)
