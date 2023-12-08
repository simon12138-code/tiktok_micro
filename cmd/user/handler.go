package main

import (
	"context"
	"tiktok/cmd/user/pack"
	"tiktok/cmd/user/service"
	user "tiktok/kitex_gen/user"
	"tiktok/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: 首先要鉴权，然后查询对应的用户信息，查到就返回，cachemiss就db然后cache-rebuild
	resp = new(user.UserInfoResponse)
	userInfo, err := service.NewUserInfoService(ctx).UserInfo(req.Token, req.UserId)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.User = pack.Do2User(userInfo)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	//调用VideoRPC服务
	return resp, nil
}

// UserCreate implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserCreate(ctx context.Context, req *user.UserCreateRequest) (resp *user.UserCreateResponse, err error) {
	// TODO: 组装业务逻辑，用户创建需要JWT生成，一次db，所以业务逻辑层无其他交互，那么直接调用service接口
	resp = new(user.UserCreateResponse)
	if len(req.UserName) == 0 || len(req.UserPassword) == 0 || len(req.UserName) > 32 || len(req.UserPassword) > 32 {
		//调用封装数据方法，封装参数类型错误
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	//每次调用业务代码会新建一个ctx,调用业务代码
	doUser, token, err := service.NewUserCreateService(ctx).UserCreate(req)
	//调用videoRPC服务，插入对应的user的favorite_count
	//错误处理
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	//封装统一返回
	resp.UserId = int64(doUser.UserId)
	resp.Token = token
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil

}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: 登录接口，校验参数->redis_cache->db->return
	if len(req.UserName) == 0 || len(req.UserPassword) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	id, token, err := service.NewUserLoginService(ctx).UserLogin(req.UserName, req.UserPassword)
	if err != nil {
		resp.UserId = -1
		resp.Token = ""
		pack.BuildBaseResp(err)
		return resp, err
	}
	//包装返回
	resp.UserId = int64(id)
	resp.Token = token
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// FriendList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FriendList(ctx context.Context, req *user.FriendListRequest) (resp *user.FriendListResponse, err error) {
	// TODO:
	return
}

// FollowerList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowerList(ctx context.Context, req *user.FriendListRequest) (resp *user.FollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowList(ctx context.Context, req *user.FollowListRequest) (resp *user.FollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationAction implements the UserServiceImpl interface.
func (s *UserServiceImpl) RelationAction(ctx context.Context, req *user.RelationActionRequest) (resp *user.RelationActionResponse, err error) {
	// TODO: Your code here...
	return
}
