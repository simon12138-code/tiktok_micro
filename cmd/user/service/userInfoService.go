package service

import (
	"context"
	"errors"
	"github.com/bytedance/gopkg/lang/fastrand"
	redispkg "github.com/go-redis/redis"
	"google.golang.org/protobuf/proto"
	"strconv"
	"tiktok/cmd/user/data/db"
	"tiktok/cmd/user/data/redis"
	"tiktok/cmd/user/do"
	"tiktok/cmd/user/pack"
	"tiktok/cmd/user/pkg/jwt"
	"tiktok/cmd/user/pkg/protobuf"
	"tiktok/pkg/errno"
)

type UserInfoDal interface {
	userInfo(id int) (*do.User, error)
}

type userInfoService struct {
	ctx context.Context
}

func NewUserInfoService(ctx context.Context) *userInfoService {
	return &userInfoService{ctx: ctx}
}

func (this *userInfoService) UserInfo(token string, id string) (*do.User, error) {
	//首先要鉴权，然后查询对应的用户信息，查到就返回，cachemiss就db然后cache-rebuild
	resp := new(do.User)
	if token == "" {
		return resp, errno.AuthorizationFailedErr
	}
	//创建一个新的验证key
	j := jwt.NewJWT()
	//解析token
	_, err := j.ParseToken(token)
	if err != nil {
		if err == jwt.TokenExpired {
			//token过期
			return resp, errno.NewErrNo(errno.TokenErrCode, "授权已过期")
		}
		//其他错误
		return resp, errno.NewErrNo(errno.TokenErrCode, "其他错误")
	}
	//
	//redis查询对应的
	res, err := redis.CheckUser(id)
	if err != nil {
		return resp, errno.RedisErr
	}
	//db
	if res == 0 {
		//cache_miss
		//db
		idint, err := strconv.Atoi(id)
		if err != nil {
			return resp, errno.ConvertErr(err).WithMessage("strconv Err")
		}
		userinfo, err := db.NewUserInfoDB().UserInfo(idint)
		if err != nil {
			return resp, errno.ConvertErr(err).WithMessage("db error:insert")
		}
		if userinfo == nil {
			return resp, errno.ConvertErr(err).WithMessage("user doesn't exist")
		}
		//cache-rebuild
		protouser := pack.Do2Protoc(userinfo)

		random := fastrand.Int()
		//插入用户信息index
		err = redis.SetUserIndex(userinfo.UserName, strconv.Itoa(userinfo.UserId), random)
		if err != nil {
			return resp, errno.RedisErr
		}
		userstring, err := proto.Marshal(protouser)
		if err != nil {
			return resp, errno.ConvertErr(err).WithMessage("protobuf error")
		}
		err = redis.SetUser(strconv.Itoa(userinfo.UserId), string(userstring), random)
		if err != nil {
			return resp, errno.RedisErr
		}
	}
	//fisrt cache hit，second cache
	userString, err := redis.GetUser(id)
	if err != nil {
		//error
		if !errors.Is(err, redispkg.Nil) {
			return resp, errno.RedisErr.WithMessage("user not in redis")
		} else {
			//cache miss
			userInfoDB := &db.UserInfoDB{}
			idint, err := strconv.Atoi(id)
			if err != nil {
				return resp, errno.ConvertErr(err).WithMessage("strconv Err")
			}
			userinfo, err := userInfoDB.UserInfo(idint)
			if err != nil {
				return resp, errno.ConvertErr(err).WithMessage("db error:insert")
			}
			if userinfo == nil {
				return resp, errno.UserNotExist
			}
			//cache-rebuild
			protouser := pack.Do2Protoc(userinfo)

			random := fastrand.Int()
			//插入用户信息index
			err = redis.SetUserIndex(userinfo.UserName, strconv.Itoa(userinfo.UserId), random)
			if err != nil {
				return resp, errno.RedisErr
			}
			userstring, err := proto.Marshal(protouser)
			if err != nil {
				return resp, errno.ConvertErr(err).WithMessage("protobuf error")
			}
			err = redis.SetUser(strconv.Itoa(userinfo.UserId), string(userstring), random)
			if err != nil {
				return resp, errno.RedisErr
			}
		}
	}
	//Unmarshal
	userinfo := &protobuf.User{}
	err = proto.Unmarshal([]byte(userString), userinfo)
	if err != nil {
		return resp, errno.ConvertErr(err).WithMessage("protobuf error")
	}
	resp = pack.Protoc2Do(userinfo)
	return resp, nil
}
