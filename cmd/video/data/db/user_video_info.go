package db

import (
	"golang.org/x/net/context"
	"tiktok/cmd/video/do"
	"tiktok/cmd/video/model"
	"tiktok/cmd/video/pack"
)

type UserVideoInfoDB struct {
	ctx context.Context
}

func NewUserVideoInfoDB(ctx context.Context) *UserVideoInfoDB {
	return &UserVideoInfoDB{
		ctx: ctx,
	}
}

func (this *UserVideoInfoDB) UsersInfo(users_id []int) ([]*do.UserVideoInfo, error) {

	var usermodels []model.UserVideoInfo
	result := DB.Where("user_id IN ?", users_id).Find(&usermodels)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return pack.Db2Do(usermodels), nil
}
