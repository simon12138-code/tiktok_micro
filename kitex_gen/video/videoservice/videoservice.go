// Code generated by Kitex v0.6.2. DO NOT EDIT.

package videoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	video "tiktok/kitex_gen/video"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"UserVideoInfo": kitex.NewMethodInfo(userVideoInfoHandler, newVideoServiceUserVideoInfoArgs, newVideoServiceUserVideoInfoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func userVideoInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceUserVideoInfoArgs)
	realResult := result.(*video.VideoServiceUserVideoInfoResult)
	success, err := handler.(video.VideoService).UserVideoInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceUserVideoInfoArgs() interface{} {
	return video.NewVideoServiceUserVideoInfoArgs()
}

func newVideoServiceUserVideoInfoResult() interface{} {
	return video.NewVideoServiceUserVideoInfoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UserVideoInfo(ctx context.Context, req *video.UsersVideoInfoRequest) (r *video.UsersVideoInfoResponse, err error) {
	var _args video.VideoServiceUserVideoInfoArgs
	_args.Req = req
	var _result video.VideoServiceUserVideoInfoResult
	if err = p.c.Call(ctx, "UserVideoInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
