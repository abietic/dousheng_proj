// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userinfoservice

import (
	"context"
	userinfo "dousheng/publish/kitex_gen/userinfo"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetUserInfo(ctx context.Context, req *userinfo.UesrInfoRequest, callOptions ...callopt.Option) (r *userinfo.UserInfoResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserInfoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserInfoServiceClient struct {
	*kClient
}

func (p *kUserInfoServiceClient) GetUserInfo(ctx context.Context, req *userinfo.UesrInfoRequest, callOptions ...callopt.Option) (r *userinfo.UserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, req)
}
