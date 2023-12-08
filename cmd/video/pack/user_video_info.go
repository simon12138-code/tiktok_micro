package pack

import (
	"tiktok/cmd/video/do"
	"tiktok/cmd/video/model"
	"tiktok/kitex_gen/video"
)

// pack 完成model 到 service user的转变（定义了两个结构体，类似gin中的form与model）
// User pack user info
func Db2Do(u []model.UserVideoInfo) []*do.UserVideoInfo {
	if u == nil {
		return nil
	}
	res := make([]*do.UserVideoInfo, len(u))
	for i, e := range u {
		res[i] = &do.UserVideoInfo{
			UserId:         e.UserId,
			FavoriteCount:  e.FavoriteCount,
			FavoritedCount: e.FavoritedCount,
			WorkCount:      e.WorkCount,
		}
	}
	return res
}

func Do2DTO(doUsersInfo []*do.UserVideoInfo) []*video.UserVideoInfo {
	dToUsersInfo := make([]*video.UserVideoInfo, len(doUsersInfo))
	for i, e := range doUsersInfo {
		dToUsersInfo[i].UserId = int64(e.UserId)
		dToUsersInfo[i].FavoritedCount = int64(e.FavoritedCount)
		dToUsersInfo[i].FavoriteCount = int64(e.FavoriteCount)
		dToUsersInfo[i].WorkCount = int64(e.WorkCount)
	}
	return dToUsersInfo
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
