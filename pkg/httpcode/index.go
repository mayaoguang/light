package httpcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"light/pkg/logging"
	"reflect"
	"time"

	"github.com/kataras/iris/v12"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/validator.v2"
)

type Req struct {
	ctx       iris.Context
	body      []byte
	requestId string
	Log       logging.Encoder
}

// Ok 正确的json返回.
func (slf *Req) Ok(data interface{}) {
	slf.Code(Code200, nil, data)
}

// ParamError json返回参数错误.
func (slf *Req) ParamError(err error) {
	slf.Code(ParamErr, err, nil)
}

// ServiceError 通用错误处理.
func (slf *Req) ServiceError(err error) {
	slf.Code(ServiceErr, err, nil)
}

// Code 自定义code码返回.
func (slf *Req) Code(code ErrCode, err error, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	if err != nil {
		slf.Log.Warnf("api: %s, param: %s, err: %v", slf.ctx.Request().RequestURI, slf.body, err)
	}
	var runTime string
	if startTime, ok := slf.ctx.Values().Get(CtxStartTime).(time.Time); ok {
		runTime = time.Now().Sub(startTime).String()
	}

	slf.ctx.Header(CtxRequestId, slf.requestId)
	_ = slf.ctx.JSON(map[string]interface{}{"code": code.Code, "message": code.Msg, "run": runTime, "data": data})
}

// NewRequest 解析post传参.
func NewRequest(ctx iris.Context, params interface{}) (r *Req, err error) {
	uid := uuid.NewV4().String()
	r = &Req{
		ctx:       ctx,
		requestId: uid,
		Log:       logging.GetEncoder().WithField(CtxRequestId, uid),
	}
	if params != nil {
		defer func() {
			if err != nil {
				r.ParamError(err)
			}
		}()

		body, err := ctx.GetBody()
		if err != nil {
			return r, fmt.Errorf("api: %s GetBody err: %v", ctx.Request().RequestURI, err)
		}
		if len(body) > 0 {
			if err = json.Unmarshal(body, params); err != nil {
				return r, fmt.Errorf("api: %s Unmarshal err: %v, body: %s", ctx.Request().RequestURI, err, body)
			}
			r.body = body
		}
		// 页数限制
		if page := reflect.Indirect(reflect.ValueOf(params)).FieldByName("PageNum"); page.IsValid() && page.Int() > MaxPage {
			return r, fmt.Errorf("page num greater than limit")
		}
		// 条数限制
		if size := reflect.Indirect(reflect.ValueOf(params)).FieldByName("PageSize"); size.IsValid() && size.Int() > MaxSize {
			return r, fmt.Errorf("page size greater than limit ")
		}
		// 参数校验
		if err = validator.Validate(params); err != nil {
			return r, fmt.Errorf("validate param err: %v", err)
		}
	}
	return
}

// ToExcel 数据导出excel.
func (slf *Req) ToExcel(titleList []string, dataList interface{}, fileName string) (err error) {
	buf, err := ExportExcel(titleList, dataList)
	if err != nil {
		return
	}
	content := bytes.NewReader(buf.Bytes())
	slf.ctx.ServeContent(content, fileName, time.Now())
	return nil
}

// ToSecondaryTitleExcel 导出二级标题.
func (slf *Req) ToSecondaryTitleExcel(firstTitle []string, secondTitle [][]string, dataList interface{}, fileName string) (err error) {
	buf, err := ExportSecondaryTitleExcel(firstTitle, secondTitle, dataList)
	if err != nil {
		return
	}
	content := bytes.NewReader(buf.Bytes())
	slf.ctx.ServeContent(content, fileName, time.Now())
	return nil
}
