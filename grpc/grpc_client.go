package main

import (
	"context"
	"github.com/Litekube/network-controller/grpc/grpc_client"
	"github.com/Litekube/network-controller/grpc/pb_gen"
	"github.com/Litekube/network-controller/utils"
)

var logger = utils.GetLogger()

var client *grpc_client.GrpcClient

func Init() {
	client = &grpc_client.GrpcClient{
		Ip:          "101.43.253.110",
		Port:        "6440",
		GrpcCertDir: "/root/go_project/network-controller-usage/certs/test/",
		CAFile:      "ca.pem",
		CertFile:    "client.pem",
		KeyFile:     "client-key.pem",
	}
	err := client.InitGrpcClientConn()
	logger.Info(err)
}

func main() {
	Init()
	tokenResp, err := GetToken()
	if err != nil {
		logger.Errorf("fail to call GetToken err: %+v", err)
		return
	}
	token := tokenResp.Token
	logger.Infof("register token: %+v", token)

	// run network client first, then execute two methods below
	//checkResp, err := CheckConnState(token)
	//if err != nil {
	//	logger.Errorf("fail to call CheckConnState err: %+v", err)
	//	return
	//}
	//logger.Info("get bind ip:%+s, conn state:%+v", checkResp.BindIp, checkResp.ConnState)
	//
	//unRegisResp, err := UnRegister(token)
	//if err != nil {
	//	logger.Errorf("fail to call UnRegister err: %+v", err)
	//	return
	//}
	//logger.Infof("if succeed to unRegister: %+v", unRegisResp.Result)
}

func GetToken() (*pb_gen.GetTokenResponse, error) {
	req := &pb_gen.GetTokenRequest{}

	resp, err := client.C.GetToken(context.Background(), req)
	logger.Info(resp)
	logger.Info(err)

	return resp, nil
}

func CheckConnState(token string) (*pb_gen.CheckConnResponse, error) {
	req := &pb_gen.CheckConnStateRequest{
		Token: token,
	}

	resp, err := client.C.CheckConnState(context.Background(), req)
	logger.Info(resp)
	logger.Info(err)
	return resp, err
}

func UnRegister(token string) (*pb_gen.UnRegisterResponse, error) {
	req := &pb_gen.UnRegisterRequest{
		Token: token,
	}

	resp, err := client.C.UnRegister(context.Background(), req)
	logger.Info(resp)
	logger.Info(err)
	return resp, err
}
