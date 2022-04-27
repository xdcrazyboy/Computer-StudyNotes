
# 开始学习Go

## go的独特之处
**它的主要目标是“兼具Python 等动态语言的开发速度和C/C++等编译型语言的性能与安全性”**

### 为并发而生
Go语言的并发是基于 goroutine 的，goroutine 类似于线程，但并非线程。可以将 goroutine 理解为一种虚拟线程。Go 语言运行时会参与调度 goroutine，并将 goroutine 合理地分配到每个 CPU 中，最大限度地使用CPU性能。开启一个goroutine的消耗非常小（大约2KB的内存），你可以轻松创建数百万个goroutine。

goroutine的特点：

    1.`goroutine`具有可增长的分段堆栈。这意味着它们只在需要时才会使用更多内存。
    2.`goroutine`的启动时间比线程快。
    3.`goroutine`原生支持利用channel安全地进行通信。
    4.`goroutine`共享数据结构时无需使用互斥锁。


### 优势
**简单易学**：

- 语法简洁
  - 其语法在C语言的基础上进行了大幅的简化，去掉了不需要的表达式括号，循环也只有 for 一种表示方法，就可以实现数值、键值等各种遍历。
- 代码风格统一
- 开发效率高
  - Go语言实现了开发效率与执行效率的完美结合，让你像写Python代码（效率）一样编写C代码（性能）。


**自带GC**


**静态编译**


**简单的思想，没有继承，多态，类等**

### go适合做什么
* 服务端开发
* 分布式系统，微服务
* 网络编程
* 区块链开发
* 内存KV数据库，例如boltDB、levelDB
* 云平台


广泛用于人工智能、云计算开发、容器虚拟化、⼤数据开发、数据分析及科学计算、运维开发、爬虫开发、游戏开发等领域。


### 主要特征
1.自动立即回收。
2.更丰富的内置类型。
3.函数多返回值。
4.错误处理。
5.匿名函数和闭包。
6.类型和接口。
7.并发编程。
8.反射。
9.语言交互性。


### 一些go独有的特点

1. 函数可以返回任意数量的返回值
```go
package main

import "fmt"

func swap(x, y string) (string, string, string) {
	return y, x, x + y
}

func main() {
	a, b, c := swap("hello", "world")
	fmt.Println(a, b, c)
}
```

2. 短变量声明：在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。

函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用

```go
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```

3. 可见性
1）声明在函数内部，是函数的本地值，类似private
2）声明在函数外部，是对当前包可见(包内所有.go文件都可见)的全局值，类似protect
3）声明在函数外部且首字母大写是所有包可见的全局值,类似public


## 经常疑惑的点

### GOPATH

- 下载的第三方包源代码文件放在$GOPATH/src目录下， 
- 产生的二进制可执行文件放在 $GOPATH/bin目录下，
- 生成的中间缓存文件会被保存在 $GOPATH/pkg 下

### 包 package
- go 里面一个目录为一个package, 一个package级别的func, type, 变量, 常量, 这个package下的所有文件里的代码都可以随意访问, 不需要首字母大写
- **同目录**下的两个文件如hello.go和hello2.go中的package 定义的名字要是同一个，不同的话，是会报错的 ==> 所以main方法要单独放一个文件

### init函数和main函数
**init函数**： 用于包(package)的初始化
    1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等

    2 每个包可以拥有多个init函数

    3 包的每个源文件也可以拥有多个init函数

    4 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)

    5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序

    6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用


**main函数**：Go语言程序的默认入口函数(主函数)
func main(){ 
	//body
}


**异同**：
- 相同点：
    - 两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用。
- 不同点：
    - init可以应用于任意包中，且可以重复定义多个。
    - main函数只能用于main包中，且只能定义一个。


**执行顺序**：
* 对同一个go文件的init()调用顺序是从上到下的。

* 对同一个package中不同文件是按文件名字符串比较“从小到大”顺序调用各文件中的init()函数。

* 对于不同的package，如果不相互依赖的话，按照main包中"先import的后调用"的顺序调用其包中的init()，如果package存在依赖，则先调用最早被依赖的package中的init()，最后调用main函数。

* 如果init函数中使用了println()或者print()你会发现在执行过程中这两个不会按照你想象中的顺序执行。这两个函数官方只推荐在测试环境中使用，对于正式环境不要使用


# 《build-web-application-with-golang》
>https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/01.3.md

## Go命令

```go
The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        work        workspace maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildconstraint build constraints
        buildmode       build modes
        c               calling between Go and C
        cache           build and test caching
        environment     environment variables
        filetype        file types
        go.mod          the go.mod file
        gopath          GOPATH environment variable
        gopath-get      legacy GOPATH go get
        goproxy         module proxy protocol
        importpath      import path syntax
        modules         modules, module versions, and more
        module-get      module-aware go get
        module-auth     module authentication using go.sum
        packages        package lists and patterns
        private         configuration for downloading non-public code
        testflag        testing flags
        testfunc        testing functions
        vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.
```

### go build
目的：用于编译代码。在包的编译过程中，若有必要，会同时编译与之相关联的包。

**命令介绍**

* 如果是普通包，就像我们在1.2节中编写的mymath包那样，当你执行go build之后，它不会产生任何文件。如果你需要在$GOPATH/pkg下生成相应的文件，那就得执行go install。

* 如果是main包，当你执行go build之后，它就会在当前目录下生成一个可执行文件。如果你需要在$GOPATH/bin下生成相应的文件，需要执行go install，或者使用go build -o 路径/a.exe。

* 如果某个项目文件夹下有多个文件，而你只想编译某个文件，就可在go build之后加上文件名，例如go build a.go；go build命令默认会编译当前目录下的所有go文件。

* 你也可以指定编译输出的文件名。例如1.2节中的mathapp应用，我们可以指定go build -o astaxie.exe，
  * 默认情况是你的package名(非main包)，或者是第一个源文件的文件名(main包)。
  * 默认生成的可执行文件名是文件夹名。）
* go build会忽略目录下以“_”或“.”开头的go文件。
* go build的时候会选择性地编译以系统名结尾的文件（Linux、Darwin、Windows、Freebsd）。例如Linux系统下面编译只会选择array_linux.go文件,其它系统命名后缀文件全部忽略，例如array_windows.go。


**参数的介绍**
-o 指定输出的文件名，可以带上路径，例如 go build -o a/b/c
-i 安装相应的包，编译+go install
-a 更新全部已经是最新的包的，但是对标准包不适用
-n 把需要执行的编译命令打印出来，**但是不执行**，这样就可以很容易的知道底层是如何运行的
-p n 指定可以并行可运行的编译数目，默认是CPU数目
-race 开启编译的时候自动检测数据竞争的情况，目前只支持64位的机器
-v 打印出来我们正在编译的包名
-work 打印出来编译时候的临时文件夹名称，并且如果已经存在的话就不要删除
-x 打印出来执行的命令，其实就是和-n的结果类似，只是这个会执行
-ccflags 'arg list' 传递参数给5c, 6c, 8c 调用
-compiler name 指定相应的编译器，gccgo还是gc
-gccgoflags 'arg list' 传递参数给gccgo编译连接调用
-gcflags 'arg list' 传递参数给5g, 6g, 8g 调用
-installsuffix suffix 为了和默认的安装包区别开来，采用这个前缀来重新安装那些依赖的包，-race的时候默认已经是-installsuffix race,大家可以通过-n命令来验证
-ldflags 'flag list' 传递参数给5l, 6l, 8l 调用
-tags 'tag list' 设置在编译的时候可以适配的那些tag

### go clean
目的：**移除当前源码包和关联源码包里面编译生成的文件**


这些文件包括：
```
_obj/            旧的object目录，由Makefiles遗留
_test/           旧的test目录，由Makefiles遗留
_testmain.go     旧的gotest文件，由Makefiles遗留
test.out         旧的test记录，由Makefiles遗留
build.out        旧的test记录，由Makefiles遗留
*.[568ao]        object文件，由Makefiles遗留

DIR(.exe)        由go build产生
DIR.test(.exe)   由go test -c产生
MAINFILE(.exe)   由go build MAINFILE.
```


常用于清除编译文件，然后GitHub递交源码，在本机测试的时候这些编译文件都是和系统相关的，但是对于源码管理来说没必要。


`go clean -i -n`

**参数介绍**

-i 清除关联的安装的包和可运行文件，也就是通过go install安装的文件
-n 把需要执行的清除命令打印出来，但是不执行，这样就可以很容易的知道底层是如何运行的
-r 循环的清除在import中引入的包
-x 打印出来执行的详细命令，其实就是-n打印的执行版本


### go fmt <文件名>.go
- **格式化代码文档**，不符合标准格式编译不通过。 
- 一般编译器都支持保存时自动格式化，背后就是调用go fmt，而go fmt就是调用的gofmt
- 需要参数-w，否则格式化结果不会写入文件
- gofmt -w -l src，可以格式化整个项目。


gofmt的**参数介绍**

-l 显示那些需要格式化的文件
-w 把改写后的内容直接写入到文件中，而不是作为结果打印到标准输出。
-r 添加形如“a[b:len(a)] -> a[b:]”的重写规则，方便我们做批量替换
-s 简化文件中的代码
-d 显示格式化前后的diff而不是写入文件，默认是false
-e 打印所有的语法错误到标准输出。如果不使用此标记，则只会打印不同行的前10个错误。
-cpuprofile 支持调试模式，写入相应的cpufile到指定的文件


### go get
目的：**动态获取远程代码包**


- 目前支持的有BitBucket、GitHub、Google Code和Launchpad，需要安装对应的源码管理工具，且加入环境变量，比如安装了Git。
- 这个命令在内部实际上分成了两步操作：
  - 第一步是下载源码包
  - 第二步是执行go install
- 其实go get支持自定义域名的功能，具体参见go help remote。


**参数介绍**：

-d 只下载不安装
-f 只有在你包含了-u参数的时候才有效，不让-u去验证import中的每一个都已经获取了，这对于本地fork的包特别有用
-fix 在获取源码之后先运行fix，然后再去做其他的事情
-t 同时也下载需要为运行测试所需要的包
-u 强制使用网络去更新包和它的依赖包
-v 显示执行的命令


### go install
实际执行了2步：
- 第一步是生成结果文件(可执行文件或者.a包)，一般就是build，所以支持build的编译参数。
- 第二步会把编译好的结果移到$GOPATH/pkg或者$GOPATH/bin。


### go test
目的：自动读取源码目录下面名为`*_test.go`的文件，生成并运行测试用的可执行文件. 默认不需要任何参数。


常用的**参数**：
-bench regexp 执行相应的benchmarks，例如 -bench=.
-cover 开启测试覆盖率
-run regexp 只运行regexp匹配的函数，例如 -run=Array 那么就执行包含有Array开头的函数
-v 显示测试的详细命令


### go tool
go tool下面下载聚集了很多命令，常用fix和vet

- go tool fix . 用来修复以前老版本的代码到新版本，例如go1之前老版本的代码转化到go1,例如API的变化
- go tool vet directory|files 用来分析当前目录的代码是否都是正确的代码
  >例如是不是调用fmt.Printf里面的参数不正确，例如函数里面提前return了然后出现了无用代码之类的。


### go generate
用于在编译前自动化生成某类代码。

- go generate和go build是完全不一样的命令，通过分析源码中特殊的注释，然后执行相应的命令。
- 这些命令都是很明确的，**没有任何的依赖**在里面。
- 而且大家在用这个之前心里面一定要有一个理念，**这个go generate是给你用的，不是给使用你这个包的人用的，是方便你来生成一些代码的**。


举个例子：
我们经常会使用`yacc`来生成代码，那么我们常用这样的命令：

>go tool yacc -o gopher.go -p parser gopher.y

-o 指定了输出的文件名， -p指定了package的名称。这是一个单独的命令，如果我们想让go generate来触发这个命令，那么就可以在**当前目录的任意一个xxx.go文件**里面的**任意位置**增加一行如下的注释：

>//go:generate go tool yacc -o gopher.go -p parser gopher.y

这里我们注意了，`//go:generate`是**没有任何空格**的，固定格式，在扫描源码文件的时候就是根据这个来判断的。

所以我们可以通过如下的命令来生成，编译，测试。如果gopher.y文件有修改，那么就重新执行go generate重新生成文件就好。
```
$ go generate
$ go build
$ go test
```

### godoc
目的：打印附于Go语言程序实体上的文档
- go doc已经废弃，现在是godoc。 

- go get golang.org/x/tools/cmd/godoc


通过命令在命令行执行 godoc -http=:端口号 比如godoc -http=:8080。然后在浏览器中打开127.0.0.1:8080，你将会看到一个golang.org的本地copy版本，通过它你可以查询pkg文档等其它内容。如果你设置了GOPATH，在pkg分类下，不但会列出标准包的文档，还会列出你本地GOPATH中所有项目的相关文档，这对于经常被墙的用户来说是一个不错的选择。



### 其他命令
go version 查看go当前的版本
go env 查看当前go的环境变量
go list 列出当前全部安装的package
go run 编译并运行Go程序
go fix 把Go语言源码文件中的旧版本代码修正为新版本的代码
go vet 是一个用于检查Go语言源码中静态错误的简单工具


# Go语言基础语法

##  hello world

### 关键字 25个
break    default      **func**    interface    **select**
case     **defer**        go      **map**          **struct**
**chan**     else         **goto**    package      switch
const    fallthrough  if      **range**        type
continue for          import  return       var

### 保留字 37个
- Constants: true  false  iota  nil

- Types:
```
int  int8  int16  int32  int64  
uint  uint8  uint16  uint32  uint64  uintptr
float32  float64  complex128  complex64
bool  byte  rune  string  error
```

- Functions:   
```
make  len  cap  new  append  copy  close  delete
complex  real  imag
panic  recover
```

### 写个hello world
- 每一个可独立运行的Go程序，必定包含一个package main，在这个main包中必定包含一个入口函数main，而这个函数既没有参数，也没有返回值。
- Go使用package（和Python的模块类似）来组织代码。
- main.main()函数(这个函数位于主包）是每一个独立的可运行程序的入口点。
- Go使用UTF-8字符串和标识符(因为UTF-8的发明者也就是Go的发明者之一)，所以它天生支持多语言。


## 基础

### 变量
>**Go把变量类型放在变量名后面**

- 使用**var**关键字是Go最基本的定义变量方式: `var variableName type`
  - 定义多个变量: `var vname1, vname2, vname3 type`
  - 定义变量并初始化值: `var variableName type = value`
  - 同时初始化多个变量: `var vname1, vname2, vname3 type= v1, v2, v3`
- 简化定义——忽略类型：`var vname1, vname2, vname3 = v1, v2, v3`
- 再简化：`vname1, vname2, vname3 := v1, v2, v3`  
  - `:=这`个符号直接取代了var和type,这种形式叫做简短声明.
  - 只能用在函数内部；
  - 在函数外部使用则会无法编译通过，所以一般用**var**方式来定义全局变量。


- `_`（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。 **有什么用呢？后面解释**。 例如：`_, b := 34, 35`，就是丢弃34.
- Go对于已声明但未使用的变量会在编译阶段报错，比如下面的代码就会产生一个错误：声明了i但未使用。


### 常量
 **定义语法**：
```go
const constantName = value
//如果需要，也可以明确指定常量的类型：
const Pi float32 = 3.1415926 
```


**常见例子**：
```go
const Pi = 3.1415926
const i = 10000
const MaxThread = 10
const prefix = "astaxie_"
```


>**特别之处**: 可以指定相当多的小数位数(例如200位)， 若指定給float32自动缩短为32bit，指定给float64自动缩短为64bit

### 内置基础类型

#### Boolean

布尔值的类型为`bool`，值是true或false，**默认**为`false`。
```go
//示例代码
var isActive bool  // 全局变量声明
var enabled, disabled = true, false  // 忽略类型的声明
func test() {
	var available bool  // 一般声明
	valid := false      // 简短声明
	available = true    // 赋值操作
}
```


#### 数值类型

- 整数分为有符号和无符号两种：
  - 有符号：int，int8, int16, int32, int64，rune（=int32）
  - 无符号：uint, uint8, uint16, uint32, uint64, byte（=uint8）
  - 
>这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错。 **没法强转嘛？**

```go
var a int8

var b int32

c:=a + b
```

- 浮点数的类型有float32和float64两种（没有float类型），默认是float64。
- 复数： complex64，complex128（64位实数+64位虚数），默认类型128。
```go
// 复数的形式为RE + IMi，其中RE是实数部分，IM是虚数部分，而最后的i是虚数单位
var c complex64 = 5+5i
//output: (5+5i)
fmt.Printf("Value is: %v", c)
```


#### 字符串
字符串是用一对双引号（`""`）或反引号（``）括起来定义，它的类型是string。

- 字符串是不可变的，下面代码会报错：cannot assign to s[0]
```go
var s string = "hello"
s[0] = 'c'
```
- 真想修改可以迂回：
```go
s := "hello"
c := []byte(s)  // 将字符串 s 转换为 []byte 类型
c[0] = 'c'
s2 := string(c)  // 再转换回 string 类型
fmt.Printf("%s\n", s2)
```
- 字符串可以做切片，用作修改
```go
s := "hello"
s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
fmt.Printf("%s\n", s)
```
- ` 括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。下面变量输出的话也是分行的。
```go
m := `hello
	world`
```


#### array  定长数组
定义：
```go
var arr [n]type
```

读取和赋值
```go
var arr [10]int  // 声明了一个int类型的数组
arr[0] = 42      // 数组下标是从0开始的
arr[1] = 13      // 赋值操作
fmt.Printf("The first element is %d\n", arr[0])  // 获取数据，返回42
fmt.Printf("The last element is %d\n", arr[9]) //返回未赋值的最后一个元素，默认返回0
```


- 长度是数组类型的一部分，长度不同那就是不一样的类型。 长度不可改变。
- **数组之间的赋值**是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。
  - 是否可以把一个数组直接赋值给另一个数组？ 长度不一样可以不？ 一样可以不？


#### slice  动态数组


#### map



#### 错误类型
Go内置有一个error类型，专门用来处理错误信息，Go的package里面还专门有一个包errors来处理错误：
```go
err := errors.New("emit macho dwarf: elf header corrupted")
if err != nil {
	fmt.Print(err)
}
```

#### 内置接口error
```go
    type error interface { //只要实现了Error()函数，返回值为String的都实现了err接口

            Error()    String

    }

```

#### 分组声明 用括号（）
同时声明多个常量、变量，或者导入多个包时，可采用分组的方式进行声明
```go
import "fmt"
import "os"

const i = 100
const pi = 3.1415
const prefix = "Go_"

var i int
var pi float32
var prefix string
```
分组规整为：
```go
import(
	"fmt"
	"os"
)

const(
	i = 100
	pi = 3.1415
	prefix = "Go_"
)

var(
	i int
	pi float32
	prefix string
)
```

#### iota枚举
这个关键字用来声明enum的时候采用，它默认开始值是0，const中每增加一行加1


目前感觉没啥卵用，先不介绍了。

### 内置函数
Go 语言拥有一些**不需要进行导入**操作就可以使用的内置函数。
- 它们有时可以针对不同的类型进行操作，例如：len、cap 和 append
- 或必须用于系统级的操作，例如：panic。
- 因此，它们需要直接获得编译器的支持。

```
append          -- 用来追加元素到数组、slice中,返回修改后的数组、slice
close           -- 主要用来关闭channel
delete            -- 从map中删除key对应的value
panic            -- 停止常规的goroutine  （panic和recover：用来做错误处理）
recover         -- 允许程序定义goroutine的panic动作
real            -- 返回complex的实部   （complex、real imag：用于创建和操作复数）
imag            -- 返回complex的虚部
make            -- 用来分配内存，返回Type本身(只能应用于slice, map, channel)
new                -- 用来分配内存，主要用来分配值类型，比如int、struct。返回指向Type的指针
cap                -- capacity是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
copy            -- 用于复制和连接slice，返回复制的数目
len                -- 来求长度，比如string、array、slice、map、channel ，返回长度
print、println     -- 底层打印函数，在部署环境中建议使用 fmt 包
```

### 运算符

**算数运算符**
 + 、 - 、 * 、 / 、%  没啥区别

- 注意： ++（自增）和--（自减）在Go语言中是单独的语句，并不是运算符。



**关系运算符**
 == 、 != 、 > 、>= 、 < 、 <=  ,也没啥区别


**逻辑运算符**
&&、 ll 、 ! ,好像也没啥区别，会阻断，没有单个的？单个的是位运算符


**位运算符**
位运算符对整数在内存中的二进制位进行操作， 很少用
运算符 | 描述 
---------|----------
& |	参与运算的两数各对应的二进位相与。（两位均为1才为1）
l | 参与运算的两数各对应的二进位相或。（两位有一个为1就为1）
^ | 参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。（两位不一样则为1）
<< | 左移n位就是乘以2的n次方。“a<<b”是把a的各二进位全部左移b位，高位丢弃，低位补0。
\>>	| 右移n位就是除以2的n次方。“a>>b”是把a的各二进位全部右移b位。


**赋值运算符**
没啥区别：+= 、&=

### 


# Go并发
**并发**在微观层面，任务不是同时运行。**并行**是多个任务同时运行。

**线程**也叫轻量级进程，通常一个进程包含若干个线程。
- 线程可以利用进程所拥有的资源。
- 在引入线程的操作系统中，通常都是把**进程作为分配资源的基本单位**，而把**线程作为独立运行和独立调度的基本单位**
	>比如音乐进程，可以一边查看排行榜一边听音乐，互不影响。

## Goroutine介绍


# 网络编程


# 项目实战

## 适合企业开发者的目录结构
![](../img/go入门/go入门_2022-04-26-19-17.png)

```

```