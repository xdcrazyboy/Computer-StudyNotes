[TOC]



## 经常疑惑的点

### GOPATH

- 下载的第三方包源代码文件放在$GOPATH/src目录下， 
- 产生的二进制可执行文件放在 $GOPATH/bin目录下，
- 生成的中间缓存文件会被保存在 $GOPATH/pkg 下


- 配置变量GOPATH会跟安装默认的冲突不？ 
  >是配置到~/.bashrc里。然后source一下.


- `go install xxx`文件后，提示安装到了最初go安装默认的位置，但是已经指定GOPATH，为啥还会安装到那？  提示：`open /usr/local/go/bin/mathapp: permission denied`。 难道还要指定GOROOT？ 我希望安装到的是自定义的GOPATH下的src/bin目录下
>**解决办法**：
>(a) 如果你没有设置你的GOBIN env变量,你可以在GOROOT/bin中获得Go编译器二进制文件,而你的二进制文件将在GOPATH/bin中.(我个人喜欢这种二进制分离.)
>(b) 如果你将GOBIN设置为任何东西,那么Go二进制文件和你的二进制文件都将转到GOBIN.



### 包 package
- go 里面一个目录为一个package, 一个package级别的func, type, 变量, 常量, 这个package下的所有文件里的代码都可以随意访问, 不需要首字母大写
- **同目录**下的两个文件如hello.go和hello2.go中的package 定义的名字要是同一个，不同的话，是会报错的 ==> 所以main方法要单独放一个文件


### init函数和main函数

**init函数**： 用于包(package)的初始化


  1. init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
  2. 每个包可以拥有多个init函数
  3. 包的每个源文件也可以拥有多个init函数
  4. 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)
  5. 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
  6. init函数**不能被其他函数调用**，而是在main函数执行之前，自动被调用。 手动显示调用init会收到编译错误：`undefined: init`


**main函数**：Go语言程序的默认入口函数(主函数)


 1. 可执行程序的 main 包必须定义 main 函数，否则 Go 编译器会报错。
 2. 在启动了多个 Goroutine的 Go 应用中，main.main 函数将在 Go 应用的主 Goroutine 中执行。如果住goroutine结束、返回，则其他子goroutine都会结束。


**异同**：
- 相同点：
    - 两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用。
- 不同点：
    - init可以应用于任意包中，且可以重复定义多个。
    - main函数只能用于main包中，且只能定义一个。


**执行顺序**
![go包初始化顺序](../../Computer-StudyNotes/img/《Go%20疑难杂症》/Go包初始化顺序.jpg)

- 如果 main 包依赖的包中定义了 init 函数，或者是 main 包自身定义了 init 函数，那么 Go 程序在这个包初始化的时候，就会自动调用它的 init 函数，因此这些 **init 函数的执行就都会发生在 main 函数之前**。

* 对同一个go文件的多个init()调用顺序是**从上到下的**。

* 对同一个package中不同文件是按文件名字符串比较“从小到大”顺序调用各文件中的init()函数。

* 对于不同的package，如果不相互依赖的话，按照main包中"先import的后调用"的顺序调用其包中的init()，如果package存在依赖，则先调用最早被依赖的package中的init()，**最后调用main函数**。

* 如果init函数中使用了println()或者print()你会发现在执行过程中这两个不会按照你想象中的顺序执行。这两个函数官方只推荐在测试环境中使用，对于正式环境不要使用



# 看go语言圣经的疑惑
##  defer修饰的方法到底什么时候执行？

### 看了下csdn有小伙伴遇到一样的疑问

https://blog.csdn.net/lj779323436/article/details/109696343

但最近看《go程序设计语言》一书关于defer这一块的介绍时，书中写了一个demo，用defer实现了进入函数的打印以及出函数的打印和函数花费的时间，现把代码贴出来：
```go
//gopl.io/ch5/trace

func bigSlowOperation() {
    defer trace("bigSlowOperation")() // don't forget the extra parentheses
    // ...lots of work…
    time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
    start := time.Now()
    log.Printf("enter %s", msg)
    return func() { 
        log.Printf("exit %s (%s)", msg,time.Since(start)) 
    }
}

func main(){
	bigSlowOperation()
}

```

仔细一看，发现defer后面是一个函数（方法），而且这个函数的返回值也是函数，按照常规的思维方式，defer修饰的代码应该在函数（方法）结束的时候才执行，也就是time.Sleep(10 * time.Second)这句代码执行之后再才轮到defer修饰的trace方法执行，这样trace方法中的log.Printf("enter %s", msg)岂不是在程序执行10s后才执行吗？这如何实现杠进入bigSlowOperation就打印呢？下面我们跑一下
```
//打印
2015/11/18 09:53:26 enter bigSlowOperation
2015/11/18 09:53:36 exit bigSlowOperation (10.000589217s)

```

**如果defer不带返回值呢？**
不带return的defer后面的函数确实是在原函数结束的时候才执行


所以该题主得初步出结论：**defer后面的函数里面只要有return语句，则只有这个return的语句才会在原函数结束时执行；**


我又试了下，想看看返回不是函数，而是其他值，比如int是不是还会如此， 看下面评论。

### 我的评论

看这块的时候我也遇到这个坑了，不过你的结论没有准确覆盖。
你有没有发现书上的例子return的是一个函数？你发现了。 但是你仔细看defer声明背后带了个圆括号？：
defer trace("bigSlowOperation")()

如果我把这个函数改一下，改成返回int，那带圆括号是会报错的，可能是返回func()和圆括号是语法匹配？这个坑我还没找到解释。 但是， 我可以把圆括号去掉改成下面这个：
```go
func bigSlowOperation() {
	defer trace("xxx") // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
	log.Println("ready end...")
}
 
func trace(msg string) int {
	v := 2
	if msg == "xxx" {
		v++
		log.Printf("v:%d", v)
		return v
	}
	log.Printf("enter %s, v:%d", msg, v)
	return v
}
 
// !-main
 
func main() {
	bigSlowOperation()
}
```

结果发现：
```
2022/05/14 01:43:42 ready end...
2022/05/14 01:43:42 v:3
```
初步结论：defer后面函数虽然带了return语句，return前的代码也不会运行，而是一起等到原函数结束后才运行。 目前只有返回func的defer函数（注意不要漏掉圆括号）的场景，才可以实现return前的代码经过defer时就运行，return后的func是原函数结束后执行。



而书中两个提到两句话可能是解开问题的关键（虽然我还是没弄懂，刚学习一周多）：

1. 你只需要在调用普通函数或方法前加上关键字defer，就完成了defer所需要的语法。当执行到该条语句时，函数和参数表达式得到计算，但直到包含该defer语句的函数执行完毕时，defer后的函数才会被执行。
- 执行到该条语句--应该就是经过defer时，
- 函数和参数得到计算---这是核心点，这里的函数是指啥？ 参数又是指啥，我感觉函数应该就是最开始的例子，参数的例子可能我还得试一下。
- defer后的函数才会被执行---这一句的函数跟上面的应该不是一个东西？（我严重怀疑是翻译的问题，决定是看看英文）

2. 需要注意一点：不要忘记defer语句后的圆括号，否则本该在进入时执行的操作会在退出时执行，而本该在退出时执行的，永远不会被执行。



### 于是我就去看了英文版本
下面2个网址可以在线阅读：

https://www.doc88.com/p-6189947807970.html?r=1

https://www.shuzhiduo.com/A/Vx5Mvo47dN/

Syntactically, a defer statement is an ordinary function or method call prefixed by the keyword defer. 
**The function and argument expressions are evaluated** when **the statement is executed**, but **the actual call is deferred** until **the function** that contains the defer statement has **finished**, whether normally, by executing a return statement or falling off the end, or abnormally, by panicking.
 Any number of calls may be deferred; they are executed in the reverse of the order in which they were deferred.


The defer statement can also be used to pair "on entry" and "on exit" actions when debugging a complex function. 

The bigSlowOperation function below calls trace immediately, which does the "on entry* action then returns a fuction value that, when called, does the corresponding "on exit" action.

By deferring a call to the returned function in this way, we can instrument the entry point and all exit points of a function in a single statement and even pass values, like the start time, between the two actions. 

But don't forget the final parentheses in the defer statement, or the "on entry" action will happen on exit and the on-exit action won't happen at all!



## 圣经翻译疑惑：给命名类型指定方法有啥限制？
`Distance() float64`




# 疑惑

## 如何知道类型实现了哪些接口？

有个疑问，go里面一个类型实现了接口所有的方法，才算该接口类型，但并没有语法显式 说明这个类型实现了哪个接口(例如java中有implements), 这样看别人代码的时候，碰到一 个类型，无法知道这个类型是不是实现了一个接口，除非类型和接口写在一个文件，然后 还要自己一个一个方法去对比。有比较快的方法可以知道当前类型实现了哪些接口么?
