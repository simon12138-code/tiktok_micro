package model

import (
	"tiktok/cmd/video/config"
)

// model
type UserVideoInfo struct {
	//gorm.Model
	UserId         int `json:"user_id" gorm:"primaryKey"`
	FavoriteCount  int `json:"favorite_count"`
	FavoritedCount int `json:"favorited_count"`
	WorkCount      int `json:"work_count"`
}

func (u *UserVideoInfo) TableName() string {
	return config.Settings.TableName.TableName
}
