## 微服务治理

### 高可用
- 服务熔断降级限流机制 熔断降级的概念

        1.服务降级
        2.服务限流
        3.服务熔断
        4.兜底，plan b
    
- 框架调用方式解耦方式 Kit 或 Istio 或 Micro 
    
        框架和解耦grpc,protobuf
    
- [服务注册、发现 RPC调用框架](服务注册和发现/readme.md)
    
        服务注册和发现consul zookeeper etcd
    
  
- 结构化数据库

        mysql
             
- 网关 (kong gateway).

- Docker部署管理 Kubenetters. 自动集成部署 CI/CD 实践.自动扩容机制规则.
    
        docker k3s k8s
- Trasport 数据传输(序列化和反序列化).

        json protobuf

- 压测 优化.
    
        容量设计，性能监控，报警
        
### 集群和高可用
    
    mysql,redis,kafka,aerospike,console,zookeeper,etcd
           
### 高并发
- 异步、消峰
            
        kafka nsq robotmq redis
        
- 多级缓存。高性能存储。非结构化数据库

        redis aerospike mongo es
        缓存预热、缓存穿透、缓存雪崩   
                 
### 监控反馈，排查问题

- Logging 日志sls收集
- 链路监控 和prometheus.
- Metrics 指针对每个请求信息的仪表盘化.
