package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"tiktok/cmd/user/config"
	"tiktok/kitex_gen/video"
	"tiktok/kitex_gen/video/videoservice"
	"time"
)

var videoClient videoservice.Client

func initVideoRpc() {
	etcdAddress := fmt.Sprintf("%s:%d", config.Settings.Etcdinfo.Host, config.Settings.Etcdinfo.Port)
	r, err := etcd.NewEtcdResolver([]string{etcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		config.Settings.RPCinfo[0].Name,
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

//func CreateUser(ctx context.Context, req *user.CreateUserRequest) error {
//	resp, err := userClient.CreateUser(ctx, req)
//	if err != nil {
//		return err
//	}
//	if resp.BaseResp.StatusCode != 0 {
//		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
//	}
//	return nil
//}

func GetVideoInfo(ctx context.Context, req *video.UsersVideoInfoRequest) ([]*video.UserVideoInfo, error) {
	resp, err := videoClient.UserVideoInfo(ctx, req)
	if err != nil {
		return []*video.UserVideoInfo{}, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return []*video.UserVideoInfo{}, err
	}
	return resp.UsersVideoInfo, nil
}
