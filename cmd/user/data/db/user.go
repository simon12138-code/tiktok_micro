// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package db

import (
	"context"
	"tiktok/cmd/user/config"
	"tiktok/cmd/user/do"
	"tiktok/cmd/user/pack"

	//"github.com/opentracing/opentracing-go"

	"gorm.io/gorm"
)

// model
type User struct {
	gorm.Model
	UserId          int    `json:"id"`
	Avatar          string `json:"avatar"`
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	FollowCount     int    `json:"follow_count"`
	FollowerCount   int    `json:"follower_count"`
}

func (u *User) TableName() string {
	return config.Settings.TableName.UserTableName
}

//
//// MGetUsers multiple get list of user info
//func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
//	res := make([]*User, 0)
//	if len(userIDs) == 0 {
//		return res, nil
//	}
//	//WithContext方法可以创建一个新的DBsession
//	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
//		return nil, err
//	}
//	return res, nil
//}
//
//// CreateUser create user info
//func CreateUser(ctx context.Context, users []*User) error {
//	//span, _ := opentracing.StartSpanFromContext(ctx, "createUser")
//	//defer span.Finish()
//	//返回DB封装的错误类型
//	return DB.WithContext(ctx).Create(users).Error
//}
//
//// QueryUser query list of user info
//func QueryUser(ctx context.Context, userName string) ([]*User, error) {
//	//span, _ := opentracing.StartSpanFromContext(ctx, "QueryUser")
//	//defer span.Finish()
//	res := make([]*User, 0)
//	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
//		return nil, err
//	}
//	return res, nil
//}

type UserCreateDB struct {
}

func (this *UserCreateDB) CheckUser(ctx context.Context, userName string) (bool, error) {
	var user User
	rows := DB.WithContext(ctx).Where(&User{UserName: userName}).Find(&user)
	if rows.Error != nil {
		return false, rows.Error
	}
	if rows.RowsAffected < 1 {
		return false, nil
	}
	return true, nil
}

func (this *UserCreateDB) UserCreate(user *do.User) error {
	usermodel := &User{
		UserName: user.UserName,
		Password: user.PassWord,
	}
	result := DB.Create(&usermodel)
	if result.RowsAffected < 1 || result.Error != nil {
		return result.Error
	}
	user.UserId = usermodel.UserId
	return nil
}

type UserInfoDB struct {
}

func NewUserInfoDB() *UserInfoDB {
	return &UserInfoDB{}
}
func (this *UserInfoDB) UserInfo(id int) (*do.User, error) {
	usermodel := &User{
		UserId: id,
	}
	result := DB.Where(&User{UserId: id}).First(&usermodel)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return pack.User(usermodel), nil
}

type UserLoginDB struct {
}

func NewUserLoginDB() *UserLoginDB {
	return &UserLoginDB{}
}
func (this *UserLoginDB) GetUser(username string) (*do.User, error) {
	usermodel := &User{
		UserName: username,
	}
	result := DB.Where(&User{UserName: username}).First(&usermodel)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return pack.User(usermodel), nil
}
