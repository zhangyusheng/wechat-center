// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.5

package wechat

import (
	context "context"
	"fmt"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type MessageHTTPServer interface {
	Auth(context.Context, *AuthRequest) (*AuthReply, error)
	ReplyMessage(context.Context, *ReplyMessageRequest) (*ReplyMessageReply, error)
}

func RegisterMessageHTTPServer(s *http.Server, srv MessageHTTPServer) {
	r := s.Route("/")
	r.GET("/wechat/message/reply", _Message_ReplyMessage0_HTTP_Handler(srv))
	r.GET("/wechat/auth", _Message_Auth0_HTTP_Handler(srv))
}

func _Message_ReplyMessage0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReplyMessageRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.wechat.Message/ReplyMessage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReplyMessage(ctx, req.(*ReplyMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReplyMessageReply)
		return ctx.Result(200, reply)
	}
}

func _Message_Auth0_HTTP_Handler(srv MessageHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.wechat.Message/Auth")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Auth(ctx, req.(*AuthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthReply)
		fmt.Println(reply)
		return ctx.Result(200, in.Echostr)
	}
}

type MessageHTTPClient interface {
	Auth(ctx context.Context, req *AuthRequest, opts ...http.CallOption) (rsp *AuthReply, err error)
	ReplyMessage(ctx context.Context, req *ReplyMessageRequest, opts ...http.CallOption) (rsp *ReplyMessageReply, err error)
}

type MessageHTTPClientImpl struct {
	cc *http.Client
}

func NewMessageHTTPClient(client *http.Client) MessageHTTPClient {
	return &MessageHTTPClientImpl{client}
}

func (c *MessageHTTPClientImpl) Auth(ctx context.Context, in *AuthRequest, opts ...http.CallOption) (*AuthReply, error) {
	var out AuthReply
	pattern := "/wechat/auth"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.wechat.Message/Auth"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MessageHTTPClientImpl) ReplyMessage(ctx context.Context, in *ReplyMessageRequest, opts ...http.CallOption) (*ReplyMessageReply, error) {
	var out ReplyMessageReply
	pattern := "/wechat/message/reply"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.wechat.Message/ReplyMessage"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
