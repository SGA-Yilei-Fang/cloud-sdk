package service

import (
	"context"

	api "github.com/SGA-Bicheng-Zhang/cloud-sdk/api/grpc"
	"github.com/SGA-Bicheng-Zhang/cloud-sdk/dao"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"

	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
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
	// defer func() { fmt.Printf("req: %v, resp: %v", req, resp) }()
	// resp = &api.AliyunSendMessageResponse{
	// 	Code:    200,
	// 	Message: "ok",
	// }
	// return

	resp_fail := &api.AliyunSendMessageResponse{
		Code:    -1,
		Message: "fail",
	}

	resp_ok := &api.AliyunSendMessageResponse{
		Code:    200,
		Message: "ok",
	}

	client, _err := CreateClient(tea.String("LTAI5tSUJu8pymy6WQCKJrF3"), tea.String("9BFGFNoQyp7gBndvzfZsQdZuDM0Hy0"))
	if _err != nil {
		return resp_fail, _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(req.Phone),
		SignName:      tea.String("索尼无锡软件中心"),
		TemplateCode:  tea.String("SMS_243231247"),
		TemplateParam: tea.String("{\"code\":\"1234\"}"),
	}

	runtime := &util.RuntimeOptions{}

	mess, _err := client.SendSmsWithOptions(sendSmsRequest, runtime)
	if _err != nil {
		return resp_fail, _err
	} else {
		console.Log(util.ToJSONString(tea.ToMap(mess)))
		return resp_ok, _err

	}

}

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}

	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	// _result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	if _err != nil {
		panic(_err)
	}

	return _result, _err
}
