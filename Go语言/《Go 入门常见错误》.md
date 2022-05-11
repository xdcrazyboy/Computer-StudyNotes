
# 基础类型

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




# 多线程