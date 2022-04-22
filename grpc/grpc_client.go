package main

import (
	"context"
	"encoding/base64"
	"github.com/Litekube/litekube-vpn/grpc/grpc_client"
	"github.com/Litekube/litekube-vpn/grpc/pb_gen"
	"github.com/Litekube/litekube-vpn/utils"
	certutil "github.com/rancher/dynamiclistener/cert"
	"time"
)

var logger = utils.GetLogger()
var client *grpc_client.GrpcClient

func Init() {
	client = &grpc_client.GrpcClient{
		Ip:          "101.43.253.110",
		Port:        "6440",
		GrpcCertDir: "/Users/zhujianxing/GoLandProjects/litekube-vpn/certs/test1/",
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

	// run vpn client first, then execute two methods below
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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resp, err := client.C.GetToken(ctx, req)
	logger.Info(resp)
	logger.Info(err)
	if err != nil {
		return resp, err
	}

	caBytes, err := base64.StdEncoding.DecodeString(resp.GrpcCaCert)
	certBytes, err := base64.StdEncoding.DecodeString(resp.GrpcClientCert)
	keyBytes, err := base64.StdEncoding.DecodeString(resp.GrpcClientKey)
	certutil.WriteCert("/Users/zhujianxing/GoLandProjects/litekube-vpn/certs/test/ca.pem", caBytes)
	certutil.WriteCert("/Users/zhujianxing/GoLandProjects/litekube-vpn/certs/test/client.pem", certBytes)
	certutil.WriteKey("/Users/zhujianxing/GoLandProjects/litekube-vpn/certs/test/client-key.pem", keyBytes)
	return resp, nil
}

func CheckConnState(token string) (*pb_gen.CheckConnResponse, error) {
	req := &pb_gen.CheckConnStateRequest{
		Token: token,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resp, err := client.C.CheckConnState(ctx, req)
	logger.Info(resp)
	logger.Info(err)
	return resp, err
}

func UnRegister(token string) (*pb_gen.UnRegisterResponse, error) {
	req := &pb_gen.UnRegisterRequest{
		Token: token,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resp, err := client.C.UnRegister(ctx, req)
	logger.Info(resp)
	logger.Info(err)
	return resp, err
}
