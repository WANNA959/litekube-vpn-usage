# litekube-vpn使用demo

* [litekube-vpn使用demo](#litekube-vpn使用demo)
    * [vpn部分](#vpn部分)
    * [grpc部分](#grpc部分)
    * [一次完成执行流程](#一次完成执行流程)

## vpn部分

- server/client执行详见vpm/vpn.go
    - 根据yml配置文件的不同，可以自动鉴别server or client并启动
    - 配置文件字段含义见server.yml/client.yml注册

```shell
# vpn & grpc server启动
go run vpn/vpn.go server.yml

# vpn client启动
go run vpn/vpn.go client.yml
```

## grpc部分

- 详见grpc/grpc_client.go
    - 运行前应该先执行 `go run vpn/vpn.go server.yml`

```shell
cd grpc
# 注册 得到token=a9f683a2d05b4957，以及返回的grpc+vpn证书
go test -v -run TestGetToken

# 查询连接状态
go test -v -run TestCheckConnState 

# 取消注册（取消绑定 & 断连
go test -v -run TestUnRegister 
```

## 一次完成执行流程

```shell
# 启动vpn & grpc server
go run vpn/vpn.go server.yml

# 注册 得到token=a9f683a2d05b4957，以及返回的grpc+vpn证书
cd grpc
go test -v -run TestGetToken

# 根据token修改client.yml，启动client
go run vpn/vpn.go client.yml

# 查询连接状态
go test -v -run TestCheckConnState 

# 取消注册（取消绑定 & 断连
go test -v -run TestUnRegister 
```

