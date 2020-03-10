## 常用工具类

### log/panic.go
	输出panic错误信息

### net/client.go
    发送http请求(GET/POST/HEAD/PUT/DELETE)

### net/http.go
	获取request请求真实ip或port

### net/net.go
	获取设备网卡Ipv4地址列表
	
### rand/rand.go
	以时间戳为种子, 生成随机的整数/浮点数
	
### vers/vers.go
	版本比对, 支持大于/小于
	
### zip/snappy.go
	压缩算法, snappy算法（const中配置,大于指定大小使用压缩）
	
### zip/gzip.go
	压缩算法, gzip算法（const中配置,大于指定大小使用压缩）
