### Raft
-
    所有节点都是对等的，有3中角色leader,follower,candidate（预备leader）
    term 任期
    
#### 作为leader
- 每个任期内都只会有一个leader
- leader 只能追加而不能重写或者删除日志
- 如果leader接受到一个比自己term的AppendEntries时就知道自己要退位了，自动转为follower

#### 作为follower
- 当一个follower发现leader不可用是，会发起投票竞选leader此时为（candidate）
- 如果多个follower 同时竞选导致票数不超过一半，会随机sleep，重新发起竞选
- 如果follower日志新于candidate，则投反对票

#### 关于日志
- 如果一条日志被commit，那么大于该日志任期的日志都应该包含这个日志
- 如果两个日志条目的index和term都相同，则两个如果日志中，两个条目及它们之前的日志条目也完全相同

#### 什么情况下会产生脑裂

- leader宕机时，配置刚好被修改了，配置还没下发到部分机器，老机器有可能会因为老配置选举产生2个leader
- raft把配置修改也作为log append到log中，
- 如果老leader被隔离到较小的分区中，新分区会产生一个新leader产生脑裂，大多数节点会更新到新的term，老leader会无法commit一个新的提交

#### raft如何处理一个请求

- client向集群的leader发送get/put请求
- leader添加命令到自己打日志
- 发送log到followers
- follower添加日志到自己的日志
- leader收到大多数的节点返回时，log进入commit状态
- leader返回给客户端

#### 可以同时接受多少个节点宕机

- 2n+1 时容错n;2n时容错n-1

#### 如果请求commit之前，leader丧失领导权会发生什么

- leader丧失领导权会同步失败，client重新发起请求

#### raft的主要作用

- 选举产生新leader
- 确保存在宕机节点的情况下，保持同步复制的正确性

#### client与哪个节点交互？

 - 当client发起请求的时候，只能和leader节点交互，如果不知道哪个是leader，就随便找一个节点，然后发出请求，如果不是leader节点就会返回leader节点的index。

#### 如果一个客户端PUT/GET RPC超时还没收到回复会怎么样？
     Call()调用返回false；
     如果服务器挂掉，或者请求丢失掉 ==> 重发
     如果服务器执行，但是回复丢失 ==> 重发是危险的

-  存在的问题：后两种对于client来说看到的一样，都是没有回复，如果已经执行了，client也需要收到执行的结果。如果重发就存在多次执行的可能。  
- 解决办法：做重复RPC的检测，我们让KV service有能力检测是否已经收到过此请求，对于每一个请求client会生成一个ID，在发起请求的时候会把ID包含在RPC中发送给leader，如果leader收到了相同ID的RPC请求就认为是重复的。
