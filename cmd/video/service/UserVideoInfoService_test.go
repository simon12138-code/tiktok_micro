package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/video/config"
	"tiktok/cmd/video/data/db"
	"tiktok/pkg/errno"
)

func TestUsersVideoInfoService_UsersVideoInfo(t *testing.T) {
	config.InitConfig()
	db.InitDB()
	input1 := []int64{1, 2, 3, 5}
	input2 := []int64{1, 2, 3, 4}
	go func() {
		expectOutput := errno.NewErrNo(errno.ServiceErrCode, "some id not exists")
		res, err := NewUsersVideoInfoService(context.Background()).UsersVideoInfo(input1)
		assert.Equal(t, expectOutput, err)
		println(res)
	}()
	go func() {
		expectOutput := 4
		res, _ := NewUsersVideoInfoService(context.Background()).UsersVideoInfo(input2)
		assert.Equal(t, expectOutput, len(res))
		println(res)
	}()

}
