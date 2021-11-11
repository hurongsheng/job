### linux
#### select、poll和epoll的总结对比

    https://blog.csdn.net/qq_35976351/article/details/85228002
    
    select、poll和epoll是多路IO复用函数，是对IO事件的集中代理。
    select提供了事件发生通知。主程序需要遍历所有事件获得事件
    epoll会通知到具体是哪些事件有数据反馈
