
### Kafka

#### 简介 [参考](https://blog.csdn.net/lingbo229/article/details/80761778)

    分布式、支持分区、多副本
    基于zookeeper协调的分布式消息系统
    
#### 特性

    - 高吞吐、低延时：每秒可以处理数十万条消息，延时几毫秒
    - 支持集群热扩展
    - 持久性、可靠性：消息被持久化到本地磁盘，并支持数据备份防止数据丢失
    - 高并发：支持数千个客户端同时读写
    - 容错性：允许节点失败（小于N）

#### 场景

    - 日志收集
    - 消息系统
    
#### 集群高可用机制

    集群管理用zookeeper管理，参考zookeeper选举、恢复机制
   
#### 概念
    - consumer group 
        每个consumer组里的consumer共同消费一个topic的消息
    - partition 分区
        同一个分区取数是有序的
        但是多个consumer消费的时候处理不一定有序
        每次只允许一个consumer group 中的consumer从一个partition中取数
        如果希望并发则增加多个partition（分布式消费）
    
     
        
    