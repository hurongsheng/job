
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
    - 用户行为收集
    - 流式处理
    
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
    - offset
        由于 parttion是需要顺序读取，所以需要一个offset。
        对于high level api。offset存储于zookeeper中
        对于low level api，offset 由客户端自行维护
    -Partition Replica
        副本数量，不能多于broker数量
        producer 向partion leader发送消息，partion leader同步给partion follwer。并依据配置，有几个flower收到消息即给ack
        如果某个Replica宕机，且这个Replica在ack列表中，则会等到超时，然后从ack列表中移除。若不在ack列表中则无影响
    
#### 最佳实践

    consumer group 的 consumer数量等于partition数量，处理不过来时，同时增加
        
#### 消息投递可靠性
    一个消息如何算投递成功，Kafka提供了三种模式：
    - 第一种是啥都不管，发送出去就当作成功，这种情况当然不能保证消息成功投递到broker； partion ack =0
    - 第二种是Master-Slave模型，只有当Master和所有ack列表接收到消息时，才算投递成功，这种模型提供了最高的投递可靠性，但是损伤了性能； ack=-1
    - 第三种模型，即只要Master确认收到消息就算投递成功；实际使用时，根据应用特性选择，绝大多数情况下都会中和可靠性和性能选择第三种模型  ack=1