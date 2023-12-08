package pack

import (
	"tiktok/cmd/user/data/db"
	"tiktok/cmd/user/do"
	"tiktok/cmd/user/pkg/protobuf"
	"tiktok/kitex_gen/user"
)

// pack 完成model 到 service user的转变（定义了两个结构体，类似gin中的form与model）
// User pack user info
func User(u *db.User) *do.User {
	if u == nil {
		return nil
	}
	return &do.User{
		UserId:          u.UserId,
		UserName:        u.UserName,
		Avatar:          u.Avatar,
		FollowerCount:   u.FollowerCount,
		FollowCount:     u.FollowCount,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
	}
}

func Protoc2Do(u *protobuf.User) *do.User {
	if u == nil {
		return nil
	}
	return &do.User{
		UserId:          int(u.UserId),
		UserName:        u.UserName,
		Avatar:          u.Avatar,
		FollowerCount:   int(u.FollowerCount),
		FollowCount:     int(u.FollowCount),
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
	}
}

func Do2Protoc(userinfo *do.User) *protobuf.User {
	return &protobuf.User{
		UserId:          int64(userinfo.UserId),
		UserName:        userinfo.UserName,
		Signature:       userinfo.Signature,
		Avatar:          userinfo.Signature,
		BackgroundImage: userinfo.BackgroundImage,
	}
}

func Do2User(doUser *do.User) *user.User {
	return &user.User{
		Id:              int64(doUser.UserId),
		UserName:        doUser.UserName,
		FollowCount:     int64(doUser.FollowCount),
		FollowerCount:   int64(doUser.FollowerCount),
		BackgroundImage: doUser.BackgroundImage,
	}

}

func Protoc2User(protocUser *protobuf.User) *user.User {
	return &user.User{
		Id:              protocUser.UserId,
		UserName:        protocUser.UserName,
		Avatar:          protocUser.Avatar,
		BackgroundImage: protocUser.BackgroundImage,
		Signature:       protocUser.Signature,
		FollowCount:     protocUser.FollowCount,
		FollowerCount:   protocUser.FollowerCount,
	}
}

//// Users pack list of user info
//func Users(us []*db.User) []*user.User {
//	users := make([]*user.User, 0)
//	for _, u := range us {
//		if user2 := User(u); user2 != nil {
//			users = append(users, user2)
//		}
//	}
//	return users
//}
