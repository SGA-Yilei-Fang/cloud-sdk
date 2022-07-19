package service

import (
	"context"
	"fmt"

	api "github.com/SGA-Bicheng-Zhang/cloud-sdk/api/grpc"
	"github.com/SGA-Bicheng-Zhang/cloud-sdk/dao"
)

var _ api.AliyunServer = (*AliyunService)(nil)

type AliyunService struct {
	api.UnimplementedAliyunServer
	dao *dao.AliyunDao
}

func NewAliyunService(d *dao.AliyunDao) *AliyunService {
	return &AliyunService{dao: d}
}

func (s *AliyunService) SendMessage(ctx context.Context, req *api.AliyunSendMessageRequest) (resp *api.AliyunSendMessageResponse, err error) {
	// TODO: implment this method.
	defer func() { fmt.Printf("req: %v, resp: %v", req, resp) }()
	resp = &api.AliyunSendMessageResponse{
		Code:    200,
		Message: "ok",
	}
	return
}
