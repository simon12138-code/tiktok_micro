package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jinzhu/copier"
	"tiktok-api/internal/api"
	"tiktok-api/internal/model/request"
	"tiktok-api/internal/model/response"
	"tiktok-api/pkg/myerr"
	"tiktok-common/errcode"
	"tiktok/kitex_gen/user"

	res2 "tiktok-common/response"
)

type HandlerUser struct {
}

func NewHandlerUser() *HandlerUser {
	return &HandlerUser{}
}

func (h *HandlerUser) Register(ctx context.Context, c *app.RequestContext) {
	res := res2.NewResponse(c)
	// 1.接收参数 参数模型
	req := &request.UserCreateRequest{}
	if err := c.Bind(req); err != nil {
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	// 2.校验参数 判断参数是否合法
	if err := req.Verify(); err != nil {
		res.Reply(myerr.ErrUserNameORPassWord)
		return
	}
	// 3.调用user rpc服务 获取响应
	params := &user.UserCreateRequest{}
	_ = copier.Copy(params, req)
	result, err := api.UserClient.UserCreate(ctx, params)
	if err != nil {
		res.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	// 4.返回结果
	resp := &response.UserCreateResponse{}
	_ = copier.Copy(resp, result)
	res.Reply(nil, resp)
}

func (h *HandlerUser) Login(ctx context.Context, c *app.RequestContext) {
	res := res2.NewResponse(c)
	// 1.接收参数 参数模型
	req := &request.UserLoginRequest{}
	if err := c.Bind(req); err != nil {
		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	// 2.校验参数 判断参数是否合法
	if err := req.Verify(); err != nil {
		res.Reply(myerr.ErrUserNameORPassWord)
		return
	}
	// 3.调用user rpc服务 获取响应
	params := &user.UserLoginRequest{}
	//调包进行do转换
	_ = copier.Copy(params, req)
	result, err := api.UserClient.UserLogin(ctx, params)
	if err != nil {
		res.Reply(errcode.ErrServer.WithDetails(err.Error()))
		return
	}
	// 4.返回结果
	resp := &response.LoginResponse{}
	_ = copier.Copy(resp, result)
	res.Reply(nil, resp)
}

//func (h *HandlerUser) UserIndex(ctx context.Context, c *app.RequestContext) {
//	res := res2.NewResponse(c)
//	// 1.接收参数 参数模型
//	req := &request.UserIndexRequest{}
//	if err := c.Bind(req); err != nil {
//		res.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
//		return
//	}
//	content, ok := token.GetTokenContent(c)
//	if !ok {
//		res.Reply(errcode.ErrServer)
//		return
//	}
//	req.MyUserId = content.ID
//	// 2.调用user rpc服务 获取响应
//	params := &user.UserIndexRequest{}
//	_ = copier.Copy(params, req)
//
//	result, err := api.UserClient.UserIndex(ctx, params)
//	if err != nil {
//		res.Reply(errcode.ErrServer.WithDetails(err.Error()))
//		return
//	}
//	// 4.返回结果
//	resp := &response.UserIndexResponse{}
//	_ = copier.Copy(resp, result)
//	res.Reply(nil, resp)
//}
