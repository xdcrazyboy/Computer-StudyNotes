
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

