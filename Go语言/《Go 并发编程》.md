
## Go的MPG线程模型

### 协程
**进程**是一个具有**独立功能的程序**关于某个数据集合的一次**动态执行过程**，是操作系统进行**资源分配和调度的基本单位**，是应用程序运行的载体。

而**线程**则是程序执行过程中一个**单一的顺序控制流程**，是 **CPU 调度和分派的基本单位**。

线程是比进程更小的独立运行基本单位，一个进程中可以拥有一个或者以上的线程，这些线程共享进程所持有的资源，在 CPU 中被调度执行，共同完成进程的执行任务。 



**用户态和内核态，内核空间和用户空间**

根据**资源访问权限**的不同，操作系统会把内存空间分为内核空间和用户空间：
- 内核空间的代码能够直接访问计算机的底层资源，如 CPU 资源、I/O 资源等，为用户空间的代码提供计算机底层资源访问能力；
- 用户空间为上层应用程序的活动空间，无法直接访问计算机底层资源，需要借助“系统调用”“库函数”等方式调用内核空间提供的资源。


线程也可以分为内核线程和用户线程。

- **内核线程**由操作系统管理和调度，是内核调度实体，它能够直接操作计算机底层资源，可以充分利用 CPU 多核并行计算的优势，但是线程切换时需要 CPU 切换到内核态，存在一定的开销，可创建的线程数量也受到操作系统的限制。
- **用户线程**由用户空间的代码创建、管理和调度，无法被操作系统感知。用户线程的数据保存在用户空间中，切换时无须切换到内核态，切换开销小且高效，可创建的线程数量理论上只与内存大小相关。


**协程是一种用户线程**，属于**轻量级**线程。

优势：
- 协程的调度，完全由用户空间的代码控制；
- 协程拥有自己的寄存器上下文和栈，并存储在用户空间；
- 协程切换时无须切换到内核态访问内核空间，切换速度极快。

缺点：  
- 但这也给开发人员带来较大的技术挑战：开发人员需要在用户空间处理协程切换时上下文信息的保存和恢复、栈空间大小的管理等问题。

Go 是为数不多**在语言层次实现协程并发**的语言，它采用了一种特殊的两级线程模型：MPG 线程模型

### MPG线程模型

- M，即 machine，相当于内核线程在 Go 进程中的映射，它与内核线程一一对应，代表真正执行计算的资源。在 M 的生命周期内，它只会与一个内核线程关联。

- P，即 processor，代表 Go 代码片段执行所需的上下文环境。M 和 P 的结合能够为 G 提供有效的运行环境，它们之间的结合关系不是固定的。P 的最大数量决定了 Go 程序的并发规模，由 runtime.GOMAXPROCS 变量决定。

- G，即 goroutine，是一种轻量级的用户线程，是对代码片段的封装，拥有执行时的栈、状态和代码片段等信息。

# Context

## 什么是Context
Go 语言中用来设置截止日期、同步信号，传递请求相关值的结构体。上下文与 Goroutine 有比较密切的关系。


该接口定义了四个需要实现的方法，其中包括：
```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```
1. Deadline — 返回 context.Context 被取消的时间，也就是完成工作的截止日期；
2. Done — 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭，多次调用 Done 方法会返回同一个 Channel；
3. Err — 返回 context.Context 结束的原因，它只会在 Done 方法对应的 Channel 关闭时返回非空的值；
   1. 如果 context.Context 被取消，会返回 Canceled 错误；
   2. 如果 context.Context 超时，会返回 DeadlineExceeded 错误；
4. Value — 从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据；


context 包中提供的 context.Background、context.TODO、context.WithDeadline 和 context.WithValue 函数会返回实现该接口的私有结构体。

## 为什么需要Context
- 在 Goroutine 构成的树形结构中对信号进行同步以减少计算资源的浪费是 context.Context 的最大作用。



**背景知识**：waitgroup和channel

WaitGroup 和信道(channel)是常见的 2 种并发控制的方式。

如果并发启动了多个子协程，需要等待所有的子协程完成任务，WaitGroup 非常适合于这类场景，例如下面的例子：

```go
var wg sync.WaitGroup

func doTask(n int) {
	time.Sleep(time.Duration(n))
	fmt.Printf("Task %d Done\n", n)
	wg.Done()
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doTask(i + 1)
	}
	wg.Wait()
	fmt.Println("All task Done")
}

```
- wg.Wait() 会等待所有的子协程任务全部完成，所有子协程结束后，才会执行 wg.Wait() 后面的代码.
- WaitGroup并**不能主动通知子协程退出**。
  
  
  假如开启了一个定时轮询的子协程，**有没有什么办法，通知该子协程退出呢**？
>select + chan 的机制
```go

var stop chan bool

func reTask(name string) {
	for {
		select {
		case <-stop:
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	stop = make(chan bool)
	go reTask("worker1")
	time.Sleep(3 * time.Second)
	stop <- true
	time.Sleep(3 * time.Second)
}
```
>子协程使用 for 循环定时轮询，如果 stop 信道有值，则退出，否则继续轮询。


更复杂的场景如何做并发控制呢？
比如子协程中开启了新的子协程，或者需要同时控制多个子协程。这种场景下，select+chan的方式就显得力不从心了。


Go 语言提供了 Context 标准库可以解决这类场景的问题，Context 的作用和它的名字很像，上下文，即子协程的下上文。Context 有两个主要的功能：

* 通知子协程退出（正常退出，超时退出等）；
* 传递必要的参数。

## contex.WithCancel
`context.WithCancel()`创建**可取消的Context对象**，即可以主动通知子协程退出。

### 控制单个协程

使用Context改写上面的例子，效果与select+chan相同。
```go

func reTask(ctx context.Context, name string) {
	for {
		select {
		// 在子协程中，使用 select 调用 <-ctx.Done() 判断是否需要退出。
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// context.Backgroud() 创建根 Context，通常在 main 函数、初始化和测试代码中创建，作为顶层 Context。
	// context.WithCancel(parent) 创建可取消的子 Context，同时返回函数 cancel
	ctx, cancel := context.WithCancel(context.Background())
	go reTask(ctx, "worker1")
	time.Sleep(3 * time.Second)
	// 主协程中，调用 cancel() 函数通知子协程退出。
	cancel()
	time.Sleep(3 * time.Second)
}
```

### 控制多个协程

调用 `cancel()` 函数后该 `Context` 控制的所有子协程都会退出。

```go

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go reTask(ctx, "worker1")
	go reTask(ctx, "worker2")

	time.Sleep(3 * time.Second)
	// 为每个子协程传递相同的上下文 ctx 即可，调用 cancel() 函数后该 Context 控制的所有子协程都会退出。
	cancel()
	time.Sleep(3 * time.Second)
}
```


### context.WithValue
如果需要**往子协程中传递参数**，可以使用 `context.WithValue()`。
```go
type Options struct {
	Interval time.Duration
}

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			op := ctx.Value("options").(*Options)
			time.Sleep(op.Interval * time.Second)
		}

	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	vCtx := context.WithValue(ctx, "options", &Options{1})

	go reqTask(vCtx, "worker1")
	go reqTask(vCtx, "worker2")

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}
```

- `context.WithValue()` 创建了一个基于 ctx 的子 Context，并携带了值 options。
- 在子协程中，使用 `ctx.Value("options"`) 获取到传递的值，读取/修改该值。


### context.WithTimeout
如果需要控制子协程的执行时间，可以使用 `context.WithTimeout` 创建具有**超时通知机制**的 `Context` 对象。
```go
ctx, cancel := context.WithCancel(context.Background())
```
- WithTimeout()的使用与 WithCancel() 类似，多了一个参数，用于设置超时时间。
- 因为超时时间设置为 2s，但是 main 函数中，3s 后才会调用 cancel()，因此，在调用 cancel() 函数前，子协程因为超时已经退出了。



### context.WithDeadline
- 超时退出可以控制子协程的最长执行时间，那 `context.WithDeadline()` 则可以控制子协程的**最迟退出时间**。

```go

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")

			time.Sleep(1 * time.Second)
		}

	}
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))

	go reqTask(ctx, "worker1")
	go reqTask(ctx, "worker2")

	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	cancel()
	time.Sleep(3 * time.Second)
}
```

- WithDeadline 用于设置截止时间。在这个例子中，将截止时间设置为1s后，cancel() 函数在 3s 后调用，因此子协程将在调用 cancel() 函数前结束。
- 在子协程中，可以通过 ctx.Err() 获取到子协程退出的错误原因。