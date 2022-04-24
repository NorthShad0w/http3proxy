# http3proxy

## 介绍

通过cdn和http3出网进行代理的工具

## 使用方法

### 编译

修改`conn/rhttp3conn.go`文件

264行`ServerName`参数为域前置SNI域名

277行`req.Host`参数为真实域名

linux使用下列命令编译生成可执行文件

```bash
# linux版本
CGO_ENABLED=0 go build -ldflags "-w -s"
strip http3proxy
upx --best ./http3proxy

# windows版本
GOOS="windows" GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-w -s"
```

### 运行

#### cdn配置

客户端到cdn加密udp/443，cdn到服务端不加密tcp/80

#### 服务端运行

服务端置于cdn后，开启80端口监听请求

```bash
sudo ./http3proxy -type server -proto rhttp -listen 0.0.0.0:80 -nolog 1 -noprint 1
```

#### 客户端运行

```bash
./http3proxy -name "test" -type reverse_socks5_client -server <cdn-ip>:443 -fromaddr 0.0.0.0:<服务端上开启的socks5端口> -proto rhttp3 -proxyproto tcp -nolog 1 -noprint 1
```

#### 使用代理

socks5代理开在服务端指定的端口上


