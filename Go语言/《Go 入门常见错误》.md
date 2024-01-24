[TOC]
# 基础

## 基础类型

### 切片

**长度和容量的区别**
- 切片的长度就是它所包含的元素个数。
- 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。


**用make创建切片**
- 第二个参数指定长度；指定它的容量，需向 make 传入第三个参数
```go
a := make([]int, 5)  // len(a)=5
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```





#### range遍历
- 只有一个值时，**默认取的是index而不是value**。 建议使用` _, v := range arr`
```go
for i := range arrs {
    fmt.Printf("i:%d\n", i)
    v += i
}
```

## 声明和定义

### make 和 new的区别
- **相同点**：都在堆上分配内存，但它们行为不同，适用于不同类型。

**不同点**：
- new()
  - 为值类型分配内存
  - new(T)为每个新的类型T分配一片内存，初始化为0并且返回类型为*T的内存地址 • 适用于值类型如数组、结构体
- make()
  * 为引用类型分配内存并初始化，返回的是类型本身，因为其就是引用类型
  * make(T)返回一个类型为T的初始值
  * 只适用于3种内建的引用类型:切片、map 、 channel



### 常见的坑
- 不能用简短声明方式来单独为一个变量重复声明， := 左侧至少有一个新变量，才允许多变量的重复声明：


**不小心覆盖了变量**
- 对从动态语言转过来的开发者来说，简短声明很好用，这可能会让人误会 := 是一个赋值操作符。
- 如果你在新的代码块中像下边这样误用了 :=，编译不会报错，但是变量不会按你的预期工作：
  
```go
func main() {
	x := 1
	println(x)		// 1
	{
		println(x)	// 1
		x := 2
		println(x)	// 2	// 新的 x 变量的作用域只在代码块内部
	}
	println(x)		// 1
}
```

- 复制代码这是 Go 开发者常犯的错，而且不易被发现。
- 可使用 vet 工具来诊断这种变量覆盖，Go 默认不做覆盖检查，添加 -shadow 选项来启用：
> go tool vet -shadow main.go
>  main.go:9: declaration of "x" shadows declaration at main.go:5
- 注意 vet 不会报告全部被覆盖的变量，可以使用 go-nyet 来做进一步的检测：
> $GOPATH/bin/go-nyet main.go
> main.go:10:3:Shadowing variable `x`




## 打印语句Printf的格式转换
- 不同类型打印值
  - `%s` 字符串
  - `%d`  十进制整数
  - `%x, %o, %b`  十六进制，八进制，二进制整数。
  - `%f, %g, %e`  浮点数： 3.141593 3.141592653589793 3.141593e+00
  * `%t`          布尔：true或false
  * `%c`          字符（rune） (Unicode码点)
  * `%s`          字符串
* 常用特别
  * `%v`          变量的自然形式（natural format）
  * `%T`          变量的类型
  * `#`
  * 带`+`         输出字段名
* 其他不常用
  * `%q`          带双引号的字符串"abc"或带单引号的字符'c'
  * `%%`          字面上的百分号标志（无操作数）
  * `%*s`         其中的*会在字符串之前填充**一些空格**。`fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)`


**其他说明**
- 后缀f指format，ln指line。 
  - 以字母`f`结尾的格式化函数: 如`log.Printf`和`fmt.Errorf`，都采用fmt.Printf的格式化准则。
  - 以`ln`结尾的格式化函数: 则遵循Println的方式，以跟`%v`差不多的方式格式化参数，并在最后添加一个换行符




# 函数、方法

## 错误处理
error类型是一个接口类型，它的定义为: `type error interface { Error() string }`


Ø 定义错误
• errors.New函数接收错误信息创建
    err := errors.New(“error message”) • fmt 创建错误对象
    fmt.Error(“error message”)

Øpanic:用于主动抛出错误,在调试程序时， 通过panic 来打印堆栈，方便定位错误
Ø recover:用来捕获 panic 抛出的错误， 阻 止panic继续向上传递

### defer
- 函数中使用defer，如果在defer声明之前抛出异常， defer不会执行，因为defer都没有压入到栈空间中;
- 如果在defer声明之后抛出异常， defer会被执行. 
- 如果有多个defer按照栈的规则：**先入后出**


# 接口、结构体

## 接口


### 接口类型转换
go 的在 interface 类型转换的时候， 不是使用类型的转换， 而是使用
```
t,ok := i.(T)
//具体例子
func getName(params ...interface{}) string {
    var stringSlice []string
    for _, param := range params {
        stringSlice = append(stringSlice, param.(string))
    }   
    return strings.Join(stringSlice, "_")
}
  
func main() {
    fmt.Println(getName("redis", "slave", "master"))
}
```



# 多线程