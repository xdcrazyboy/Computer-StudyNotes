

# 基础部分

## 第一章 总览Go语言能做啥
通过十几个程序介绍了用Go语言如何实现类似读写文件、文本格式化、创建图像、网络客户端和服务器通讯等日常工作。

### 入门


- 缺少了必要的包或者导入了不需要的包，程序都无法编译通过。Go语言编译过程没有警告信息，争议特性之一
- 不需要添加分号，除非一行有多填语句。编译器会主动把特定符号后的换行符转换为分号，因此换行符添加的位置会影响Go代码的正确解析。
  - 行末出现：比如行末是标识符、整数、浮点数、虚数、字符或字符串文字、关键字break、continue、fallthrough或return中的一个、运算符和分隔符++、--、)、]或}中的一个），会自动插入分号分割符。 
  - 以+结尾的话不会被插入分号分隔符，所以 a + b， 可以在 + 号后面换行，不能在加号前面，因为 a后面换行会被插入分号，那编译就报错了。


#### 命令行参数
- **输入源：**自于程序外部：文件、网络连接、其它程序的输出、敲键盘的用户、命令行参数或其它类似输入源。
- **包导入顺序**并不重要；gofmt工具格式化时**按照字母顺序对包名排序**。
- i--给i减1。它们是语句，而不像C系的其它语言那样是表达式。所以j = i++非法，而且++和--都**只能放在变量名后面**，因此--i也非法。


**命令行参数**


- 程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量。
- os.Args变量是一个字符串（string）的切片（slice）
- os.Args的第一个元素：os.Args[0]，是命令本身的名字；其它的元素则是程序启动时传给它的参数


- for循环，唯一的循环，多种形式。
```go
for initialization; condition; post {
    // zero or more statements
}
```
  - initialization语句是可选的，在循环开始前执行。initalization如果存在，必须是一条简单语句（simple statement），即，短变量声明、自增语句、赋值语句或函数调用。
  - condition是一个布尔表达式（boolean expression），其值在每次循环迭代开始时计算。如果为true则执行循环体语句。
  - post语句在循环体执行结束后执行，**之后再次**对condition求值。
- for循环的这三个部分每个都可以省略，如果省略initialization和post，分号也可以省略（相当于 while）：
```go
// a traditional "while" loop
for condition {
    // ...
}
```
- 省略掉condition，变成for{} ，无限循环，可以用 break，return终止。
- for循环的另一种形式，**在某种数据类型的区间（range）上遍历**，如字符串或切片。

```go
// Echo2 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
```
  - 每次循环迭代，range产生一对值；索引以及在该索引处的元素值。
  - 这个例子不需要索引，但range的语法要求，要处理元素，必须处理索引
    - 一种思路是把索引赋值给一个临时变量（如temp）然后忽略它的值，但Go语言不允许使用无用的局部变量（local variables）
    - go语言提供了一种解决方案：`空标识符（blank identifier）`，即`_`（也就是下划线）。
      - 空标识符可用于在任何语法需要变量名但程序逻辑不需要的时候
```go
// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("methodName:" + os.Args[0])

	for i, arg := range os.Args[1:] {
		fmt.Println(strconv.FormatInt(int64(i), 10) + ":" + arg)
	}

}
```

#### 查找重复的行
**dup 版本一****:

打印标准输入中多次出现的行，以重复次数开头。该程序将引入if语句，map数据类型以及bufio包

```go
// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    //Scanner类型是该包最有用的特性之一，它读取输入并将其拆成行或单词；通常是处理行形式的输入最简单的方法。
    input := bufio.NewScanner(os.Stdin)
    //每次调用input.Scan()，即读入下一行，并移除行末的换行符；
    //读取的内容可以调用input.Text()得到。
    //Scan函数在读到一行时返回true，不再有输入时返回false。
    for input.Scan() {
        //下面语句等价于： line := input.Text(); counts[line] = counts[line] + 1
        counts[input.Text()]++
        //map中不含某个键时不用担心，首次读到新行时，等号右边的表达式counts[line]的值将被计算为其类型的零值，对于int即0
    }
    // NOTE: ignoring potential errors from input.Err()
    for line, n := range counts {
        if n > 1 {
            //%d表示以十进制形式打印一个整型操作数
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
```

**Printf的格式转换**：
```
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数。
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true或false
%c          字符（rune） (Unicode码点)
%s          字符串
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
```
>后缀f指format，ln指line
- 以字母`f`结尾的格式化函数: 如`log.Printf`和`fmt.Errorf`，都采用fmt.Printf的格式化准则。
- 以`ln`结尾的格式化函数: 则遵循Println的方式，以跟`%v`差不多的方式格式化参数，并在最后添加一个换行符
- 带#
- 带+，输出字段名



**dup版本二**

读取标准输入或是使用os.Open打开各个具名文件，并操作它们。

```go
// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            //第一个值是被打开的文件(*os.File）
            f, err := os.Open(arg)
            //如果err等于内置值nil（译注：相当于其它语言里的NULL），那么文件被成功打开
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
    // NOTE: ignoring potential errors from input.Err()
}

```
**说明：**
- map是一个由make函数创建的数据结构的引用。
- map作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本），被调用函数对map底层数据结构的任何修改，调用者函数都可以通过持有的map引用看到。
- 在我们的例子中，countLines函数向counts插入的值，也会被main函数看到。
>（译注：类似于C++里的引用传递，实际上指针是另一个指针了，但内部存的值指向同一块内存）


**dup版本三**
一口气把全部输入数据读到内存中，一次分割为多行，然后处理它们。


引入了ReadFile函数（来自于io/ioutil包），其读取指定文件的全部内容，strings.Split函数把字符串分割成子串的切片。

```go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
        //ReadFile函数返回一个字节切片（byte slice），必须把它转换为string，才能用strings.Split分割。
        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
```


### GIF动画
生成的图形名字叫利萨如图形（Lissajous figures）。

下面代码引入新的结构，包括const声明，struct结构体类型，复合声明。
```go
// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
    "time"
)
//引入包带过多单词时，通常我们只需要用最后那个单词表示这个包就可以
//[]color.Color{...} 复合声明,slice切片
var palette = []color.Color{color.White, color.Black}

const (
    //整个包都可用，常量声明的值必须是一个数字值、字符串或者一个固定的boolean值。
    whiteIndex = 0 // first color in palette
    blackIndex = 1 // next color in palette
)

func main() {
    // The sequence of images is deterministic unless we seed
    // the pseudo-random number generator using the current time.
    // Thanks to Randall McPherson for pointing out the omission.
    
    //调用lissajous函数，用它来向标准输出流打印信息
    rand.Seed(time.Now().UTC().UnixNano())
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    //把常量声明定义在函数体内部，那么这种常量就只能在函数体内用
    const (
        cycles  = 5     // number of complete x oscillator revolutions
        res     = 0.001 // angular resolution
        size    = 100   // image canvas covers [-size..+size]
        nframes = 64    // number of animation frames
        delay   = 8     // delay between frames in 10ms units
    )

    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    //复合声明，生成的是struct结构体，其内部变量LoopCount字段会被设置为nframes；而其它的字段会被设置为各自类型默认的零值
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase difference
    //外层循环会循环64次，每一次都会生成一个单独的动画帧。
    for i := 0; i < nframes; i++ {
        //它生成了一个包含两种颜色的201*201大小的图片，白色和黑色。
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            //内层循环设置两个偏振值。x轴偏振使用sin函数。
            x := math.Sin(t)
            //y轴偏振也是正弦波，但其相对x轴的偏振是一个0-3的随机值，初始偏振值是一个零值，随着动画的每一帧逐渐增加。
            y := math.Sin(t*freq + phase)
            //循环会一直跑到x轴完成五次完整的循环。每一步它都会调用SetColorIndex来为(x,y)点来染黑色。
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
                blackIndex)
        }
        phase += 0.1
        //将结果append到anim中的帧列表末尾，并设置一个默认的80ms的延迟值
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    //循环结束后所有的延迟值被编码进了GIF图片中，并将结果写入到输出流
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

```


### 并发访问url
先来个非并发版本的：
类似curl的最基本功能，fetch访问指定的url，并将响应的body打印出来
```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		status := resp.Status
		fmt.Println(status)
		//函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，使用这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout
		out, err := io.Copy(os.Stdout, resp.Body)
		// b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", out)
	}
}
```


**并发**：
fetchall的特别之处在于它会同时去获取所有的URL，所以这个程序的总执行时间不会超过执行时间最长的那一个任务。

```go
// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        //goroutine是一种函数的并发执行方式，而channel是用来在goroutine之间进行参数传递。
        //main函数本身也运行在一个goroutine中
        go fetch(url, ch) // start a goroutine
    }
    for range os.Args[1:] {
        fmt.Println(<-ch) // receive from channel ch
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

/**
* - 当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，
*     直到另一个goroutine从这个channel里接收或者写入值，这样两个goroutine才会继续执行channel操作之后的逻辑。
* - 每一个fetch函数在执行时都会往channel里发送一个值（ch <- expression），主函数负责接收这些值（<-ch）。
* - 这个程序中我们用main函数来接收所有fetch函数传回的字符串，可以避免在goroutine异步执行还没有完成时main函数提前退出。
*/
func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }
    //ioutil.Discard输出流可以把这个变量看作一个垃圾桶，可以向里面写一些不需要的数据
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close() // don't leak resources
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
```


### Web服务
```go
// Server2 is a minimal "echo" and counter server.
package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

var mu sync.Mutex
var count int

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/showCount", showCounter)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    count++
    mu.Unlock()

    fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
    for k, v := range r.Header {
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    }
    fmt.Fprintf(w, "Host = %q\n", r.Host)
    fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    for k, v := range r.Form {
        fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
    }
}

// counter echoes the number of calls so far.
func showCounter(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    fmt.Fprintf(w, "Count %d\n", count)
    mu.Unlock()
}
```

### 本章要点

**控制流**：if、for、switch


**命名类型**


**指针**
- 指针是一种直接存储了变量的内存地址的数据类型
  - 指针是可见的内存地址
  - &操作符可以返回一个变量的内存地址
  - *操作符可以获取指针指向的变量内容
    - 要注意区分表示指针类型的*，比如： 类型 *T 是指向 T 类型值的指针，其零值为 nil。


- 不像C语言那样不受约束，也不想其他语言那样沦为单纯的“引用”，折中。
  - 没有指针运算，不能对指针进行加减。


**方法和接口**
- 方法：是和命名类型关联的一类函数
  - **go特别**：方法可以被关联到任意一种命名类型
- 接口：一种抽象类型，这种类型可以让我们以同样的方式来处理不同的固有类型，
  - 不用关心它们的具体实现，而只需要关注它们提供的方法


**包（packages）**
Go语言提供了一些很好用的package，并且这些package是可以扩展的。

- 可以在 https://golang.org/pkg 和 https://godoc.org 中找到标准库和社区写的package。
- godoc这个工具可以让你直接在本地命令行阅读标准库的文档。


**注释**
- 在源文件的开头写的注释是这个源文件的文档
- 每一个函数之前写一个说明函数行为的注释也是一个好习惯。

---

## 第二章 基本实体
元素结构、变量、新类型定义、包和文件、以及作用域等概念

### 命名
- 大小写敏感
- 25个关键字不能用于命名， 30多个预定义名字可以用，但要避免过渡使用语义混乱。
- 可见性
  - 在函数内部定义，那么它就只在函数内部有效
- 在函数外部定义，那么将在当前包的所有文件中都可以访问
- 函数外部定义大写包级名字（包级函数名本身也是包级名字，大写函数也就是公共函数），可以被外部的包访问。
  - 例如fmt包的Printf函数就是导出的，可以在fmt包外部访问。
  - 包本身的名字一般总是用小写字母
- 长度无限制，但是断点好，特别是局部变量。
- 推荐驼峰，而不是下划线


### 声明
Go语言主要有四种类型的声明语句：var、const、type和func，分别对应**变量、常量、类型和函数实体**对象的声明。

- 包一级的各种类型的声明语句的顺序无关紧要（译注：函数内部的名字则必须先声明之后才能使用）

**变量声明**：
- `var 变量名字 类型 = 表达式`
- 在Go语言中不存在未初始化的变量
  - 数字0，空字符串，false
  - 接口或引用类型（包括slice、指针、map、chan和函数）变量对应的零值是nil
  - 数组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值。
  - go语言程序员应该让一些聚合类型的零值也具有意义，这样可以保证不管任何类型的变量总是有一个合理有效的零值状态。
- **初始化**
  - 在包级别声明的变量会在main入口函数执行前完成初始化。
  - 局部变量将在声明语句被执行到的时候完成初始化。
- **短声明**
  - 在函数内部，有一种称为简短变量声明语句的形式可用于声明和初始化局部变量。
  - 它以“名字 := 表达式”形式声明变量，变量的类型根据表达式来自动推导。
  - 例如：`t := 0.0`
  - var形式的声明语句往往是用于需要显式指定变量类型的地方，或者因为变量稍后会被重新赋值而初始值无关紧要的地方。例如：`var err error`，后面会进行重新赋值。
  - 请记住“:=”是一个变量声明语句，而“=”是一个变量赋值操作。
    - 简短变量声明左边的变量可能并不是全部都是刚刚声明的。
    - 如果有一些**已经在相同的词法域声明过了**，那么简短变量声明语句对这些已经声明过的变量就只有赋值行为了。 
    - 简短变量声明语句中**必须至少**要声明一个**新**的变量.
    - 如果变量是在外部词法域声明的，那么简短变量声明语句将会在当前词法域重新声明一**个新的变量**
```go
f, err := os.Open(infile)
// 解决的方法是第二个简短变量声明语句改用普通的多重赋值语句。就是用=号
f, err := os.Create(outfile) // compile error: no new variables
```


**指针**
如果用“var x int”声明语句声明一个x变量：
- `&x`表达式（取x变量的内存地址），将产生一个指向该整数变量的指针，指针对应的**数据类型是*int**，“指向int类型的指针”。
- 如果指针名字为p，可以说p指针指向变量x，或者p指针保存了x变量的内存地址。
- `*p`:**对应p指针指向的变量的值**
  - 因为*p对应一个变量，所以该表达式也可以出现在赋值语句的左边，表示更新指针所指向的变量的值。
```go
x := 1
p := &x         // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2          // equivalent to x = 2
fmt.Println(x)  // "2"
```
- 任何类型的指针的零值都是nil。
  - 如果p指向某个有效变量，那么p != nil测试为真。
  - 指针之间也是可以进行相等测试的，只有当它们**指向同一个变量**或**全部是nil时才相等**。


- 每次我们对一个变量取地址，或者复制指针，我们都是为原变量创建了新的别名。
- 指针特别有价值的地方在于我们可以不用名字而访问一个变量
- 但是这是一把双刃剑：要找到一个变量的所有访问者并不容易，我们必须知道变量全部的别名（译注：这是Go语言的垃圾回收器所做的工作）


指针是实现标准库中flag包的关键技术，它使用命令行参数来设置对应变量的值，而这些对应命令行标志参数的变量可能会零散分布在整个程序中。
- -n用于忽略行尾的换行符，-s sep用于指定分隔字符（默认是空格）
- 程序中的sep和n变量分别是指向对应命令行标志参数变量的指针，因此必须用\*sep和*n形式的指针语法间接引用它们
```go
// Echo4 prints its command-line arguments.
package main

import (
    "flag"
    "fmt"
    "strings"
)
//调用flag.Bool函数会创建一个新的对应布尔型标志参数的变量。
//第一个是命令行标志参数的名字“n”，然后是该标志参数的默认值（这里是false），最后是该标志参数对应的描述信息。
//如果用户在命令行输入了一个无效的标志参数，或者输入-h或-help参数，那么将打印所有标志参数的名字、默认值和描述信息。
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
//程序中的sep和n变量分别是指向对应命令行标志参数变量的指针，因此必须用*sep和*n形式的指针语法间接引用它们
func main() {
    //用于更新每个标志参数对应变量的值（之前是默认值）
    flag.Parse()
    //对于非标志参数的普通命令行参数可以通过调用flag.Args()函数来访问，返回值对应一个字符串类型的slice
    fmt.Print(strings.Join(flag.Args(), *sep))
    if !*n {
        fmt.Println()
    }
}
```


**new函数**
另一个创建变量的方法是调用内建的new函数。
- 表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值
- 然后返回变量地址，**返回的指针类型为*T**

```go
p := new(int)   // p, *int 类型, 指向匿名的 int 变量
```


- 每次调用new函数都是返回一个新的变量的地址，因此下面两个地址是不同的：
```go
p := new(int)
q := new(int)
fmt.Println(p == q) // "false"
```
>如果两个类型都是空的，也就是说类型的大小是0，例如struct{}和[0]int，有可能有相同的地址（依赖具体的语言实现）


由于new只是一个预定义的函数，它并不是一个关键字，因此我们可以将new名字重新定义为别的类型。例如下面的例子：
```go
//由于new被定义为int类型的变量名，因此在delta函数内部是无法使用内置的new函数的。
func delta(old, new int) int { return new - old }
```


**变量的生命周期**
- 对于在包一级声明的变量来说，它们的生命周期和整个程序的运行周期是一致的。
- 局部变量的生命周期则是动态的：每次从创建一个新变量的声明语句开始，直到该变量不再被引用为止，然后变量的存储空间**可能被回收**.


**Go语言的自动垃圾收集器是如何知道一个变量是何时可以被回收的呢？**
- 基本的实现思路是，从每个包级的变量和每个当前运行函数的**每一个局部变量开始**，通过指针或引用的访问路径遍历，是否可以找到该变量。
- 因为一个变量的有效周期只取决于是否可达，因此一个循环迭代内部的局部变量的生命周期可能超出其局部作用域。局部变量可能在函数返回之后依然存在。


- 编译器会**自动选择**在栈上还是在堆上分配局部变量的存储空间。 不是根据是用var还是new声明觉得的，而是
  - 下面f函数的x在函数退出后依然可以通过包一级的global变量找到，虽然它是在函数内部定义的；用Go语言的术语说，这个**x局部变量从函数f中逃逸**了
  - 当g函数返回时，变量*y将是不可达的，也就是说可以马上被回收的。因此，*y并没有从函数g中逃逸，编译器可以选择在栈上分配*y的存储空间（译注：也可以选择在堆上分配，然后由Go语言的GC回收这个变量的内存空间），虽然这里用的是new方式。
```go
var global *int

func f() {
    var x int
    x = 1
    global = &x
}

func g() {
    y := new(int)
    *y = 1
}
```
- 虽然不需要显式地分配和释放内存，但是要编写高效的程序你依然需要了解变量的生命周期
>例如，如果将指向短生命周期对象的指针保存到具有长生命周期的对象中，特别是保存到全局变量时，会阻止对短生命周期对象的垃圾回收（从而可能影响程序的性能）。


### 赋值
- 自增和自减是语句，而不是表达式，因此x = i++之类的表达式是错误的
- 元组赋值：允许同时更新多个变量的值，在赋值之前，赋值语句右边的所有表达式将会先进行求值，然后再统一更新左边对应变量的值。
  - 对于处理有些同时出现在元组赋值语句左右两边的变量很有帮助，例如我们可以这样**交换两个变量**的值：
```go
x, y = y, x

a[i], a[j] = a[j], a[i]

//计算最大公约数
func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}


//计算斐波那契额数列的第N个数
func fib(n int) int {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x+y
    }
    return x
}
```
- 如果map查找（§4.3）、类型断言（§7.10）或通道接收（§8.4.2）出现在赋值语句的右边，它们都可能会产生两个结果，有一个额外的布尔结果表示操作是否成功
```go
v, ok = m[key]             // map lookup
v, ok = x.(T)              // type assertion
v, ok = <-ch               // channel receive

//也有只产生一个结果的情形
v = m[key]                // map查找，失败时返回零值
v = x.(T)                 // type断言，失败时panic异常
v = <-ch                  // 管道接收，失败时返回零值（阻塞不算是失败）

_, ok = m[key]            // map返回2个值
_, ok = mm[""], false     // map返回1个值
_ = mm[""]                // map返回1个值
```


**可赋值性**
- 不管是隐式还是显式地赋值，在赋值语句左边的变量和右边最终的求到的值必须有相同的数据类型。更直白地说，只有右边的值对于左边的变量是可赋值的，赋值语句才是允许的。
- 大部分的类型必须完全匹配，nil可以赋值给任何指针或引用类型的变量。
- 常量则有更灵活的赋值规则，因为这样可以避免不必要的显式的类型转换。
- 对于两个值是否可以用==或!=进行相等比较的能力也和可赋值能力有关系：对于任何类型的值的相等比较，第二个值必须是对第一个值类型对应的变量是可赋值的



### 类型
- 变量或表达式的类型定义了对应存储值的属性特征。例如
  - 数值在内存的存储大小（或者是元素的bit个数），
  - 它们在内部是如何表达的，
  - 是否支持一些操作符，
  - 以及它们自己关联的方法集等。
- 一些变量有有着相同的内部结构，但是却表示完全不同的概念，通过创建不一样的类型名称，分割开。
  - 声明类型： `type 类型名字 底层类型`
  - 类型声明语句一般出现在包一级，因此如果新创建的类型名字的首字符大写，则在包外部也可以使用。


**将不同温度单位定义为不同的类型**
- 它们不可以被相互比较或混在一个表达式运算。
- 需要一个类似Celsius(t)或Fahrenheit(t)形式的显式转型操作才能将float64转为对应的类型
>Celsius(t)和Fahrenheit(t)是类型转换操作，它们并不是函数调用。
- 好处
  - 可以避免一些像无意中使用不同单位的温度混合计算导致的错误
```go
/ Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
    AbsoluteZeroC Celsius = -273.15 // 绝对零度
    FreezingC     Celsius = 0       // 结冰点温度
    BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }


//如果两个值有着不同的类型，则不能直接进行比较
var c Celsius
var f Fahrenheit
fmt.Println(c == 0)          // "true"
fmt.Println(f >= 0)          // "true"
fmt.Println(c == f)          // compile error: type mismatch
//Celsius(f)是类型转换操作，它并不会改变值，仅仅是改变值的类型而已。测试为真的原因是因为c和g都是零值。
fmt.Println(c == Celsius(f)) // "true"!
```


**类型转换操作**：
- 类型转换不会改变值本身，但是会使它们的语义发生变化
- 对于每一个类型T，都有一个对应的类型转换操作T(x)，用于将x转为T类型（译注：如果T是指针类型，可能会需要用小括弧包装T，比如(*int)(0)）
- 只有当两个类型的底层基础类型相同时，才允许这种转型操作，或者是两者都是指向相同底层结构的指针类型，这些转换只改变类型而不会影响值本身。
- 数值类型之间的转型也是允许的，并且在字符串和一些特定类型的slice之间也是可以转换的，
  - 可能改变值的表现：将一个浮点数转为整数将丢弃小数部分，将一个字符串转为[]byte类型的slice将拷贝一个字符串数据的副本
  - 在**任何情况下**，**运行时不会发生转换失败的错误**（译注: 错误只会发生在编译阶段）。


**命名类型**：
- 可以提供书写方便
- 还可以为该类型的值定义新的行为。（这些行为表示为一组关联到该类型的函数集合，我们称为**类型的方法集**。）


**String()方法，类似toString**：当使用fmt包的打印方法时，将会优先使用该类型对应的String方法返回的结果打印。
```go
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

c := FToC(212.0)
fmt.Println(c.String()) // "100°C"
fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
```


### 包和文件
- 每个包还有一个包名，包名一般是短小的名字（并不要求包名是唯一的），包名在包的声明处指定。按照惯例，一个包的名字和包的导入路径的最后一个字段相同，例如gopl.io/ch2/tempconv包的名字一般是tempconv。


**包的初始化**
- 首先是解决包级变量的依赖顺序，然后按照包级变量声明出现的顺序依次初始化
- 如果包中含有多个.go源文件，它们将按照发给编译器的顺序进行初始化，Go语言的构建工具首先会将.go文件根据文件名排序，然后依次调用编译器编译。
- 对于在包级别声明的变量，如果有初始化表达式则用表达式初始化.复杂的一般用一个特殊的**init初始化函数**来简化初始化工作。每个文件都可以包含多个init初始化函数
```go
func init() { /* ... */ }
```
  - 每个文件中的init初始化函数，在程序开始执行时按照它们声明的顺序被自动调用
  - 每个包在解决依赖的前提下，以导入声明的顺序初始化，每个包只会被初始化一次。
  - 初始化工作是自下而上进行的，main包最后被初始化。这种方式可以确保在main函数执行之前，所有依赖的包都已经完成初始化工作了。


### 作用域
- 对于导入的包，例如tempconv导入的fmt包，则是对应源文件级的作用域，因此只能在当前的文件中访问导入的fmt包，当前包的其它源文件无法访问在当前源文件导入的包。
- 当编译器遇到一个名字引用时，它会对其定义进行查找，查找过程从最内层的词法域向全局的作用域进行。
- 如果查找失败，则报告“未声明的名字”这样的错误。如果该名字在内部和外部的块分别声明过，则**内部块的声明首先被找到**。
- 下面**第二个if语句嵌套在第一个内部**，因此第一个if语句条件初始化词法域声明的变量在第二个if中也可以访问。switch语句的每个分支也有类似的词法域规则：条件部分为一个隐式词法域，然后是每个分支的词法域。
```go
if x := f(); x == 0 {
    fmt.Println(x)
} else if y := g(x); x == y {
    fmt.Println(x, y)
} else {
    fmt.Println(x, y)
}
fmt.Println(x, y) // compile error: x and y are not visible here
```
- 下面代码：变量f的作用域只在if语句内，因此后面的语句将无法引入它，这将导致编译错误。你可能会收到一个局部变量f没有声明的错误提示，具体错误信息依赖编译器的实现。
```go
if f, err := os.Open(fname); err != nil { // compile error: unused: f
    return err
}
f.ReadByte() // compile error: undefined f
f.Close()    // compile error: undefined f
```
- 通常需要在if之前声明变量，这样可以确保后面的语句依然可以访问变量：
```go
f, err := os.Open(fname)
if err != nil {
    return err
}
f.ReadByte()
f.Close()
```

**有个例子：要特别注意短变量声明语句的作用域范围。**考虑下面的程序，它的目的是获取当前的工作目录然后保存到一个包级的变量中。
```go
var cwd string

func init() {
    cwd, err := os.Getwd() // NOTE: wrong!
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working directory = %s", cwd)
}
```
  - 虽然cwd在外部已经声明过，但是:=语句还是将cwd和err**重新声明为新的局部变量**。因为内部声明的cwd将屏蔽外部的声明，因此上面的代码并**不会正确更新包级声明的cwd变量**。
  - 全局的cwd变量依然是没有被正确初始化的，而且看似正常的日志输出更是让这个BUG更加隐晦。
- 有许多方式可以避免出现类似潜在的问题。最直接的方法是**通过单独声明err变量**，来**避免使用:=的简短声明方式**：
```go
var cwd string

func init() {
    var err error
    cwd, err = os.Getwd()
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
}
```


## 第三章 基础数据类型：数字、布尔值、字符串和常量

### 整型

- 有两种一般**对应特定CPU平台机器字大小**的有符号和无符号整数**int和uint**；其中int是应用最广泛的数值类型。这两种类型都有同样的大小，32或64bit，但是我们不能对此做任何的假设；因为不同的编译器即使在相同的硬件平台上可能产生不同的大小。
- Unicode字符**rune类型是和int32等价**的类型，通常用于**表示一个Unicode码点**。这两个名称可以互换使用。
- 同样**byte也是uint8类型的等价类型**，byte类型一般用于**强调数值是一个原始的数据**而不是一个小的整数。
- 还有一种无符号的整数类型uintptr，没有指定具体的bit大小但是足以容纳指针。uintptr类型只有在底层编程时才需要，特别是Go语言和C语言函数库或操作系统接口相交互的地方。
- 不管它们的具体大小，int、uint和uintptr是**不同类型的兄弟类型**。其中int和int32也是不同的类型，即使int的大小也是32bit，在需要将int当作int32类型的地方需要一个显式的类型转换操作，反之亦然。


**取模%**
- 在Go语言中，%取模运算符的符号和被取模数的符号总是一致的，因此-5%3和-5%-3结果都是-2。
- 除法运算符/的行为则依赖于操作数是否全为整数，比如5.0/4.0的结果是1.25，但是5/4的结果是1，因为整数除法会向着0方向截断余数


**溢出**：
- 一个算术运算的结果，不管是有符号或者是无符号的，如果需要更多的bit位才能正确表示的话，就说明计算结果是溢出了。超出的高位的bit位部分将被丢弃。如果原始的数值是有符号类型，而且最左边的bit位是1的话，那么最终结果可能是负的，例如int8的例子：

```go
var u uint8 = 255
fmt.Println(u, u+1, u*u) // "255 0 1"

var i int8 = 127
fmt.Println(i, i+1, i*i) // "127 -128 1"
```


- 位操作运算符`^`作为二元运算符时是按位异或（XOR），当用作一元运算符时表示按位取反；也就是说，它返回一个每个bit位都取反的数。
- 位操作运算符`&^`用于按位置零（AND NOT）：如果对应y中bit位为1的话，表达式z = x &^ y结果z的对应的bit位为0，否则z对应的bit位等于x相应的bit位的值。


下面的代码演示了如何**使用位操作解释uint8类型值的8个独立的bit位**。
它使用了Printf函数的%b参数打印二进制格式的数字；其中%08b中08表示打印至少8个字符宽度，**不足的前缀部分用0填充**。
```go
var x uint8 = 1<<1 | 1<<5  //00000010 | 00100000
var y uint8 = 1<<1 | 1<<2

fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
```


- **无符号数**往往只有在**位运算**或其它特殊的运算场景才会使用，就像bit集合、分析二进制文件格式或者是哈希和加密操作等。它们通常**并不用于仅仅是表达非负数量的场合**。


**类型转换**：对于每种类型T，**如果转换允许的话，类型转换操作T(x)将x转换为T类型**。


**fmt.Printf格式**： 
- 当使用fmt包打印一个数值时，我们可以用`%d`、`%o`或`%x`参数控制输出的进制格式，就像下面的例子：
```go
o := 0666
fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
x := int64(0xdeadbeef)
fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
// Output:
// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
```
请注意fmt的两个使用技巧。通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数，
- 但是%之后的\[1]副词告诉Printf函数再次使用**第一个操作数**。
- 第二，%后的#副词告诉Printf在用`%o`、`%x`或`%X`输出时生成0、0x或0X前缀。


字符使用%c参数打印，或者是用%q参数打印带单引号的字符：
```go
ascii := 'a'
unicode := '国'
newline := '\n'
fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
```


用Printf函数的%g参数打印浮点数，将采用更紧凑的表示形式打印，并提供足够的精度，但是对应表格的数据，使用%e（带指数）或%f的形式打印可能更合适.

所有的这三个打印形式都可以指定打印的宽度和控制打印精度。
```go
for x := 0; x < 8; x++ {
    //打印e的幂，打印精度是小数点后三个小数精度和8个字符宽度
    fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
}
```


### 浮点数
Go语言提供了两种精度的浮点数，float32和float64。


通常应该**优先使用float64类型**，因为float32类型的累计计算误差很容易扩散，并且float32能精确表示的正整数并不是很大。
>（译注：因为float32的有效bit位只有23个，其它的bit位用于指数和符号；当整数大于23bit能表达的范围时，float32的表示将出现误差）：
```go
var f float32 = 16777216 // 1 << 24
fmt.Println(f == f+1)    // "true"!
```


### 复数
- Go语言提供了两种精度的复数类型：complex64和complex128，分别对应float32和float64两种浮点数精度。
- 内置的complex函数用于构建复数，内建的**real和imag函数**分别返回复数的实部和虚部：
```go
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y)                 // "(-5+10i)"
fmt.Println(real(x*y))           // "-5"
fmt.Println(imag(x*y))           // "10"
```
- 如果一个浮点数面值或一个十进制整数面值后面跟着一个i，例如3.141592i或2i，它将构成一个复数的虚部，复数的实部是0：`fmt.Println(1i * 1i) // "(-1+0i)", i^2 = -1`
- 复数也可以用==和!=进行相等比较。只有两个复数的实部和虚部都相等的时候它们才是相等的. 风险跟浮点数比较类似。


### 布尔类型
- 布尔值可以和&&（AND）和||（OR）操作符结合，并且**有短路行为**：如果运算符左边值已经可以确定整个布尔表达式的值，那么运算符右边的值将不再被求值。
- **&&的优先级比||高**
>（助记：&&对应逻辑乘法，||对应逻辑加法，乘法比加法优先级要高），下面形式的布尔表达式是不需要加小括弧的：
```go
if 'a' <= c && c <= 'z' ||
    'A' <= c && c <= 'Z' ||
    '0' <= c && c <= '9' {
    // ...ASCII letter or digit...
}
```
- 布尔值并**不会隐式转换为数字值0或1**，反之亦然。必须使用一个显式的if语句辅助转换


### 字符串
* 一个字符串是一个不可改变的字节序列。
* 字符串可以包含任意的数据，包括byte值0，但是通常是用来包含人类可读的文本。
* 内置的`len`函数可以返回一个字符串中的**字节数目**（**不是rune字符数目**），索引操作s[i]返回第i个字节的字节值，i必须满足0 ≤ i< len(s)条件约束。
  * 第i个字节并不一定是字符串的第i个字符，因为对于非ASCII字符的UTF8编码会要两个或多个字节。
* 子字符串操作s[i:j]基于原始的s字符串的第i个字节开始到第j个字节（并不包含j本身）生成一个新字符串。
- 字符串**可以用==和<进行比较**；比较通过**逐个字节比较**完成的，因此比较的结果是**字符串自然编码的顺序**。



**字符串不可变？**
- 我们也可以给一个字符串变量分配一个新字符串值。下面代码：并不会导致原始的字符串值被改变，但是变量s将因为+=语句持有一个新的字符串值，但是t依然是包含原先的字符串值。
```go
s := "left foot"
t := s
s += ", right foot"

fmt.Println(s) // "left foot, right foot"
fmt.Println(t) // "left foot"
```
- s[0] = 'L' // compile error: cannot assign to s[0]



#### **原生字符串**
一个原生的字符串面值形式是`...`，使用反引号代替双引号。在原生的字符串面值中，没有转义操作；全部的内容都是字面的意思，包含退格和换行，因此一个程序中的原生字符串面值可能跨越多行。
- 常用于编写正则表达式。


#### **编码：Unicode和UTF-8**
- Unicode码点对应Go语言中的rune整数类型（译注：rune是int32等价类型）。
- UTF8是一个将Unicode码点编码为字节序列的变长编码。UTF8编码是由Go语言之父Ken Thompson和Rob Pike共同发明的，现在已经是Unicode的标准。
- UTF8编码使用1到4个字节来表示每个Unicode码点，ASCII部分字符只使用1个字节，常用字符部分使用2或3个字节表示。
  - 每个符号编码后第一个字节的高端bit位用于表示编码总共有多少个字节。
    - 如果第一个字节的高端bit为0，则表示对应7bit的ASCII字符，ASCII字符每个字符依然是一个字节，和传统的ASCII编码兼容。
    - 如果第一个字节的高端bit是110，则说明需要2个字节；后续的每个高端bit都以10开头， 后面更大的也是类似操作。
    - 1110需要三个字节，11110需要四个字节，后面字节都是以10开头。


变长的编码无法直接通过索引来访问第n个字符，但是UTF8编码获得了很多额外的优点：
- 更加紧凑，完全兼容ASCII码，并且可以自动同步。
- 可以通过向前回朔最多3个字节就能确定当前字符编码的开始字节的位置。
- 它也是一个前缀编码，所以当从左向右解码时不会有任何歧义也并不需要向前查看
- 没有任何字符的编码是其他字符编码的子串，因此搜索一个字符时只要搜索它的字节编码序列即可。
- UTF8编码的顺序和Unicode码点的顺序一致，因此可以直接排序UTF8编码序列。
- 同时因为没有嵌入的NUL(0)字节，可以很好地兼容那些使用NUL作为字符串结尾的编程语言。


Go语言的源文件采用UTF8编码，并且Go语言处理UTF8编码的文本也很出色。unicode包提供了诸多处理rune字符相关功能的函数（比如区分字母和数字，或者是字母的大写和小写转换等），unicode/utf8包则提供了用于rune字符序列的UTF8编码和解码的功能。


**字符串中字节数和字符数**：字符串包含13个字节，以UTF8形式编码，但是只对应9个Unicode字符：
```go
import "unicode/utf8"

s := "Hello, 世界"
fmt.Println(len(s))                    // "13"
fmt.Println(utf8.RuneCountInString(s)) // "9"
```
- Go语言的range循环在处理字符串的时候，会自动隐式解码UTF8字符串。
```go
for i, r := range "Hello, 世界" {
    fmt.Printf("%d\t%q\t%d\n", i, r, r)
}
```
- UTF8字符串作为交换格式是非常方便的，但是在程序内部采用rune序列可能更方便，因为rune大小一致，支持数组索引和方便切割。


#### 字符串处理和转换

**字符串和Byte切片**


- 标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。
  - strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能。
  - bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。
  - strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。
  - unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。每个函数有一个单一的rune类型的参数，然后返回一个布尔值。


- 实现一个将path文件路径简化成文件名（去掉前面目录和后缀）：
```go
func basename(s string) string {
    slash := strings.LastIndex(s, "/") // -1 if "/" not found
    s = s[slash+1:]
    if dot := strings.LastIndex(s, "."); dot >= 0 {
        s = s[:dot]
    }
    return s
}
```


- **字符串和字节slice之间可以互相转换**
```go
s := "abc"
b := []byte(s)
s2 := string(b)
```
  - 需要确保在变量b被修改的情况下，原始的s字符串也不会改变。
- strings包中的六个函数：（bytes包中也对应的六个函数，区别就是类型换成了字节slice类型）

```go
func Contains(s, substr string) bool
func Count(s, sep string) int
func Fields(s string) []string
func HasPrefix(s, prefix string) bool
func Index(s, sep string) int
func Join(a []string, sep string) string

```

- bytes包还提供了Buffer类型用于字节slice的缓存。一个Buffer开始是空的，但是随着string、byte或[]byte等类型数据的写入可以动态增长，一个bytes.Buffer变量并不需要初始化，因为零值也是有效的
```go
// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
    var buf bytes.Buffer
    buf.WriteByte('[')
    for i, v := range values {
        if i > 0 {
            buf.WriteString(", ")
        }
        fmt.Fprintf(&buf, "%d", v)
    }
    buf.WriteByte(']')
    return buf.String()
}

func main() {
    fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}
```
>bytes.Buffer类型有着很多实用的功能，我们在第七章讨论接口时将会涉及到，我们将看看如何将它用作一个I/O的输入和输出对象，例如当做Fprintf的io.Writer输出对象，或者当作io.Reader类型的输入源对象。


**字符串与数字的转换**



由strconv包提供这类转换功能。
- 将一个整数转为字符串，一种方法是用fmt.Sprintf返回一个格式化的字符串；另一个方法是用strconv.Itoa(“整数到ASCII”)：
```go
  x := 123
y := fmt.Sprintf("%d", x)
fmt.Println(y, strconv.Itoa(x)) // "123 123"

```
- FormatInt和FormatUint函数可以**用不同的进制来格式化数字**：
```go
fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
```

- `fmt.Printf`函数的%b、%d、%o和%x等参数提供功能往往比strconv包的Format函数方便很多，特别是在需要**包含有附加额外信息**的时候：
```go
s := fmt.Sprintf("x=%b", x) // "x=1111011"
```

- 如果要将一个**字符串解析为整数**，可以使用strconv包的**Atoi或ParseInt**函数，还有用于解析无符号整数的ParseUint函数：
```go
x, err := strconv.Atoi("123")             // x is an int
y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
```
  - ParseInt函数的第三个参数是用于指定整型数的大小；例如16表示int16，0则表示int。
  - 在任何情况下，返回的结果y总是int64类型，你可以通过强制类型转换将它转为更小的整数类型。



### 常量
- 存储在常量中的数据类型只可以是布尔型、数字型(整数型、浮点型和复数)和字符串型。
- 常量表达式的值在编译期计算，而不是在运行期。
- 如果没有显式指明类型，那么将从右边的表达式推断类型。如果转换合法的话。
    - 可以通过%T参数打印类型信息：fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
    - 无类型整数常量转换为int，它的内存大小是不确定的，但是无类型浮点数和复数常量则转换为内存大小明确的float64和complex128。
    - 如果要给变量一个不同的类型，我们必须显式地将无类型的常量转化为所需的类型，或给声明的变量指定明确的类型：var i = int8(0)


#### iota 常量生成器
常量声明可以使用iota常量生成器初始化，它用于生成一组以相**似规则初始化的常量**，但是不用每行都写一遍初始化表达式。
- 比如星期、月份、年份
- 在第一个声明的常量所在的行，**iota将会被置为0**，然后在每一个有常量声明的行加一。


#### 无类型常量
编译器为这些没有明确基础类型的数字常量**提供比基础类型更高精度的算术运算**；
>你可以认为至少有256bit的运算精度。

这里有六种未明确类型的常量类型，分别是无类型的布尔型、无类型的整数、无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串。

- 通过延迟明确常量的具体类型，无类型的常量不仅可以提供更高的运算精度，而且可以直接**用于更多的表达式而不需要显式的类型转换。**


- 对于常量面值，**不同的写法可能会对应不同的类型**。
>例如0、0.0、0i和\u0000虽然有着相同的常量值，但是它们分别对应无类型的整数、无类型的浮点数、无类型的复数和无类型的字符等不同的常量类型


---


## 第四章 复合类型：数组和结构体

从简单的数组、字典、切片到动态列表

### 定长数组 array
数组是一个由**固定长度**的**特定类型**元素组成的序列，一个数组可以由零个或多个元素组成。
>因为长度固定，而且没有任何添加或删除数组元素的方法。
>Go语言中很少直接使用数组。 除了像SHA256这类需要处理特定大小数组的特例外。 一般用slice，但是要先理解数组。


- 索引下标的范围是从0开始到数组长度减1的位置。内置的**len函数**将返回数组中**元素的个数**。
- 在数组字面值中，如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算。
- `var r [3]int = [3]int{1, 2}`, 优先赋值前面的，r|[2]为默认值0
- 长度是数组类型的组成部分，长度不同的数组，类型是不一样的。


**初始化**：
```go
var r [3]int = [3]int{1, 2}
fmt.Println(r[2]) // "0"

s := [...]int{99: -1}
```


**比较**：
如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的
```go
import "crypto/sha256"

func main() {
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    //%x 十六进制的格式打印, %t布尔类型， %T对应的数据类型
    fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
    // Output:
    // 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
    // 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
    // false
    // [32]uint8
}
```


**数组参数，值传递**：
当调用一个函数的时候，函数的每个调用参数将会被赋值给函数内部的参数变量，所以函数参数变量接收的是一个复制的副本，并不是原始调用的变量。**而数组是值类型**，所以复制的数组改变不会影响外面。 如果是引用类型，那改变的就是复制进去的指针地址对应的值。


当然，我们可以**显式地传入一个数组指针**，那样的话函数通过指针对数组的任何修改都可以直接反馈到调用者。
```go
func zero(ptr *[32]byte) {
    *ptr = [32]byte{}
}
```




### 可变数组 slice
- 每个元素类型相同，没有固定长度。
- 是否可称为引用类型？ 有说可以叫指针结构的包装，比叫引用类型更严谨。


#### 获取子序列
- 用s[i]访问单个元素，用s[m:n]获取子序列。Go言里也采用左闭右开形式，0 ≤ m ≤ n ≤ len(s)，包含n-m个元素。

如果切片操作超出cap(s)的上限将导致一个panic异常，但是超出len(s)则是意味着扩展了slice，因为**新slice的长度会变大**：
```go
months := [...]string{1: "January", /* 省略掉定义，自行补充 */, 12: "December"}
summer := months[6:9]  //len:3 , cap:7

fmt.Println(summer[:20]) // panic: out of range
endlessSummer := summer[:5] // extend a slice (within capacity)
fmt.Println(endlessSummer)  // "[June July August September October]"
```

- []byte是字节类型**切片**
- 复制一个slice只是对底层的数组创建了一个新的slice别名


#### 难点一： 长度len 和 容量cap
**一个切片的容量总是固定的。** slice的切片操作s[i:j]，其中0 ≤ i≤ j≤ cap(s)，用于创建一个**新的**slice。
- 容量一般是从slice的**开始位置到底层数据的结尾**位置。
- 内置的len和cap函数分别返回slice的长度和容量


例子：
```go
s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
s4 := s3[3:6]
```

s3的长度和容量都是 8
s4的长度（大小）是3，**s4的容量是多少？**
- 切片的容量代表了它的底层数组的长度，但这仅限于使用make函数或者切片值字面量初始化切片的情况。
- 更通用的规则是: 一个切片的容量可以被看作是透过这个窗口最多可以看到的底层数组中元素的个数。
  - 而在底层数组不变的情况下，切片代表的**窗口可以向右扩展，直至其底层数组的末尾**。
    - 这里底层数组是最底层，哪怕slicea 从arr而来，sliceB从sliceA而来。


#### 比较
* slice之间不能比较，因此我们不能使用==操作符来判断两个slice是否含有全部相等元素。
>不过标准库提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相等（[]byte）.

- 对于其他类型的slice，我们必须自己展开每个元素进行比较,为啥不支持实现呢？
  - 第一个原因，一个slice的元素是间接引用的，一个slice甚至可以包含自身（译注：当slice声明为[]interface{}时，slice的元素可以是自身）
  - 第二个原因，因为slice的元素是间接引用的，一个固定的slice值（译注：指slice本身的值，不是元素的值）在不同的时刻可能包含不同的元素，因为**底层数组的元素可能会被修改**。


* slice唯一合法的比较操作是和nil比较，与任意类型的nil值一样，我们可以用`[]int(nil)`类型转换表达式来生成一个对应类型slice的nil值.
* 如果你需要测试一个slice是否是空的，使用len(s) == 0来判断，而不应该用s == nil来判断。


#### make生成slice
- 内置的make函数创建一个指定元素类型、长度和容量的slice。
- 容量部分可以省略，在这种情况下，容量将等于长度。
- 在底层，make创建了一个匿名的数组变量，然后返回一个slice；只有通过返回的slice才能引用底层匿名的数组变量。


#### append函数
内置的append函数用于向slice追加元素：`sliceA = append(sliceA, r)`。下面是第一个版本的appendInt函数，专门用于处理[]int类型的slice：
```go
func appendInt(x []int, y int) []int {
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x) {
        // There is room to grow.  Extend the slice.
        z = x[:zlen]
    } else {
        // There is insufficient space.  Allocate a new array.
        // Grow by doubling, for amortized linear complexity.
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2 * len(x)
        }
        z = make([]int, zlen, zcap)
        //copy函数将返回成功复制的元素的个数,等于两个slice中较小的长度
        copy(z, x) // a built-in function; see text
    }
    z[len(x)] = y
    return z
}
```
- 内置的append函数则可以追加多个元素，甚至追加一个slice。
- 内置的append函数可能使用比appendInt更复杂的内存扩展策略。
>因此，通常我们并不知道append调用是否导致了内存的重新分配，因此我们也**不能确认新的slice和原始的slice是否引用的是相同的底层数组空间**。同样，我们不能确认在原先的slice上的操作**是否会影响到新的slice**。

因此，通常是将append返回的结果直接赋值给输入的slice变量：`sliceA = append(sliceA, r)`
- 更新slice变量不仅对调用append函数是必要的，实际上对应任何可能导致长度、容量或底层数组变化的操作都是必要的。
- 要正确地使用slice，需要记住尽管底层数组的元素是间接访问的，但是slice对应结构体本身的指针、长度和容量部分是直接访问的。
- 要更新这些信息需要像上面例子那样一个显式的赋值操作。


输入的slice和输出的slice共享一个底层数组。这可以避免分配另一个数组，不过原来的数据将可能会被覆盖，正如下面两个打印语句看到的那样：
```go
func main(){
    data := []string{"one", "", "three"}
    fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
    fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
}

func nonempty(strings []string) []string {
    i := 0
    for _, s := range strings {
        if s != "" {
            strings[i] = s
            i++
        }
    }
    return strings[:i]
}

```

#### 为啥要弄这种滑动窗口式的设计


#### slice[2:] 省略掉的是len(slice)
省略掉的：默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]


从Go1.2开始slice支持了三个参数的slice：
```go
var array [10]int
sliceA := array[2:4]   //slice的容量是8
sliceB = array[2:4:7]   //这个7是底层数组的位置，表示该切片最多到这个位置。上面这个的容量就是7-2，即5。这样这个产生的新的slice就没办法访问最后的三个元素。
sliceB[2] 会抛异常，超出边界。 长度只有2.
```



#### 扩容：
- 它并不会改变原来的切片，而是会生成一个容量更大的切片，然后将把原有的元素和新元素一并拷贝到新切片中。
  - 拷贝耗费性能嘛？ 原来的切片是否保留？回收啥时候回收？
  - 扩容2倍，当原长度大于或等于1024时，Go 语言将会以原容量的 1.25倍作为新容量的基准(以下新容量基准)


**切片的底层数组什么时候会被替换?**
- 确切地说，一个切片的底层数组永远不会被替换。
- 为什么?虽然在扩容的时候 Go 语言一定会生成新的底层数组，但是它也同时生成了新的切片。
- 没有为切片替换底层数组这一说，扩容是直接换新的切片。。。
- 在无需扩容时，append函数返回的是**指向原底层数组的新切片**，而在需要扩容时，append函数返回的是**指向新底层数组的新切片**。


**知识点**
1. 初始时两个切片引用同一个底层数组，在后续操作中对某个切片的操作超出底层数组的容 量时，这两个切片引用的就不是同一个数组了


### map

#### 初始化和取值
map[keyType]valueType
- map的key，可以是int，可以是string及所有完全定义了==与!=操作的类型
  - 虽然浮点数类型也是支持相等运算符比较的，但是**将浮点数用做key类型则是一个坏的想法**，最坏的情况是可能出现的NaN和任何浮点数都不相等。
- 值则可以是任意类型,但是键之间、值之间类型要相同。
- 不要使用new，永远用make来构造map


```go
ages := make(map[string]int)
ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
fmt.Println(ages["alice"]) // "32"
delete(ages, "alice") // remove element ages["alice"]
```


- 即使map中不存在“bob”下面的代码也可以正常工作，因为ages["bob"]失败时将返回0。 如何
```go
ages["bob"] = ages["bob"] + 1 // happy birthday!
```
- map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，它们的行为和一个空的map类似。
- 但是向一个nil值的map存入元素将导致一个panic异常：
```go
ages["carol"] = 21 // panic: assignment to entry in nil map
```
- 在向map存数据前必须先创建map。 声明了一个nil的map，如何再创建？ 再写一遍`ages = make(map[string]int, 10)`？
- map下标获取值可选获取两个值： `age, ok := ages["bob"] 和  age := ages["bob"]`都是合法的？


#### 用map实现set
Go程序员将这种忽略value的map当作一个**字符串集合**。
```go
func main() {
    seen := make(map[string]bool) // a set of strings
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        line := input.Text()
        if !seen[line] {
            seen[line] = true
            fmt.Println(line)
        }
    }

    if err := input.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
        os.Exit(1)
    }
}
```


### 结构体
结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。
```go
type Employee struct {
    ID        int
    Name      string
    Address   string
    DoB       time.Time
    Position  string
    Salary    int
    ManagerID int
}

var dilbert Employee
```

**访问**：
- 直接点操作符访问和赋值：`dilbert.Salary -= 5000`
- 对成员取地址，然后通过指针访问：
```go
position := &dilbert.Position
*position = "Senior " + *position
```


**点操作符也可以和指向结构体的指针一起工作：**
```go
var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"
//上一句相当于下面，为啥呢？？？
(*employeeOfTheMonth).Position += " (proactive team player)"

```


**调用函数返回的是值，并不是一个可取地址的变量，所以不能直接用点操作符赋值**：
```go
func EmployeeByID(id int) *Employee { /* ... */ }

fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"
id := dilbert.ID
//这个赋值操作是可行的，但如果把函数方法EmployeeByID的返回值改为Employee，那下面赋值语句就会报错。
EmployeeByID(id).Salary = 0 // fired for... no real reason

```


**命名为S的结构体成员可以包含 *S指针成员**

一个命名为S的结构体类型将不能再包含S类型的成员：因为一个聚合的值不能包含它自身。（该限制同样适用于数组。）
但是S类型的结构体可以包含*S指针类型的成员，这可以让我们创建递归的数据结构，比如链表和树结构等。
可以使用一个二叉树来实现一个插入排序：


**空结构体**：
写成struct{}，它的大小为0，也不包含任何信息，但是有时候依然是有价值的：


#### 结构体字面量、声明初始化
- 写法一：只有值，按照类型和顺序一一对应。 
  - 缺点：如果做了调整，就需要修改代码。
  - 一般用在： 定义结构体的包内部使用、或者较小的结构体重使用，这些结构体成员排列比较规则
  - 比如：image.Point(x,y)、color。RGBA(red, green, blue, alpha)
- 写法二：以成员名字和相应的值来初始化，可以包含**部分**或全部的成员
  - 好处：没写的成员默认用零值，顺序不重要，切新增字段，老代码也就是默认零值，不用可以不改也不报错。


**使用规则**：
- 两种写法不能混用
- 不能在外部包中用写法一来偷偷初始化结构体中**未导出**的成员。 **TODO：未验证**
```go
package p
type T struct{ a, b int } // a and b are not exported

package q
import "p"
var _ = p.T{a: 1, b: 2} // compile error: can't reference a, b
var _ = p.T{1, 2}       // compile error: can't reference a, b
```
- 函数对结构体进行修改操作，必须传入指针；
>因为在Go语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。
- 初始化一个结构体变量（下面三种写法等价）：
```go
//这个写法可以直接在表达式中使用，比如一个函数调用
pp := &Point{1, 2}

pp := new(Point)
*pp = Point{1, 2}
```


#### 结构体嵌入和匿名成员
```go
type Point struct {
    X, Y int
}

type Circle struct {
    Center Point
    Radius int
}

type Wheel struct {
    Circle Circle
    Spokes int
}
```
- 结构体类型清晰，但是访问每个成员变量变得繁琐，需要多级。w.Circle.Center.Y = 8


**如何解决访问繁琐的问题？**
- **匿名成员**：只声明一个成员对应的数据类型而不指名成员的名字
  - 匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针。
  - 
```go
type Point struct {
    X, Y int
}

type Circle struct {
    Point
    Radius int
}

type Wheel struct {
    Circle
    Spokes int
}
```
- **匿名嵌入的特性**是啥？ 可以让我们直接访问叶子属性：`w.X = 8 //equivalent to w.Circle.Point.Y = 8`
  - 匿名成员Circle和Point都有自己的名字——就是**命名的类型名字**——但是这些名字在点操作符中是可选的
  - 不能同时包含两个类型相同的匿名成员，这会导致名字冲突。
  - 包内Point和Circle匿名成员都是导出的。即使它们不导出（比如结构名改成小写字母开头的point和circle）。但是在包外部，因为circle和point没有导出，不能访问它们的成员
- 但是匿名成员就没法用下面方式进行声明初始化了：
```go
w = Wheel{8, 8, 5, 20}                       // compile error: unknown fields
w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields
```
需要写成下面这样：
```go
w = Wheel{Circle{Point{8, 8}, 5}, 20}

w = Wheel{
    Circle: Circle{
        Point:  Point{X: 8, Y: 8},
        Radius: 5,
    },
    Spokes: 20, // NOTE: 这个逗号是必须的
}

// Printf函数中%v参数包含的#副词，它表示用和Go语言类似的语法打印值。对于结构体类型来说，将包含每个成员的名字。
fmt.Printf("%#v\n", w)  //Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}
```


**匿名成员还有其他好处嘛？**为什么要嵌入一个没有任何子成员类型的匿名成员类型呢？
- 匿名成员并不要求是结构体类型；其实任何命名的类型都可以作为结构体的匿名成员。
- **匿名类型的方法集**：简短的点运算符语法除了访问匿名成员嵌套的成员，还可以访问它们的方法。这个机制可以用于**将一些有简单行为的对象组合成有复杂行为的对象**。组合是Go语言中面向对象编程的核心


### JSON
Go语言对于这些标准格式的编码和解码都有良好的支持，由标准库中的encoding/json、encoding/xml、encoding/asn1等包提供支持
>Protocol Buffers的支持由 github.com/golang/protobuf 包提供），并且这类包都有着相似的API接口。


**JSON与go基础类型、对象的对应表示**：
```go
boolean         true
number          -273.15
string          "She said \"Hello, BF\""
array           ["gold", "silver", "bronze"]
object          {"year": 1980,
                 "event": "archery",
                 "medals": ["gold", "silver", "bronze"]}
```


**定义**
```go
type Movie struct {
    Title  string
    Year   int  `json:"released"`  //一个结构体成员Tag是和在编译阶段关联到该成员的元信息字符串
    Color  bool `json:"color,omitempty"` //json串这个字段的field是 color； 
                                        // omitempty 当Go语言结构体成员为空或零值时不生成该JSON对象(只是这个字段不生成)
    Actors []string
}

//该函数有两个额外的字符串参数用于表示每一行输出的前缀和每一个层级的缩进：
data, err := json.MarshalIndent(movies, "", "    ")

//下面的代码将JSON格式的电影数据解码为一个结构体slice，结构体中只有Title成员。
//通过定义合适的Go语言数据结构，我们可以选择性地解码JSON中感兴趣的成员。
var titles []struct{ Title string }
if err := json.Unmarshal(data, &titles); err != nil {
    log.Fatalf("JSON unmarshaling failed: %s", err)
}
fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
```
- Go => JSON, 叫**编组**（marshaling）
- JSON => Go，叫**解码**（Unmarshal）



**应用**
- 许多web服务都提供JSON接口，通过HTTP接口发送JSON格式请求并返回JSON格式的信息。
  - 第一步：定义结构
  - 第二步：定义方法，处理结构体，编码解码
  - 第三步：调用方法
```go
//gopl.io/ch4/github
const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

//gopl.io/ch4/issues
func main() {
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%d issues:\n", result.TotalCount)
    for _, item := range result.Items {
        fmt.Printf("#%-5d %9.9s %.55s\n",
            item.Number, item.User.Login, item.Title)
    }
}
```
-  https://developer.github.com/v3/ 
-  https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述。下载每个链接（只下载一次）然后创建一个离线索引。编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。
-  检索和下载 https://omdbapi.com/ 上电影的名字和对应的海报图像。编写一个poster工具，通过命令行输入的电影名字，下载对应的海报。


### 文本和HTML模板

- 一个模板是一个字符串或一个文件，连包含了一个或多个由双花括号包含的`{{action}}`对象。
- 大部分的字符串只是按字面值打印，但是对于actions部分将触发其它的行为.
- 模板语言包含通过选择结构体的成员、调用函数或方法、表达式控制流if-else语句和range循环语句，还有其它实例化模板等诸多特性

#### 模板字符串
```go
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`
```
- 对于每一个action，都有一个当前值的概念，对应点操作符，写作“.”。
- 当前值“.”**最初被初始化为调用模板时的参数**，在当前例子中对应github.IssuesSearchResult类型的变量。
- 模板中`{{range .Items}}和{{end}}`对应一个循环action，每次迭代的当前值对应当前的Items元素的值。
- `|`操作符表示将前一个表达式的结果作为后一个函数的输入，类似于UNIX中管道的概念。
  >`printf`一个基于fmt.Sprintf实现的内置函数，所有模板都可以直接使用。 

```go
// 先创建并返回一个模板；
report, err := template.New("report").
    Funcs(template.FuncMap{"daysAgo": daysAgo}). // Funcs方法将daysAgo等自定义函数注册到模板中，并返回模板；
    Parse(templ)  //调用Parse函数分析模板
if err != nil {
    log.Fatal(err)
}

// template.Must辅助函数可以简化这个致命错误的处理：它接受一个模板和一个error类型的参数，检测error是否为nil（如果不是nil则发出panic异常），然后返回传入的模板。
var report = template.Must(template.New("issuelist").
    Funcs(template.FuncMap{"daysAgo": daysAgo}).
    Parse(templ))

//使用github.IssuesSearchResult作为输入源、os.Stdout作为输出源来执行模板：
func main() {
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }
    if err := report.Execute(os.Stdout, result); err != nil {
        log.Fatal(err)
    }
}
```
  

#### HTML模板
- html/template包已经自动将特殊字符转义
- 如果不想被转义，可以把对应字符串定义到信任的template.HTML字符串类型，最终生成的html文件就不会转义这个字段的字符串。
```go
import "html/template"

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
```


## 第五章 函数、错误处理、panic、recover、有defer语句。
>引用类型包括指针（§2.3.2）、切片（§4.2)）、字典（§4.3）、函数（§5）、通道（§8）,它们都是对程序中一个变量或状态的间接引用。这意味着对任一引用类型数据的修改都会影响所有该引用的拷贝。


**简写**：
```go
//以下语句等价
func f(i, j, k int, s, t string)                 { /* ... */ }
func f(i int, j int, k int,  s string, t string) { /* ... */ }

func first(x int, _ int) int { return x }
fmt.Printf("%T\n", first) // "func(int, int) int" 类型
```

- 函数的类型被称为函数的签名。
  - 如果两个函数**形式参数列表和返回值列表中的变量类型**一一对应，那么这两个函数被认为**有相同的类型或签名**。
  - 形参和返回值的**变量名不影响函数签名**，也不影响它们是否可以以省略参数类型的形式表示。
>每一次函数调用都必须按照声明顺序为所有参数提供实参（参数值）。
>在函数调用时，Go语言**没有默认参数值**，也没有任何方法可以通过参数名指定形参，因此形参和返回值的变量名对于函数调用者而言没有意义。

- **实参通过值的方式传递**，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。
- 但是，如果**实参包括引用类型**，如指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的间接引用**被修改**。

- 没有函数体的函数声明：表示该函数不是以Go实现的


### 递归
```go
//!+
func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}


// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}


// visit appends to links each link found in n, and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

```

- 在findlinks中，我们必须确保resp.Body被关闭，释放网络资源。虽然**Go的垃圾回收机制**会回收不被使用的内存，但是这**不包括操作系统层面的资源，比如打开的文件、网络连接**。因此我们必须显式的释放这些资源。
- 一个函数内部可以将另一个有多返回值的函数调用作为返回值，下面的例子展示了与findLinks有相同功能的函数。
```go
func findLinksLog(url string) ([]string, error) {
    log.Printf("findLinks %s", url)
    return findLinks(url)
}
```
- 如果一个函数**所有的返回值都有显式的变量名**，那么该函数的**return语句可以省略操作数*。这称之为bare return。
  - 但是使得代码难以被理解。
  - 用在卫语句挺合适，还没赋值的就是默认零值。 Go会将返回值 words和images在函数体的开始处，根据它们的类型，将其初始化为0。
```go
func CountWordsAndImages(url string) (words, images int, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        err = fmt.Errorf("parsing HTML: %s", err)
        return
    }
    words, images = countWordsAndImages(doc)
    return
}
func countWordsAndImages(n *html.Node) (words, images int) { /* ... */ }
```

### 错误
>一个良好的程序永远不应该发生panic异常。

- 如果导致失败的原因只有一个，额外的返回值可以是一个布尔值，通常被命名为**ok**。
  >比如，cache.Lookup失败的唯一原因是key不存在。 比如map获取值。
- error类型可能是nil或者non-nil。
  - nil意味着函数运行成功，non-nil表示失败。
  - 对于non-nil的error类型，我们可以通过调用error的Error函数或者输出函数获得字符串类型的错误信息。`fmt.Printf("%v", err)`
  - 当函数返回non-nil的error时，其他的返回值是未定义的（undefined），这些未定义的返回值应该被忽略。
  - 有少部分函数在发生错误时，仍然会返回一些有用的返回值。比如，当读取文件发生错误时，Read函数会返回可以读取的字节数以及错误信息。 **应该是先处理这些不完整的数据，再处理错误**
  - Go语言将函数运行失败时返回的错误信息当做一种预期的值，而不是异常。 而Go对于异常处理是针对哪些未被预料到的错误。panic？bug？
  - Go使用控制流机制（如if和return）处理错误，这使得编码人员能更多的关注错误处理。


**错误处理策略**

1. 传播错误
  - 要增加必要的上下文，再传播到上游。
    ```go
    if err != nil {
        return nil, fmt.Errorf("parsing %s as HTML: %v", url,err)
    }    
    ```
  - 由于错误信息经常是以链式组合在一起的，所以错误信息中应**避免大写和换行符**。
  - 要注意**错误信息表达的一致性**，即相同的函数或同包内的同一组函数返回的错误在构成和处理方式上是相似的。
    >一般而言，被调用函数f(x)会将调用信息和参数信息作为发生错误时的上下文放在错误信息中并返回给调用者，调用者需要**添加一些错误信息中不包含的信息**，比如添加url到html.Parse返回的错误中。
2. 重新尝试失败的操作
  - 什么情况下用？如果错误的发生是偶然性的，或由不可预知的问题导致的。
  - 注意事项？ 限定重试时间间隔或重试次数
```go
func WaitForServer(url string) error {
    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)
    for tries := 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)
        if err == nil {
            return nil // success
        }
        log.Printf("server not responding (%s);retrying…", err)
        time.Sleep(time.Second << uint(tries)) // exponential back-off
    }
    return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
```

3. 输出错误信息并**结束程序**
   - 注意的是，这种策略只应在main中执行。
   - 对库函数而言，应仅向上传播错误，除非该错误意味着程序内部包含不一致性，即遇到了bug，才能在库函数中结束程序。

```go
// (In function main.)
if err := WaitForServer(url); err != nil {
    fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
    os.Exit(1)
}

//等价于
if err := WaitForServer(url); err != nil {
    log.Fatalf("Site is down: %v\n", err)
}
```


4. 只输出错误信息。 不中断也不往上传，使用log打印错误
```go
if err := Ping(); err != nil {
    log.Printf("ping failed: %v; networking disabled",err)
}
```
>log包中的所有函数会为没有换行符的字符串增加换行符。


5. 直接忽略掉错误。
   - 当你决定忽略某个错误时，你应该清晰地写下你的意图
   - 比如删除临时目录，有定时任务兜底，你主动删除操作失败也无所谓，就可以不处理删除失败的情况。


**文件结尾错误（EOF）**


- io包保证任何由文件结束引起的读取失败都返回同一个错误——io.EOF
```go
if err == io.EOF {
        break // finished reading
    }
```


**函数值**


- 函数值：被作为参数的函数，可以带有参数值（就是带着行为的状态？）。
- 函数类型的零值是nil。调用值为nil的函数值会引起panic错误。
- 函数值可以与nil比较，但是函数值之间是不可比较的，也不能用函数值作为map的key。

```go
 var f func(int) int
 f(3) // 此处f的值为nil, 会引起panic错误

  if f != nil {
        f(3)
   }
```
- 通过行为来参数化函数：下面的`strings.Map`对字符串中的每个字符调用add1函数，并将每个add1函数的返回值组成一个新的字符串返回给调用者。
```go
func add1(r rune) rune { return r + 1 }

fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
```


### 匿名函数
>拥有函数名的函数只能在包级语法块中被声明
**匿名函数**： 通过函数字面量（是一种表达式，就是函数声明func后不带函数名，而是直接func()）语法可以在任何表达式中表示一个函数值。 
```go
strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
```

- 更为**重要**的是，通过这种方式定义的函数可以访问完整的**词法环境**（lexical environment），这意味着**在函数中定义的内部函数可以引用该函数的变量**。
>不是很理解这个意味着有啥特别？ 啥场景？ 很难嘛？之前不支持？ **看代码吧**
```go
// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}
func main() {
    f := squares()
    fmt.Println(f()) // "1"
    fmt.Println(f()) // "4"
    fmt.Println(f()) // "9"
    fmt.Println(f()) // "16"
}
```
- 为啥每次调用值不一样?squares（）里的局部变量不是每次调用`f()匿名函数`就会归零？
  - 第二次调用squares时，会生成第二个x变量，并返回一个新的匿名函数。新匿名函数操作的是第二个x变量。 为啥第二个x变量初始值是1？
- **squares的例子证明，函数值不仅仅是一串代码，还记录了状态**。
- 在squares中定义的匿名内部函数可以访问和更新squares中的局部变量，这意味着**匿名函数和squares中，存在变量引用**。
- 这就是**函数值**属于引用类型和函数值不可比较的原因。Go使用闭包（closures）技术实现**函数值**，Go程序员也**把函数值叫做闭包**。
  >这函数值是啥，越来越糊涂了？
- 这个例子展示了：**变量的生命周期不由它的作用域决定**：squares返回后，变量x仍然隐式的存在于f中。


**网页抓取的核心问题就是如何遍历图**。
- 在topoSort的例子中，已经展示了深度优先遍历，在网页抓取中，我们会展示如何用广度优先遍历图。


**警告：捕获迭代变量**

>介绍Go词法作用域的一个陷阱

需求：创建一些目录，然后将目录删除。
- 没问题的代码：
  - 为什么要在循环体中用循环变量d赋值一个新的局部变量，为什么要在循环体中用循环变量d赋值一个新的局部变量？ 后面一种是错误的
```go
//没问题的
var rmdirs []func()  //方法切片
for _, d := range tempDirs() {
    dir := d // NOTE: necessary!
    os.MkdirAll(dir, 0755) // creates parent directories too
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir)
    })
}

// ...do some work…
for _, rmdir := range rmdirs {
    rmdir() // clean up
}
```
- 有问题的
```go
var rmdirs []func()
for _, dir := range tempDirs() {
    os.MkdirAll(dir, 0755)
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir) // NOTE: incorrect! 这一步是把的dir传入函数中作为函数的变量？存起来了
    })
}
// ...do some work…
for _, rmdir := range rmdirs {
    rmdir() // clean up
}
```
- for循环语句引入了新的词法块，循环变量dir在这个词法块中被声明。
- 在该循环中生成的所有**函数值**（带参数的函数？）都共享相同的循环变量
- **注意**：函数值中记录的是循环变量的**内存地址**，而不是循环变量某一时刻的值。
- 以dir为例，后续的迭代会不断更新dir的值，**当删除操作执行时，for循环已完成**，dir中存储的值等于最后一次迭代的值。
- 这意味着，每次对os.RemoveAll的调用删除的都是相同的目录。 


### 可变参数
参数数量可变的函数称为可变参数函数。
- 需要在参数列表的最后一个**参数类型之前**加上省略符号“...”
- 调用者**隐式的创建一个数组**，并将原始参数复制到数组中，再把数组的一个切片作为参数传给被调用函数。
- 可以使用range去遍历它，那如果可变参数的类型就是一个切片呢？ 相当于传了个二维切片，可以两次for循环遍历。
```go
values := []int{1, 2, 3, 4}
fmt.Println(sum(values...)) // "10"
//等价于下面的
fmt.Println(sum(1, 2, 3, 4)) // "10"
```


**用处**：
- 可变参数函数经常被用于格式化字符串。 format后面带的多个参数。


### defer函数
- 当执行到该条语句时，**函数和参数表达式得到计算**，但直到包含该defer语句的函数执行完毕时，**defer后的函数**才会被执行。 （看下面的例子，似乎是defer最后的return是最后执行，其他的都是在经过时运行？）
  >不太理解函数和参数表达式得到计算是什么意思？ defer语句里面带的函数和表达式先计算了？ 那不就是执行了嘛？ 还是主函数流程？ defer后的函数是啥
  **！！！测试发现**：defer后面的函数里面只要有return语句，那return之前的语句会执行（比如打印进入函数的时间），会停留大return位置，当原函数结束时执行return；
  - 测试的方法是**返回一个函数**符合上面的说法，如果返回int呢？ 
  - **这里要注意一个细节：** `defer trace("bigSlowOperation")()`后面的圆括号，表示前面部分返回必须为一个func然后方便带上()? 返回int这么写会报错，可以去掉圆括号，那么就不会报错，但是整个`trace("bigSlowOperation")`都会在原函数结束时执行.


- 不论包含defer语句的函数是通过return正常结束，还是由于panic导致的异常结束。
- 可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序**相反**。经过一条defer语句就入栈，最后执行时是出栈这些defer语句。
- **常被用于**处理**成对**的操作，如打开、关闭、连接、断开连接、加锁、释放锁。


调试复杂程序时，defer机制也常被**用于记录何时进入和退出函数**。
```go
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
```

>注意一点：不要忘记defer语句后的**圆括号**，否则本该在进入时执行的操作会在退出时执行，而本该在退出时执行的，永远不会被执行。


defer语句中的函数会在return语句更新返回值变量后再执行，又因为在函数中定义的匿名函数可以访问该函数**包括返回值变量在内的所有变量**.
所以，对匿名函数采用defer机制，可以使其**观察函数的返回值**。 就是在defer里打印出来？对于有许多return语句的函数而言，这个技巧很有用。


defer后面的函数可以修改 命名的返回值。


for循环里面，不断打开file，只定义一个defer，可能出现打开太多，用满文件描述符。 解决办法： 将打开操作抽象出一个函数，在函数里定义defer，那每次打开就会及时关闭，再打开下一个。


### Panic异常
- 一般会将panic异常和日志信息一并记录。
- 直接调用内置的panic函数也会引发panic异常；panic函数接受任何值作为参数。
- 当某些不应该发生的场景发生时，我们就应该调用panic。比如，当程序到达了某条逻辑上不可能到达的路径：
>在健壮的程序中，任何可以预料到的错误，如不正确的输入、错误的配置或是失败的I/O操作都应该被优雅的处理

**panic的适用场景与其他语言Exception的区别**
- panic一般用于严重错误，如程序内部的逻辑不一致。 优先使用错误处理机制，而不是panic。


**在Go的panic机制中，延迟函数的调用在释放堆栈信息之前。**
- 如何使程序从panic异常中恢复，阻止程序的崩溃。?
>为了方便诊断问题，runtime包允许程序员输出堆栈信息。在下面的例子中，我们通过在main函数中延迟调用printStack输出堆栈信息。 就是defer捕获，然后优雅的输出panic等错误信息
```go
func main() {
    defer printStack()
    f(3)
}
func printStack() {
    var buf [4096]byte
    n := runtime.Stack(buf[:], false)
    os.Stdout.Write(buf[:n])
}
```


### Recover捕获异常
> 一般不应该对panic异常做任何处理
如果想从异常中恢复，或者说在程序奔溃前做一些操作，可以考虑使用recover()

- 比如： 当web服务器遇到不可预料的严重问题时，在崩溃前应该将所有的连接关闭；


如果在deferred函数中调用了内置函数recover，并且定义该defer语句的函数发生了panic异常，recover会使程序从panic中恢复，并返回panic value。
- 这个panic value就是recover函数返回值，由上游 调用 panic(value)传入

* 导致panic异常的函数不会继续运行，但能正常返回。
- 在未发生panic时调用recover，recover会返回nil。

- 不应该试图去恢复其他包或者由他人开发的函数引起的panic。
- 公有的API应该将函数的运行失败作为error返回，而不是panic。
- 只恢复应该被恢复的panic异常：
  - web服务器遇到处理函数导致的panic时会调用recover，输出堆栈信息，继续运行。
  - 为了标识某个panic是否应该被恢复，我们可以将panic value设置成特殊类型。
    - 在recover时对panic value进行检查，如果发现panic value是特殊类型，就将这个panic作为error处理。
```go
func soleTitle(doc *html.Node) (title string, err error) {
    type bailout struct{}
    defer func() {
        switch p := recover(); p {
        case nil:       // no panic
        case bailout{}: // "expected" panic
            err = fmt.Errorf("multiple title elements")
        default:
            panic(p) // unexpected panic; carry on panicking
        }
    }()
    // Bail out of recursion if we find more than one nonempty title.
    forEachNode(doc, func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "title" &&
            n.FirstChild != nil {
            if title != "" {
                panic(bailout{}) // multiple titleelements
            }
            title = n.FirstChild.Data
        }
    }, nil)
    if title == "" {
        return "", fmt.Errorf("no title element")
    }
    return title, nil
}
```

---

## 第六章 方法（method）
一个方法则是一个一个和特殊类型关联的函数。


### 方法声明
**在函数声明时**，在其名字之前放上一个变量，**即是一个方法**。

这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法。

```go
package geometry

import "math"

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```
- 上面的代码里那个`(p Point)`附加的参数p，叫做**方法的接收器**（receiver）.命名使用类型的第一个字母，简约。
- p.Distance的表达式叫做**选择器**，因为他会选择合适的对应p这个对象的Distance方法来执行。
- 选择器也会被用来选择一个struct类型的字段，比如p.X。
>由于方法和字段都是在同一命名空间，所以如果我们在这里声明一个X方法的话，编译器会报错，因为在调用p.X时会**有歧义.** **方法和字段不能同名？** 但是方法却可以同名？比如下面的：`func (path Path) Distance()`与上面的在一个包下，由于接收器不同，也可以同名而不报错。**两个Distance方法有不同的类型**。他们两个方法之间没有任何关系。
```
// A Path is a journey connecting the points with straight lines.
type Path []Point
// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
    sum := 0.0
    for i := range path {
        if i > 0 {
            sum += path[i-1].Distance(path[i])
        }
    }
    return sum
}
```
>但是**函数的签名其实就是函数的参数列表和结果列表的统称**，它定义了可用来鉴别不同函数的那些特征，同时也定义了我们与函数交互的方式。
上面同名方法算函数嘛？ 如果算，那签名是一致的，名字也一样，只有接收器不一样，不会报错？  
- **进入了误区**： 函数同名不能出现一个包下，哪怕签名不一样也不行。 跟参数名字冲突也不行？ 但是接收器不同那就是可以？

>因为**每种类型都有其方法的命名空间**，我们在用Distance这个名字的时候，不同的Distance调用指向了不同类型里的Distance方法。


也就是说，一个包内有一个方法的命名空间，而每种类型有自己的方法命名空间。 
- 在每个命名空间内部 方法、函数、字段都不能重名，调用会有歧义。
- 在不同命名空间可以出现重名（哪怕都在一个文件里，也可以）


Path是一个命名的slice类型，而不是Point那样的struct类型，然而我们依然可以为它定义方法。**这也是Go与其他语言不同**
- 可以给同一个包内的任意命名类型定义方法，只要这个命名类型的底层类型不是指针或者interface。


### 基于指针对象的方法
- 一般会约定如果Point这个类有一个指针作为接收器的方法，那么所有Point的方法都必须有一个指针接收器，即使是那些并不需要这个指针接收器的函数。
- 在声明方法时，如果一个类型名本身是一个指针的话，是不允许其出现在接收器中的。



如果接收器p是一个Point类型的变量，并且其方法需要一个Point指针作为接收器，可以用下面这种简短的写法：
```go
p.ScaleBy(2)
```
编译器会隐式地帮我们用&p去调用ScaleBy这个方法。这种简写方法只适用于“变量”。临时变量的地址获取不到，所以不行：`Point{1, 2}.ScaleBy(2)`


在每一个合法的方法调用表达式中，也就是下面三种情况里的任意一种情况都是可以的：

1. 要么接收器的实际参数和其形式参数是相同的类型，比如两者都是类型T或者都是类型*T：
```go
Point{1, 2}.Distance(q) //  Point
pptr.ScaleBy(2)         // *Point
```
2. 或者接收器实参是类型T，但接收器形参是类型*T，这种情况下编译器会隐式地为我们取变量的地址：
```go
p.ScaleBy(2) // implicit (&p)
```
3. 或者接收器实参是类型*T，形参是类型T。编译器会隐式地为我们解引用，取到指针指向的实际变量：
```go
pptr.Distance(q) // implicit (*pptr)
```


- 不管你的method的receiver是指针类型还是非指针类型，都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换。
- 在声明一个method的receiver该是指针还是非指针类型时，你需要考虑两方面的因素，第一方面是这个对象本身是不是特别大，如果声明为非指针变量时，调用会产生一次拷贝；第二方面是如果你用指针类型作为receiver，那么你一定要注意，这种指针类型指向的始终是一块内存地址，就算你对其进行了拷贝。


**Nil也是一个合法的接收器类型**
- 一个map取值的知识点
  - 直接写nil.Get("item")的话是无法通过编译的，因为nil的字面量编译器无法判断其准确类型。
  - 尝试更新一个空map会报panic
```go
m = nil
fmt.Println(m.Get("item")) // ""
m.Add("item", "3")         // panic: assignment to entry in nil map
```


### 通过嵌入结构体来扩展类型
- 被嵌入类型可以直接访问（匿名）嵌入的类型的字段和方法，不需要调用嵌入类型。  `cp.Point.X` 可以简写成 `cp.Y`
- 类型中内嵌的匿名字段也可能是一个命名类型的指针，这种情况下字段和方法会被间接地引入到当前的类型中（访问时也需要先调用该对象再访问其字段）
- 当编译器解析一个选择器到方法时，比如p.ScaleBy，它会首先去找直接定义在这个类型里的ScaleBy方法，然后找被ColoredPoint的内嵌字段们引入的方法，然后去找Point和RGBA的内嵌字段引入的方法，然后一直递归向下找。
  - 在同一级出现一样的就会报错，有歧义，编译器不知道选择哪个。


### 方法值和方法表达式
没太看懂啥用


### Bit数组
- Go语言里的集合一般会用map[T]bool这种形式来表示，T代表元素类型。
- 表示非负整数时，使用bit数组，当集合的第i位被设置时，我们才说这个集合包含元素i。


**fmt会直接调用用户定义的String方法**
- 当时有个需要注意的地方：
  - 直接fmt.Println(x)，会调用x类型的String(),如果只定义了x指针接收器的String()方法，那这里就会以原始的方式打印。 所以在这种情况下&符号是不能忘的（fmt.Println(&x)）。在我们这种场景下，你把String方法绑定到IntSet对象上，而不是IntSet指针上可能会更合适一些。
  - 当然，如果这样写 `fmt.Println(x.String())`，编译器会隐式地在x前插入&操作符，也能正确调用到x的指针方法。


### 封装
**优点**
- 首先，因为调用方不能直接修改对象的变量值，其只需要关注少量的语句并且只要弄懂少量变量的可能的值即可。
- 第二，隐藏实现的细节，可以防止调用方依赖那些可能变化的具体实现，这样针对实现可以做很多优化，只要不破坏对外暴露的api。
- 第三，也是最重要的，阻止了外部调用方对对象内部的值任意地进行修改。


在命名一个**getter方法**时，我们通常会**省略掉前面的Get前缀**。
- 这种简洁上的偏好也可以推广到各种类型的前缀比如Fetch，Find或者Lookup。


## 第七章 接口
>接口类型是对其它类型行为的抽象和概括；
- Go语言的接口类型——满足隐式实现的。也就是说，我们没有必要对于给定的具体类型定义所有满足的接口类型。
  - 好处一：可以让你创建一个新的接口类型满足已经存在的具体类型却不会去改变这些类型的定义；
  - 好处二：当我们使用的类型来自于不受我们控制的包时这种设计尤其有用。


### 接口约定
- 它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；它们只会表现出它们自己的方法。
- 当我们看到一个接口类型的值时，不知道它是什么，只知道通过它的方法来做什么。


**举个例子：**
- fmt.Printf，它会把结果写到标准输出
- fmt.Sprintf，它会把结果以字符串的形式返回
- 它们都使用了另一个函数fmt.Fprintf来进行封装。
```go
package fmt

func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
func Printf(format string, args ...interface{}) (int, error) {
    return Fprintf(os.Stdout, format, args...)
}
func Sprintf(format string, args ...interface{}) string {
    var buf bytes.Buffer
    Fprintf(&buf, format, args...)
    return buf.String()
}
```
- **注意**： Fprintf的前缀F表示文件（File）也表明格式化输出结果应该被写入第一个参数提供的文件中。
  - 在Printf函数中的第一个参数os.Stdout是*os.File类型；
  - 在Sprintf函数中的第一个参数&buf是一个指向可以写入字节的内存缓冲区，然而它并不是一个文件类型尽管它在某种意义上和文件类型相似。
  - **其实**，只要第一个参数实现了io.Writer接口类型的方法就行。io.Writer类型定义了函数Fprintf和这个函数调用者之间的**约定**。
```go
// Writer is the interface that wraps the basic Write method.
type Writer interface {
    // Write writes len(p) bytes from p to the underlying data stream.
    // It returns the number of bytes written from p (0 <= n <= len(p))
    // and any error encountered that caused the write to stop early.
    // Write must return a non-nil error if it returns n < len(p).
    // Write must not modify the slice data, even temporarily.
    //
    // Implementations must not retain p.
    Write(p []byte) (n int, err error)
}
```


### 接口类型
>一个实现了这些方法的**具体类型**是这个接口类型的**实例**。

io.Writer类型是用得最广泛的接口之一，因为它提供了所有类型的写入bytes的抽象，包括文件类型，内存缓冲区，网络链接，HTTP客户端，压缩工具，哈希等等。
**Go语言有单方法接口的命名习惯**，比如Reader，Closer。


可以内嵌组合这些接口，当然，方式有多种，下面三种都是等价的。
```go
//组合内嵌
type ReadWriter interface {
    Reader
    Writer
}
//不用内嵌
type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}
//混搭
type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Writer
}
```


### 实现接口的条件
Go的程序员经常会简要的把一个具体的类型描述成一个特定的接口类型。举个例子，`*bytes.Buffer`是`io.Writer`；`*os.Files`是`io.ReadWriter`。


如果类型实现了接口x，那就能直接赋值给接口x。
```go
var w io.Writer
w = os.Stdout           // OK: *os.File has Write method
w = new(bytes.Buffer)   // OK: *bytes.Buffer has Write method
w = time.Second         // compile error: time.Duration lacks Write method
```


先解释一个类型持有一个方法的表示当中的细节
- 对于每一个命名过的具体类型T；它的一些方法的接收者是类型T本身然而另一些则是一个*T的指针。
- 在T类型的参数上调用一个*T的方法是合法的，只要这个参数是一个变量；（编译器隐式的获取了它的地址）
  - 但请注意：这说明T类型的值没有拥有全部*T指针的方法，它就可能只实现了部分的接口。 
  - 举个例子：IntSet类型的String方法的接收者是一个指针类型，所以我们不能在一个不能寻址的IntSet值上调用这个方法。
    ```go
    type IntSet struct { /* ... */ }
    func (*IntSet) String() string
    var _ = IntSet{}.String() // compile error: String requires *IntSet receiver

    ```
  - 但是可以在一个IntSet变量上调用这个方法
    ```go
    var s IntSet
    var _ = s.String() // OK: s is a variable and &s has a String method 
    ```
- 也就是说只有 *IntSet类型实现了fmt.Stringer接口
```go
var _ fmt.Stringer = &s // OK
var _ fmt.Stringer = s  // compile error: IntSet lacks String method
```


**空接口： interface{}**
- 空接口类型对实现它的类型没有要求，所以我们可以将任意一个值赋给空接口类型。
  - 当然不能直接对它持有的值做操作，因为interface{}没有任何方法。
  - 可以用类型断言来获取interface{}中值的方法。


如果我们发现我们需要以同样的方式处理Audio和Video，我们可以定义一个Streamer接口来代表它们之间相同的部分而不必对已经存在的类型做改变。 
```go
type Audio interface {
    Stream() (io.ReadCloser, error)
    RunningTime() time.Duration
    Format() string // e.g., "MP3", "WAV"
}
type Video interface {
    Stream() (io.ReadCloser, error)
    RunningTime() time.Duration
    Format() string // e.g., "MP4", "WMV"
    Resolution() (x, y int)
}
type Streamer interface {
    Stream() (io.ReadCloser, error)
    RunningTime() time.Duration
    Format() string
}
```


### 7.4 flag.Value接口
一个标准的接口类型flag.Value是怎么帮助命令行标记定义新的符号的？


### 接口值
在一个接口值中，类型部分代表与之相关类型的描述符。 类型描述符--比如类型的名称和方法

一个接口的零值就是它的类型type和值value的部分都是nil。


**比较**
- 接口值可以使用==和!＝来进行比较。
- 两个接口值**相等**仅当它们都是nil值，或者它们的**动态类型相同**并且**动态值**也相等(根据这个动态值类型对应的==操作相等，要求必须是可比较的）。
- 接口值是可比较的，所以它们可以用在map的键或者作为switch语句的操作数。
- 如果两个接口值的动态类型相同，但是这个动态类型是不可比较的（比如切片），将它们进行比较就会失败并且panic:
- 接口介于安全的可比较类型和完全不可比较类型之间
  - 在比较接口值或者包含了接口值的聚合类型时，我们必须要意识到潜在的panic
  - 在使用接口作为map的键或者switch的操作数也要注意
  - 只能比较你非常确定它们的动态值是可比较类型的接口值。
  - **可以使用使用fmt包的%T动作获取接口值的动态类型。**（内部使用反射来获取接口动态类型的名称）


**注意：一个包含nil指针的接口不是nil接口**
- 主要区别就是接口的type是否也为nil？
```go
const debug = true

func main() {
    var buf *bytes.Buffer
    if debug {
        buf = new(bytes.Buffer) // enable collection of output
    }
    f(buf) // NOTE: subtly incorrect!
    if debug {
        // ...use buf...
    }
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
    // ...do something...
    if out != nil {  //它的动态类型是*bytes.Buffer，是一个包含空指针值的非空接口
        out.Write([]byte("done!\n"))  // panic: nil pointer dereference， 因为 out的动态值为空。
    }
}
```
- 解决方案：就是将main函数中的变量buf的类型改为io.Writer，因此可以避免一开始就将一个不完整的值赋值给这个接口。

**为啥？？？**


### sort.Interface接口
Go语言的sort.Sort函数**不会对具体的序列和它的元素做任何假设**。
- 使用了一个接口类型sort.Interface来指定通用的排序算法和可能被排序到的序列类型之间的约定。
  - 一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式；
```go
package sort

type Interface interface {
    Len() int
    Less(i, j int) bool // i, j are indices of sequence elements
    Swap(i, j int)
}
```


### http.Handler接口

- net/http包提供了一个请求多路器ServeMux来简化URL和handlers的联系。一个ServeMux将一批http.Handler聚集到一个单一的http.Handler中。
```go
func main() {
    db := database{"shoes": 50, "socks": 5}
    mux := http.NewServeMux()
    mux.Handle("/list", http.HandlerFunc(db.list))
    mux.Handle("/price", http.HandlerFunc(db.price))
    log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
    for item, price := range db {
        fmt.Fprintf(w, "%s: %s\n", item, price)
    }
}
//省略price方法

```
  - db.list是一个实现了handler类似行为的函数, 它不满足http.Handler接口并且不能直接传给mux.Handle。
  - 语句http.HandlerFunc(db.list)是一个转换而非一个函数调用，因为http.HandlerFunc是一个类型。


```go
package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```
- HandlerFunc显示了在Go语言接口机制中一些不同寻常的特点:
  - 它是一个实现了接口http.Handler的方法的函数类型。
  - ServeHTTP方法的行为是调用了它的函数本身。因此**HandlerFunc是一个让函数值满足一个接口的适配器**，这里函数和这个接口**仅有的方法有**相同的函数签名。

- 为了方便，net/http包提供了一个全局的ServeMux实例DefaultServerMux和包级别的http.Handle和http.HandleFunc函数。
- 现在，为了**使用DefaultServeMux作为服务器的主handler**，我们不需要将它传给ListenAndServe函数；nil值就可以工作。
  - 下面代码就把`mux := http.NewServeMux()`这一步省了，除非需要多个服务器监听不同的端口，然后再构建不同的ServeMux去调用ListenAndServe。 但大部分情况只需要一个web服务器。
```go
func main() {
    db := database{"shoes": 50, "socks": 5}
    http.HandleFunc("/list", db.list)
    http.HandleFunc("/price", db.price)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
```


>go语言目前还没有一个全为的web框架，Go语言标准库中的构建模块就已经非常灵活以至于这些框架都是不必要的。此外，尽管在一个项目早期使用框架是非常方便的，但是它们带来额外的复杂度会使长期的维护更加困难。


### error接口
```go
type error interface {
    Error() string
}
```
创建一个error最简单的方法就是调用errors.New函数，它会根据传入的错误信息返回一个新的error。整个errors包仅只有4行：

```go
package errors
//每个New函数的调用都分配了一个独特的和其他错误不相同的实例。哪怕是一样的错误信息。
func New(text string) error { return &errorString{text} }

//使用结构而不是直接暴露字符串，为了保护它表示的错误避免粗心（或有意）的更新
type errorString struct { text string }

func (e *errorString) Error() string { return e.text }
```

- 用得更多的是fmt.Errorf，它还会处理字符串格式化。
```
func Errorf(format string, args ...interface{}) error {
    return errors.New(Sprintf(format, args...))
}
```


### 类型断言 x.(T)
类型断言是一个使用在接口值上的操作，形如**x.(T)**。
- x表示一个接口的类型和T表示一个类型。
- 检查它操作对象x的动态类型是否和断言的类型T匹配。
  - 如果断言的类型T是一个**具体类型**，然后类型断言检查x的动态类型是否和T相同。
    - 检查成功，断言的结果是 x的动态值，类型是T。 也就是说具体类型的类型断言从它的操作对象中获得具体的值。
    - 检查失败，抛出panic
  - 如果断言的类型T是一个**接口类型**，？？？（莫名其妙的翻译）
    - 如果检查成功了，断言结果为类型T，不过保留了接口值内部的动态类型和值的部分？
    - 如果失败了？
- 如果断言操作的对象是一个nil接口值，那么不论被断言的类型是什么这个类型断言都会失败。
```go
//第一个类型断言后，w和rw都持有os.Stdout,它们都有一个动态类型*os.File
var w io.Writer
//w只公开了Write方法
w = os.Stdout
//rw变量还公开了它的Read方法
rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
w = new(ByteCounter)
rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
```
- 几乎不需要对一个更少限制性的接口类型（更少的方法集合）做断言。
- 判断接口值的是否是某动态类型，然后根据结果去做一些操作，一般会使用返回两个结果断言： 
  - 第一个结果是表示断言得到的类型。如果失败了，就会等于被断言类型的零值
  - 第二个结果是ok bool，表示判断是否成功。


**通过类型断言识别错误类型**


举例：os包中文件操作返回的错误原因，有三种经常的错误必须进行不同的处理：文件已经存在（对于创建操作），找不到文件（对于读取操作），和权限拒绝。
- 如何对错误值表示的失败进行分类？
  - 直接判断是否包含特定子字符串是不健壮的。
  - 更可靠的是使用一个专门的类型来描述结构化的错误。 比如os包中的PathError、LinkError。
    - 调用方需要使用类型断言来检测错误的具体类型以便将一种失败和另一种区分开；具体的类型可以比字符串提供更多的细节。
```go
_, err := os.Open("/no/such/file")
fmt.Println(err) // "open /no/such/file: No such file or directory"
fmt.Printf("%#v\n", err)
// Output: &os.PathError{Op:"open", Path:"/no/such/file", Err:0x2}

//有几个特定的方法可以对错误类型进行判断
func IsExist(err error) bool
func IsNotExist(err error) bool
func IsPermission(err error) bool

//使用
os.IsNotExist(err)
```
- 区别错误通常必须在失败操作后，错误传回调用者前进行。


**通过类型断言查询接口**


- 有一个允许字符串高效写入的WriteString方法；这个方法会避免去分配一个临时的拷贝。
- 但是我们不确定某个io.Writer类型的变量是否拥有这个方法，可以定义一个只有这个方法的新接口，然后使用类型断言检测w的动态类型是否满足这个新接口。
```go
func writeString(w io.Writer, s string) (n int, err error) {
    type stringWriter interface {
        WriteString(string) (n int, err error)
    }
    if sw, ok := w.(stringWriter); ok {
        return sw.WriteString(s) // avoid a copy
    }
    return w.Write([]byte(s)) // allocate temporary copy
}

func writeHeader(w io.Writer, contentType string) error {
    if _, err := writeString(w, "Content-Type: "); err != nil {
        return err
    }
    if _, err := writeString(w, contentType); err != nil {
        return err
    }
    // ...
}
```
- 这个例子的神奇之处在于，没有定义了WriteString方法的标准接口，也没有指定它是一个所需行为的标准接口。
- 一个具体类型只会通过它的方法决定它是否满足stringWriter接口，而不是任何它和这个接口类型所表达的关系。


定义一个特定类型的方法隐式地获取了对特定行为的协约。对于Go语言的新手，特别是那些来自有强类型语言使用背景的新手，可能会发现它缺乏显式的意图令人感到混乱，但是在实战的过程中这几乎不是一个问题。除了空接口interface{}，接口类型很少意外巧合地被实现。


### 类型分支
接口有两种使用方式：
- 第一种，以io.Reader，io.Writer，fmt.Stringer，sort.Interface，http.Handler和error为典型，一个接口的方法表达了实现这个接口的具体类型间的相似性，但是隐藏了代码的细节和这些具体类型本身的操作。**重点在于方法上，而不是具体的类型上。**
- 第二种，利用一个接口值可以持有各种具体类型值的能力，将这个接口当成这些类型的联合。使用类型断言用来动态地区别这些类型。 重点在于具体的类型满足这个接口，而不在于接口的方法，且没有隐藏任何信息。我们将以这种方式使用的接口描述为discriminated unions（可辨识联合）。


- 一个类型分支像普通的switch语句一样，它的运算对象是`x.(type)`——它使用了关键词字面量type——并且每个case有一到多个类型。
- 对于bool和string情况的逻辑需要通过类型断言访问提取的值，所以对于这个断言的值需要用一个临时变量存起来，方便使用。
- 在每个单一类型的case内部，变量x和这个case的类型相同。
```go
func sqlQuote(x interface{}) string {
    switch x := x.(type) {
    case nil:
        return "NULL"
    case int, uint:
        return fmt.Sprintf("%d", x) // x has type interface{} here.
    case bool:
        if x {
            return "TRUE"
        }
        return "FALSE"
    case string:
        return sqlQuoteString(x) // (not shown)
    default:
        panic(fmt.Sprintf("unexpected type %T: %v", x, x))
    }
}
```


### 示例：基于标记的XML解码
* encoding/xml包也提供了一个更低层的基于标记的API用于XML解码。
* 在基于标记的样式中，解析器消费输入并产生一个标记流；
* 四个主要的标记类型
  * StartElement，EndElement，CharData，和Comment
  * 每一个都是encoding/xml包中的具体类型。


### 接口补充说明
- 接口只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要。遵循这条规则必定会从任意特定的实现细节中抽象出来。结果就是有更少和更简单方法的更小的接口，而小的接口更容易满足。
- 对于接口设计的一个好的标准就是 ask only for what you need（只考虑你需要的东西）



---

## 第八章 并发编程（一）基于顺序通信进程（CSP）
>本章讲解goroutine和channel，其支持“顺序通信进程”（communicating sequential processes）或被简称为CSP。CSP是一种现代的并发编程模型，在这种编程模型中值会在不同的运行实例（goroutine）中传递，尽管大多数情况下仍然是被限制在单一实例中。

- 在语法上，go语句是一个普通的函数或方法调用前加上关键字go。 `go f()`
- go语句会使其语句中的函数在一个新创建的goroutine中运行。而go语句本身会迅速地完成。
- 主函数返回时，所有的goroutine都会被直接打断，程序退出。
- 除了从主函数退出或者直接终止程序之外，没有其它的编程方法能够让一个goroutine来打断另一个的执行. 可以使用通信去让其主动退出。


### 并发示例1和2

**示例1：并发的Clock服务**

- 第一个例子是一个顺序执行的时钟服务器，它会每隔一秒钟将当前时间写到客户端：
- 但是那样客户端必须等待第一个客户端完成工作，这样服务端才能继续向后执行；因为我们这里的服务器程序同一时间只能处理一个客户端连接。
- 我们这里对服务端程序做一点小改动，使其支持并发：在handleConn函数调用的地方增加go关键字，让每一次handleConn的调用都进入一个独立的goroutine。
```go

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}
```


**示例2：并发的Echo服务**
```go
//reverb1
func echo(c net.Conn, shout string, delay time.Duration) {
    fmt.Fprintln(c, "\t", strings.ToUpper(shout))
    time.Sleep(delay)
    fmt.Fprintln(c, "\t", shout)
    time.Sleep(delay)
    fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
    input := bufio.NewScanner(c)
    for input.Scan() {
        echo(c, input.Text(), 1*time.Second)
        //如果增加go,才能实现 上一次还没说完三次，下一次也会开始说。 不然就只能顺序说，只有上一句话说完，才能开始下一句。
        go echo(c, input.Text(), 1*time.Second)
    }
    // NOTE: ignoring potential errors from input.Err()
    c.Close()
}

//netcat2
func main() {
    conn, err := net.Dial("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    go mustCopy(os.Stdout, conn)
    mustCopy(conn, os.Stdin)
}
```

### 8.4 Channels
- 如果说goroutine是Go语言程序的并发体的话，那么channels则是它们之间的**通信机制**
- 一个channel是一个通信机制，它可以让一个goroutine通过它给另一个goroutine发送值信息。


**初始化**

make函数初始化一个channel：
```go
ch := make(chan int) //可以发送int类型数据的channel
```


**比较**

- 两个相同类型的channel可以使用==运算符比较。
- 如果两个channel引用的是相同的对象，那么比较的结果为真。
- 一个channel也可以和nil进行比较。


**操作**

- 一个channel有发送和接受两个主要操作，都是通信行为。
- 发送： `ch <- x  // a send statement`
- 接收:一个不使用接收结果的接收操作也是合法的。
  - `x = <-ch`
  - `<-ch`  
- **关闭**:`close(ch)`。 关闭后
  - 基于该channel的任何发送操作都将导致panic异常。
  - 对一个已经被close过的channel进行接收操作依然可以接受到之前已经成功发送的数据；
  - 如果channel中已经没有数据的话将产生一个零值的数据。


**缓存通道**

- 以最简单方式调用make函数创建的是一个无缓存的channel，但是我们也可以指定第二个整型参数，对应channel的容量。
```go
ch = make(chan int)    // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3
```


**不带缓存的Channels**

- 一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels上执行接收操作，当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的语句。
- 基于无缓存Channels的发送和接收操作将**导致两个goroutine做一次同步操作**。因为这个原因，无缓存Channels有时候也被称为**同步Channels**。



## 第九章 并发编程（二）传统的基于共享变量
>在多goroutine之间的共享变量，并发问题的分析手段，以及解决这些问题的基本模式。


## 第十章 包机制和包的组织结构
>这一章还展示了如何有效地利用Go自带的工具，使用单个命令完成编译、测试、基准测试、代码格式化、文档以及其他诸多任务。


- Go语言标准包200多个（查看命令：go list std | wc -l）
- 目前互联网上已经发布了非常多的Go语言开源包，它们可以通过 http://godoc.org 检索


Go语言的**闪电**般的**编译速度主要得益于**三个语言特性。
1. 第一点，所有导入的包必须在每个文件的开头显式声明，这样的话编译器就没有必要读取和分析整个源文件来判断包的依赖关系。
2. 第二点，禁止包的环状依赖，因为没有循环依赖，包的依赖关系形成一个有向无环图，每个包可以被独立编译，而且很可能是被并发编译。
3. 第三点，编译后包的目标文件不仅仅记录包本身的导出信息，目标文件同时还记录了包的依赖关系。
>在编译一个包的时候，编译器只需要读取每个直接导入包的目标文件，而不需要遍历所有依赖的的文件（译注：很多都是重复的间接依赖）。


**导入路径**：为了避免冲突，所有非标准库包的导入路径建议以所在组织的互联网域名为前缀；而且这样也有利于包的检索。


**包声明**：默认的包名就是包导入路径名的最后一段，因此即使两个包的导入路径不同，它们依然可能有一个相同的包名。


关于默认包名一般采用导入路径名的最后一段的约定也有三种例外情况：
- main包：名字为main的包是给go build，构建命令一个信息，这个包编译完之后必须调用连接器生成一个可执行程序。
- _test.go结尾的文件：并且这些源文件声明的包名也是以_test为后缀名的。这种目录可以包含两种包：
  - 一种是普通包，
  - 另一种则是测试的外部扩展包。
  - 所有以_test为后缀包名的测试外部扩展包都由go test命令独立编译，普通包和测试的外部扩展包是相互独立的。
- 带版本号：例如“gopkg.in/yaml.v2”。这种情况下包的名字并不包含版本号后缀，而是yaml。


**导入声明**：
- 如果我们想同时导入两个有着名字相同的包，例如math/rand包和crypto/rand包，那么导入声明必须至少为一个同名包指定一个新的包名以避免冲突。这叫做**导入包的重命名**。
```go
import(
    "crypto/rand"
    mrand "math/rand" // alternative name mrand avoids conflict
)
```
- 如果文件中已经有了一个名为path的变量，那么我们可以将“path”标准包重命名为pathpkg。
- 每个导入声明语句都明确指定了当前包和被导入包之间的依赖关系。**如果遇到包循环导入的情况，Go语言的构建工具将报告错误。**


**包的匿名导入**
- 导入了又不用会编译报错。
  - 导入的意义是啥？ 导入包后会做一些预处理，我只需要这些处理后的效果：它会计算包级变量的初始化表达式和执行导入包的init初始化函数。
    - 例如图像处理image包有这种场景：主程序只需要匿名导入特定图像驱动包就可以用image.Decode解码对应格式的图像了。
    - 例如数据库，匿名导入相应的数据库驱动包，直接就可以用。
```go
import (
    "database/sql"
    _ "github.com/lib/pq"              // enable support for Postgres
    _ "github.com/go-sql-driver/mysql" // enable support for MySQL
)

db, err = sql.Open("postgres", dbname) // OK
db, err = sql.Open("mysql", dbname)    // OK
db, err = sql.Open("sqlite3", dbname)  // returns error: unknown driver "sqlite3"
```
  - 怎么规避报错？ 匿名导入，就是重命名为下划线 _ 。


**包的命名原则**

- 一般使用短小的，但也要易于理解无歧义。 ioutils够简洁了，就不需要命名为util。
- 避免包名使用常用作局部变量的名字。例如path
- 一般采用单数
- 设计变量名时，考虑与包名的混用。 所以不需要在变量名里包含包名的重复意思。 比如包名bos，变量名就不需要bosName
- 只暴露一个主要的数据结构和与它相关的方法，还有一个以New命名的函数用于创建实例。


### 包的工具
**GOPATH**


当需要切换到不同工作区的时候，只要更新GOPATH就可以了。 `export GOPATH=$HOME/gobook`

GOPATH对应的工作区目录有三个子目录。
- src子目录用于存储源代码
- pkg子目录用于保存编译后的包的目标文件
- bin子目录用于保存编译后的可执行程序


**GOROOT**


GOROOT用来指定Go的安装目录，还有它自带的标准库包的位置。
- 用户一般不需要设置GOROOT，默认情况下Go语言安装工具会将其设置为安装的目录路径。


**其他环境变量**
- GOOS环境变量用于指定目标操作系统（例如android、linux、darwin或windows）
- GOARCH环境变量用于指定处理器的类型，例如amd64、386或arm等。


**下载包**


- go get命令获取的代码是真实的本地存储仓库，而不仅仅只是复制源文件，因此你依然可以使用版本管理工具比较本地代码的变更或者切换到其它的版本。
  -  -u 表示下载最新版本。
- 进入文件目录，然后获取版本号，这里地址其实是有个转换的， https://golang.org/x/net/html 包含了如下的元数据，它告诉Go语言的工具当前包真实的Git仓库托管地址.
```
$ cd $GOPATH/src/golang.org/x/net
$ git remote -v
origin  https://go.googlesource.com/net (fetch)
origin  https://go.googlesource.com/net (push)

$./fetch https://golang.org/x/net/html | grep go-import
<meta name="go-import"
      content="golang.org/x/net git https://go.googlesource.com/net">
```


**编译build**


```
$ cd anywhere
$ go build gopl.io/ch1/helloworld
```
- go install命令和go build命令很相似，但是它会保存每个包的编译成果，而不是将它们都丢弃。
- 被编译的包会被保存到\$GOPATH/pkg目录下，目录路径和 src目录路径对应，可执行程序被保存到$GOPATH/bin目录。
- go install命令和go build命令都**不会重新编译没有发生变化的包**.
- go build -i命令将安装每个目标所依赖的包。


**分系统**
- 如果一个文件名包含了一个操作系统或处理器类型名字，例如net_linux.go或asm_amd64.s，Go语言的构建工具将只在对应的平台编译这些文件。
- 在包声明和包注释前面，可以增加参数告诉build只在特定系统编译或者不编译这个文件：
```
// +build linux darwin 只编译
// +build ignore   不编译

```


**包文档**
- 包中每个**导出**的成员和包声明前都应该包含目的和用法说明的注释。
- 文档注释一般是完整的句子，第一行摘要说明，以被注释者的名字开头。 
- 注释中的参数直接用定义的名字就行，不需要额外的引号或者标记注明。
- 包注释
  - 注释之后紧跟着包，这个注释就是包注释，只能有一个，多个文件同样包的注释会合并。
  - 如果包注释过长，可以单独放在一个文件里，一般叫做doc.go。
  - 文档要简洁不可忽视。 多看标准库。
- `go doc`命令，该命令打印其后所指定的实体的声明与文档注释，该实体可能是一个包、某个具体的包成员、一个方法
  - 不需要输入完整的包导入路径或正确的大小写
```
$ go doc time
package time // import "time"
。。。
$ go doc time.Since
$ go doc time.Duration.Seconds

```
- godoc的在线服务 https://godoc.org ，包含了成千上万的开源包的检索工具。
- 也可以在自己的工作区目录运行godoc服务。运行下面的命令，然后在浏览器查看 http://localhost:8000/pkg 页面：`godoc -http :8000`
  - 其中-analysis=type和-analysis=pointer命令行标志参数用于打开文档和代码中关于静态分析的结果。


**内部包**
  
作用：希望在内部子包之间共享一些通用的处理包，或者我们只是想实验一个新包的还并不稳定的接口，暂时只暴露给一些受限制的用户使用。

怎么用：Go语言的构建工具对**包含internal名字的路径段**的包导入路径做了特殊处理。这种包叫internal包，一个internal包只能被和internal目录有同一个父目录的包所导入。
- 例如，net/http/internal/chunked内部包只能被net/http/httputil或net/http包导入，但是不能被net/url包导入


**查询包**：go list命令可以查询可用包的信息。 `go list github.com/go-sql-driver/mysql`
- 用"..."表示匹配任意的包的导入路径。导出工作区所有包。
- 某个主题相关的所有包:`go list ...xml...`
- -json命令行参数表示用JSON格式打印每个包的元信息: `go list -json hash`
- 命令行参数-f则允许用户使用text/template包（§4.6）的模板语言定义输出文本的格式。 
  - `go list -f '{{join .Deps " "}}' strconv`  用join模板函数将结果链接为一行
  - 打印compress子目录下所有包的导入包列表: `go list -f '{{.ImportPath}} -> {{join .Imports " "}}' compress/...`


## 第十一章 单元测试
Go语言的工具和标准库中集成了轻量级的测试功能，避免了强大但复杂的测试框架。测试库提供了一些基本构件，必要时可以用来构建复杂的测试构件。


- 文件名是：被测文件_test.go，这个后缀的源文件在执行go build时不会被构建成包的一部分。
- 函数名：TestAdd(t *testing.T)
- 参数：
  - 功能测试：`t *testing.T`
  - 性能测试： `b *testing.B`
  - 示例测试： 不限制
- 功能
  - 功能测试（测试函数）：以Test为函数名前缀的函数，用于测试程序的一些逻辑行为是否正确；
  - 性能测试（基准测试）：以Benchmark为函数名前缀的函数，它们用于衡量一些函数的性能；go test命令会多次运行基准测试函数以计算一个平均的执行时间。
  - 示例函数：以Example为函数名前缀的函数，提供一个由编译器保证正确性的示例文档。
- `go test`命令如果没有参数指定包那么将默认采用当前目录对应的包。
- `go test`命令会遍历所有的*_test.go文件中符合上述命名规则的函数，生成一个临时的main包用于调用相应的测试函数，接着构建并运行、报告测试结果，最后清理测试中生成的临时文件。


### 测试函数（功能测试）

- 每个测试函数必须导入testing包。函数有如下签名。
```go
func TestName(t *testing.T) {
    // ...
}
```




## 第十二章 反射

一种程序在运行期间审视自己的能力。反射是一个强大的编程工具，不过要谨慎地使用；这一章利用反射机制实现一些重要的Go语言库函数，展示了反射的强大用法。



## 第十三章 底层编程的细节

在必要时，可以使用unsafe包绕过Go语言安全的类型系统。