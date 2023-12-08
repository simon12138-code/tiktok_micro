package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// 提供一个抽象的接口，各个不同的包去实现具体的路由启动，注册到全局变量中，然后统一启动
var routers []Router

type Router interface {
	Route(r *server.Hertz)
}

func Register(root ...Router) {
	routers = append(routers, root...)
}

func InitRouter(r *server.Hertz) {
	r.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{"msg": "pong"})
	})
	for _, root := range routers {
		root.Route(r)
	}
}
