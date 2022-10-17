# http2udp

## 原理

使用 HTTP Post 方法发送一个

Body 为 Json 的请求，格式如下：

```json
{
  "ip": "213.24.24.24",
  "port": "8228",
  "msg": "up"
}
```

`http2udp` 就会向 `213.24.24.24:8228` 发送一个内容为 `up` 的 UDP 包

## 使用

```bash
./http2udp -h

# Usage of ./http2udp:
#  -p int
#        port (default 8080)
```


```shell
./http2udp -p 80
```

