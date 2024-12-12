package about

import (
	"context"

	"gomall/app/frontend/biz/service"
	"gomall/app/frontend/biz/utils"
	common "gomall/app/frontend/hertz_gen/frontend/common"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// About .
// @router /about [GET]
func About(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewAboutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	//c.HTML(consts.StatusOK, "about", utils.WrapResponse(ctx, c, resp))
	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}