# 3月14日

## 案例一： gc日志跟业务日志停止时间不匹配？ 看gc日志的方法： 22764450


gc日志不太对呀，gc时间跟业务日志停止时间段不太一样
先看业务日志：

```java

[WARN ] 2022-03-14 14:53:23,305 [RMI TCP Connection(382177)-10.164.34.27] from ip:[10.164.34.27],RemoteInvocation: method name 'getCpcIdeaListByCpcGrpIdListAndStatusForSpecial'; parameter types [java.lang.Long, java.util.List, java.util.List, java.util.List],target:[com.sun.proxy.$Proxy343],loginUser null!
[INFO ] 2022-03-14 14:53:45,037 [org.springframework.scheduling.concurrent.ScheduledExecutorFactoryBean#0-3] Timer task, the program is running.....
[INFO ] 2022-03-14 14:53:45,039 [Thread-920] Clear timed cache
[INFO ] 2022-03-14 14:53:45,044 [resin-23-SendThread(10.160.14.154:2181)] Client session timed out, have not heard from server in 21824ms for sessionid 0x17ddcb5bc628138, closing socket connection and attempting reconnect
[INFO ] 2022-03-14 14:53:45,044 [resin-23-SendThread(10.160.14.154:2181)] Client session timed out, have not heard from server in 23539ms for sessionid 0x17ddcb5bc628134, closing socket connection and attempting reconnect
[INFO ] 2022-03-14 14:53:45,044 [resin-23-SendThread(10.160.14.166:2181)] Client session timed out, have not heard from server in 23558ms for sessionid 0x27ddcb5bc726940, closing socket connection and attempting reconnect
[WARN ] 2022-03-14 14:53:45,048 [RMI TCP Connection(381748)-10.160.185.53] from ip:[10.160.185.53],RemoteInvocation: method name 'getPlanById'; parameter types [java.lang.Long, java.lang.Long],target:[com.sun.proxy.$Proxy330],loginUser null!
[WARN ] 2022-03-14 14:53:45,057 [RMI TCP Connection(381793)-10.143.185.26] from ip:[10.143.185.26],RemoteInvocation: method name 'getAccountLevelDtoByAccountId'; parameter types [java.lang.Long],target:[com.sun.proxy.$Proxy333],loginUser null!
```
可以看到，2022-03-14 14:53:23 ~ 2022-03-14 14:53:45 之间是没有业务日志的，也就是说属于STW时间
再看gc日志：

账号ID 22764450 部分full gc
apollo 10.164.128.210	2022-03-14 14:53:35 977	2022-03-14 14:54:03 493

```java
2022-03-14T14:53:15.236+0800: 318865.535: [GC (Allocation Failure) 
2022-03-14T14:53:15.237+0800: 318865.536: [ParNew
Desired survivor size 697925632 bytes, new threshold 6 (max 6)
- age   1:    6639472 bytes,    6639472 total
- age   2:   10362208 bytes,   17001680 total
- age   3:    7893584 bytes,   24895264 total
: 11630323K->189685K(12268352K), 5.3521316 secs] 19190776K->12536392K(40579904K), 5.3548955 secs] [Times: user=85.02 sys=0.00, real=5.36 secs]

2022-03-14T14:53:23.324+0800: 318873.623: [GC (Allocation Failure) 
2022-03-14T144:53:23.325+0800: 318873.624: [ParNew (promotion failed): 10753166K->10693429K(122
268352K), 3.2176676 secs]
2022-03-14T14:53:26.543+0800: 318876.842: [CMS: 14130222
6K->3745616K(28311552K), 18.4838208 secs] 23099874K->3745616K(40579904K), [Metass
pace: 139290K->139290K(167936K)], 21.7104200 secs] [Times: user=66.03 sys=0.00,  
real=21.71 secs]

2022-03-14T14:53:59.444+0800: 318909.743: [GC (Allocation Failure) 
2022-03-14T144:53:59.445+0800: 318909.744: [ParNew
Desired survivor size 697925632 bytes, new threshold 6 (max 6)
- age   1:  290037208 bytes,  290037208 total
: 10905216K->323809K(12268352K), 1.6613804 secs] 14650832K->5852947K(40579904K),,
 1.6645653 secs] [Times: user=25.89 sys=0.00, real=1.66 secs]

2022-03-14T14:54:45.030+0800: 318955.329: [GC (Allocation Failure) 
2022-03-14T144:54:45.034+0800: 318955.333: [ParNew
Desired survivor size 697925632 bytes, new threshold 6 (max 6)
- age   1:  263349136 bytes,  263349136 total
- age   2:  206822944 bytes,  470172080 total
: 11229025K->559302K(12268352K), 0.3198510 secs] 16758163K->6088440K(40579904K), 0.3250050 secs] [Times: user=4.83 sys=0.07, real=0.32 secs]
```



## 案例二 业务日志跟gc日志也有误差，但时间类似。 耗时还有其他原因： 23222668

从下面业务日志看，耗时确实这么长

- 查库？

- 看业务日志 14:17:36 直接跳到 14:17:44， 垃圾回收了8s?

```java
[INFO ] 2022-03-14 14:17:47,705 [RMI TCP Connection(374801)-10.143.189.65] StopWatch 'getCpcKeyListPageByQueryAndAdType': running time (millis) = 23582
-----------------------------------------
ms     %     Task name
-----------------------------------------
22045  093%  endQueryCpcList
01529  006%  QueryReport
00008  000%  HandlePageInfo

[INFO ] 2022-03-14 14:17:47,705 [RMI TCP Connection(374801)-10.143.189.65] 23222668     requestId=0     [Outer Call]    com.sogou.bizdev.cpc.key.provider.CpcKeyProvider.getCpcListByQuery,args:[ Long:23222668 KeyQueryDto:{"accountId":23222668,"colorTags":[],"cpcGrpId":0,"cpcIds":[],"cpcPlanId":0,"filterStatus":0,"keyMatchTypes":[],"keyStatusList":[],"maxMaxPrice":0,"minMaxPrice":0,"mobileQualityDegrees":[],"orderBy":21,"pagedBean":{"curPageNum":1,"dataList":[],"isFirstPage":false,"isLastPage":false,"nextPage":0,"pageSize":50,"prePage":0,"startSerial":0,"totalPages":0,"totalRecNum":0},"pcQualityDegrees":[],"queryMatchType":0,"queryWord":"","statisticDataQueryDto":{"endDate":{"date":13,"day":0,"hours":0,"minutes":0,"month":2,"seconds":0,"time":1647100800000,"timezoneOffset":-480,"year":122},"materialType":0,"startDate":{"date":13,"day":0,"hours":0,"minutes":0,"month":2,"seconds":0,"time":1647100800000,"timezoneOffset":-480,"year":122},"timeType":6}} SearchAdUser:com.sogou.bizdev.bizlog.dto.SearchAdUser@3706962d],processing time: 23805
```

- gc日志

```java
apollo 10.160.78.59	2022-03-14 14:17:23 659

```java
2022-03-14T14:17:04.176+0800: 316556.392: [GC (Allocation Failure) 
2022-03-14T14:17:04.177+0800: 316556.392: [ParNew
Desired survivor size 697925632 bytes, new threshold 1 (max 6)
- age   1: 1146176680 bytes, 1146176680 total
: 12158181K->1151620K(12268352K), 1.5177358 secs] 27777349K->18933725K(40579904K), 1.5199294 secs] [Times: user=23.09 sys=0.81, real=1.52 secs]
2022-03-14T14:17:05.711+0800: 316557.926: [GC (CMS Initial Mark) [1 CMS-initial-mark: 17782104K(28311552K)] 18933935K(40579904K), 0.0123023 secs] [Times: user=0.06 sys=0.01, real=0.02 secs]
2022-03-14T14:17:05.725+0800: 316557.940: [CMS-concurrent-mark-start]
2022-03-14T14:17:06.225+0800: 316558.440: [CMS-concurrent-mark: 0.485/0.500 secs] [Times: user=5.19 sys=0.37, real=0.50 secs]
2022-03-14T14:17:06.226+0800: 316558.441: [CMS-concurrent-preclean-start]
2022-03-14T14:17:06.273+0800: 316558.489: [CMS-concurrent-preclean: 0.047/0.047 secs] [Times: user=0.31 sys=0.04, real=0.05 secs]
2022-03-14T14:17:06.274+0800: 316558.489: [CMS-concurrent-abortable-preclean-start]
 CMS: abort preclean due to time 2022-03-14T14:17:11.288+0800: 316563.504: [CMS-concurrent-abortable-preclean: 4.983/5.014 secs] [Times: user=26.29 sys=1.40, real=5.01 secs]
2022-03-14T14:17:11.303+0800: 316563.518: [GC (CMS Final Remark) [YG occupancy: 6327134 K (12268352 K)]2022-03-14T14:17:11.303+0800: 316563.518: [GC (CMS Final Remark) 2022-03-14T14:17:11.304+0800: 316563.519: [ParNew (promotion failed): 6327134K->5311815K(12268352K), 1.0276246 secs] 24109239K->23094088K(40579904K), 1.0297185 secs] [Times: user=2.12 sys=0.12, real=1.03 secs]
2022-03-14T14:17:12.333+0800: 316564.549: [Rescan (parallel) , 1.5637997 secs]2022-03-14T14:17:13.897+0800: 316566.113: [weak refs processing, 0.0079816 secs]2022-03-14T14:17:13.905+0800: 316566.121: [class unloading, 0.0930767 secs]2022-03-14T14:17:13.998+0800: 316566.214: [scrub symbol table, 0.0218572 secs]2022-03-14T14:17:14.020+0800: 316566.236: [scrub string table, 0.0022893 secs][1 CMS-remark: 17782272K(28311552K)] 23094088K(40579904K), 2.7495797 secs] [Times: user=26.68 sys=0.37, real=2.75 secs]
2022-03-14T14:17:14.054+0800: 316566.270: [CMS-concurrent-sweep-start]
2022-03-14T14:17:36.993+0800: 316589.209: [CMS-concurrent-sweep: 17.895/22.939 secs] [Times: user=83.75 sys=17.18, real=22.94 secs]
2022-03-14T14:17:37.040+0800: 316589.255: [GC (Allocation Failure) 2022-03-14T14:17:37.041+0800: 316589.257: [ParNew: 11041517K->11041517K(12268352K), 0.0000416 secs]2022-03-14T14:17:37.041+0800: 316589.257: [CMS: 2561685K->2062560K(28311552K), 7.8205440 secs] 13603202K->2062560K(40579904K), [Metaspace: 135050K->135050K(180224K)], 7.8238814 secs] [Times: user=6.35 sys=1.49, real=7.83 secs]
2022-03-14T14:17:49.565+0800: 316601.780: [GC (Allocation Failure) 2022-03-14T14:17:49.566+0800: 316601.781: [ParNew
Desired survivor size 697925632 bytes, new threshold 6 (max 6)
- age   1:  340847128 bytes,  340847128 total
: 10905216K->366409K(12268352K), 0.1260961 secs] 12967776K->2428970K(40579904K), 0.1287021 secs] [Times: user=1.88 sys=0.11, real=0.13 secs]
```
```