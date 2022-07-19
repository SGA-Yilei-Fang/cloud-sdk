package server

import (
	pb "github.com/SGA-Bicheng-Zhang/cloud-sdk/api/grpc"
	"github.com/SGA-Bicheng-Zhang/cloud-sdk/dao"
	"github.com/SGA-Bicheng-Zhang/cloud-sdk/service"
	"google.golang.org/grpc"
)

// NewServer register grpc server.
func NewServer() *grpc.Server {
	// create new dao.
	aliyunDao := dao.NewAliyunDao()
	tencentCloudDao := dao.NewTencentCloudDao()

	// create new services.
	aliyunSvc := service.NewAliyunService(aliyunDao)
	tencentSvc := service.NewTencentCloudService(tencentCloudDao)

	grpcServer := grpc.NewServer()

	// register server here...
	pb.RegisterAliyunServer(grpcServer, aliyunSvc)
	pb.RegisterTencentCloudServer(grpcServer, tencentSvc)

	return grpcServer
}
