

namespace go userinfo

include "../base.thrift"


struct UesrInfoRequest {
    1:i64 visitor_user_id
    2:i64 host_user_id
}


struct UserInfo {
    1:string avatar
    2:string background_image
    3:string signature
    4:string username
    5:i64 follow_count
    6:i64 follower_count
    7:i64 total_favorited
    8:i64 favorite_count
    9:i64 work_count
    10:bool is_follow
}

struct UserInfoResponse {
    1:UserInfo user_info
    2:base.BaseResponse base_resp
}



service UserInfoService {
    UserInfoResponse GetUserInfo(1:UesrInfoRequest req)
}