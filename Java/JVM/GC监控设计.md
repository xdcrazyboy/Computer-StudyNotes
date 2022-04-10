

## GC监控指标（主要关注STW的几个时间）

### ParNew：

出现ParNew后，紧跟着的Time里的real时间。 如`[Times: user=1.24 sys=0.01, real=0.09 secs]` 的0.09s。
```java
2022-03-07T11:41:43.731+0800: 67.107: [ParNew
Desired survivor size 697925632 bytes, new threshold 6 (max 6)
- age   1:  108073056 bytes,  108073056 total
- age   2:   93901384 bytes,  201974440 total
- age   3:   47244104 bytes,  249218544 total
- age   4:   64773416 bytes,  313991960 total
: 9955998K->404612K(12268352K), 0.0912009 secs] 9955998K->404612K(40579904K), 0.0931721 secs] [Times: user=1.24 sys=0.01, real=0.09 secs]
```

### CMS 关注两个STW的步骤
- **CMS Initial Mark**： 一般比较短，主要识别`CMS Initial Mark`，同上，也是找紧跟其后的 real时间
```java
[GC (CMS Initial Mark) [1 CMS-initial-mark: 0K(16128000K)] 2111086K(24837120K), 0.2589686 secs] [Times: user=0.33 sys=0.12, real=0.26 secs]
```

- **CMS Final Remark**: 也是获取紧跟其后的real时间， 目前发现两种情况，一种简单的（如下第一条）。一种中间会包含一次ParNew（这里各记录各的，一个时间寄到两个类型就行。）
```java
2022-03-29T19:11:11.228+0800: 17.707: [GC (CMS Final Remark) [YG occupancy: 6657732 K (8709120 K)]2022-03-29T19:11:11.228+0800: 17.707: [Rescan (parallel) , 0.3269928 secs]2022-03-29T19:11:11.555+0800: 18.034: [weak refs processing, 0.0004249 secs]2022-03-29T19:11:11.556+0800: 18.035: [class unloading, 0.0064031 secs]2022-03-29T19:11:11.562+0800: 18.041: [scrub symbol table, 0.0612921 secs]2022-03-29T19:11:11.624+0800: 18.102: [scrub string table, 0.0007930 secs][1 CMS-remark: 0K(16128000K)] 6657732K(24837120K), 0.3975450 secs] [Times: user=0.76 sys=0.01, real=0.40 secs]


2022-03-14T14:17:11.303+0800: 316563.518: [GC (CMS Final Remark) [YG occupancy: 6327134 K (12268352 K)]
2022-03-14T14:17:11.303+0800: 316563.518: [GC (CMS Final Remark) 2022-03-14T14:17:11.304+0800: 316563.519: [ParNew (promotion failed): 6327134K->5311815K(12268352K), 1.0276246 secs] 24109239K->23094088K(40579904K), 1.0297185 secs] [Times: user=2.12 sys=0.12, real=1.03 secs]
```

上面说的总共记录成2个指标： ParNew、CMS（两类加起来）

> 如果这两个指标在**一分钟内持续>3s**就报警（暂定这个时间，后续可能调整）


## 比较严重的GC（包含full GC）
>下面集中GC情况，只要出现，**不管耗时**长短（一般都不会太短），**都报警**。

  - **Full GC** ： 出现[Full GC 这个词就记录后面的real时间。
  -**promotion failed**: 一般形如`[ParNew (promotion failed)` ,出现关键词 **promotion failed** 就记录后面的real，记录成此类型。（ParNew也可能重复记录，先不用去重）
  - **concurrent mode failure**： 出现关键词 **concurrent mode failure** 就记录后面的real时间。



**解释**后面两种情况：
- 当**young gc**的时候，把eden和survivor里的都还存活的对象，统一移到另一个survivor区中时，发现装不下了，就需要把部分对象，放到老年代中去，结果**老年代空间也不足**，这种场景呢，叫做**promotion failed**。
- 在**promotion failed**的前提下，**老年代恰好还正在full gc**，那么就会有图1红框5中的字样提示，**concurrent mode failure**。

