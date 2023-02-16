
# 一、并发理论基础

## 并发程序幕后的故事

CPU、内存、I/O 设备速度存在差异，如何解决？


为了合理利用 CPU 的高性能，平衡这三者的速度差异，计算机体系结构、操作系统、编译程序都做出了贡献，主要体现为：

- CPU 增加了**缓存**，以均衡与内存的速度差异；
- 操作系统增加了**进程、线程**，以分时复用 CPU，进而均衡 CPU 与 I/O 设备的速度差异；
- 编译程序**优化指令执行次序**，使得缓存能够得到更加合理地利用。


速度是快了很多，但并发程序很多诡异问题的根源也在这里。


### 源头之一：缓存导致的可见性问题

- 单核CPU缓存是一个，可见性没问题
- 多核CPU缓存不一样，但内存是同一个。 导致覆写


### 源头之二：线程切换带来的原子性问题


任务切换的时机大多数是在时间片结束的时候，我们现在基本都使用高级语言编程，高级语言里一条语句往往需要多条 CPU 指令完成，例如上面代码中的count += 1，至少需要三条 CPU 指令。

1. 指令 1：首先，需要把变量 count 从内存加载到 CPU 的寄存器；
2. 指令 2：之后，在寄存器中执行 +1 操作；
3. 指令 3：最后，将结果写入内存（缓存机制导致可能写入的是 CPU 缓存而不是内存）。


操作系统做任务切换，可以发生在任何一条 CPU 指令执行完，是 CPU 指令，而不是高级语言里的一条语句。


我们潜意识里面觉得 count+=1 这个操作是一个不可分割的整体，就像一个原子一样，线程的切换可以发生在 count+=1 之前，也可以发生在 count+=1 之后，但就是不会发生在中间。

我们把一个或者多个操作在 CPU 执行的过程中不被中断的特性称为原子性。


CPU 能保证的原子操作是 CPU 指令级别的，而不是高级语言的操作符，这是违背我们直觉的地方。因此，很多时候我们需要在高级语言层面保证操作的原子性。


### 源头之三：编译优化带来的有序性问题

有序性指的是程序按照代码的先后顺序执行。

编译器为了优化性能，有时候会改变程序中语句的先后顺序，例如程序中：
```java
a=6；
b=7；
```
编译器优化后可能变成
```java
b=7；a=6；
```
在这个例子中，编译器调整了语句的顺序，但是不影响程序的最终结果。

不过有时候编译器及解释器的优化可能导致意想不到的 Bug。


举例：**利用双重检查创建单例对象**
```java
public class SingletonDemo {
    private SingletonDemo instance = null;

    public static SingletonDemo getInstance(){
        if(instance == null){
            //锁的是类
            synchronized(SingletonDemo.class){
                if(instance == null){
                    return new SingletonDemo();
                }
            }
        }
        return instance;
    }
}
```

- 假设两个线程A/B都要获取实例，
  - 都发现`instance==null`，
  - 然后想加锁，这时候只有一个可以加锁成功。
  - 假设是A加锁成功，B就卡在加锁那一步。
  - 然后A就**new了一个实例**，释放锁。 
  - B加锁成功，判断`instance!=null`，就直接跳出去`return instance`了。
    - 这一步有歧义：线程在synchronized块中，发生线程切换，锁是不会释放的。 所以这里情况不会发生。
    - B也可以在第一个判空就发现instance != null，而此时A进行到给instance赋地址但未初始化，发生了时间片切换，但不会释放锁。 B无法获取锁，但发现不为null，直接返回未初始化的数据。

这看上去一切都很完美，无懈可击，但实际上这个 getInstance() 方法并不完美。问题出在哪里呢？出在**new 操作**上：

我们以为的 new 操作应该是：

1. 分配一块内存 M；
2. 在内存 M 上初始化 Singleton 对象；
3. 然后 M 的地址赋值给 instance 变量。

但是实际上优化后的执行路径却是这样的：

1. 分配一块内存 M；
2. 将 M 的地址赋值给 instance 变量；
3. 最后在内存 M 上初始化 Singleton 对象。

这样的顺序调整就可以出现：
- 假设线程 A 先执行 getInstance() 方法，当执行完指令 2 时恰好发生了线程切换，切换到了线程 B 上；
- 如果此时线程 B 也执行 getInstance() 方法，那么线程 B 在执行第一个判断时会发现 **instance != null** ，所以直接返回 instance
- 而此时的 **instance 是没有初始化过**的，如果我们这个时候访问 instance 的成员变量就可能触发**空指针异常**。


**静态内部类的单例模式方法**：

```java
public class Singleton{

    private static class SingletonHandler{
        private static singleton = new Singleton();
    }

    private MySingleton(){};

    public Singleton getInstance(){
        return SingletonHandler.singleton;
    }
}

```


