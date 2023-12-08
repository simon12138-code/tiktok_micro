package db

import (
	"context"
	"fmt"
	"testing"
	"tiktok/cmd/video/config"
)

func TestUsersInfo(t *testing.T) {
	//初始化函数
	config.InitConfig()
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	config.Settings.Mysqlinfo.Name, config.Settings.Mysqlinfo.Password, config.Settings.Mysqlinfo.Host,
	//	config.Settings.Mysqlinfo.Port, config.Settings.Mysqlinfo.DBName)
	InitDB()
	//开始测试
	users := []int{1, 2, 3, 4}
	res, err := NewUserVideoInfoDB(context.Background()).UsersInfo(users)
	if err != nil {
		panic(err)
	}
	for _, e := range res {
		fmt.Printf("%+v\n", *e)
	}

}
