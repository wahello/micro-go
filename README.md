# micro-go
```
    高并发和微服务kit实战
```
###

+ 服务注册与发现

```
    简单的字符串组合和相差案例
    1、组合字符串
    2、相差字符串
    3、健康检查
    调用案例
    http://127.0.0.1:10085/op/Concat/asda/asdasd
    http://127.0.0.1:10085/op/Diff/asda/asdasd
```

- 远程过程调用rpc


+ 分布式配置中心


+ 微服务网关


+ 微服务的容错处理与负载均衡


+ 统一认证与授权


+ 分布式链路追踪


### 目录结构
####transport层: 项目提供服务的方式（HTTP服务）
    主要负责网络传输，例如处理HTTP、gRPC、Thrift 相关逻辑
####endpoint层: 用于接受请求并返回响应
    主要负责request/response格式的转换，以及公用拦截器相关的逻辑
    并且提供对日志、限流、熔断、链路追踪和服务监控等扩展能力
####service层: 业务代码实现层
    主要负责于业务逻辑
```
Go-kit提供一下功能
- 熔断器 Circuit breaker
- 限流器 Rate limiter
- 日志 Logging
- Prometheus 统计 Metrics
- 请求跟踪 Request tracing
- 服务发现和负载均衡 
```


## 综合实战 -- 秒杀系统
  
