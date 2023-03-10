package httpcode

import "time"

// http request
const (
	CallTimeOut  = 30 * time.Second // 请求超时
	MaxPage      = 100              // 最大请求页数
	MaxSize      = 100              // 最大请求条数
	CtxStartTime = "startTime"      // 运行起始时间
	CtxRequestId = "trace-id"       // 请求唯一id
)
