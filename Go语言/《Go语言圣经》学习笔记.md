

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


## 第六章 方法




## 第七章 接口


### 接口类型转换




## 第八章 并发编程（一）基于顺序通信进程（CSP）
使用goroutines和channels处理并发编程



## 第九章 并发编程（二）传统的基于共享变量



## 第十章 包机制和包的组织结构
>这一章还展示了如何有效地利用Go自带的工具，使用单个命令完成编译、测试、基准测试、代码格式化、文档以及其他诸多任务。


- Go语言标准包200多个（查看命令：go list std | wc -l）
- 目前互联网上已经发布了非常多的Go语言开源包，它们可以通过 http://godoc.org 检索


Go语言的**闪电**般的**编译速度主要得益于**三个语言特性。
1. 第一点，所有导入的包必须在每个文件的开头显式声明，这样的话编译器就没有必要读取和分析整个源文件来判断包的依赖关系。
2. 第二点，禁止包的环状依赖，因为没有循环依赖，包的依赖关系形成一个有向无环图，每个包可以被独立编译，而且很可能是被并发编译。
3. 第三点，编译后包的目标文件不仅仅记录包本身的导出信息，目标文件同时还记录了包的依赖关系。
>在编译一个包的时候，编译器只需要读取每个直接导入包的目标文件，而不需要遍历所有依赖的的文件（译注：很多都是重复的间接依赖）。



## 第十一章 单元测试
Go语言的工具和标准库中集成了轻量级的测试功能，避免了强大但复杂的测试框架。测试库提供了一些基本构件，必要时可以用来构建复杂的测试构件。



## 第十二章 反射

一种程序在运行期间审视自己的能力。反射是一个强大的编程工具，不过要谨慎地使用；这一章利用反射机制实现一些重要的Go语言库函数，展示了反射的强大用法。



## 第十三章 底层编程的细节

在必要时，可以使用unsafe包绕过Go语言安全的类型系统。