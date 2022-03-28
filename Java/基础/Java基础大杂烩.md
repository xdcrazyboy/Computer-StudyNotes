

# 你对Java平台的理解

## 特点、特性

### 第一印象

>思维深入，且系统化
- Write Once，run anywhere，跨平台。 因为字节码-虚拟机
- 垃圾回收GC，自动内存分配和回收。
- JRE：Java运行环境，包含JVM和Java类库，以及一些模块。
- JDK： JRE的一个超集，还提供编译器、各类诊断工具。
- Java是大部分解释执行，但是JIT即时编译技术，热点代码提前编译成机器码-这属于编译执行。


**很多点**
- 语言特性：泛型、Lamabda等
- 基础类库：
  - 集合
  - IO/NIO、网络、utils
  - 并发、安全
- JVM
  - 类加载机制、常用JDK版本特点区别
  - 垃圾回收基本原理，常见垃圾收集器：SerialGC、Parallel GC、CMS、G1
  - 工具：编译器、运行时环境、安全工具、诊断、监控工具。
    - 辅助工具，如jlink、jar、jdeps
    - 编译器，javac、sjavac
    - 诊断工具：jmap、jstack、jconsole、jhsdb、jcmd
  - 解释和编译混合（mixed）： 
    - C1对应client模式（适用于启动速度敏感的应用，比如普通Java桌面应用）
    - C2对应server模式（适用于长时间运行的服务端应用）



## 多态&父子类
protected 需要从以下两个点来分析说明：

子类与基类在同一包中：被声明为 protected 的变量、方法和构造器能被同一个包中的任何其他类访问；

子类与基类不在同一包中：那么在子类中，子类实例可以访问其从基类继承而来的 protected 方法，而不能访问基类实例的protected方法。

protected 可以修饰数据成员，构造方法，方法成员，不能修饰类（内部类除外）。


