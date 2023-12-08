package main

import (
	"context"
	"tiktok/cmd/video/pack"
	"tiktok/cmd/video/service"
	video "tiktok/kitex_gen/video"
	"tiktok/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// UserVideoInfo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UserVideoInfo(ctx context.Context, req *video.UsersVideoInfoRequest) (resp *video.UsersVideoInfoResponse, err error) {
	// TODO: 根据用户id查找UserVideoInfo，先查redis，再查db，如果都不存在报错，如果cache miss就rebuild
	// 第一版只查数据库
	res, err := service.NewUsersVideoInfoService(ctx).UsersVideoInfo(req.UserId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		resp.UsersVideoInfo = []*video.UserVideoInfo{}
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UsersVideoInfo = pack.Do2DTO(res)
	return resp, nil
}
