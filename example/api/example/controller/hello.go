package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"light/example/api/example/model"
	"light/pkg/httpcode"
)

// Hello 请求示例.
func Hello(ctx iris.Context) {
	req := &model.HelloReq{}
	r, err := httpcode.NewRequest(ctx, req)
	if err != nil {
		return
	}
	res := model.HelloRes{
		Name: fmt.Sprintf("hello, %s", req.Name),
	}
	r.Log.Info(res.Name)
	r.Ok(res)
}

// Export 导出示例.
func Export(ctx iris.Context) {
	r, err := httpcode.NewRequest(ctx, nil)
	if err != nil {
		return
	}
	dataList := []*model.ExportRes{
		{1, "1", 2},
		{2, "3", 4},
	}
	r.ToExcel([]string{"Id", "年龄"}, dataList, "test")
}
