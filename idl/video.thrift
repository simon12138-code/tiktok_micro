namespace go video

struct BaseResp {
  1:i64 status_code;
  2:string status_message;
  3:i64 service_time;
}

struct UserVideoInfo{
    1:i64 user_id;
    2:i64 favorite_count;
    3:i64 favorited_count;
    4:i64 work_count;
}

struct UsersVideoInfoRequest{
    1:list<i64> user_id;
}

struct UsersVideoInfoResponse{
    1:BaseResp base_resp;
    2:list<UserVideoInfo> users_video_info;
}

service VideoService{
    UsersVideoInfoResponse UserVideoInfo(1:UsersVideoInfoRequest req)
}