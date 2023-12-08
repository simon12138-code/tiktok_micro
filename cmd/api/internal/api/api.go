package api

import (
	"tiktok/kitex_gen/chat/chatservice"
	"tiktok/kitex_gen/comment/commentservice"
	"tiktok/kitex_gen/favor/favorservice"

	"tiktok/kitex_gen/interaction/interactionservice"
	"tiktok/kitex_gen/user/userservice"
	"tiktok/kitex_gen/video/videoservice"
)

var (
	UserClient        userservice.Client
	InteractionClient interactionservice.Client
	VideoClient       videoservice.Client
	ChatClient        chatservice.Client
	FavorClient       favorservice.Client
	Comment           commentservice.Client
)
