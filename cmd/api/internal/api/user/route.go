package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"log"
	"tiktok-api/pkg/middleware"
	"tiktok-api/pkg/router"
)

// 实现路由接口，配置用户服务相关的路由
type RouterUser struct {
}

func init() {
	log.Println("init User router success")
	ru := &RouterUser{}
	router.Register(ru)
}

func (*RouterUser) Route(r *server.Hertz) {
	InitRpcUserClient()
	//初始化grpc的客户端连接
	h := NewHandlerUser()
	r.POST("/douyin/user/register/", h.Register)
	r.POST("/douyin/user/login/", h.Login)

	g := r.Group("", middleware.MustUser())
	{
		//g.GET("/douyin/user/", h.UserIndex)
	}

}
