
## 看圣经遇到的一个奇怪问题



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