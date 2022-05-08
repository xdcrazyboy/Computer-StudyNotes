
# 基础语法

## 指针
- 指针用来保存变量的地址。
```go
var x =  5
var p *int = &x
fmt.Printf("x = %d",  *p) // x 可以用 *p 访问

type Result struct {
	Num, Ans int
}

type Cal int

func (cal *Cal) Square(num int) *Result {
	return &Result{
		Num: num,
		Ans: num * num,
	}
}

func main() {
	cal := new(Cal)
	result := cal.Square(6)
    // cal type: *main.Cal
    // result type: *main.Result

}
```
* * 运算符，也称为解引用运算符，用于访问地址中的值。
* ＆运算符，也称为地址运算符，用于返回变量的地址。
- &Result 等于 new(Result) 结果是对象的指针
- *Result 表示一个Result对象的指针


## 类型

1. **如何高效拼接字符串**
>使用 strings.Builder，最小化内存拷贝次数


2. 什么是 rune 类型
   >Unicode是ASCII的超集，包含世界上书写系统中存在的所有字符，并为每个代码分配一个标准编号（称为Unicode CodePoint），在 Go 语言中称之为 rune，是 int32 类型的别名。

Go 语言中，字符串的底层表示是 byte (8 bit) 序列，而非 rune (32 bit) 序列。例如下面的例子中 语 和 言 使用 UTF-8 编码后各占 3 个 byte，因此 len("Go语言") 等于 8.
```go
fmt.Println(len("Go语言")) // 8
//将字符串转换为 rune 序列
fmt.Println(len([]rune("Go语言"))) // 4
```


## defer
- 多个 defer 语句，遵从后进先出(Last In First Out，LIFO)的原则，最后声明的 defer 语句，最先得到执行。


- defer 在 return 语句之后执行，但在函数退出之前，defer 可以修改返回值（有名返回值）。
  - 如果是未命名返回值：执行 return 语句后，Go 会创建一个临时变量保存返回值，因此，defer 语句修改了局部变量 i，并没有修改返回值。
  - 有名返回值： 对于有名返回值的函数，执行 return 语句时，并不会再创建临时变量保存，因此，defer 语句修改会对返回值产生了影响。


## tag的作用
- tag 可以理解为 struct 字段的**注解**，可以用来**定义字段的一个或多个属性**。
- 框架/工具可以**通过反射获取到某个字段定义的属性**，采取相应的处理方式。
- tag 丰富了代码的语义，增强了灵活性。
```go
package main

import "fmt"
import "encoding/json"

type Stu struct {
	Name string `json:"stu_name"`
	ID   string `json:"stu_id"`
	Age  int    `json:"-"`
}

func main() {
	buf, _ := json.Marshal(Stu{"Tom", "t001", 18})
	fmt.Printf("%s\n", buf)
}
```
这个例子使用 tag 定义了结构体字段与 json 字段的转换关系，Name -> stu_name, ID -> stu_id，忽略 Age 字段。很方便地实现了 Go 结构体与不同规范的 json 文本之间的转换。**输出**：
```json
{"stu_name":"Tom","stu_id":"t001"}

```


## 字符串

**字符串打印时，%v 和 %+v 的区别，**
>%v 和 %+v 都可以用来打印 struct 的值，区别在于 %v 仅打印各个字段的值，%+v 还会打印各个字段的名称。


## 用常量表示枚举值
```go
type StuType int32

const (
	Type1 StuType = iota
	Type2
	Type3
	Type4
)

func main() {
	fmt.Println(Type1, Type2, Type3, Type4) // 0, 1, 2, 3
}
```

## 常量赋值
```go
func main() {
	const N = 100
    //无类型常量，赋值给其他变量时，如果字面量能够转换为对应类型的变量，则赋值成功
	var x int = N
    fmt.Println(x)
    //有类型的常量 ,赋值给其他变量时，需要类型匹配才能成功
	const M int32 = 100
	var y int = M
	fmt.Println(y)
}
```


## 空 struct{} 的用途
- 使用空结构体 struct{} 可以节省内存，一般作为占位符使用，表明这里并不需要一个值。
```go
fmt.Println(unsafe.Sizeof(struct{}{})) // 0
```

- 使用 map 表示集合时，**只关注 key**，value 可以使用 struct{} 作为占位符。
```go
  type Set map[string]struct{}
```

- 使用信道(channel)控制并发时，我们只是需要一个信号，但并不需要传递值，这个时候，也可以使用 struct{} 代替
```go
func main() {
	ch := make(chan struct{}, 1)
	go func() {
		<-ch
		// do something
	}()
	ch <- struct{}{}
	// ...
}
```

- 声明只包含方法的结构体。
```go
type Lamp struct{}

func (l Lamp) On() {
        println("On")

}
func (l Lamp) Off() {
        println("Off")
}
```

# 并发编程

## 什么是协程泄露(Goroutine Leak)？
协程泄露是指协程创建后，长时间得不到释放，并且还在不断地创建新的协程，最终导致内存耗尽，程序崩溃。常见的导致协程泄露的场景有以下几种：

- 缺少接收器，导致发送阻塞
- 缺少发送器，导致接收阻塞
- 两个或两个以上的协程在执行过程中，由于竞争资源或者由于彼此通信而造成阻塞，这种情况下，也会导致协程被阻塞，不能退出。

# 垃圾回收（GC）

Go 语言采用的是**标记清除**算法。并在此基础上使用了**三色标记法**和**写屏障技术**，提高了效率。


标记清除收集器是跟踪式垃圾收集器，其执行过程可以分成标记（Mark）和清除（Sweep）两个阶段：

* 标记阶段 — 从根对象出发查找并标记堆中所有存活的对象；
* 清除阶段 — 遍历堆中的全部对象，回收未被标记的垃圾对象并将回收的内存加入空闲链表。


标记清除算法的一大问题是**在标记期间，需要暂停程序**（Stop the world，STW），标记结束之后，用户程序才可以继续执行。

**为了能够异步执行，减少 STW 的时间，Go 语言采用了三色标记法**。三色标记算法将程序中的对象分成白色、黑色和灰色三类。

* 白色：不确定对象。
* 灰色：存活对象，子对象待处理。
* 黑色：存活对象。


1. 标记开始时，所有对象加入白色集合（这一步需 STW ）。
2. 首先将根对象标记为灰色，加入灰色集合。
3. 垃圾搜集器取出一个灰色对象，将其标记为黑色，并将其指向的对象标记为灰色，加入灰色集合。
4. 重复这个过程，直到灰色集合为空为止，标记阶段结束。
5. 那么白色对象即可需要清理的对象，而黑色对象均为根可达的对象，不能被清理。

三色标记法**因为多了一个白色的状态来存放不确定对象**，**所以后续的标记阶段可以并发地执行**。**为啥？**


当然并发执行的代价是可能会造成一些遗漏，因为那些早先被标记为黑色的对象可能目前已经是不可达的了。
所以三色标记法是一个 false negative（假阴性）的算法。

三色标记法并发执行仍存在一个问题，即在 GC 过程中，对象指针发生了改变。  
>A (黑) -> B (灰) -> C (白) -> D (白)
从B到C时，C删除了D的引用， 而B又引用了D，这个时候B已经标记为黑色了，不会再扫描其指向的对象。 D怎么办？

为了解决这个问题，Go**使用了内存屏障技术**，它是**在用户程序读取对象、创建新对象以及更新对象指针时执行的一段代码**，类似于一个钩子。
垃圾收集器使用了写屏障（Write Barrier）技术，**当对象新增或更新时，会将其着色为灰色**。（将B置为灰色还是将D置为灰色？或者一起？）


一次完整的 GC 分为四个阶段：

1）标记准备(Mark Setup，需 STW)，打开写屏障(Write Barrier)
2）使用三色标记法标记（Marking, 并发）
3）标记结束(Mark Termination，需 STW)，关闭写屏障。
4）清理(Sweeping, 并发)