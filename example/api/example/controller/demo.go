package controller

import (
	"github.com/kataras/iris/v12"
	"light/example/api/example/model"
	"light/pkg/httpcode"
)

// DemoCreate demo创建.
func (slf *DemoController) DemoCreate(ctx iris.Context) {
	var (
		req    = &model.DemoCreateReq{}
		r, err = httpcode.NewRequest(ctx, req)
	)

	if err != nil {
		return
	}

	if err = demoService.DemoCreate(req); err != nil {
		r.ServiceError(err)
		return
	}

	r.Ok(nil)
}

// DemoRecords demo列表.
func (slf *DemoController) DemoRecords(ctx iris.Context) {
	var (
		req    = &model.DemoRecordReq{}
		res    = &model.DemoRecordResp{List: []*model.DemoRecordItem{}}
		r, err = httpcode.NewRequest(ctx, req)
	)

	if err != nil {
		return
	}

	if err = demoService.DemoRecords(req, res); err != nil {
		r.ServiceError(err)
		return
	}

	r.Ok(res)
}

// DemoInfo demo详情.
func (slf *DemoController) DemoInfo(ctx iris.Context) {
	var (
		req    = &model.DemoInfoReq{}
		res    = &model.DemoInfoResp{}
		r, err = httpcode.NewRequest(ctx, req)
	)

	if err != nil {
		return
	}

	if err = demoService.DemoInfo(req, res); err != nil {
		r.ServiceError(err)
		return
	}

	r.Ok(res)
}

// DemoUpdate demo更新.
func (slf *DemoController) DemoUpdate(ctx iris.Context) {
	var (
		req    = &model.DemoUpdateReq{}
		r, err = httpcode.NewRequest(ctx, req)
	)

	if err != nil {
		return
	}

	if code, err := demoService.DemoUpdate(req); err != nil {
		r.Code(code, err, nil)
		return
	}

	r.Ok(nil)
}

// DemoDelete demo删除.
func (slf *DemoController) DemoDelete(ctx iris.Context) {
	var (
		req    = &model.DemoDeleteReq{}
		r, err = httpcode.NewRequest(ctx, req)
	)

	if err != nil {
		return
	}

	if err = demoService.DemoDelete(req); err != nil {
		r.ServiceError(err)
		return
	}

	r.Ok(nil)
}
