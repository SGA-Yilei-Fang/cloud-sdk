package service

import (
	"context"
	"fmt"

	api "github.com/SGA-Bicheng-Zhang/cloud-sdk/api/grpc"
	"github.com/SGA-Bicheng-Zhang/cloud-sdk/dao"
)

var _ api.TencentCloudServer = (*TencentCloudService)(nil)

type TencentCloudService struct {
	api.UnimplementedTencentCloudServer
	dao *dao.TencentCloudDao
}

func NewTencentCloudService(d *dao.TencentCloudDao) *TencentCloudService {
	return &TencentCloudService{dao: d}
}

func (s *TencentCloudService) SendMessage(ctx context.Context, req *api.TencentSendMessageRequest) (resp *api.TencentSendMessageResponse, err error) {
	// TODO: implment this method.
	defer func() { fmt.Printf("req: %v, resp: %v", req, resp) }()
	resp = &api.TencentSendMessageResponse{
		Code:    200,
		Message: "ok",
	}
	return
}
