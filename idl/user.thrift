namespace go user

struct BaseResp {
  1:i64 status_code;
  2:string status_message;
  3:i64 service_time;
}

struct User{
    1:i64 id;
    2:string user_name;
    3:i64 follow_count;
    4:i64 follower_count;
    5:bool is_follow;
    6:string background_image;
    7:string avatar;
    8:string signature;
    9:i64 total_favorited;
    10:i64 work_count;
    11:i64 favorite_acount;
}

struct UserInfoRequest {
    1:string user_id;
    2:string token;
}

struct UserInfoResponse {
    1:BaseResp base_resp;
    2:User user;
}

struct UserLoginRequest{
    1:string user_name;
    2:string user_password;
}

struct UserLoginResponse{
    1:BaseResp base_resp;
    2:i64 user_id;
    3:string token;
}

struct UserCreateRequest{
    1:string user_name;
    2:string user_password;
}

struct UserCreateResponse{
    1:BaseResp base_resp;
    2:i64 user_id;
    3:string token;
}

struct RelationActionRequest{
    1:string token;
    2:string to_user_id;
    3:string action_type;
}

struct RelationActionResponse{
    1:BaseResp base_resp;
}

struct FriendListRequest{
    1:i64 user_id;
    2:string token;
}

struct FriendListResponse{
    1:BaseResp base_resp;
    2:list<User> user_list;
}

struct FollowerListRequest{
    1:string user_id;
    2:string token;
}

struct FollowerListResponse{
    1:BaseResp base_resp;
    2:list<User> user_list;
}

struct FollowListRequest{
    1:string user_id;
    2:string token;
}

struct FollowListResponse{
    1:BaseResp base_resp;
    2:list<User> user_list;
}


service UserService {
    UserInfoResponse UserInfo (1:UserInfoRequest req)
    TokenResponse TokenVerify (1:TokenRequest req)
    UserCreateResponse UserCreate (1:UserCreateRequest req)
    UserLoginResponse UserLogin (1:UserLoginRequest req)
    FriendListResponse FriendList(1:FriendListRequest req)
    FollowerListResponse FollowerList(1:FriendListRequest req)
    FollowListResponse FollowList(1:FollowListRequest req)
    RelationActionResponse RelationAction(1:RelationActionRequest req)
}
