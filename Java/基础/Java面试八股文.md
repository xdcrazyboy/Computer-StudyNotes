
# Java面试指南

# 常见面试题

1. Java的第一印象、特点、区别于其他语言的？ 见[《Java基础大杂烩》](./Java基础大杂烩.md)
2. 异常：Error、Exception、Throwable？见[《异常处理》](./异常处理.md)
3. 谈谈 final、finally、 finalize 有什么不同?
    - final：是可以用来修饰类、方法、变量。明确语义和意图，不可修改，也保证安全。减少同步开销，省去一些防御性拷贝的必要。
      - 类： 不可继承扩展
        - 在java.lang 包下面的很多类，相当 一部分都被声明成为 final class?在第三方类库的一些基础类中同样如此，这可以有效避免 API 使用者更改基础功能
      - 变量： 不可修改。 
        - final 字段对性能的影 响，大部分情况下，并没有考虑的必要。
        - final 不是 immutable!：
      - 方法： 不可重写（override）
    - finally：是Java 保证重点代码一定要被执行的一种机制。
      - 经常用在try-catch-finally中类似JDBC关闭连接、保证unlock锁等动作，不过关闭资源推荐使用try-with-resources语句。
    - finalize： 是基础类java.lang.Object的一个方法，设计目的是保证对象在被垃圾收集前完成特定资源的回收。（不在推荐使用）
      - 不推荐原因：你无法保证 finalize 什么时候执行，执行的是否符合预期。使用不当会 影响性能，导致程序死锁、挂起等。
    >面试官还可以考察你对性能、并发、对象生命周期或垃圾 收集基本过程等方面的理解.


