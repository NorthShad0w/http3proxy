# http3proxy

## 介绍

通过cdn和http3出网进行代理的工具

## 使用方法

### 编译

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

必须和客户端放在同一文件夹下，secret不要变  
config.json
```
{
    "secret":"C{ZM2<%4H!)$kQ8cuaV?",
    "cdn_ip_port":"172.67.179.130:443",
    "sni_name":"proxy.blankofchina.cn.com",
    "host_name":"proxy.blankofchina.cn.com"
}
```

```bash
./http3proxy -nolog 1 -noprint 1
```

#### 使用代理

socks5代理开在服务端指定的端口上


