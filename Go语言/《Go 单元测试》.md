
[TOC]

# 标准库 testing

## 测试规范

### 约定俗成
- Go 语言推荐测试文件和源代码文件放在一块，测试文件以 `_test.go` 结尾。
- 测试用例名称一般命名为 Test 加上待测试的方法名。
- 测试用的参数有且只有一个，在这里是 t *`testing.T`。
- 基准测试(benchmark)的参数是 `*testing.B`，TestMain 的参数是 `*testing.M` 类型。
- 运行 `go test`，该 package 下所有的测试用例都会被执行。
  - `-v` 参数会显示每个用例的测试结果，另外 `-cover `参数可以查看覆盖率。
- 如果只想运行**其中的一个用例**，例如 TestAdd，可以用 `-run` 参数指定，该参数支持通配符 `*`，和部分正则表达式，例如 `^`、`$`
  - **go test -run TestAdd -v**



## 子测试(Subtests)
子测试是 Go 语言内置支持的，可以在某个测试用例中，根据测试场景使用 t.Run创建不同的子测试用例：
```go
// calc_test.go

func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Fatal("fail")
		}

	})
	t.Run("neg", func(t *testing.T) {
		if Mul(2, -3) != -6 {
			t.Fatal("fail")
		}
	})
}
```