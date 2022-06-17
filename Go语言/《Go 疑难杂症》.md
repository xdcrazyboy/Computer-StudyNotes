[TOC]

# Go Module 那些道道

## 如何升级修改版本

- **查看版本**： `o list -m -versions github.com/sirupsen/logrus`
- **修改版本**：
  - 方法1： 以在项目的 module 根目录下，执行带有版本号的 go get 命令：`go get github.com/sirupsen/logrus@v1.7.0`
  - 方法2： 修改go.mod文件，然后tidy一下：
    - `go mod edit -require=github.com/sirupsen/logrus@v1.7.0`
    - `go mod tidy`


### 如何添加主版本号大于1的依赖（在代码中import）
在 Go Module 构建模式下，当依赖的主版本号为 0 或 1 的时候，我们在 Go 源码中导入依赖包，**不需要在包的导入路径上增加版本号**，也就是：

```go
import github.com/user/repo/v0 等价于 import github.com/user/repo
import github.com/user/repo/v1 等价于 import github.com/user/repo
```

- 如果新旧版本的包使用相同的导入路径，那么新包与旧包是兼容的。 反过来说，如果不兼容，那剧需要采用不同的导入路径。


如果引入的主版本大于1的依赖（比如v2.0.0），那么就不能直接使用`github.com/user/repo`,因为这是默认0/1，这个与2是不兼容的。 需要向下面这样导入：
```go
import github.com/user/repo/v2/xxx
```
 - 也就是在声明它的导入路径的基础上，加上版本号信息。
 - 然后要从新下载最新的：`go get github.com/go-redis/redis/v7`

### 升级依赖版本到一个不兼容版本

跟上面类似，修改版本号，然后重新下载。
```go

import (
  _ "github.com/go-redis/redis/v8"
  "github.com/google/uuid"
  "github.com/sirupsen/logrus"
)

//
$go get github.com/go-redis/redis/v8
```

### 移除一个依赖

- 在业务代码中删除依赖后，直接build不会删除不用的依赖，因为如果源码满足成功构建的条件，go build 命令是不会“多管闲事”地清理 go.mod 中多余的依赖项的。
- 运行下 `go mod tidy`就行， go mod tidy 会自动分析源码依赖，而且将不再使用的依赖从 go.mod 和 go.sum 中移除。


### 特殊情况：使用vendor

**什么情况下还需要用vendor？**


- 在一些不方便访问外部网络，并且对 Go 应用构建性能敏感的环境，比如在一些内部的持续集成或持续交付环境（CI/CD）中，使用 vendor 机制可以实现与 Go Module 等价的构建。


**怎么用mod模式下用vendor？**

- Go Module 构建模式下，我们再也无需手动维护 vendor 目录下的依赖包了，Go 提供了可以快速建立和更新 vendor 的命令：
  - `go mod vendor` 项目根目录，创建vendor目录。
    - go mod vendor 命令在 vendor 目录下，创建了一份这个项目的依赖包的副本
    - 并且通过 vendor/modules.txt 记录了 vendor 下的 module 以及版本。
- 如果我们要基于 vendor 构建，而不是基于本地缓存的 Go Module 构建，我们需要在 go build 后面加上 `-mod=vendor` 参数。
- 在 Go 1.14 及以后版本中，如果 Go 项目的顶层目录下存在 vendor 目录，那么 go build **默认也会优先基于 vendor 构建**，除非你给 go build 传入 -mod=mod 的参数。


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
