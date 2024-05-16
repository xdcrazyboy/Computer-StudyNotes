
# 进阶入门

## go程序是如何跑起来的？

## 一些go独有的特点

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
3）声明在函数外部且首字母大写是所有包可见的全局值,类似public.
>如果类型/接口/方法/函数/字段的首字母大写，则是 Public 的，对其他 package 可见，如果首字母小写，则是 Private 的，对其他 package 不可见。



# 类型系统



# 高级结构




# 并发


