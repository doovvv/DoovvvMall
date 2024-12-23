package service

import (
	"context"
	common "gomall/app/frontend/hertz_gen/frontend/common"
	"gomall/app/frontend/infra/rpc"
	rpcproduct "gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	// var resp = make(map[string]any)
	// items := []map[string]any{
	// 	{"Name": "bag", "Price": 100, "Picture": "static/image/bag.png"},
	// 	{"Name": "pen", "Price": 100, "Picture": "static/image/pen.png"},
	// }
	// resp["Items"] = items
	// resp["Title"] = "SYSU 100th mall"
	products, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"Items": products.Products,
		"Title": "Hot sale",
	}, nil
}
