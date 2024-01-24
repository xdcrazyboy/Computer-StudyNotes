
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
- `t.Error/t.Errorf` 遇错不停，还会继续执行其他的测试用例
- `t.Fatal/t.Fatalf` 遇错即停。


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


对于子测试场景，更推荐（tabl-driven tests）的写法：
- 所有用例的数据组织在切片 cases 中，看起来就像一张表，借助循环创建子测试。
- 好处
  - 新增用例非常简单，只需给 cases 新增一条测试数据即可。
  - 测试代码可读性好，直观地能够看到每个子测试的参数和期待的返回值。
  - 用例失败时，报错信息的格式比较统一，测试报告易于阅读。
```go
//  calc_test.go
func TestMul(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}
```

## 帮助函数（helpers）
对一些重复的逻辑，抽取出来作为公共的帮助函数(helpers)，可以增加测试代码的可读性和可维护性。 相当于抽出一些公共私有方法。


例如把上面的子测试的逻辑抽出来：
```go
// calc_test.go
package main

import "testing"

type calcCase struct{ A, B, Expected int }

func createMulTestCase(t *testing.T, c *calcCase) {
	t.Helper()
	if ans := Mul(c.A, c.B); ans != c.Expected {
		t.Fatalf("%d * %d expected %d, but %d got",
			c.A, c.B, c.Expected, ans)
	}

}

func TestMul(t *testing.T) {
	createMulTestCase(t, &calcCase{2, 3, 6})
	createMulTestCase(t, &calcCase{2, -3, -6})
	createMulTestCase(t, &calcCase{2, 0, 1}) // wrong case
}
```
报错：
```
--- FAIL: TestMul3 (0.00s)
    xxx/calc_test.go:44: 2 * 0 expected 1, but 0 got
```

- 如果报错出现在帮助函数里面，只打印帮助函数的信息，第一时间很难确定是由哪个用例出现错误的。 所以， Go 语言在 1.9 版本中引入了 `t.Helper()`，用于标注该函数是帮助函数，**报错时将输出帮助函数调用者的信息**，而不是帮助函数的内部信息。
新的报错：
```
--- FAIL: TestMul3 (0.00s)
    xxx/calc_test.go:53: 2 * 0 expected 1, but 0 got
```


**两条建议**
- 不要返回错误， 帮助函数内部直接使用 t.Error 或 t.Fatal 即可，在用例主逻辑中不会因为太多的错误处理代码，影响可读性。
- 调用 t.Helper() 让报错信息更准确，有助于定位。


## setup 和 teardown
如果在同一个测试文件中，每一个测试用例运行前后的逻辑是相同的，一般会写在 setup 和 teardown 函数中。
```go
func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
}

func Test2(t *testing.T) {
	fmt.Println("I'm test2")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
```

* 在这个测试文件中，包含有2个测试用例，Test1 和 Test2。
* 如果测试文件中包含函数 TestMain，那么生成的测试将调用 TestMain(m)，而不是直接运行测试。
* 调用 m.Run() 触发所有测试用例的执行，并使用 os.Exit() 处理返回的状态码，如果不为0，说明有用例失败。
* 因此可以在调用 m.Run() 前后做一些额外的准备(setup)和回收(teardown)工作。


## 网络测试

### TCP/HTTP
假设需要测试某个 API 接口的 handler 能够正常工作，例如 helloHandler
```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
```

**测试**：
```go
func handlerError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

func TestConn(t *testing.T) {
    // 监听一个未被占用的端口，并返回 Listener。
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	handlerError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		var b []byte = []byte("hello " + r.URL.RawQuery)
		w.Write(b)
	})
	go http.Serve(ln, nil)

    // 尽量不对 http 和 net 库使用 mock，这样可以覆盖较为真实的场景
	resp, err := http.Get("http://" + ln.Addr().String() + "/hello?world")
	handlerError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handlerError(t, err)

	if string(body) != "hello world" {
		t.Fatal("expect hello world, but got", string(body))
	}
}
```

### httptest
针对 http 开发的场景，使用标准库 net/http/httptest 进行测试更为高效。
```go
func TestConn(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	bytes, _ := ioutil.ReadAll(w.Result().Body)

	if string(bytes) != "hello world" {
		t.Fatal("expected hello world, but got", string(bytes))
	}
}
```
使用 httptest 模拟请求对象(req)和响应对象(w)，达到了相同的目的。


### Benchmark基准测试
* 函数名必须以 Benchmark 开头，后面一般跟待测试的函数名
* 参数为 b *testing.B。
* 执行基准测试时，需要添加 -bench 参数。
- 如果在运行前基准测试需要一些耗时的配置，则可以使用 b.ResetTimer() 先重置定时器。 类似sleep
```go
func BenchmarkHello(b *testing.B) {
    ... // 耗时操作
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}
```

基准测试报告，每一列值对应的含义：
```
 go test -benchmem -bench .
...
                   int 迭代次数  time.Duration基准测试花费的时间    一次迭代处理的字节数   总的分配内存的次数   总的分配内存的字节数
BenchmarkHello-16   15991854         71.6 ns/op                   5 B/op          1 allocs/op
...
```

- 使用 RunParallel 测试并发性能

```go
func BenchmarkParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
```

# go mock

## 简介

当待测试的函数/对象的依赖关系很复杂，并且有些依赖不能直接创建，例如数据库连接、文件I/O等。这种场景就非常适合使用 `mock/stub` 测试。


[gomock](https://github.com/golang/mock) 是官方提供的 mock 框架，同时还提供了 `mockgen` 工具用来辅助生成测试代码。


## 一个简单的Demo
```go
type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}
	return -1
}
```

- 假设 DB 是代码中负责与数据库交互的部分，测试用例不能创建真实的数据库连接，如果我们需要测试 GetFromDB 这个函数内部的逻辑，就需要 mock 接口 DB。
  - 第一步：使用 mockgen 生成 db_mock.go。 一般传递三个参数。包含需要被mock的接口得到源文件source，生成的目标文件destination，包名package.
  > mockgen -source=db.go -destination=db_mock.go -package=main
  - 第二步：新建 db_test.go，写测试用例。
```go
func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
```
- 这个测试用例有2个目的:
  - 一是 使用 `ctrl.Finish()` 断言 `DB.Get()`被是否被调用，如果没有被调用，后续的 mock 就失去了意义；
  - 二是 测试方法 GetFromDB() 的逻辑是否正确(如果 DB.Get() 返回 error，那么 GetFromDB() 返回 -1)。
  - NewMockDB() 的定义在 db_mock.go 中，由 mockgen 自动生成。


## 打桩（stubs）
在上面的例子中，**当** Get() 的**参数为** Tom，**则返回** error，这**称之为打桩(stub)**，有明确的参数和返回值是最简单打桩方式。

### 比较参数（Eq、Any、Not、Nil）

```go
m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
m.EXPECT().Get(gomock.Any()).Return(630, nil)
m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil) 
m.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil")) 
```

* Eq(value) 表示与 value 等价的值。
* Any() 可以用来表示任意的入参。
* Not(value) 用来表示非 value 以外的值。
* Nil() 表示 None 值。


### 返回值（Return, DoAndReturn）

```go
m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
m.EXPECT().Get(gomock.Any()).Do(func(key string) {
    t.Log(key)
})
m.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
    if key == "Sam" {
        return 630, nil
    }
    return 0, errors.New("not exist")
})
```
* Return 返回确定的值
* Do Mock 方法被调用时，要执行的操作吗，忽略返回值。
* DoAndReturn 可以动态地控制返回值。


### 调用次数（Times）
```go
func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).Times(2)
	GetFromDB(m, "ABC")
	GetFromDB(m, "DEF")
}
```
- Times() 断言 Mock 方法被调用的次数。
- MaxTimes() 最大次数。
* MinTimes() 最小次数。
* AnyTimes() 任意次数（包括 0 次）。


### 调用顺序(InOrder)
```go
func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	o1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	o2 := m.EXPECT().Get(gomock.Eq("Sam")).Return(630, nil)
	gomock.InOrder(o1, o2)
	GetFromDB(m, "Tom")
	GetFromDB(m, "Sam")
}
```


### 如何编写可 mock 的代码
* `mock`作用的是接口，因此将依赖抽象为接口，而不是直接依赖具体的类。
* 不直接依赖的实例，而是**使用依赖注入**降低耦合性。比如下面这种情况，对 DB 接口的 mock 并不能作用于 GetFromDB() 内部，这样写是没办法进行测试的。
```
func GetFromDB(key string) int {
    //这个依赖不是注入进来的，而是自己实例化的，就没法mock
	db := NewDB()
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}
```