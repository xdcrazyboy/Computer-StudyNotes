

## 命令篇


打印线程的栈信息，制作线程Dump。

jstack <进程ID> >> <输出文件>

jstack 2316 >> c:\thread.txt


**dump内存情况**


- 打印存活的对象大小和个数
```java
jmap -histo:live <pid>

jmap -histo:live 64421 > live.log
```

- 二进制方式存储堆文件
>注意要在进程用户下，或者有权限用户。 
>然后命令找不到可以试试java按照目录全路径
```java
//jmap -dump:format=b,file=文件名.hprof <进程ID>
jmap -dump:format=b,file=/opt/wkt/wkt1.hprof 64421
```



## 工具篇





## 理论篇





## 实战篇

