### context作用，原理

    context主要用于父子任务之间的同步取消信号,超时控制
    本质上是一种协程调度的方式。
    另外在使用context时有两点值得注意：
        - 上游任务仅仅使用context通知下游任务不再需要，但不会直接干涉和中断下游任务的执行，由下游任务自行决定后续的处理操作，也就是说context的取消操作是无侵入的；
        - context是线程安全的，因为context本身是不可变的（immutable），因此可以放心地在多个协程中传递使用。
    
    1.传递共享数据
    2.取消 gorotinue
    3.防止gorotinue泄露