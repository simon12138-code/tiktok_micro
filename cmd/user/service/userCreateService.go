package service

import (
	"context"
	"github.com/bytedance/gopkg/lang/fastrand"
	"google.golang.org/protobuf/proto"
	"strconv"
	"tiktok/cmd/user/data/db"
	"tiktok/cmd/user/data/redis"
	"tiktok/cmd/user/do"
	"tiktok/cmd/user/pkg/jwt"
	"tiktok/cmd/user/pkg/protobuf"
	"tiktok/cmd/user/pkg/utils"
	"tiktok/kitex_gen/user"
	"tiktok/pkg/errno"
)

// 采用依赖倒置原则
type UserCreateDal interface {
	CheckUser(ctx context.Context, username string) (bool, error)
	UserCreate(user *do.User) error
}

type userCreateService struct {
	ctx context.Context
}

func NewUserCreateService(ctx context.Context) *userCreateService {
	return &userCreateService{ctx: ctx}
}
func (u *userCreateService) UserCreate(req *user.UserCreateRequest) (*do.User, string, error) {
	resp := &do.User{
		UserId: -1,
	}
	token := ""
	//先查redis
	res, err := redis.CheckUser(req.UserName)
	userCreateDB := &db.UserCreateDB{}
	if err != nil {
		return resp, token, errno.RedisErr
	}
	//如果redis中不存在
	if res == 0 {
		res, err := userCreateDB.CheckUser(u.ctx, req.UserName)
		if err != nil {
			return resp, token, errno.ConvertErr(err).WithMessage("db error:query")
		}
		if res == true {
			return resp, token, errno.UserAlreadyExistErr
		}
		//密码加密
		Password, err := utils.EncodePassword(req.UserPassword)
		if err != nil {
			return resp, token, errno.ConvertErr(err).WithMessage("encodeErr")
		}
		//先插数据库
		user := &do.User{
			UserName: req.UserName,
			PassWord: Password,
		}
		err = userCreateDB.UserCreate(user)
		if err != nil {
			return resp, token, errno.ConvertErr(err).WithMessage("db error:insert")
		}
		resp.UserId = user.UserId
		//jwt生成token
		token, err = jwt.CreateToken(user.UserId, user.UserName)
		if err != nil {
			return resp, token, errno.ConvertErr(err).WithMessage("service error:create token fail")
		}

		//插入redis,两张表都要插入，kv:username，userid，userid，userinf0，保证过期时间一致
		random := fastrand.Int()
		//插入用户信息index
		err = redis.SetUserIndex(user.UserName, strconv.Itoa(user.UserId), random)
		if err != nil {
			return resp, token, errno.RedisErr
		}
		protouser := &protobuf.User{UserId: int64(user.UserId), UserName: user.UserName}
		userinfo, err := proto.Marshal(protouser)
		if err != nil {
			return resp, token, errno.ConvertErr(err).WithMessage("protobuf error")
		}
		//插入用户信息全部内容
		err = redis.SetUser(strconv.Itoa(user.UserId), string(userinfo), random)
		if err != nil {
			return resp, token, errno.RedisErr
		}
	}
	return resp, token, errno.UserAlreadyExistErr
}
