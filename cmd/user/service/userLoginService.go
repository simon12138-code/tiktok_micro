package service

import (
	"context"
	"errors"
	redispkg "github.com/go-redis/redis"
	"google.golang.org/protobuf/proto"
	"tiktok/cmd/user/data/db"
	"tiktok/cmd/user/data/redis"
	"tiktok/cmd/user/do"
	"tiktok/cmd/user/pkg/jwt"
	"tiktok/cmd/user/pkg/protobuf"
	"tiktok/cmd/user/pkg/utils"
	"tiktok/pkg/errno"
)

type UserLoginService struct {
	ctx context.Context
}

func NewUserLoginService(ctx context.Context) *UserLoginService {
	return &UserLoginService{
		ctx: ctx,
	}
}

type UserLoginDal interface {
	GetUser(username string) (*do.User, error)
}

func (this *UserLoginService) UserLogin(userName string, userPassword string) (int, string, error) {
	//登录接口，校验参数->密码加密->redis_cache->db->return
	userid := -1
	userToken := ""
	var userinfo *do.User
	//密码加密
	Password, err := utils.EncodePassword(userPassword)
	if err != nil {
		return userid, userToken, errno.ConvertErr(err).WithMessage("encodeErr")
	}
	userId, err := redis.GetUserIndex(userName)
	if err != nil {
		if !errors.Is(err, redispkg.Nil) {
			return userid, userToken, errno.RedisErr
		} else {
			//cache miss,search db
			userinfo, err = db.NewUserLoginDB().GetUser(userName)
			if err != nil {
				return userid, userToken, errno.ConvertErr(err).WithMessage("db err")
			}
			//db内容不存在
			if userinfo == nil {
				return userid, userToken, errno.UserNotExist
			}
			if userinfo.PassWord != Password {
				return userid, userToken, errno.ServiceErr.WithMessage("password error")
			}
			//验证成功，生成token
			userToken, err = jwt.CreateToken(userinfo.UserId, userinfo.UserName)
			if err != nil {
				return userid, "", errno.ConvertErr(err).WithMessage("service error:create token fail")
			}
			//成功返回
			return userid, userToken, nil
		}

	}
	//缓存命中，然后从对应的次缓存中获取信息
	userString, err := redis.GetUser(userId)
	if err != nil {
		if !errors.Is(err, redispkg.Nil) {
			return userid, userToken, err
		}
		//次缓存未命中，db寻找
		userinfo, err = db.NewUserLoginDB().GetUser(userName)
		if err != nil {
			return userid, userToken, errno.ConvertErr(err).WithMessage("db err")
		}
		//db内容不存在
		if userinfo == nil {
			return userid, userToken, errno.UserNotExist
		}
		//验证成功，生成token
		userToken, err = jwt.CreateToken(userinfo.UserId, userinfo.UserName)
		if err != nil {
			return userid, "", errno.ConvertErr(err).WithMessage("service error:create token fail")
		}
		//成功返回
		return userid, userToken, nil
	}
	userProtoc := &protobuf.User{}
	err = proto.Unmarshal([]byte(userString), userProtoc)
	if err != nil {
		return userid, userToken, errno.ConvertErr(err).WithMessage("proto error")
	}
	if userinfo.PassWord != Password {
		return userid, userToken, errno.ServiceErr.WithMessage("password error")
	}
	//验证成功，生成token
	userToken, err = jwt.CreateToken(userinfo.UserId, userinfo.UserName)
	if err != nil {
		return userid, "", errno.ConvertErr(err).WithMessage("service error:create token fail")
	}
	//成功返回
	return userid, userToken, nil
}
