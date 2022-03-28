
# 收集器

## CMS
### 遇到问题：CMS收集过程和日志分析
垃圾回收新生代和老年代的垃圾收集器组合： ParNew and CMS

mark-sweep分为多个阶段，其中一大部分阶段GC的工作是和Application threads的工作同时进行的（当然，gc线程会和用户线程竞争CPU的时间），默认的GC的工作线程为你服务器物理CPU核数的1/4；

>当你的服务器是多核同时你的目标是低延时，那该GC的搭配则是你的不二选择。

#### 什么是CMS
"Concurrent Mark and Sweep" 是CMS的全称，官方给予的名称是：“Mostly Concurrent Mark and Sweep Garbage Collector”;
年轻代：采用 stop-the-world [mark-copy](https://plumbr.eu/handbook/garbage-collection-algorithms/removing-unused-objects/copy) 算法；
年老代：采用 Mostly Concurrent [mark-sweep](https://plumbr.eu/handbook/garbage-collection-algorithms/removing-unused-objects/sweep) 算法；
设计目标：年老代收集的时候避免长时间的暂停；

CMS：在JDK 5发布时，HotSpot推出了一款在强交互应用中几乎可称为具有划时代意义的垃圾收集器——CMS收集器。是HotSpot虚拟机中第一款真正意义上 
CMS作为老年代的收集器，却无法与JDK 1.4.0中已经存在的新生代收集器Parallel Scavenge配合工作（一个面向高吞吐，一个面向低延时，而且G1这些不在是分代设计），所以在JDK 5中使用CMS来收集老年代的时候，新生代只能选择ParNew或者Serial收集器中的一个。ParNew收集器是激活CMS后（使用-XX：+UseConcMarkSweepGC选项）的默认新生代收集器，也可以使用-XX：+/-UseParNewGC选项来强制指定或者禁用它

#### Par New
自JDK 9开始，ParNew加CMS收集器的组合就不再是官方推荐的服务端模式下的收集器解决方案了。官方希望它能完全被G1所取代，甚至还取消了ParNew加SerialOld以及Serial加CMS这两组收集器组合的支持（其实原本也很少人这样使用），并直接取消了-XX：+UseParNewGC参数，这意味着ParNew和CMS从此只能互相搭配使用，再也没有其他收集器能够和它们配合了。读者也可以理解为从此以后，ParNew合并入CMS，成为它专门处理新生代的组成部分

#### 日志初体验

##### Minor GC

```log
2022-03-07T11:41:14.315+0800: 37.690: [GC (Allocation Failure) 
2022-03-07T11:41:14.315+0800: 37.691: [ParNew
Desired survivor size 697925632 bytes, new threshold 6 (max 6)
- age   1:  115229256 bytes,  115229256 total
- age   2:   47712552 bytes,  162941808 total
- age   3:   64836536 bytes,  227778344 total
: 11387196K->434763K(12268352K), 0.1768324 secs] 11387196K->434763K(40579904K), 0.1777042 secs] [Times: user=1.60 sys=0.28, real=0.18 secs]

```
* 2022-03-07T11:41:14.315+0800 – GC发生的时间；
* 37.690 – GC开始，相对JVM启动的相对时间，单位是秒；
* GC – 区别MinorGC和FullGC的标识，这次代表的是MinorGC;
* Allocation Failure – **MinorGC的原因**，在这个case里边，**由于年轻代不满足申请的空间，**因此触发了MinorGC;
* ParNew – 收集器的名称，它预示了年轻代使用一个**并行的** mark-copy stop-the-world 垃圾收集器；
* Desired survivor size 697925632 bytes, new threshold 6 (max 6) - 
* age   1:  115229256 bytes,  115229256 total
* 11387196K->434763K – 收集前后**年轻代的**使用情况；
* (12268352K) – **整个年轻代的容量**；
* 0.1768324 secs – Duration for the collection w/o final cleanup.
* 11387196K->434763K – 收集前后整个**堆**的使用情况；
* (40579904K) – **整个堆的容量**；
* 0.1777042 secs – ParNew**收集器标记**和**复制年轻代活着的对象**所花费的时间（包括和老年代通信的开销、对象晋升到老年代时间、垃圾收集周期结束一些最后的清理对象等的花销）；
* [Times: user=1.60 sys=0.28, real=0.18 secs] – GC事件在不同维度的耗时，具体的用英文解释起来更加合理:
  * user – Total CPU time that was consumed by Garbage Collector threads during this collection
  * sys – Time spent in OS calls or waiting for system event
  * real – Clock time for which **your application was stopped**. With **Parallel GC** this number should be close to (**user time + system time) divided by the number of threads** used by the Garbage Collector. In this particular case 8 threads were used. Note that due to some activities not being parallelizable, it always exceeds the ratio by a certain amount.

**分析一下对象晋级问题**：
收集前： 整个堆使用：11387196K， 年轻代使用：11387196K， 老年代使用：整个堆11387196K - 年轻代11387196K = 0？ 
  - 整个老年代容量：（40579904K - 12268352K） 比率是 Eden 1: Survivor 2: Old 7  
收集后： 整个堆使用：434763K， 年轻代使用：434763K， 老年代使用：0
收集清理了： 年轻代使用：11387196K - 年轻代使用434763K = 
  - 说明这次垃圾回收没有晋升到老年代的： 而且有434763K 对象的age + 1。


#### Full GC？ CMS不属于Full GC， 只收集老年代

```log
2022-03-07T11:41:38.496+0800: 61.871: [GC (CMS Initial Mark) [1 CMS-initial-mark: 0K(28311552K)] 8020070K(40579904K), 0.1120113 secs] [Times: user=1.27 sys=0.00, real=0.11 secs]
2022-03-07T11:41:38.609+0800: 61.985: [CMS-concurrent-mark-start]
2022-03-07T11:41:38.618+0800: 61.994: [CMS-concurrent-mark: 0.009/0.009 secs] [Times: user=0.07 sys=0.02, real=0.01 secs]
2022-03-07T11:41:38.619+0800: 61.994: [CMS-concurrent-preclean-start]
2022-03-07T11:41:38.674+0800: 62.050: [CMS-concurrent-preclean: 0.055/0.055 secs] [Times: user=0.18 sys=0.03, real=0.05 secs]
2022-03-07T11:41:38.679+0800: 62.054: [CMS-concurrent-abortable-preclean-start]
 CMS: abort preclean due to time 
2022-03-07T11:41:43.723+0800: 67.098: [CMS-concurrent-abortable-preclean: 4.580/5.044 secs] [Times: user=10.26 sys=2.00, real=5.04 secs]
2022-03-07T11:41:43.730+0800: 67.105: [GC (CMS Final Remark) [YG occupancy: 9955998 K (12268352 K)]
2022-03-07T11:41:43.730+0800: 67.106: [GC (CMS Final Remark) 
2022-03-07T11:41:43.731+0800: 67.107: [ParNew
Desired survivor size 697925632 bytes, new threshold 6 (max 6)
- age   1:  108073056 bytes,  108073056 total
- age   2:   93901384 bytes,  201974440 total
- age   3:   47244104 bytes,  249218544 total
- age   4:   64773416 bytes,  313991960 total
: 9955998K->404612K(12268352K), 0.0912009 secs] 9955998K->404612K(40579904K), 0.0931721 secs] [Times: user=1.24 sys=0.01, real=0.09 secs]
2022-03-07T11:41:43.824+0800: 67.199: [Rescan (parallel) , 0.0196641 secs]
2022-03-07T11:41:43.844+0800: 67.219: [weak refs processing, 0.0081472 secs]
2022-03-07T11:41:43.852+0800: 67.227: [class unloading, 0.0277837 secs]
2022-03-07T11:41:43.880+0800: 67.255: [scrub symbol table, 0.0171561 secs]
2022-03-07T11:41:43.897+0800: 67.272: [scrub string table, 0.0013598 secs][1 CMS-remark: 0K(28311552K)] 404612K(40579904K), 0.1760652 secs] [Times: user=1.59 sys=0.02, real=0.17 secs]
2022-03-07T11:41:43.908+0800: 67.283: [CMS-concurrent-sweep-start]
2022-03-07T11:41:43.913+0800: 67.288: [CMS-concurrent-sweep: 0.005/0.005 secs] [Times: user=0.01 sys=0.00, real=0.01 secs]
2022-03-07T11:41:43.913+0800: 67.289: [CMS-concurrent-reset-start]
2022-03-07T11:41:44.055+0800: 67.430: [CMS-concurrent-reset: 0.125/0.141 secs] [Times: user=0.30 sys=0.01, real=0.14 secs]
```

**具体步骤**：
1. Initial Mark： **stop-the-world：需要暂停用户进程，需尽量避免、缩短，简称STW**
   -  目标
      -  标记老年代中所有的GC Root
      -  标记被年轻代中活着的对象引用的对象
   - 日志分析：
     - 日志：`2022-03-07T11:41:38.496+0800: 61.871: [GC (CMS Initial Mark) [1 CMS-initial-mark: 0K(28311552K)] 8020070K(40579904K), 0.1120113 secs] [Times: user=1.27 sys=0.00, real=0.11 secs]`
     - GC时间开始时间： 相对于JVM启动时间的相对时间：
     - [收集阶段， 开始收集所有的GC Roots和直接引用到的对象]
     - 0K(28311552K) ： 老年代当前使用情况（老年代容量）
     - 8020070K(40579904K) ： 整个堆当前使用情况（整个堆的容量）
     - 0.1120113 secs ： 时间？
     - [Times: user=1.27 sys=0.00, real=0.11 secs] 同上解释过了


2. Concurrent Mark
   - 目标：
     - 遍历整个老年代并且标记所有存活对象， 从上一步找到的GC Roots开始。
     - **并发标记**的特点是和应用程序线程同时运行
     - 并不是老年代的多有存活对象都会被标记，因为标记的同时应用程序会改变一些对象的引用等
   - 日志分析：
     - 日志：
    ```log
    2022-03-07T11:41:38.609+0800: 61.985: [CMS-concurrent-mark-start]
    2022-03-07T11:41:38.618+0800: 61.994: [CMS-concurrent-mark: 0.009/0.009 secs] [Times: user=0.07 sys=0.02, real=0.01 secs]      
    ```
    - 开始， 该阶段会遍历整个老年代并且标记活着的对象。
    - 该阶段持续的时间


3. Concurrent Preclean
   - 目标：前一个阶段在并行运行时，一些对象的引用已经发生了变化，当这些应用发生辩护的时候，JVM会标记堆的这个区域为 Dirty Card（包含被标记且改变了的对象，卡表？）
     - 此阶段那些能够从Dirty Card对象到达的对象也会被标记，标记完后dirty标记就会被清楚。
   - 日志分析：
     - 日志：
    ```log
    2022-03-07T11:41:38.619+0800: 61.994: [CMS-concurrent-preclean-start]
    2022-03-07T11:41:38.674+0800: 62.050: [CMS-concurrent-preclean: 0.055/0.055 secs] [Times: user=0.18 sys=0.03, real=0.05 secs]     
    ```
    - 这个阶段：负责标记前一个阶段标记后又发送改变的对象。
    - 其他同上


4. Concurrent Abortable Preclean 可终止的并发预清理
   - 目标：
     - 尝试着去承担STW的Final Remark阶段足够多的工作。
     - 此阶段持续的时间依赖很多因素：由于这个阶段是重复的做相同的事情，知道发生aboart的条件之一才会停止（比如：重复次数、多少工作量、持续的时间等条件）。
   - 日志分析：
     - 日志：
    ```log
    2022-03-07T11:41:38.679+0800: 62.054: [CMS-concurrent-abortable-preclean-start]CMS: abort preclean due to time 
    2022-03-07T11:41:43.723+0800: 67.098: [CMS-concurrent-abortable-preclean: 4.580/5.044 secs] [Times: user=10.26 sys=2.00, real=5.04 secs] secs]     
    ```
    - 可终止的并发预清理，主要目的是试图尽可能缩短下一步Final Remark的时间（需要STW）
    - 4.580/5.044 secs 默认情况下，此阶段最长可持续5s
    - 其他同上


5. Final Remark： **需要STW**
   - 目标：
     - 标记整个老年代的所有存活对象
     - 由于之前的预处理是并发的，它可能跟不上应用程序改变的速度，这个时候就需要STW来完成最后的校准工作。
     - 通常CMS尽量在年轻代足够干净的时候进入Final Remark阶段，目的是消除紧接着的连续多个STW阶段？
   - 日志分析：
     - 日志：
    ```log
    2022-03-07T11:41:43.730+0800: 67.105: [GC (CMS Final Remark) [YG occupancy: 9955998 K (12268352 K)]
    2022-03-07T11:41:43.730+0800: 67.106: [GC (CMS Final Remark)  
    <!-- 中间插入了一个ParNew的日志，待确认是否误入，还是新版本remark就是加了这条日志 -->
    2022-03-07T11:41:43.731+0800: 67.107: [ParNew
    Desired survivor size 697925632 bytes, new threshold 6 (max 6)
    - age   1:  108073056 bytes,  108073056 total
    - age   2:   93901384 bytes,  201974440 total
    - age   3:   47244104 bytes,  249218544 total
    - age   4:   64773416 bytes,  313991960 total
    : 9955998K->404612K(12268352K), 0.0912009 secs] 9955998K->404612K(40579904K), 0.0931721 secs] [Times: user=1.24 sys=0.01, real=0.09 secs]
    2022-03-07T11:41:43.824+0800: 67.199: [Rescan (parallel) , 0.0196641 secs]
    2022-03-07T11:41:43.844+0800: 67.219: [weak refs processing, 0.0081472 secs]
    2022-03-07T11:41:43.852+0800: 67.227: [class unloading, 0.0277837 secs]
    2022-03-07T11:41:43.880+0800: 67.255: [scrub symbol table, 0.0171561 secs]
    2022-03-07T11:41:43.897+0800: 67.272: [scrub string table, 0.0013598 secs][1 CMS-remark: 0K(28311552K)] 404612K(40579904K), 0.1760652 secs] [Times: user=1.59 sys=0.02, real=0.17 secs]
    ```
    - 收集阶段，这个阶段会标记老年代所有存活对象，包括那些在并发标记阶段更改的或者新建的引用对象
    - [YG occupancy: 9955998 K (12268352 K)] ： 年轻代当前占用情况和容量
    - 在停止应用之前，先清理一下年轻代，因为配了： XX:+CMSScavengeBeforeRemark
    - [Rescan (parallel) , 0.0196641 secs] : 这个阶段在**应用停止的阶段**完成存活对象的标记工作，这一步会扫描年轻代
    - [weak refs processing, 0.0081472 secs] ： 子阶段一， 处理弱引用
    - [class unloading, 0.0277837 secs] ： 子阶段二， 卸载那些不适用的类
    - [scrub symbol table, 0.0171561 secs] ： 子阶段三， 清理symbol table
    - [scrub string table, 0.0013598 secs] ： 子阶段四，that is cleaning up symbol and string tables which hold class-level metadata and internalized string respectively
    - [1 CMS-remark: 0K(28311552K)] ： 此阶段后 老年代使用内存大小和容量
    - 404612K(40579904K), 0.1760652 secs] ： 此阶段后整个堆使用内存大小和容量
    - 其他时间同上


>通过以上五个阶段的标记，老年代所有存活的对象已经被标记，并且现在要通过垃圾回收采用清扫的方式，回收哪些不在使用的对象了。


6. Concurrent Sweep
   - 目标：
     - 并发移除那些不用的对象，回收他们占用的空间供未来使用。
   - 日志分析：
     - 日志：
    ```log
    2022-03-07T11:41:43.908+0800: 67.283: [CMS-concurrent-sweep-start]
    2022-03-07T11:41:43.913+0800: 67.288: [CMS-concurrent-sweep: 0.005/0.005 secs] [Times: user=0.01 sys=0.00, real=0.01 secs]   
    ```
    - 同上


7. Concurrent Reset
   - 目标：
     - 重新设置CMS算法内部的数据结构，准备下一个CMS生命周期使用
   - 日志分析：
     - 日志：
    ```log
    2022-03-07T11:41:43.913+0800: 67.289: [CMS-concurrent-reset-start]
    2022-03-07T11:41:44.055+0800: 67.430: [CMS-concurrent-reset: 0.125/0.141 secs] [Times: user=0.30 sys=0.01, real=0.14 secs]   
    ```
    - 同上

### CMS优化建议

1. 一般CMS的GC耗时 80%都在remark阶段，如果发现remark阶段停顿时间很长，可以尝试添加该参数：`-XX:+CMSScavengeBeforeRemark`
   - 用来开启或关闭在 CMS-remark 阶段之前的清除（Young GC）尝试。如果开启，在CMS开始前，会进行一次年轻代的清理，也就是为啥我看很多CMS的remark前都有一次ParNew（并行清理年轻代垃圾的收集器）
   - 为啥提前清理年轻代，可以减少CMS的remark阶段？
    >由于 YoungGen 存在引用 OldGen 对象的情况，因此 CMS-remark 阶段会将 YoungGen 作为 OldGen 的 “GC ROOTS” 进行扫描，防止回收了不该回收的对象。而配置 -XX:+CMSScavengeBeforeRemark 参数，在 CMS GC 的 CMS-remark 阶段开始前先进行一次 Young GC，有利于减少 Young Gen 对 Old Gen 的无效引用，降低 CMS-remark 阶段的时间开销。


2. CMS是基于标记-清除算法的，只会将标记为为存活的对象删除，并不会移动对象整理内存空间，**会造成内存碎片**，这时候我们需要用到这个参数:`XX:CMSFullGCsBeforeCompaction=n`
   - CMS GC要决定是否在full GC时做压缩，会依赖几个条件，下面三种条件的**任意一种**成立都会让CMS决定这次做full GC时**要做压缩**。
     1. UseCMSCompactAtFullCollection 与 `CMSFullGCsBeforeCompaction` 是搭配使用的；前者目前默认就是true了，也就是关键在后者上。
     2. 用户调用了System.gc()，而且DisableExplicitGC没有开启。
     3. young gen报告接下来如果做增量收集会失败；简单来说也就是young gen预计old gen没有足够空间来容纳下次young GC晋升的对象。


3. 执行CMS GC的过程中，同时业务线程也在运行，当年轻带空间满了，执行ygc时，需要将存活的对象放入到老年代，而此时老年代空间不足。
   - 产生可能原因：
     - CMS还没有机会回收老年带产生的，
     - 或者在做Minor GC的时候，新生代Survivor空间放不下，需要放入老年代，而老年代也放不下而产生concurrent mode failure。
   -  确定发生concurrent mode failure的原因是因为碎片造成的，还是Eden区有大对象直接晋升老年代造成的？
     - **优化方法**：一般有大量的对象晋升老年代容易导致这个错，有优化空间，要保证大部分对象尽肯能的再新生代gc掉。
   - **经验判断**：在进行Minor GC时，Survivor Space放不下，对象只能放入老年代，而此时老年代也放不下造成的，多数是由于**老年代有足够的空闲空间，但是由于碎片较多**，新生代要转移到老年带的对象比较大,找不到一段连续区域存放这个对象导致的promotion failed。


4. 过早提升与提升失败

- 过早提升（Premature Promotion）：在 Minor GC 过程中，Survivor Unused 可能不足以容纳 Eden 和另一个 Survivor 中的存活对象， 那么多余的将被移到老年代。
  - 这会导致老年代中短期存活对象的增长， 可能会引发严重的性能问题。 
  - **早提升的原因**：
    1. Survivor空间太小，容纳不下全部的运行时短生命周期的对象，如果是这个原因，可以尝试将Survivor调大，否则端生命周期的对象提升过快，导致老年代很快就被占满，从而引起频繁的full gc；
    2. 对象太大，Survivor和Eden没有足够大的空间来存放这些大象；
- 提升失败（Promotion Failure）：如果老年代满了， Minor GC 后会进行 Full GC， 这将导致遍历整个堆。
  - **提升失败原因**
    - 当提升的时候，发现老年代也没有足够的连续空间来容纳该对象。为什么是没有足够的连续空间而不是空闲空间呢？老年代容纳不下提升的对象有两种情况，多数情况是后者：
      1. 老年代空闲空间不够用了；
      2. 老年代虽然空闲空间很多，但是碎片太多，没有连续的空闲空间存放该对象；
     - 解决方法
    1. 如果是因为内存碎片导致的大对象提升失败，cms需要进行空间整理压缩；
    2.  如果是因为提升过快导致的，说明Survivor空闲空间不足，那么可以尝试调大Survivor；
    3.  如果是因为**老年代空间不够导致的，尝试将CMS触发的阈值调低**；


5. 导致回收停顿时间变长原因
   
>linux使用了swap，内存换入换出（vmstat），尤其是开启了大内存页的时候，因为swap只支持4k的内存页，大内存页的大小为2M，大内存页在swap的交换的时候需要先将swap中4k内存页合并成一个大内存页再放入内存或将大内存页切分为4k的内存页放入swap，合并和切分的操作会导致操作系统占用cup飙高，用户cpu占用反而很低；除了swap交换外，网络io（netstat）、磁盘I/O （iostat）在 GC 过程中发生会使 GC 时间变长。

如果是以上原因，就要去查看gc日志中的Times耗时：

`[Times: user=0.00 sys=0.00, real=0.00 secs]`

- user是用户线程占用的时间，sys是系统线程占用的时间，如果是io导致的问题，会有两种情况

  1. user与sys时间都非常小，但是real却很长，如下：`[ Times: user=0.51 sys=0.10, real=5.00 secs ]`
     >user+sys的时间远远小于real的值，这种情况说明停顿的时间并不是消耗在cup执行上了，不是cup肯定就是io导致的了，所以这时候要去检查系统的io情况。
  2. sys时间很长，user时间很短，real几乎等于sys的时间，如下：`[ Times: user=0.11 sys=31.10, real=33.12 secs ]`
     >这时候其中一种原因是开启了大内存页，还开启了swap，大内存进行swap交换时会有这种现象；


6. 增加线程数

- CMS默认启动的**回收线程数目**是 (ParallelGCThreads + 3)/4) ，这里的ParallelGCThreads是**年轻代的并行收集**线程数；

  - 年轻代的并行收集线程数默认是(ncpus <= 8) ? ncpus : 3 + ((ncpus * 5) / 8)；
  - 如果要直接设定CMS回收线程数，可以通过`-XX:ParallelCMSThreads=n`，注意这个**n不能超过cpu线程数**，需要注意的是增加gc线程数，就会和应用争抢资源


CMS并发GC不是“full GC”。

HotSpot VM里对concurrent collection和full collection有明确的区分。所有带有“FullCollection”字样的VM参数都是跟真正的full GC相关，而跟CMS并发GC无关的，**cms收集算法只是清理老年代**。


#### 参考文章

- [CMS日志分析](https://www.cnblogs.com/zhangxiaoguang/p/5792468.html)
- 《深入理解Java虚拟机-JVM高级特性与最佳实践》第三版
- [简书-爱吃糖果的：CMS日志分析](https://www.jianshu.com/p/03fac9502311)


# GC算法

## 基础常识

### 如何判断哪些对象是需要回收的

#### 引用计数法

现在已经不用了

#### 可达性分析法
通过GC Roots的对象作为起始点，向下搜索，无法到达的就认为这些对象不可用。

**什么是GC Roots**：
- 一般都是哪些**堆外指向对内**的引用
  - JVM栈中引用的对象
  - 方法去中静态属性引用的对象
  - 方法区中常量引用的对象
  - 本地方法栈中引用的对象