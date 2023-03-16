package main

import (
	"context"
	"dousheng/common/errno"
	publish "dousheng/publish/kitex_gen/publish"
	"dousheng/publish/pack"
	"dousheng/publish/service"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishVideo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishVideo(ctx context.Context, req *publish.PublishVideoRequest) (resp *publish.PublishVideoResponse, err error) {
	// TODO: Your code here...
	resp = new(publish.PublishVideoResponse)
	res, err := service.NewCreateVideoService(ctx).CreateVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoId = int64(res)
	return
}

// ListVideos implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) ListVideos(ctx context.Context, req *publish.ListVideosRequest) (resp *publish.ListVideosResponse, err error) {
	// TODO: Your code here...
	resp = new(publish.ListVideosResponse)
	res, err := service.NewPublishListService(ctx).GetPublishList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = res
	return
}
