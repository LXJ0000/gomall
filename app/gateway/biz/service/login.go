package service

import (
	"context"

	auth "github.com/LXJ0000/gomall/app/gateway/hertz_gen/gateway/auth"
	"github.com/LXJ0000/gomall/app/gateway/infra/rpc"
	"github.com/LXJ0000/gomall/app/gateway/utils"
	"github.com/LXJ0000/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp *auth.LoginResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	response, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	utils.MustHandleError(err)

	return &auth.LoginResp{
		UserId: response.UserId,
	}, nil
}
