package redis

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/go-redis/redis"
	"tiktok/cmd/user/config"
)

var Redis *redis.Client

func InitRedis() {
	//拼接redis地址
	addr := fmt.Sprintf("%s:%d", config.Settings.Redisinfo.Host, config.Settings.Redisinfo.Port)
	//生成redis客户端
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0, //使用默认数据库
	})
	//连接redis数据库
	_, err := Redis.Ping().Result()
	//打印错误
	if err != nil {
		color.Red("[InitRedis] 链接redis异常:")
		color.Yellow(err.Error())
	}
}
