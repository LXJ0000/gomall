package service

import (
	"context"

	common "github.com/LXJ0000/gomall/app/gateway/hertz_gen/common"
	auth "github.com/LXJ0000/gomall/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type SendSmsCodeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSendSmsCodeService(Context context.Context, RequestContext *app.RequestContext) *SendSmsCodeService {
	return &SendSmsCodeService{RequestContext: RequestContext, Context: Context}
}

func (h *SendSmsCodeService) Run(req *auth.SendSmsCodeReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
