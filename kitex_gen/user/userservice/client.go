// Code generated by Kitex v0.6.2. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "tiktok/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UserInfo(ctx context.Context, req *user.UserInfoRequest, callOptions ...callopt.Option) (r *user.UserInfoResponse, err error)
	UserCreate(ctx context.Context, req *user.UserCreateRequest, callOptions ...callopt.Option) (r *user.UserCreateResponse, err error)
	UserLogin(ctx context.Context, req *user.UserLoginRequest, callOptions ...callopt.Option) (r *user.UserLoginRequest, err error)
	FriendList(ctx context.Context, req *user.FriendListRequest, callOptions ...callopt.Option) (r *user.FriendListResponse, err error)
	FollowerList(ctx context.Context, req *user.FriendListRequest, callOptions ...callopt.Option) (r *user.FollowerListResponse, err error)
	FollowList(ctx context.Context, req *user.FollowListRequest, callOptions ...callopt.Option) (r *user.FollowListResponse, err error)
	RelationAction(ctx context.Context, req *user.RelationActionRequest, callOptions ...callopt.Option) (r *user.RelationActionResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) UserInfo(ctx context.Context, req *user.UserInfoRequest, callOptions ...callopt.Option) (r *user.UserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserInfo(ctx, req)
}

func (p *kUserServiceClient) UserCreate(ctx context.Context, req *user.UserCreateRequest, callOptions ...callopt.Option) (r *user.UserCreateResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserCreate(ctx, req)
}

func (p *kUserServiceClient) UserLogin(ctx context.Context, req *user.UserLoginRequest, callOptions ...callopt.Option) (r *user.UserLoginRequest, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserLogin(ctx, req)
}

func (p *kUserServiceClient) FriendList(ctx context.Context, req *user.FriendListRequest, callOptions ...callopt.Option) (r *user.FriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FriendList(ctx, req)
}

func (p *kUserServiceClient) FollowerList(ctx context.Context, req *user.FriendListRequest, callOptions ...callopt.Option) (r *user.FollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowerList(ctx, req)
}

func (p *kUserServiceClient) FollowList(ctx context.Context, req *user.FollowListRequest, callOptions ...callopt.Option) (r *user.FollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowList(ctx, req)
}

func (p *kUserServiceClient) RelationAction(ctx context.Context, req *user.RelationActionRequest, callOptions ...callopt.Option) (r *user.RelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, req)
}
