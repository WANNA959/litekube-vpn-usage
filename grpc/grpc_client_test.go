package main

import (
	"encoding/base64"
	certutil "github.com/rancher/dynamiclistener/cert"
	"os"
	"testing"
)

func TestGetToken(t *testing.T) {
	Init()
	tokenResp, err := GetToken()
	if err != nil {
		logger.Errorf("fail to call GetToken err: %+v", err)
		return
	}
	token := tokenResp.Token
	logger.Infof("register token: %+v", token)

	caBytes, err := base64.StdEncoding.DecodeString(tokenResp.GrpcCaCert)
	certBytes, err := base64.StdEncoding.DecodeString(tokenResp.GrpcClientCert)
	keyBytes, err := base64.StdEncoding.DecodeString(tokenResp.GrpcClientKey)
	certutil.WriteCert("/root/go_project/litekube-vpn-usage/certs/test1/ca.pem", caBytes)
	certutil.WriteCert("/root/go_project/litekube-vpn-usage/certs/test1/client.pem", certBytes)
	certutil.WriteKey("/root/go_project/litekube-vpn-usage/certs/test1/client-key.pem", keyBytes)

	caBytes, err = base64.StdEncoding.DecodeString(tokenResp.VpnCaCert)
	certBytes, err = base64.StdEncoding.DecodeString(tokenResp.VpnClientCert)
	keyBytes, err = base64.StdEncoding.DecodeString(tokenResp.VpnClientKey)
	certutil.WriteCert("/root/go_project/litekube-vpn-usage/certs/test2/ca.pem", caBytes)
	certutil.WriteCert("/root/go_project/litekube-vpn-usage/certs/test2/client.pem", certBytes)
	certutil.WriteKey("/root/go_project/litekube-vpn-usage/certs/test2/client-key.pem", keyBytes)

}

func TestCheckConnState(t *testing.T) {
	Init()
	checkResp, err := CheckConnState(os.Args[1])
	if err != nil {
		logger.Errorf("fail to call CheckConnState err: %+v", err)
		return
	}
	logger.Infof("get bind ip:%+v, conn state:%+v", checkResp.BindIp, checkResp.ConnState)
}

func TestUnRegister(t *testing.T) {
	Init()
	unRegisResp, err := UnRegister(os.Args[1])
	if err != nil {
		logger.Errorf("fail to call UnRegister err: %+v", err)
		return
	}
	logger.Infof("if succeed to unRegister: %+v", unRegisResp.Result)
}
