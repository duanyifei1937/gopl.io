# ch8
* clock1: 启动8000，每个1s输出当前时间，写入conn; 单线程执行；并发1；
* netcat1: 连接clock1, 将conn内数据terminal显示；

* clock2: goroute clock1;支持并发；每一个连接一个goroutine;

---
* reverb: 
    conn --> output; 回声三次输出；
* netcat2:
    主goroutine:            stdint --> conn;
    second goroutine:       conn --> stdout;
    当主退出，程序退出，即使second未退出；
    需要通道等待两边同时退出；

    优化tcp半关闭；



