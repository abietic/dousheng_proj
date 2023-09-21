

namespace go publish

include "../base.thrift"


struct PublishVideoRequest {
    1:string video_url
    2:string cover_url
    3:string title
    4:i64 user_id
}

struct ListVideosRequest {
    1:i64 visitor_user_id
    2:i64 host_user_id
}

struct VideoInfo {
    1:i64 video_id
    2:string play_url
    3:string cover_url
    4:string title
    5:i64 favorite_count
    6:i64 comment_count
    7:bool is_favorite
}

struct PublishVideoResponse {
    1:i64 video_id
    2:base.BaseResponse base_resp
}

struct ListVideosResponse {
    1:list<VideoInfo> video_list
    2:base.BaseResponse base_resp
}


service PublishService {
    PublishVideoResponse PublishVideo(1:PublishVideoRequest req)
    ListVideosResponse ListVideos(1:ListVideosRequest req)
}