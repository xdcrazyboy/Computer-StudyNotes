

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
- 有许多方式可以避免出现类似潜在的问题。最直接的方法是通过单独声明err变量，来**避免使用:=的简短声明方式**：
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
 位操作运算符`&^`用于按位置零（AND NOT）：如果对应y中bit位为1的话，表达式z = x &^ y结果z的对应的bit位为0，否则z对应的bit位等于x相应的bit位的值。

## 第四章 复合类型：数组和结构体

从简单的数组、字典、切片到动态列表

### 定长数组 array


### 可变数组 slice
- 是否可称为引用类型？ 有说可以叫指针结构的包装，比叫引用类型更严谨。
- 用s[i]访问单个元素，用s[m:n]获取子序列。Go言里也采用左闭右开形式，0 ≤ m ≤ n ≤ len(s)，包含n-m个元素。

#### 难点一： 长度len 和 容量cap
**一个切片的容量总是固定的。**


例子：
```go
s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
s4 := s3[3:6]
```

s3的长度和容量都是 8
s4的长度/大小是3，**s4的容量是多少？**
- 切片的容量代表了它的底层数组的长度，但这仅限于使用make函数或者切片值字面量初始化切片的情况。
- 更通用的规则是: 一个切片的容量可以被看作是透过这个窗口最多可以看到的底层数组中元素的个数。
  - 而在底层数组不变的情况下，切片代表的**窗口可以向右扩展，直至其底层数组的末尾**。
    - 这里底层数组是最底层，哪怕slicea 从arr而来，sliceb从slicea而来。


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




### map[keyType]valueType
- map的key，可以是int，可以是string及所有完全定义了==与!=操作的类型
- 值则可以是任意类型



## 第五章 函数、错误处理、panic、recover、有defer语句。
>引用类型包括指针（§2.3.2）、切片（§4.2)）、字典（§4.3）、函数（§5）、通道（§8）,它们都是对程序中一个变量或状态的间接引用。这意味着对任一引用类型数据的修改都会影响所有该引用的拷贝。


# Go语言特性
>接口、并发、包、测试和反射等语言特性。

Go语言的面向对象机制与一般语言不同。
- 它没有类层次结构，甚至可以说没有类；
- 仅仅通过组合（而不是继承）简单的对象来构建复杂的对象。
- 方法不仅可以定义在结构体上，而且，可以定义在任何用户自定义的类型上；
- 并且，具体类型和抽象类型（接口）之间的关系是隐式的，

>所以很多类型的设计者可能并不知道该类型到底实现了哪些接口。

## 第六章 方法



## 第七章 接口


## 第八章 并发编程（一）基于顺序通信进程（CSP）

使用goroutines和channels处理并发编程

## 第九章 并发编程（二）传统的基于共享变量


## 第十章 包机制和包的组织结构


这一章还展示了如何有效地利用Go自带的工具，使用单个命令完成编译、测试、基准测试、代码格式化、文档以及其他诸多任务。

## 第十一章 单元测试
Go语言的工具和标准库中集成了轻量级的测试功能，避免了强大但复杂的测试框架。测试库提供了一些基本构件，必要时可以用来构建复杂的测试构件。


## 第十二章 反射

一种程序在运行期间审视自己的能力。反射是一个强大的编程工具，不过要谨慎地使用；这一章利用反射机制实现一些重要的Go语言库函数，展示了反射的强大用法

## 第十三章 底层编程的细节

在必要时，可以使用unsafe包绕过Go语言安全的类型系统。