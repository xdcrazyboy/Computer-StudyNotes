## 强引用、软引用、弱引用、幻象引用有什么区别?


# 集合

## ArrayList、Vector、LinkedList有何区别？

- Vector 是 Java 早期提供的**线程安全**的动态数组，如果不需要线程安全，并不建议选择， 毕竟同步是有额外开销的。扩容会创建新的数组，并拷贝原来的数组数据。 扩容1倍。 
- ArrayList 是应用更加广泛的动态数组实现，它本身不是线程安全的，所以性能要好很多。扩容50%。
- LinkedList 是双向链表，不需要调整容量，**不是线程安全**。


如果事先可以估计到，应用操作是偏向于插入、删除，还是随机访问较多，就可以针对性的进行选择。

* **TreeSet** 支持自然顺序访问，但是添加、删除、包含等操作要相对低效(log(n) 时 间)。
* **HashSet** 则是利用哈希算法，理想情况下，如果哈希散列正常，可以提供常数时间的添 加、删除、包含等操作，但是它不保证有序。
* **LinkedHashSet**，内部构建了一个记录插入顺序的双向链表，因此提供了按照插入顺序 遍历的能力，与此同时，也保证了常数时间的添加、删除、包含等操作，这些操作性能略 低于 HashSet，因为需要维护链表的开销。
  * 在遍历元素时，HashSet 性能受自身容量影响，所以初始化时，除非有必要，不然不要 将其背后的 HashMap 容量设置过大。
  * 而对于 LinkedHashSet，由于其内部链表提供的 方便，遍历性能只和元素多少有关系。


**集合排序**

 Java 提供的默认排序算法，具体是什么排序方 式以及设计思路等。

这个问题本身就是有点陷阱的意味，因为需要区分是 Arrays.sort() 还是 Collections.sort() (底层是调用 Arrays.sort());什么数据类型;多大的数据集(太小的数据集，复杂排序 是没必要的，Java 会直接进行二分插入排序)等。
对于原始数据类型，目前使用的是所谓双轴快速排序(Dual-Pivot QuickSort)，是一 种改进的快速排序算法，早期版本是相对传统的快速排序，你可以阅读源码。


而对于对象数据类型，目前则是使用TimSort，思想上也是一种归并和二分插入排序 (binarySort)结合的优化排序算法。TimSort 并不是 Java 的独创，简单说它的思路是 查找数据集中已经排好序的分区(这里叫 run)，然后合并这些分区来达到排序的目的。


另外，Java 8 引入了并行排序算法(直接使用 parallelSort 方法)，这是为了充分利用现 代多核处理器的计算能力，底层实现基于 fork-join 框架(专栏后面会对 fork-join 进行相 对详细的介绍)，当处理的数据集比较小的时候，差距不明显，甚至还表现差一点;但是， 当数据集增长到数万或百万以上时，提高就非常大了，具体还是取决于处理器和系统环境。


##  Hashtable、HashMap、TreeMap 有什么不同?

- Hashtable： 同步，不支持null键和值。
- HashMap： 不同步，支持 null 键和值。 存取时间接近常数。
- TreeMap 则是基于红黑树的一种提供**顺序**访问的 Map，它的 get、 put、remove 之类操作都是 O(log(n))的时间复杂度 


**Map**

**继承关系**
- Dictionary：HashTable：Properties
- AbstractMap： HashMap：LinkedHashMap、TreeMap、EnumMap
- EnumMap、HashMap、SortedMap


HashMap 的性能表现非常依赖于哈希码的有 效性，请务必掌握 **hashCode 和 equals **的一些基本约定，比如:

- equals 相等，hashCode 一定要相等。
- 重写了 hashCode 也要重写 equals。
- hashCode 需要保持一致性，状态改变返回的哈希值仍然要一致。
- equals 的对称、反射、传递等特性。


LinkedHashMap 和 TreeMap 都可以保证某种**顺序**，但二者还是非常不同的。

- LinkedHashMap 通常提供的是**遍历顺序符合插入顺序**，它的实现是通过为条目(键值 对)维护一个双向链表。注意，通过特定构造函数，我们可以创建反映访问顺序的实例， 所谓的 put、get、compute 等，都算作“访问”。 **插入顺序就是遍历读取的顺序？ get操作也算是访问，如果get过会被放在前面**
  - 用处：我们构建一个空间占用敏感的资源池，希望可以自动将最不常被访问的对象释放掉，这就可以利用 LinkedHashMap 提供的机制来实现。
  - 构建一个具有优先级的调度系统的问题，其本质就是个 典型的优先队列场景，Java 标准库提供了基于二叉堆实现的 PriorityQueue，它们都是依赖于**同一种排序机制**，当然也包括 TreeMap 的马甲 TreeSet。


### HashMap 源码分析

**HashMap 内部实现基本点分析**

1. HashMap内部结构可以看作是数组(Node<K,V>[] table)和链表结合组成的复合结构，数组被分为一个个桶(bucket)，通过哈希值决定了 键值对在这个数组的寻址;
2. 哈希值相同的键值对，以链表形式存储，链表大小超过（8），就会被改造为树形结构。


3. putVal 方法本身逻辑非常集中，从初始化、扩容到树化，全部都和它有关。

```java
final V putVal(int hash, K key, V value, boolean onlyIfAbent, boolean evit) {
    Node<K,V>[] tab; Node<K,V> p; int , i;
    if ((tab = table) == null || (n = tab.length) = 0)
        n = (tab = resize()).length;
    if ((p = tab[i = (n - 1) & hash]) == ull)
        tab[i] = newNode(hash, key, value, nll);
    else {
        // ...
        if (binCount >= TREEIFY_THRESHOLD - 1) // -1 for first
           treeifyBin(tab, hash);
        //  ...
} }
```

- 如果表格是 null，resize 方法会负责初始化它，这从 tab = resize() 可以看出。


4. **resize方法**兼顾两个职责：
   - 创建**初始存储表格**（如上一段代码）
   - 在容量不满足需求的时候（如下一段代码），进行扩容(resize)。


- `threshold=newCap * loadFator;` 如果构建HashMap时没指定就用默认常量值。
- threshold通常以倍数进行调整(newThr = oldThr << 1)。

```java
if (++size > threshold) resize();
```

- 具体键值对在哈希表中的位置(数组 index)取决于下面的位运算: `i = (n - 1) & hash`。
  
- **hash值**
  - 为什么这里需要将高位数据移位到低位进行异或运算呢?
    >这是因为有些数据计算出的哈希值差异主要在高位，而 HashMap 里的哈希 寻址是忽略容量以上的高位的，那么这种处理就可以有效避免类似情况下的哈希碰撞。
```java
static final int hash(Object kye) {
  int h;
  return (key == null) ? 0 : (h = key.hashCode()) ^ (h >>>16;
}
```

- 扩容后，需要将老的数组中的元素重新放置到新的数组。这是扩容的主要开销来源。



**容量(capacity)和负载系数(load factor)**


容量和负载系数**决定了可用的桶的数量**，空桶太多会浪费空间，如果使用的太满则
会严重影响操作的性能。极端情况下，假设只有一个桶，那么它就退化成了链表。


既然容量和负载因子这么重要，我们**在实践中应该如何选择**呢?

- 预先设置的容量需要满足，大于“预估元素数量 / 负载因子”，同时它是 2 的幂 数，结论已经非常清晰了。

- 如果没有特别需求，不要轻易进行更改，因为 JDK 自身的默认负载因子是非常符合通用 场景的需求的。
- 
- 如果确实需要调整，建议不要设置超过 0.75 的数值，因为会显著增加冲突，降低 HashMap 的性能。
- 如果使用太小的负载因子，按照上面的公式，预设容量值也进行调整，否则可能会导致更加频繁的扩容，增加无谓的开销，本身访问性能也会受影响。




**树化**

对应逻辑主要在 putVal 和 treeifyBin 中：


```java
 final void treeifyBin(Node<K,V>[] tab, int hash) {
    int n, index; Node<K,V> e;
    if (tab == null || (n = tab.length) < MIN_TREEIFY_CAPACITY) 
      resize();
    else if ((e = tab[index - (n - 1) & hash]) != null){
      //树化逻辑
    }

```

当 bin 的数量大于 TREEIFY_THRESHOLD 时:
- 如果容量小于 MIN_TREEIFY_CAPACITY，只会进行简单的扩容。 
- 如果容量大于 MIN_TREEIFY_CAPACITY ，则会进行树化改造。


## 如何保证容器是线程安全的?ConcurrentHashMap 如何实现高 效地线程安全?


### ConcurrentHashMap 分析

**演化的过程**


#### 1. 早期Java7： 分段、分离锁
- 分离锁，也就是将内部进行分段(Segment)，里面则是 HashEntry 的数组，和 HashMap 类似，哈希相同的条目也是以链表形式存放。
- HashEntry 内部使用 volatile 的 value 字段来保证可见性，也利用了不可变对象的机制以改进利用 Unsafe 提供的底层能力，比如 volatile access，去直接完成部分操作，以最优化性能，毕竟 Unsafe 中的很多操作都是 JVM intrinsic 优化过的。
- **Segment 的数量**由所谓的 concurrentcyLevel 决定，默认是 16，也可以 在相应构造函数直接指定。注意，Java 需要它是 2 的幂数值

-  get 操作：需要保证的是可见性， 所以并没有什么同步逻辑。
   -  计算hash，找到位置
   -  Unsafe直接进行volatile access
-  put 操作： 首先是通过二次哈希避免哈希冲突，然后以 Unsafe 调用方式，直接获 取相应的 Segment，然后进行线程安全的 put 操作。
   -  put要先获取锁，锁的是segment。
   -  ConcurrentHashMap 会获取再入锁，以保证数据一致性，Segment 本身就是基于 ReentrantLock 的扩展实现，所以，在并发修改期间，相应 Segment 是被锁定的。
   -  在最初阶段，进行重复性的扫描，以确定相应 key 值是否已经在数组里面，进而决定是 更新还是放置操作。重复扫描、检测冲突是 ConcurrentHashMap 的常见技巧。
   -  扩容是单独对Segment扩容。
-  **分离锁副作用**： size 方法 计算 和初始化操作耗时。如果不进行同步，简单的计算所有 Segment 的总值，可能会因为并发 put，导致结 果不准确，但是直接锁定所有 Segment 进行计算，就会变得非常昂贵。
   -  ConcurrentHashMap 的实现是通过重试机制(RETRIES_BEFORE_LOCK，指定重 试次数 2)，来试图获得可靠值。
   -  如果没有监控到发生变化(通过对比 Segment.modCount)，就直接返回，否则获取锁进行操作。


#### 2. Java8 的一些变化：去掉分段，CAS、volatile、Unsafe

1. 其内部仍然有 Segment 定义，但仅仅是**为了保证序列化时的兼容性**而已，**不再有任何结构上的用处**。
2. 因为不再使用 Segment，**初始化操作大大简化**，修改为 **lazy-load** 形式，这样可以有效 避免初始开销。
3. 数据存储利用 volatile 来保证可见性。
4. 使用 CAS 等操作，在特定场景进行无锁并发操作。
5. 使用 Unsafe、LongAdder 之类底层手段，进行极端情况的优化。
6. 1.8以后的锁的颗粒度，是加在链表头上的


具体看数据存储内部实现：
- key是final的，因为在生命周期中，一个条目的 Key不可能变化;
- val声明为volatile，以保证其可见性。


**put方法的实现**

补充源码


- 在同步逻辑上，它使用的是 synchronized。 
  >synchronized相比于 ReentrantLock，它可以减少内存消耗，这是个 非常大的优势。

- size计算： 
  - 真正的逻辑是在 sumCount 方法中。 思路仍然和以前类似，都是分而治之的进行计数，然后求和处理，但实现却 基于一个奇怪的 CounterCell。
  - 对于 CounterCell 的操作，是基于 java.util.concurrent.atomic.LongAdder 进行 的，是一种 JVM 利用空间换取更高效率的方法，利用了Striped64内部的复杂逻辑。



## 设计模式

- 希望你写一个典型的设计模式实现。这虽然看似简单，但即使是最简单的单例，也能够综合考察代码基本功。


- 考察典型的设计模式使用，尤其是**结合标准库或者主流开源框架**，考察你对业界良好实践的掌握程度。


### 写个单例模式



2. 利用内部类持有静态对象的方式实现，其理论依据是对象初始化过程 中隐含的初始化锁。
```java
public class Singleton {
  private Singleton() {}
  private static Singleton getSingleton(){
    return Holder.singletion;
  }

  private static class Holder {
    private static Singleton single = new Singletion();
  }
}

```


3. 其实实践中未必需要如此复杂，如果我们看 Java 核心类库自己的 单例实现，比如java.lang.Runtime，你会发现: **它并没使用复杂的双检锁之类。**
   1. 静态实例被声明为 final，这是被通常实践忽略的，一定程度保证了实例不被篡改
```java
public class Runtime {
  private static final Runtime currentRuntime = new Runtime();
  private static Version version;
  //..
  public static Runtime getRuntime(){
    return currentRuntime;
  }

  private Runtime() {}
}

```


### Spring用了哪些设计模式？

- BeanFactory和ApplicationContext应用了工厂模式。
- 在 Bean 的创建中，Spring 也为不同 scope 定义的对象，提供了单例和原型等模式实现。
- AOP 领域则是使用了代理模式、装饰器模式、适配器模式等。 
- 各种事件监听器，是观察者模式的典型应用。
- 类似 JdbcTemplate 等则是应用了模板模式。


### 11. Java提供了哪些IO方式? NIO如何实现多路复用?

1. 传统的 java.io 包，它基于**流模型**实现，提供了我们最熟知的一些 IO 功能，比如 File 抽象、输入输出流等。
   1. 交互方式是同步、阻塞的方式。 读取写入流在读写动作完成之前，线程会一直阻塞在那里，他们之间调用是可靠的线性顺序。
   2. 好处是代码比较简单直观，缺点是IO效率和扩展性存在局限。
2. Java 1.4 中引入了 NIO 框架(java.nio 包)，提供了 Channel、Selector、 Buffer 等新的抽象，可以构建**多路复用的、同步非阻塞** IO 程序，同时提供了更接近操作系 统底层的高性能数据操作方式。
3. 在 Java 7 中，NIO 有了进一步的改进，也就是 NIO 2，引入了**异步非阻塞** IO 方 式，也有很多人叫它 AIO(Asynchronous IO)。
   1. 异步 IO 操作**基于事件和回调**机制
   2. 可以理解为：应用操作直接返回，而不会阻塞在那里，当后台处理完成，操作系统会通知相 应线程进行后续工作。



#### BIO、NIO、NIO 2(AIO)


**基础 API 功能与设计， InputStream/OutputStream 和 Reader/Writer 的关系和区别。**

- 输入流、输出流(InputStream/OutputStream)是用于读取或写入字节的，例如操作 图片文件。
- 而 Reader/Writer 则是用于操作字符，增加了字符编解码等功能，适用于类似从文件中 读取或者写入文本信息。
- Reader/Writer 相当于构建了应用逻辑和原始数据之间的桥梁。
- BufferedOutputStream 等带缓冲区的实现，可以避免频繁的磁盘读写，进而提高 IO 处 理效率。
- 很多 IO 工具类都实现了 Closeable 接口，因为需要进行资源的释 放。比如，打开 FileInputStream，它就会获取相应的文件描述符(FileDescriptor)， 需要利用 try-with-resources、 try-finally 等机制保证 FileInputStream 被明确关闭，进而相应文件描述符也会失效，否则将导致资源无法被释放。
- File
- RandomAccessFile
- InputStream
  - FilterInputStream --> BufferedInputStream
  - BytesArrayInputStream、ObjectInputStream、PipeInputStream
- OutputStream
  - FilterOutputStream --> BufferedOutputStream
  - BytesArrayOutputStream、ObjectOutputStream、PipeOutputStream
- Reader
  - InputStreamReader-->FileReader
  - BufferedReader、PipeReader
- Writer
  - OutputStreamWriter-->FileWriter
  - BufferedWriter/PipeWriter


**NIO、NIO 2 的基本组成。**

1.Java NIO 概览

- 首先，熟悉一下 NIO 的主要组成部分:
  - **Buffer**，高效的数据容器，除了布尔类型，所有原始数据类型都有相应的 Buffer 实现。
  - **Channel**，类似在 Linux 之类操作系统上看到的文件描述符，是 NIO 中被用来支持批量 式 IO 操作的一种抽象。File 或者 Socket，通常被认为是比较高层次的抽象，而 Channel 则是更加操作系统底层 的一种抽象
  - **Selector**，是 NIO 实现多路复用的基础，它提供了一种高效的机制，可以检测到注册在 Selector 上的多个 Channel 中，是否有 Channel 处于就绪状态，进而实现了单线程对 多 Channel 的高效管理。



**给定场景，分别用不同模型实现，分析 BIO、NIO 等模式的设计和实现原理。**

2.NIO 能解决什么问题?

 NIO 多路复用：
 - 首先，通过 Selector.open() 创建一个 Selector，作为类似调度员的角色。
 - 然后，创建一个 ServerSocketChannel，并且向 Selector 注册，通过指定 SelectionKey.OP_ACCEPT，告诉调度员，它关注的是新的连接请求。
 - Selector 阻塞在 select 操作，当有 Channel 发生接入请求，就会被唤醒。
 - NIO 则是利用了单线程轮询事件的机制，通过高效地定位就绪的 Channel，来决定做什么，**仅仅 select 阶段是阻塞的**，可以有效避免大量客户端连接时，频繁线程切换带来 的问题。
  
> **局限性**：
> - 当每个channel所进行的都是耗 时操作时，由于是同步操作，就会积压很多channel任务，从而影响性能。
> - 如果回调时客户端做了重操作，就会影响调度，导致后续的client回调缓慢。


AIO 实现异步IO，利用事件和回调处理 Accept、Read 等操作。：
- Future、CompletionHandler 
- Reactor、 Proactor 模式
- 业务逻辑的关键在于，通过指定 CompletionHandler 回调接口，在 accept/read/write 等关键节点，通过事件机制调用。


**NIO 提供的高性能数据操作方式是基于什么原理，如何使用?**

- **Linux 上依赖于 epoll**。
- Windows 上 NIO2(AIO)模式则是依赖于 iocp。


## Java 有几种文件拷贝方式?哪一种最高效?

1. 利用 java.io 类库，直接为源文件构建一个 FileInputStream 读取，然后再为目标文件构建 一个 FileOutputStream，完成写入工作。
2. 利用 java.nio 类库提供的 transferTo 或 transferFrom 方法实现。
3. Java 标准库也提供了文件拷贝方法 (java.nio.file.Files.copy)


>对于 Copy 的效率，这个其实与操作系统和配置等情况相关，总体上来说，NIO **transferTo/From 的方式可能更快**，因为它更能利用现代操作系统底层机制，避免不必要 拷贝和上下文切换。



- [ ] 不同的 copy 方式，底层机制有什么区别? 
- [ ] 为什么零拷贝(zero-copy)可能有性能优势? 
- [ ] Buffer 分类与使用。
- [ ] Direct Buffer 对垃圾收集等方面的影响与实践选择。


1. 拷贝**实现机制**分析

- **用户态空**间(User Space) : 操作系统内核、硬件驱动等运行在内核态空间，具有相对高的特权

- **内核态**空间(Kernel Space) : 给普通应用和服务使用


当我们使用输入输出流进行读写时，实际上是进行了多次上下文切换，比如应用读取数据
时，先在内核态将数据从磁盘读取到内核缓存，再切换到用户态将数据从内核缓存读取到用户缓存。写操作也类似。
>所以，这种方式会带来一定的额外开销，可能会降低 IO 效率。


基于 NIO transferTo 的实现方式，在 Linux 和 Unix 上，则会使用到**零拷贝**技术，数据传输并**不需要用户态参与**，省去了上下文切换的开销和不必要的内存拷贝，进而可能提高应用拷贝性能。


2.Java **IO/NIO 源码结构**


Java 标准库也提供了文件拷贝方法 (java.nio.file.Files.copy)
   - 有几种不太的copy方法，可以自己看源码。可以看到，copy 不仅仅是支持文件之间操作，后面两种 copy 实现，能够在方法实现里直接看到使用的是
   - InputStream.transferTo()，你可以直接看源码，其内部实现其实是 stream 在用户态的读写;
   - NIO 部分代码甚至是定义为模板而不是 Java 源文件，在 build 过程自 动生成源码。原来文件系统实际逻辑存在于 JDK 内部实现里，公共 API 其实是通过 ServiceLoader 机 制加载一系列文件系统实现，然后提供服务。

**如何提高类似拷贝等IO操作的性能**，有一些宽泛的原则:

- 在程序中，使用缓存等机制，合理减少 IO 次数(在网络通信中，如 TCP 传输，window 大小也可以看作是类似思路)。
- 使用 transferTo 等机制，减少上下文切换和额外 IO 操作。
- 尽量减少不必要的转换过程，比如编解码;对象序列化和反序列化，比如操作文本文件或者网络通信，如果不是过程中需要使用文本信息，可以考虑不要将二进制信息转换成字符串，直接传输二进制信息。


3. 掌握 **NIO Buffer**

Java 为每种原始数据类型都提供了相应的Buffer 实现(布尔除外)，所以掌握和使用 Buffer 是十分必要的，尤其是涉及 Direct Buffer 等使用，因为其在垃圾收集等方面的特殊性，更要重点掌握。


Buffer 有几个基本属性:

* capcity，它反映这个 Buffer 到底有多大，也就是数组的长度。
* position，要操作的数据起始位置。
* limit，相当于操作的限额。在读取或者写入时，limit 的意义很明显是不一样的。
  * 读取操作时，很可能将 limit 设置到所容纳数据的上限;
  * 而在写入时，则会设置容量或容 量以下的可写限度。
  * mark，记录上一次 postion 的位置，默认是 0，算是一个便利性的考虑，往往不是必须 的。


4.**Direct Buffer 和垃圾收集**

- Direct Buffer:如果我们看 Buffer 的方法定义，你会发现它定义了 isDirect() 方法，返回当前 Buffer 是否是 Direct 类型。这是因为 Java 提供了堆内和堆外(Direct) Buffer，我们可以以它的 allocate 或者 allocateDirect 方法直接创建。
- MappedByteBuffer:它将文件按照指定大小直接映射为内存区域，当程序访问这个内存 区域时将直接操作这块儿文件数据，省去了将数据从内核空间向用户空间传输的损耗。我们可以使用FileChannel.map创建 MappedByteBuffer，它本质上也是种 Direct Buffer。


Java 会尽量对 Direct Buffer 仅做本地 IO 操作，对于很多大数据量的 IO 密集操作，可能会带来非常大的性能优势，因为:

- Direct Buffer **生命周期内内存地址都不会再发生更改**，进而内核可以安全地对其进行访问，很多 IO 操作会很高效。
- 减少了堆内对象存储的可能额外维护工作，所以访问效率可能有所提高。
- Direct Buffer 创建和销毁过程中，都会比一般的堆内 Buffer 增加部分开销， 所以通常都建议用于**长期使用、数据较大**的场景。
- 使用 Direct Buffer不在堆上， 所以 Xmx 之类参数，其实并不能影响 Direct Buffer 等堆外成员所使用的内存额度，我们可以使用下面参数设置大小:`-XX:MaxDirectMemorySize=512M`
  - 这意味着我们在计算 Java 可以使用的内存大小的时 候，不能只考虑堆的需要，还有 Direct Buffer 等一系列堆外因素。如果出现内存不足，堆 外内存占用也是一种可能性。
  - 大多数垃圾收集过程中，都不会主动收集 Direct Buffer，它的垃圾收集过程是基于Cleaner(一个内部实现)和幻象引用 (PhantomReference)机制。其本身不是 public 类型，内部实现了一个 Deallocator 负 责销毁的逻辑。对它的销毁往往要拖到 full GC 的时候，所以使用不当很容易导致 OutOfMemoryError。


对于 Direct Buffer 的回收，我有几个建议:
- 在应用程序中，显式地调用 System.gc() 来强制触发。
- 另外一种思路是，在大量使用 Direct Buffer 的部分框架中，框架会自己在程序中调用释放方法，Netty 就是这么做的。
- 重复使用 Direct Buffer。


5. 跟踪和诊断 Direct Buffer 内存占用?

- 在 JDK 8 之后的版本，我们可以方便地使用 Native Memory Tracking(NMT)特性来进行诊断。启动参数：`-XX:NativeMemoryTracking={summary|detail}`
- 注意，激活 NMT 通常都会导致 JVM 出现 5%~10% 的性能下降，请谨慎考虑。

## 谈谈接口和抽象类有什么区别?

**接口**

- 接口是对行为的抽象，它是抽象方法的集合，利用接口可以达到 API 定义和实现分离的目的。
- 接口，不能实例化;不能包含任何非常量成员，任何 field 都是隐含着 public static final 的意义;
- 、没有非静态方法实现，也就是说要么是抽象方法，要么是静态方法。 
- Java 标准类库中，定义了非常多的接口，比如 java.util.List。


**抽象类**

- 抽象类是不能实例化的类，用 abstract 关键字修饰 class，其目的主要是代码重用。
- 可以有一个或者多个抽象方法，也可以没有抽象方法。
- 抽象类大多用于抽取相关 Java 类的共用方法实现或者是共同成员变量， 然后通过继承的方式达到代码复用的目的。
- Java 标准库中，比如 collection 框架，很多通 用部分就被抽取成为抽象类，例如 java.util.AbstractList。


```java
public class ArrayList<E> extends AbstractList<E> implements List<E>, RandomAccess, Cloneable, java.io.Serializable
```

**深入理解，考点**


1. 对于 Java 的基本元素的语法是否理解准确。能否定义出语法基本正确的接口、抽象类或 者相关继承实现，涉及重载(Overload)、重写(Override)更是有各种不同的题目。
2. 在软件设计开发中妥善地使用接口和抽象类。你至少知道典型应用场景，掌握基础类库重 要接口的使用;掌握设计方法，能够在 review 代码的时候看出明显的不利于未来维护的设计。
3. 掌握 Java 语言特性演进。现在非常多的框架已经是基于 Java 8，并逐渐支持更新版本， 掌握相关语法，理解设计目的是很有必要的。



**Java 不支持多继承。**

- 这种限制，在规范了代码实现的同时，也产生了一些局限性，影响着程序设计结构。 Java 类可以实现多个接口，因为接口是抽象方法的集合，所以这是声明性的，但不能通过扩展多个抽象类来重用逻辑。
- 在一些情况下存在特定场景，需要抽象出与具体实现、实例化无关的通用逻辑，或者纯调用 关系的逻辑，但是使用传统的抽象类会陷入到单继承的窘境。以往常见的做法是，实现由静 态方法组成的工具类(Utils)，比如 java.util.Collections。
  - 为接口添加任何抽象方法，相应的所有实现了这个接口的类，也必须实现新增方法，否则会出现编译错误。
  - 对于抽象类，如果我们添加非抽象方法，其子类只会享受到能力扩展，而不用担心编译出问题。
- 有一类**没有任何方法的接口**，通常叫作 Marker Interface，顾名思义，它的目的就是为了声明某些东西，比如我 们熟知的 Cloneable、Serializable 等。 类似Annotation，不过后者因为其可以指定参数和值，在表达能力上要更强大一些。


**Java 8 以后，接口也是可以有方法实现的!**
- 对 default method 的支持。
- Default method 提供了一种二进制兼容的扩展已有接口的 办法。比如，我们熟知的 java.util.Collection，它是 collection 体系的 root interface， 在 Java 8 中添加了一系列 default method，主要是增加 Lambda、Stream 相关的功 能。


**面向对象设计**

1. 基本要素：封装、继承、多态。

- **封装**
  - 目的是隐藏事务内部的实现细节，以便提高安全性和简化编程。
  - 封装提供了合理的边 界，避免外部调用者接触到内部的细节。
  - 避免太多无意义的细节浪费调用者的精力
- **继承**
  - 是代码复用的基础机制
  - 但要注意，继承可以看作是非常紧耦合的一种关系，父类代码修改，子类行为也会变动。
  - 过度滥用继承会起到反作用
- **多态**
  - 重写是父子类中相同名字和参数的方法，不同的实现;
  - 重载则是相同名字的方法，但是不同的 参数，本质上这些方法签名是不一样的。
  - 向上转型


2. 基本设计原则：SOLID

- **单一职责**(Single Responsibility)，类或者对象最好是只有单一职责，在程序设计中如 果发现某个类承担着多种义务，可以考虑进行拆分。
- **开关原则**(Open-Close, Open for extension, close for modification)，设计要对扩展开放，对修改关闭。 程序设计应保证平滑的扩展性，尽量避免因为新增同类 功能而修改已有实现，这样可以少产出些回归(regression)问题。
- **里氏替换**(Liskov Substitution)，这是面向对象的基本要素之一，进行继承关系抽象时，**凡是可以用父类或者基类的地方，都可以用子类替换**。
- **接口分离**(Interface Segregation)，我们在进行类和接口设计时，如果在一个接口里 定义了太多方法，其子类很可能面临两难，就是只有部分方法对它是有意义的，这就破坏 了程序的内聚性。
- **依赖反转**(Dependency Inversion)，实体应该依赖于抽象而不是实现。


# 并发

## synchronized和ReentrantLock有什么区别呢?

synchronized 是 Java 内建的同步机制，所以也有人称其为 Intrinsic Locking，它提供了 互斥的语义和可见性。

synchronized 可以用来修饰方法，也可以使用在特定的代码块。


ReentrantLock，通常翻译为再入锁，是 Java 5 提供的锁实现。
- 再入锁通过代码直接调用 lock() 方法获取，代码书写也更加灵 活。
- 必须要明确调用 unlock() 方法释放，不然就会一直持有该锁。


### 理解什么是线程安全

> 推荐看Brain Goetz 等专家撰写的《Java 并发编程实战》

线程安全是一个多线程环境下正确性的概念，也就是保证多线程环境 下共享的、可修改的状态的正确性，这里的状态反映在程序中其实可以看作是数据。


换个角度来看，如果**状态不是共享的**，或者**不是可修改的**，也就不存在线程安全问题，进而可以推理出保证线程安全的两个办法:

- 封装: 通过封装，我们可以将对象内部状态隐藏、保护起来。
- 不可变：  final 和 immutable。 Java 语言目前还没有真正意义上的原生不可变，但是未来也许会引入。


**线程安全**需要保证几个**基本特性**:

- **原子性**，简单说就是相关操作不会中途被其他线程干扰，一般通过同步机制实现。
- **可见性**，是一个线程修改了某个共享变量，其状态能够**立即被其他线程知晓**，通常被解释为将线程本地状态**反映到主内存**上，`volatile`就是负责保证可见性的。
- 有序性，是保证线程内串行语义，避免指令重排等。


### synchronized、ReentrantLock 等机制的基本使用与案例。

原子性： 加synchronized保护起来，使用this作为互斥单元。。。
- 如果用 javap 反编译，可以看到类似片段，利用 `monitorenter/monitorexit` 对实现了同 步的语义。

```java
synchronized (this) {
  int former = sharedState ++;
  int latter = sharedState;
  //...
}

synchronized (ClassName.class) {}
```

**ReentrantLock**

- ReentrantLock是Lock的实现类，是一个互斥的同步器，在多线程高竞争条件下，ReentrantLock比synchronized有更加优异的性能表现。
- Lock使用起来比较灵活，但是必须有释放锁的配合动作


什么是再入？ 
>它是表示当一个线程试图获取一个它已经获取的锁时，这个获取动作就自动成功，这是对锁获取粒度的一个概念，也就是锁的持有是以线程为单位而不是基于调用次数。

Java 锁实现强调再入性是为了和 pthread 的行为进行区分。

再入锁可以设置公平性(fairness)，我们可在创建再入锁时选择是否是公平的，当公平性为真时，会倾向于将锁赋予等待时间最久的
线程。`ReentrantLock fairLock = new ReentrantLock(true);`

- 公平性是减少线程“饥饿”(个别线程长期等待锁，但始终无法获取)情况发生的一个办法。
- 若要保证公平性则会引入额外开销，自然 会导致一定的吞吐量下降。所以，我建议只有当你的程序确实有公平性需要的时候，才有必 要指定它。


ReentrantLock 相比 synchronized，因为可以像普通对象一样使用，所以可以利用其提供 的各种便利方法，进行精细的同步操作，甚至是实现 synchronized 难以表达的用例，如:

- 带超时的获取锁尝试。
- 可以判断是否有线程，或者某个特定线程，在排队等待获取锁。 - 可以响应中断请求。


**条件变量**(java.util.concurrent.Condition)

- 相当于将 wait、notify、notifyAll 等操作转化为相应的对象，将复杂而晦涩的同步操作转变为直观可控的对象行为。
- 条件变量最为典型的应用场景就是标准类库中的 ArrayBlockingQueue 等。


Java6之后，在高竞争情况下，ReentrantLock 仍然有一定优势。并发在4个线程以下synchronized效果更好，越大，lock性能越好。


- [ ] 掌握 synchronized、ReentrantLock 底层实现
- [ ] 理解锁膨胀、降级;
- [ ] 理解偏斜锁、自旋锁、轻量级锁、重量级锁等概念。
- [ ] 掌握并发包中 java.util.concurrent.lock 各种不同实现和案例分析。


## synchronized 底层如何实现? 什么是锁的升级、降级?

synchronized 代码块是由一对儿 monitorenter/monitorexit 指令实现的，**Monitor对象是同步的基本实现单元**。

- 在 Java 6 之前，Monitor 的实现完全是依靠操作系统内部的**互斥锁**，因为**需要进行用户态到内核态的切换**，所以同步操作是一个无差别的重量级操作。
- 所谓锁的升级、降级，就是 JVM 优化 synchronized 运行的机制，当 JVM 检测到不同的竞争状况时，会自动切换到适合的锁实现，这种切换就是锁的升级、降级。
- 初始状态时默认是偏向锁时，线程请求先通过CAS替换mark word中threadId,如果替换成功则该线程持有当前锁。如果 替换失败，锁会升级为轻量级锁，


### 偏向锁

- 当**没有竞争**出现时，**默认会使用偏斜锁**。
- JVM 会利用 CAS 操作在对象头上的`Mark Word`部分设置**线程ID**，以表示这个对象偏向于当前线程。
- 这么做的假设是基于在很多应用场景中，**大部分对象生命周期被一个线程锁定**，使用偏斜锁可以降低无竞争开销。


如果有**另外的线程试图锁定某个已经被偏斜过的对象**，JVM 就需要撤销(revoke)偏斜锁，并**切换到轻量级锁**实现。


### 轻量级锁

- 轻量级锁依赖 CAS 操作 Mark Word来试图获取锁，如果重试成功，就使用普通的轻量级锁;否则，进一步升级为重量级锁。


>当 JVM 进入安全点(SafePoint)的时候，会检查是否有闲置的 Monitor，然后试图进行降级。

作者说---我个人认为，能够基础性地理解这些概念和机制，其实对于大多数并发编程已经足够了，毕竟大部分工程师未必会进行更底层、更基础的研发，很多时候解决的是知道与否，**真正的提高还要靠实践踩坑**。


#### 锁升级过程

synchronized 是 JVM 内部的 Intrinsic Lock，所以偏斜锁、轻量级 锁、重量级锁的代码实现，并不在核心类库部分，而是在 JVM 的代码中。


1. 首先，synchronized 的行为是 JVM runtime 的一部分，所以我们需要先找到 Runtime 相关的功能实现。通过在代码中查询类似`“monitor_enter”`或`“Monitor Enter”`，很直观的就可以定位到代码:
 
```c++
sharedRuntime.cpp/hpp，它是解释器和编译器运行时的基类。
synchronizer.cpp/hpp，JVM 同步相关的各种基础逻辑。

//在 sharedRuntime.cpp 中，下面代码体现了 synchronized 的主要逻辑。

Handle h_obj(THREAD, obj);
  if (UseBiasedLocking) {
    // Retry fast entry if bias is revoked to avoid unnecessary inflation
  ObjectSynchronizer::fast_enter(h_obj, lock, true, CHECK);
  } else {
    ObjectSynchronizer::slow_enter(h_obj, lock, CHECK);
  }

```

- UseBiasedLocking： 是一个检查，因为在JVM启动时可以指定是否启用偏向锁；
- fast_enter 是我们熟悉的完整锁获取路径;
- slow_enter 则是绕过偏斜锁，直接进入轻量级锁获取逻辑。

>偏斜锁并不适合所有应用场景，撤销操作(revoke)是比较重的行为，只有当存在较多不会真正竞争的 synchronized 块儿时，才能体现出明显改善。 有人认为当你需要大量使用并发库时，就意味着并发高也就是不需要偏向锁。 建议最好是在实践中进行测试。

偏斜锁会延缓 JIT 预热的进程，所以很多性能测试中会显式地关闭偏斜锁， 命令如下:`-XX:-UseBiasedLocking`


类似 `fast_enter` 这种实现，解释器或者动态编译器，都是拷贝`synchronizer.cpp`这段基础逻辑，所以如果我们修改这部分逻辑，要保证一致性。微小的 问题都可能导致死锁或者正确性问题。

```c++
void ObjectSynchronizer::fast_enter(Handle obj, BasicLock* lock, bool attempt_rebias, TRAPS){
  if (UseBiasedLocking) {
    if (!SafepointSynchronize::is_at_safepoint()){
      BiasedLocking::Condition cond = BiasedLocking::revoke_and_rebias(obj, attempt_reb...
      if (cond == BiasedLocking::BIAS_REVOKED_AND_REBIASED) {
        return;
      }
    } else {
      assert(!attempt_rebias, "can not rebias toward VM thread");
      BiasedLocking::revoke_at_safepoint(obj);
    }
    assert(!obj->mark()->has_bias_pattern(), "biases should be revoked by now");
  }
  slow_enter(obj, lock, THREAD);
}
```

- biasedLocking定义了偏斜锁相关操作
  - revoke_and_rebias 是获取偏斜锁的入口方法 
  - revoke_at_safepoint 则定义了当检测到安全点时的处理逻辑。
- 如果获取偏斜锁失败，则进入 slow_enter。

>这个方法里面同样检查是否开启了偏斜锁，如果关闭了偏斜锁，是不会进入这个方法的，所以算是个额外的保障性检查吧。


太多细节就不展开了，明白它是通过 **CAS 设置 Mark Word** 就完全够用了，对象头中 Mark Word 的结构，可以参考下图: 被偏斜的对象，对象头前部有个`Thread pointor`和`Epoch`。


2. **轻量级锁**

slow_enter:
```c++
void ObjectSynchronizer::slow_enter(Handle obj, BasicLock* lock, TRAPS) {
  markOop mark = obj->mark();
  if(mark->is_neutral()) {
    // 将目前的 Mark Word 复制到 Displaced Header 上
    lock->set_displaced_header(mark);
    // 利用 CAS 设置对象的 Mark Word
    if (mark == obj()->cas_set_mark((markOop) lock, mark)) {
      TEVENT(slow_enter: release stacklock);
      return;
    }
    //检查存在竞争
  } else if (mark->has_locker() && THREAD->is_lock_owned((address)mark->locker())) {
    //clean
    ock->set_displaced_header(NULL);
    return;
  }

  // 重置 Displaced Header
  lock->set_displaced_header(markOopDesc::unused_mark());
  ObjectSynchronizer::inflate(THREAD, obj(), inflate_cause_monitor_enter)->enter(THREAD);
}

//更多细节： 
// deflate_idle_monitors是分析锁降级逻辑的入口，这部分行为还在进行持续改进，因为 其逻辑是在安全点内运行，处理不当可能拖长 JVM 停顿(STW，stop-the-world)的 时间。

//fast_exit 或者 slow_exit 是对应的锁释放逻辑。
```

- 设置 Displaced Header，然后利用 cas_set_mark 设置对象 Mark Word，如果成功就成功获取轻量级锁。
- 否则 Displaced Header，然后进入锁膨胀阶段，具体实现在 inflate 方法中。


### 其他的一些特别的锁类型

1. ReadWriteLock 是一个单独的接口，它通常是代表了一对儿锁，分别对应只读和写操作，标准类库中提供了再入版本的读写锁实现(ReentrantReadWriteLock)，对应的语义和 ReentrantLock 比较相似。
2. StampedLock 竟然也是个单独的类型，从类图结构可以看出它是不支持再入性的语义的， 也就是它不是以持有锁的线程为单位。


**为什么我们需要读写锁(ReadWriteLock)等其他锁呢?**

- 虽然ReentrantLock 和 synchronized 简单实用，但是行为上有一定局限性， 通俗点说就是“太霸道”，**要么不占，要么独占**。
- 有的时候不需要大量竞争的写操作，而是以并发读取为主。
- Java 并发包提供的读写锁等扩展了锁的能力，它所基于的原理是**多个读操作是不需要互斥的**，因为读操作并不会更改数据，所以不存在互相干扰。
  - 在运行过程中，如果**读锁试图锁定时，写锁是被某个线程持有，读锁将无法获得**，而只好等待对方操作结束，这样就可以**自动保证不会读取到有争议的数据**。


读写锁看起来比 synchronized 的粒度似乎细一些，但在实际应用中，其表现也并不尽如人意，主要还是因为**相对比较大的开销**。
> 啥开销？

所以，JDK 在后期引入了 StampedLock，在提供类似读写锁的同时，还支持优化读模式。 优化读基于假设，**大多数情况下读操作并不会和写操作冲突**，其逻辑是：
  - **先试着读**
  - 然后通过**validate方法**确认是否进入了写模式
  - 如果没有进入，就成功避免了开销;
  - 如果进入，则尝试获取读锁。

>请注意：writeLock 和 unLockWrite 一定要保证成对调用。


Java 并发包内的各种同步工具，不仅仅是各种 Lock，其他的如Semaphore、CountDownLatch，甚至是早期的FutureTask等，都是基 于一种AQS框架。


**自旋锁**

- 基于大部分的锁都是使用很短的时间的， 获取不到锁等一下就可能后去到的假设。
- 当获取锁失败的时候，不进入休眠等待（操作系统层面挂起，重新唤醒有个内核切换花销），而是继续“运动”做几个空循环，再进行尝试获取锁。 超过一定次数才正式挂起。 
- **好处**： 减少线程阻塞，内核用户态上下文切换开销。 适用于在锁竞争不激烈，占用锁非常短的情况。 属于乐观锁。
- **缺点**：消耗CPU，单cpu无效，因为基于cas的轮询会占用cpu,导致无法做线程切换。



## 一个线程两次调用start()方法会出现什么情况?

**典型回答**：
>Java 的线程是不允许启动两次的，第二次调用必然会抛出 IllegalThreadStateException， 这是一种运行时异常，多次调用 start 被认为是编程错误。


### 线程生命周期的不同状态

在 Java 5 以后，线程状态被明确定义在其公共内部枚举类型 `java.lang.Thread.State` 中，分别是：
- **新建（New）**: 表示线程被创建出来还没真正启动的状态
- **就绪（Runnable）**: 表示该线程已经在 JVM 中**执行**，当然由于执行需要计算资源，它**可能是正在运行**，也可能还在**等待系统分配给它CPU片段**，在就绪队列里面排队。
  >在其他一些分析中，会额外区分一种状态 RUNNING，但是从 Java API 的角度，并不能表示出来。
- **阻塞（Blocked）**: 表示线程在等 待 Monitor lock。
  >比如，线程试图通过 synchronized 去获取某个锁，但是其他线程已 经独占了，那么当前线程就会处于阻塞状态。
- **等待（Waiting）**: 表示正在等待其他线程采取某些操作。类似的如生产者消费者模式，发现条件为满足就让线程等待（wait），条件满足通过类似notify等动作，通知消费线程继续工作。
  >Thread.join() 也会令线程进入等待状态。
- **计时等待（Time_Wait）**: 与等待状态类似，调用的是存在超时条件的方法，比如 wait 或 join 等方法的指定超时版本
- **终止（Terminated）**： 不管是意外退出还是正常执行结束，线程已经完成使命，终止 运行


### 线程到底是什么以及 Java 底层实现方式

**是什么**？

- 线程是系统调度的最小单元，一个进程可以包含多个线 程。
- 作为任务的真正运作者，有自己的栈(Stack)、寄存器(Register)、本地存储 (Thread Local)等，
- 会和进程内其他线程共享文件描述符、虚拟地址空间等。


**Java底层实现方式**

线程还分为内核线程、用户线程，Java 的线程实现其实是与虚拟机相关的。

在 Java 1.2 之后，JDK 已经抛弃了所谓的Green Thread，也就是用户调度的线程，现在的模型是**一对一映射到操作系统内核线程**。

如果我们来看 Thread 的源码，你会发现其基本操作逻辑大都是以**JNI 形式调用的本地代码**。

- Java 语言得益于精细粒度的线程和相关的并发操作，其 构建高扩展性的大型应用的能力已经毋庸置疑。
- 其复杂性也提高了并发编程的门槛，go语言的协程大大提高构建并发应用的效率
- Java 也在Loom项目中，孕育新的类似轻量级用户线程(Fiber)等机制


如何创建线程？
- 直接扩展 Thread 类，然后实例化。
- 实现一个 Runnable，将代码逻放在 Runnable 中，然后构建 Thread 并启动(start)，等 待结束(join)。
- **好处：** 不会受Java 不支持类多继承的限制，重用代码实现，当我们需要重 复执行相应逻辑时优点明显。


### 线程状态的切换，以及和锁等并发工具类的互动

有哪些因素可能影响线程的状态呢?主要有:

- **线程自身的方法**，除了 start，还有多个 join 方法，等待线程结束;yield 是告诉调度 器，主动让出CPU;  被标记为过时的 resume、stop、suspend

- **基类 Object** 提供了一些基础的 wait/notify/notifyAll 方法。
  - 如果我们持有某个对象的 Monitor 锁，调用 wait 会让当前线程处于等待状态，直到其他线程 notify 或者 notifyAll。
  - 所以，本质上是提供了 Monitor 的获取和释放的能力，是基本的线程间通信 方式。

- **并发类库**中的工具，比如 CountDownLatch.await() 会让当前线程进入等待状态，直到 latch 被基数为 0，这可以看作是线程间通信的 Signal。


有了并发包，大多数情况下，我们已经不再 需要去调用 wait/notify 之类的方法了。


**守护线程(Daemon Thread)**，有的时候应用中需要一个长期驻留的服务程序， 但是不希望其影响应用退出，就可以将其设置为守护线程，如果 JVM 发现只有守护线程存在时，**将结束进程**。 注意，**必须在线程启动之前设置。**


### 线程编程时容易踩的坑与建议

**Spurious wakeup**

>尤其是在多核 CPU 的系统中，线程等待存在一种可能，就是 在没有任何线程广播或者发出信号的情况下，线程就被唤醒，如果处理不当就可能出现诡异 的并发问题.

所以我们在等待条件过程中，建议采用下面模式来书写。
```java
// 推荐
while ( isCondition()) { 
  waitForAConfition(...); 
}

// 不推荐，可能引入 bug
if ( isCondition()) {
  waitForAConfition(...);
}

```

自旋锁”(spin-wait, busy-waiting)，也可以认为其不算是一种锁，而是一种**针对短期等待的性能优化**技术。


**慎用ThreadLocal**


 Java 提供的一种**保存线程私有信息**的机制，因为其在**整个线程生命周期内有效**，所以可以方便地在一个线程关联的不同业务模块之间传递信息，比如事务 ID、Cookie 等上下文相关信息。

 - 数据存储于线程相关的 ThreadLocalMap，其内部条目是 **弱引用**.

```java
static class Entry extends ThreadLocalMap {
  static class Entry extends WeakReference<ThreadLocal<?>> {
    Object value;
    Entry(ThreadLocal<?> k, Object v) {
      super(k);
      value = v;
    }
  }
}
```

- 当 Key 为 null 时，该条目就变成“废弃条目”，相关“value”的回收，往往依赖于几个 关键点，即 set、remove、rehash。
- 通常弱引用都会和**引用队列配合清理机制使用**，但是 ThreadLocal 是个**例外**，它并没有这么做。
  - 这意味着，**废弃项目的回收依赖于显式地触发**，否则就要等待线程结束，进而回收相应 ThreadLocalMap!
  - 这就是很多 OOM 的来源，所以通常都会建议，**应用一定要自己负责remove**，并且**不要和线程池配合**，因为 worker 线程往往是不会退出的。
>theadlocal里面的值如果是线程池的线程里面设置的，当任务完成，线程归还线程池时， 这个threadlocal里面的值是不是不会被回收?


## 第18讲 | 什么情况下Java程序会产生死锁?如何定位、修复?

死锁是一种特定的程序状态，在实体之间，由于循环依赖导致彼此一直处于等待之中，没有任何个体可以继续前进。

- 死锁不仅仅是在线程之间会发生，**存在资源独占的进程之间同样也可能出现死锁**。
- 通常来说，我们大多是聚焦在多线程场景中的死锁，指两个或多个线程之间，**由于互相持有对方需要的锁，而永久处于阻塞的状态。**


**定位**

定位死锁最常见的方式就是**利用jstack等工具获取线程栈**，然后定位互相之间的依赖关 系，进而找到死锁。如果是比较明显的死锁，往往 jstack 等就能直接定位


### 写一个可能死锁的程序，考察下基本的线程编程

```java
public class DeadLockSample extends Thread {
  private String first;
  private String second;
  public DeadLockSample(String name, String first, String second) {
    super(name);
    this.first = first;
    this.second = second;
  }

  public void run() {
    synchronized(first){
      System.out.println(this.name + " obtained: " + first);
      try {
        Thread.sleep(1000L);
        synchronized(second) {
          System.out.println(this.name + " obtained: " + second);
        }
      } catch (InterruptedException e) {
        //Do nothing
      }
    }
  }

  public static void main(String[] args) throws InterruptedException {
    String lockA = "lockA";
    String lockB = "lockB";
    DeadLockSample t1 = new DeadLockSample("Thread1", lockA, lockB);
    DeadLockSample t2 = new DeadLockSample("Thread2", lockB, lockA);
    t1.start();
    t2.start();
    t1.join();
    t2.join();
  }
}
```

先调用 Thread1 的 start，但是可能Thread2 却先打印出来了呢? 
>这就是因为线程调度依赖于(操作系统)调度器，虽然你可以通过优先级之类进行影响，但 是具体情况是不确定的。


**定位流程**

1. jps、ps看进程ID
2. 调用 jstack 获取线程栈: `jstack your_pid`
3. 分析得到的输出: ，找到处于 BLOCKED 状态的线 程，按照试图获取(waiting)的锁 ID查找，很快就定位 问题.
4. 结合代码分析线程栈信息.


**区分线程状态 -> 查看等待目标 -> 对比 Monitor 等持有状态**

### 诊断死锁的工具，分布式环境下能否用API实现？

使用 Java 提供的标准管理 API，ThreadMXBean，其直接就提供了 findDeadlockedThreads() 方法用于定位。


### 如何在编程中尽量避免一些典型场景的死锁

**发生死锁的原因**：
- 互斥条件，类似 Java 中 Monitor 都是独占的，要么是我用，要么是你用。
- 互斥条件是长期持有的，在使用结束之前，自己不会释放，也不能被其他线程抢占。
- 循环依赖关系，两个或者多个个体之间出现了锁的链条环。


**避免**：
- 尽量避免使用多个锁，并且只有需要时才持有锁。
- 如果必须使用多个锁，尽量设计好锁的获取顺序，这个说起来简单，做起来可不容易，你可以参看著名的银行家算法。
- 使用带超时的方法，为程序带来更多可控性。
- 通过静态代码分析(如 FindBugs)去查找固定的模 式，进而定位可能的死锁或者竞争情况。
  - 类加载过程发生的死锁，尤其是在框架大量使用自定义类加载时，因为往往不是在应用本身的代码库中


有时候并不是阻塞导致的死锁，只是某个线程进入了死循环，导致其他线程一直等待，这种问题如何诊断呢?

>CPU使用量飙升，用`top -Hp`看使用率高的pid，转换为16进制，去jstack搜索线程状态。


## 第19讲 | Java 并发包提供了哪些并发工具类?
  