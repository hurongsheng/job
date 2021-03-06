### 什么是GC

    垃圾回收，指对不再使用的内存分配空间的回收
    由赋值器和回收器组成
    
### 根对象是什么

    全局变量：程序在编译期就能确定的那些存在于程序整个生命周期的变量。
    执行栈：每个 goroutine 都包含自己的执行栈，这些执行栈上包含栈上的变量及指向分配的堆内存区块的指针。
    寄存器：寄存器的值可能表示一个指针，参与计算的这些指针可能指向某些赋值器分配的堆内存区块。
    
### Gc的实现方式

    1.一般有追踪和引用计数两种，一般都混合使用
    2.追踪：从根对象出发，推导扫描每个对象的引用，然后回收可回收是数量。类似老师点名
    3.引用计数：每个对象自己维护被引用的计数器，计数器归零则可回收
    
### 分代式

    分代式只指，存活时间较久的，认为可能存活更久，存活时间较短的，则可能存活的时间较短
    所以区分了永久代，老年代和年轻代
    
### go采用三色标记法

    go采用的是，无分代，不整理，并发的三色标记法
    白色：可能存活。灰色：波面。黑色：确定存活
    初始所有对象为白色，由根对象出发，推导各个节点，最终不可达的白色节点为死亡节点，需要回收
    
### STW

    stop the world 和start the world 的缩写
    STW 在垃圾回收过程中为了保证实现的正确性、防止无止境的内存增长等问题而不可避免的需要停止赋值器进一步操作对象图的一段过程。
    
### 何时出发gc

     主动调用runtime.GC，系统监控2min没有执行gc了，通过步调算法，监控内存增长，增长到一定程度后，启动gc
     
### gc调优

    如何gc带来的开销主要是stw，那么gc的调优主要是减少内存的申请，提高已申请内存的复用率，控制分配内存的速度