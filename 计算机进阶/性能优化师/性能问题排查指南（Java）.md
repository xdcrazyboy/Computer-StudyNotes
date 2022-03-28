
[TOC]
# 命令汇总备忘录
## 系统工具命令
- **top**
	- top -Hp pid 查看具体线程使用系统资源情况
- **vmstat**：指定采样周期和次数的功能性检测工具，统计内存使用情况、观察CPU使用率、swap适应情况。主要用于**观察进程上下文切换**
	- vmstat 1 3  一秒一次，统计三次。 出现是的数据解析如下
	![enter image description here](/tencent/api/attachments/s3/url?attachmentid=232697)
	- r：等待运行的进程数；
	- b: 处于非中断睡眠状态的进程数
	- swpd： 虚拟内存使用情况；
	- free： 空闲的内存
	- buff：用于缓冲的内存数
	- si:从磁盘交换到内存的交换页数量;
	- so:从内存交换到磁盘的交换页数量;
	- bi:发送到块设备的块数;
	- bo:从块设备接收到的块数;
	- in:每秒中断数;
	- cs:每秒上下文切换次数;
	- us:用户 CPU 使用时间;
	- sy:内核 CPU 系统使用时间;
	- id:空闲时间;
	- wa:等待 I/O 时间;
	- st:运行虚拟机窃取的时间
- **pidstat** ： Sysstat 中的一个组件，也是一款功能强大的性能监测工具，**深入到线程级别**。
	>可以通过命令:`yum install sysstat` 安装该监控组件
		
		- pidstat -help **查看参数**，解析如下：
			- -u:默认的参数，显示各个进程的 cpu 使用情况;
			- -r:显示各个进程的内存使用情况;
			- -d:显示各个进程的 I/O 使用情况;
			- -w:显示每个进程的上下文切换情况;
			- -p:指定进程号;
			- -t:显示进程中线程的统计信息。
		- 通过ps、jsp获取进程ID，然后 `pidstat -p 345 -r 1 3`, 获取的数据解析如下：
			- Minflt/s:任务每秒发生的次要错误，不需要从磁盘中加载页;
			- Majflt/s:任务每秒发生的主要错误，需要从磁盘中加载页;
			- VSZ:虚拟地址大小，虚拟内存使用 KB;
			- RSS:常驻集合大小，非交换区内存使用 KB。

## JDK工具命令

### jstack 堆栈分析
- jstack pid 查看线程堆栈信息，结合top -Hp pid查看现场状态，也经常用来排查一些死锁异常。
	- 	线程 ID、线程的状态(wait、sleep、running 等状态)以及是否持有锁

### jmap 堆内存
**作用**：
- 堆内存初始化配置信息以及堆内存使用情况
- 堆内存中对象的信息：包括产生了哪些对象，对象数量多少

**举例**：
- jmap -heap pid 查看堆内存初始化配置信息以及堆内存的使用情况
![enter image description here](/tencent/api/attachments/s3/url?attachmentid=232773)
- jmap -histo:live pid 查看堆内存中的对象数目、大小统计直方图，如果带live则只统计存活对象。
![enter image description here](/tencent/api/attachments/s3/url?attachmentid=232778)
- dump到文件中：`jmap -dump:format=b,file=/tmp/heap.hprof pid`， 然后将文件下载下来，使用 **MAT** 工具打开文件进行分析。

# 系统篇
参考Brendan Gregg提供的完整图谱：
![enter image description here](/tencent/api/attachments/s3/url?attachmentid=233820)

## CPU

## 内存


## 网络

## 经验之谈
- dstat命令是一个用来替换vmstat、iostat、netstat、nfsstat和ifstat这些命令的工具，是一个全能系统信息统计工具


# JVM篇

## 运行时监控
- 利用 JMC、JConsole 等工具进行运行时监控。


## 工具分析
- 利用各种工具，在运行时进行堆转储分析，或者获取各种角度的统计数据(如jstat -
gcutil 分析 GC、内存分带等)。

## GC日志分析
- GC 日志等手段，诊断 Full GC、Minor GC，或者引用堆积等


## Profiling
对于应用**Profiling**，简单来说就是利用一些侵入性的手段，收集程序运行时的**细节**，以定
位性能问题瓶颈.
>所谓的细节，就是例如内存的使用情况、最频繁调用的方法是什么，或者上下文切换的情况等
一般不建议生产系统进行 Profiling，大多数是在性能测试阶段进行。

但是，当生产系统确实存在这种需求时，也不是没有选择。我建议使用 JFR配合JMC来做 Profiling，因为它是从 Hotspot JVM 内部收集底层信息，并经过了大量优化，性能开销非常低，通常是低于 2% 的


它的使用也非常方便，你不需要重新启动系统或者提前增加配置。例如，你可以在运行时启动 JFR 记录，并将这段时间的信息写入文件:
``` 
Jcmd <pid> JFR.start duration=120s filename=myrecording.jfr
```
然后，使用 JMC 打开“.jfr 文件”就可以进行分析了，方法、异常、线程、IO 等应有尽有，其功能非常强大。如果你想了解更多细节，可以参考[相关指南](https://blog.takipi.com/oracle-java-mission-control-the-ultimate-guide/) 。


profiling收集程序运行时信息的方式主要有以下三种:
- 事件方法:对于 Java，可以采用 JVMTI(JVM Tools Interface)API 来捕捉诸如方法调用、类载入、类卸载、进入 / 离开线程等事件，然后基于这些事件进行程序行为的分析。统计抽样方法(sampling): 该方法每隔一段时间调用系统中断，然后收集当前的调用栈(call stack)信息，记录调用栈中出现的函数及这些函数的调用结构，基于这些信息得



# 实践案例

## 系统越来越慢？


### 经验
1. 找繁忙线程时，top -h , 再jstack， 再换算tid比较累，而且jstack会造成停顿。推荐用vjtools里的vjtop, 不断显示繁忙的javaj线程，不造成停顿