# xddns

[![Go Reference](https://pkg.go.dev/badge/github.com/tonyzzp/xddns.svg)](https://pkg.go.dev/github.com/tonyzzp/xddns)

--------------

命令行工具，操作dns。目前支持ali/cloudflare。
同时支持获取本机ip，然后同步到dns

### 安装
```bash
go install github.com/tonyzzp/xddns@latest
```

### 帮助

```bash
xddns --help
xddns --help set
```

### 配置文件

配置文件`xddns-config.yaml` (**搜索路径顺序:workingdir,exedir**)，内容如下

- ali: 需要dns的读写权限 (系统配置策略 AliyunDNSFullAccess)
- cloudfalre: 需要读取zones列表和修改相应zone dns记录的权限 (zone.zone,  zone.dns)

```yaml
ali:
  region: cn-shenzhen
  key_id: xxxxx
  key_secret: xxxxx

cloudflare:
  token: xxxxx
```

### 通过命令行参数指定配置文件

--config参数一定要加在子命令前面

```bash
xddns --config /my/path/config.yaml set ...
xddns --config /my/path/config.yaml update ...
```


### 直接设置dns

```bash
xddns set --type A --domain myname.com --value 192.168.1.101
xddns set --type CNAME --domain www.myname.com --value  myname.com
```


### 直接更新为本机ip
```bash
xddns update --type ipv4 --domain a.myname.com
xddns update --type ipv6 --domain a.myname.com
```

### 查看本机ip
```bash
xddns ip
```