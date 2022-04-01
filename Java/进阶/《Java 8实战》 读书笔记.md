# Java 8 的新功能

● 语言需要不断改进以跟进硬件的更新或满足程序员的期待

# 概述

主要：

● Stream API
● 向方法传递代码的技巧
● 接口中的默认方法


好处：

●  Java 8对于程序员的主要好处在于它提供了更多的编程工具和概念，能以更快，更重要的是能以更为简洁、更易于维护的方式解决新的或现有的编程问题。 
  ○ 流处理 
    ■ Unix命令行允许多个程序通过管道（|）连接在一起，比如： cat file1 file2 | tr "[A-Z]" "[a-z]" | sort | tail -3
  ○ 用行为参数化把代码传递给方法：如何区别于匿名函数
  ○ 并行与共享的可变数据

这两个要点（没有共享的可变数据，将方法和函数即代码传递给其他方法的能力）是我们平常所说的函数式编程范式的基石。

● 不能有共享的可变数据”的要求意味着，一个方法是可以通过它将参数值转换为结果的方式完全描述的； 
  ○ 换句话说，它的行为就像一个数学函数，没有可见的副作。 执行时在元素之间没有互动。

# 流

## 匿名函数：

 filterApples(inventory, (Apple a)-> a.getWeight() < 80 ||
                                              "brown".equals(a.getColor()) );

你甚至都不需要为只用一次的方法写定义；代码更干净、更清晰，因为你用不着去找自己到底传递了什么代码。但要是Lambda的长度多于几行（它的行为也不是一目了然）的话，那你还是应该用方法引用来指向一个有描述性名称的方法，而不是使用匿名的Lambda。你应该以代码的清晰度为准绳。

要是没有多核CPU，可能他们真的就到此为止了，为了更好地利用并行，Java的设计师没有这么做。Java 8中有一整套新的类集合API——Stream，它有一套函数式程序员熟悉的、类似于filter的操作，比如map、reduce，还有我们接下来要讨论的在Collections和Streams之间做转换的方法。

## 流和Collection：

● Collection主要是为了存储和访问数据，而Stream则主要用于描述对数据的计算。
● Stream允许并提倡并行处理一个Stream中的元素： 筛选一个Collection（将上一节的filterApples应用在一个List上）的最快方法常常是将其转换为Stream，进行并行处理，然后再转换回List
```java

//顺序处理
List<Apple> heavyApples = inventory.stream().filter((Apple a)-> a.getWeight() > 150).collect(toList());
//并行处理
List<Apple> heavyApples = inventory.parallelStream().filter((Apple a)-> a.getWeight() > 150).collect(toList());

```
## 默认方法

Java 8中加入默认方法主要是为了支持库设计师，让他们能够写出更容易改进的接口。

为什么要加默认方法？

比如上面的collection.stream().xxxx， 以前的collection或者说List这些都没有stream()，是java8加的，但是没有这个方法，java7运行这个代码就会编译错误，如果是我们自己的接口，增加了一个方法，而这个接口有多个实现，这些实现都需要去实现这个新的方法，不然就无法编译通过。

你如何改变已发布的接口，而不破坏已有的实现呢？
这个问题再Java 8 解决了——接口可以包含实现类没有提供实现的方法签名了！， 那谁来实现它？ 缺失的方法主体（实现）随着接口一起提供。（这就是默认实现），也就是说接口里面不再是没有实现，只有方法签名，而是可以增加一个默认方法（带着默认实现）。

● Java 8在接口声明中使用新的default关键字来表示这一点
● Java 8 List接口中有如下的默认方法实现，实现直接用List调用sort方法。

一个类可以实现多个接口，不是吗？那么，如果在好几个接口里有多个默认实现，是否意味着Java中有了某种形式的多重继承？是的，在某种程度上是这样。我们在第9章中会谈到，Java 8用一些限制来避免出现类似于C++中臭名昭著的菱形继承问题.

## 来自函数式编程的其他好思想

● 将方法和Lambda作为一等值，以及在没有可变共享状态时，函数或方法可以有效、安全地并行执行。
● 通过使用更多的描述性数据类型来避免null。
● （结构）模式匹配： 使用多态和方法重载来替代if-then-else； 
  ○ 你可以把模式匹配看作switch的扩展形式，可以同时将一个数据类型分解成元素
  ○ 函数式语言倾向于允许switch用在更多的数据类型上，包括允许模式匹配（在Scala代码中是通过match操作实现的）

## 通过行为参数化传递代码

● 行为参数化就是可以帮助你处理频繁变更的需求的一种软件开发模式

### 举个例子：选苹果

● 从一堆苹果里筛选出 
  ○ 红苹果、绿苹果...用颜色做为参数搞定。
  ○ 然后，要重的轻的，又复制一下上面的代码，把颜色改为重量
  ○ ....违反了DRY（Don't RepeatYourself，不要重复自己）法则，且你复制了大部分的代码来实现遍历库存，并对每个苹果应用筛选条件
```java

public static List<Apple> filterApplesByColor(List<Apple> inventory, String color) {
    List<Apple> result=new ArrayList<Apple>();
    for (Apple apple: inventory){
        if ( apple.getColor().equals(color) ) {
            result.add(apple);
        }
    }
    return result;
}
List<Apple> greenApples=filterApplesByColor(inventory, "green");
List<Apple> greenApples=filterApplesByColor(inventory, "red");

public static List<Apple> filterApplesByWeight(List<Apple> inventory, int weight) {
    List<Apple> result=new ArrayList<Apple>();
    for (Apple apple: inventory){
        if ( apple.getWeight() > weight ){
            result.add(apple);
        }
    }
    return result;
}

//3 、通过标记，将不同的条件区分，比如枚举类。。。hen
List<Apple> greenApples=filterApples(inventory, "green", 0, true);
List<Apple> heavyApples=filterApples(inventory, "", 150, false);

```
- 首先，客户端代码看上去糟透了。true和false是什么意思？
- 此外，这个解决方案还是不能很好地应对变化的需求

试试高度抽象，用策略模式：

● 定义一族算法，把它们封装起来（称为“策略”），然后在运行时选择一个算法
● 该怎么利用ApplePredicate的不同实现呢？你需要filterApples方法接受ApplePredicate对象，对Apple做条件测试。  
这就是行为参数化：让方法接受多种行为（或战略）作为参数，并在内部使用，来完成不同的行为。
```java

public interface ApplePredicate{
    boolean test(Apple apple);
}

public class AppleGreenColrPredict implements ApplePredicate{
    public boolean test(Apple apple){
        return "green".equals(apple.getColor());
    }
}

//4. 根据抽象条件筛选
public static List<Apple> filterApples(List<Apple> inventory, ApplePredicate p){
    List<Apple> result = new ArrayList<>();
    for (Apple apple: inventory){
        if ( p.text()){
            result.add(apple);
        }
    }
    return result;
}

List<Apple> greenApples = filterApples(inventory, new AppleGreenColrPredict());

```
上面的中方式有几个关键点：
1．传递代码/行为。

现在你可以创建不同的ApplePredicate对象，并将它们传递给filterApples方法。把filterApples方法的行为参数化
但是，
由于该filterApples方法只能接受对象，所以你必须把代码包裹在ApplePredicate对象里。你的做法就类似于在内联“传递代码”，因为你是通过一个实现了test方法的对象来传递布尔表达式的. 其实可以通过Lambda直接将表达式"green".equals(apple.getColor)传递给filterApples方法，不用定义个ApplePredicate类。

2．多种行为，一个参数。

这样可以把行为抽象出来，让你的代码适应需求的变化，但这个过程很啰嗦，因为你需要声明很多只要实例化一次的类。让我们来看看可以怎样改进。
对付啰嗦：

● 匿名类： 同时声明和实例化一个类

//5. 使用匿名类
```java
List<Apple> greenApples = filterApples(inventory, new AppleGreenColrPredict(){
    public boolean test(Apple apple){
        return "green".equals(apple.getColor());
    }
});

```
● 但匿名类还是不够好。 
  ○ 第一，它往往很笨重，因为它占用了很多空间
  ○ 第二，很多程序员觉得它用起来很让人费解

在只需要传递一段简单的代码时（例如表示选择标准的boolean表达式），你还是要创建一个对象，明确地实现一个方法来定义一个新的行为（例如Predicate中的test方法

- 第六次尝试： 使用Lambda表达式
```java
List<Apple> greenApples = filterApples(inventory, (Apple apple) ->  "green".equals(apple.getColor()));


```
- 第七次尝试：将List类型抽象化， filterApples方法还只适用于Apple。你还可以将List类型抽象化，从而超越你眼前要处理的问题
```java

public interface Predicate{
    boolean test(T t);
}

//4. 根据抽象条件筛选
public static <T> List<T> filter(List<T> list, Predicate<T> p){
    List<T> result = new ArrayList<>();
    for (T e: list){
        if ( e.test()){
            result.add(apple);
        }
    }
    return result;
}

List<Apple> greenApples = filter(inventory, (Apple apple) ->  "green".equals(apple.getColor()));

List<Integer> evenNumbers = filter(numbers, (Integer i) -> i % 2 == 0);

```
在灵活性和简洁性之间找到了最佳平衡

### 来些实例

接口只有一个实现，可以直接用匿名函数 -> Lambda

1. Comparator排序
```java

inventory.sort(new Comparator<Apple>() {
    public int compare(Apple a1, Apple a2){
        return a1.getWeight().compareTo(a2.getWeight());
    }
});

nventory.sort((Apple a1, Apple a2)-> a1.getWeight().compareTo(a2.getWeight()));

```
2. 用Runnable执行代码块
```java

Thread t=new Thread(()-> System.out.println("Hello world"));
```