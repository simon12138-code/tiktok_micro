package service

import (
	"context"

	"tiktok/cmd/video/data/db"

	"tiktok/cmd/video/do"
	"tiktok/pkg/errno"
)

type UsersVideoInfoService struct {
	ctx context.Context
}

func NewUsersVideoInfoService(ctx context.Context) *UsersVideoInfoService {
	return &UsersVideoInfoService{
		ctx: ctx,
	}
}

type UsersVideoInfoDal interface {
	GetUser(user_id int) (*do.UserVideoInfo, error)
}

func (this *UsersVideoInfoService) UsersVideoInfo(user_ids []int64) ([]*do.UserVideoInfo, error) {
	//先查缓存
	resp := []*do.UserVideoInfo{}
	users_id := make([]int, len(user_ids))
	for i, e := range user_ids {
		users_id[i] = int(e)
	}
	//usersInfo, err := redis.GetUserVideoInfo(users_id)
	//redis查询成功
	//if err != nil {
	//if !errors.Is(err, redispkg.Nil) {
	//	//redis服务错误直接返回
	//	return resp, errno.RedisErr
	//}
	// cache miss
	//搜索db
	//users_id_int := make([]int,len(usersInfo))
	//for i,e :=range	usersInfo{
	//	n,_ := strconv.Atoi(e)
	//	users_id_int[i] = n
	//}
	//操作数据库进行查询
	usersDo, err := db.NewUserVideoInfoDB(context.Background()).UsersInfo(users_id)
	if err != nil {
		return resp, errno.ConvertErr(err).WithMessage("db error")
	}
	if len(usersDo) != len(user_ids) {
		return resp, errno.NewErrNo(errno.ServiceErrCode, "some id not exists")
	}
	//查询成功
	return usersDo, nil
}
