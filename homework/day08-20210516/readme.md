1. socket clock
2. query worker
3. TCPPool
4. 使用map+locker实现一个线程安全的map key, value
5. 使用slice+locker实现一个线程安全的切片
   启动4种操作 每种操作10个例程 循环1000次对应操作
    添加元素
    删除元素
    查看元素
    修改元素

    mutex/rwmutex