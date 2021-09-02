package service

import (
	"context"
	"fmt"

	pb "wechat-center/api/wechat"
)

type MessageService struct {
	pb.UnimplementedMessageServer
}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (s *MessageService) ReplyMessage(ctx context.Context, req *pb.ReplyMessageRequest) (*pb.ReplyMessageReply, error) {
	return &pb.ReplyMessageReply{}, nil
}
func (s *MessageService) Auth(ctx context.Context, req *pb.AuthRequest) (*pb.AuthReply, error) {
	fmt.Println(req)
	return &pb.AuthReply{Echostr: req.Echostr}, nil
}
