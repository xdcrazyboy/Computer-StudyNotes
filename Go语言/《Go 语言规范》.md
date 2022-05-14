
# 如何写出优雅的Go代码

## Go语言规范

- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [《Effective Go》](https://go.dev/doc/effective_go)


### 腾讯规约

#### 不符合个人习惯的点
**特别不符合**


- xxx



**一般不符合**


- 作为**输入参数**或者数组下标时，运算符和运算数之间不需要空格，紧凑展示。

#### import规范
* 对包进行分组管理，通过空行隔开，默认分为本地包（标准库、内部包）、第三方包。
* 标准包永远位于最上面的第一组。
* 内部包是指不能被外部 import 的包，如 GoPath 模式下的包名或者非域名开头的当前项目的 GoModules 包名。
* **带域名**的包名都属于第三方包，如 git.code.oa.com/xxx/xxx，github.com/xxx/xxx，**不用区分是否是当前项目内部**的包。
* goimports 默认最少分成本地包和第三方包两大类，这两类包必须分开不能放在一起。本地包或者第三方包内部可以继续按实际情况细分不同子类。

- 不要使用相对路径引入包
- 匿名包的引用建议使用一个新的分组引入，并在匿名包上写上注释说明。

#### error处理
- error 必须是最后一个参数
- 错误描述不需要标点结尾。
- 判断错误不要if-else，直接用卫语句一个if搞定。 而且需要单独处理，不与其他变量组合逻辑判断。
```go
// 不要采用这种方式：
x, y, err := f()
if err != nil || y == nil {
    return err   // 当y与err都为空时，函数的调用者会出现错误的调用逻辑
}
```
* 【推荐】对于不需要格式化的错误，生成方式为：errors.New("xxxx")。
* 【推荐】建议go1.13 以上，error 生成方式为：fmt.Errorf("module xxx: %w", err)。

## 腾讯 Protobuf 代码规范
Protobuf作为服务数据接口的重要组成部分，某种程度上对其正确性和稳定性的**要求比对功能代码本身还高**。


## Go Mod 开发&发布公约
