# 《Go Web开发入门》

# Web开发基础


# Web开发框架

## 为什么不直接用标准库，必须使用框架？
net/http提供了基础的Web功能，即监听端口，映射静态路由，解析HTTP报文。

一些Web开发中简单的需求并**不支持**，需要手工实现。
* 动态路由：例如hello/:name，hello/*这类的规则。
* 鉴权：没有分组/统一鉴权的能力，需要在每个路由映射的handler中实现。
* 模板：没有统一简化的HTML机制。

框架的核心能力：
* 路由(Routing)：将请求映射到函数，支持动态路由。例如'/hello/:name。
* 模板(Templates)：使用内置模板引擎提供模板渲染机制。
* 工具集(Utilites)：提供对 cookies，headers 等处理机制。
* 插件(Plugin)：Bottle本身功能有限，但提供了插件机制。可以选择安装到全局，也可以只针对某几个路由生效。

## Go Gin 简明教程
Gin 是使用 Go/golang 语言实现的 HTTP Web 框架。**接口简洁，性能极高**。截止 1.4.0 版本，包含测试代码，仅14K，其中测试代码 9K 左右，也就是说框架源码仅 5K 左右。

### Gin特性
* **快速**：路由不使用反射，基于Radix树，内存占用少。

* **中间件**：HTTP请求，可先经过一系列中间件处理，例如：Logger，Authorization，GZIP等。这个特性和 NodeJs 的 Koa 框架很像。中间件机制也极大地提高了框架的可扩展性。

* **异常处理**：服务始终可用，不会宕机。Gin 可以捕获 panic，并恢复。而且有极为便利的机制处理HTTP请求过程中发生的错误。

* **JSON**：Gin可以解析并验证请求的JSON。这个特性对Restful API的开发尤其有用。

* **路由分组**：例如将需要授权和不需要授权的API分组，不同版本的API分组。而且分组可嵌套，且性能不受影响。

* **渲染内置**：原生支持JSON，XML和HTML的渲染。


### 安装Gin
```
go get -u -v github.com/gin-gonic/gin
```
-v：打印出被构建的代码包的名字
-u：已存在相关的代码包，强行更新代码包及其依赖包


### Hello Gin
```go
package main

import "github.com/gin-gonic/gin"

func main() {
    //生成了一个实例，这个实例即 WSGI 应用程序
	r := gin.Default()
    //声明了一个路由，告诉 Gin 什么样的URL 能触发传入的函数
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin")
	})
    //让应用运行在本地服务器上，默认监听端口是 _8080_，可以传入参数设置端口(注意要带冒号)
	r.Run(":8089") // listen and serve on 0.0.0.0:8080
}
```

### 路由
路由方法有 GET, POST, PUT, PATCH, DELETE 和 OPTIONS，还有Any，可匹配以上任意类型的请求。

- 无参数
- 解析路径参数
- 获取Query参数
- 获取Post参数
- GET 和 POST 混合
- Map字典参数
- 重定向

```go
//1. 无参数
	//生成了一个实例，这个实例即 WSGI 应用程序
	r := gin.Default()
	//声明了一个路由，告诉 Gin 什么样的URL 能触发传入的函数
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin")
	})

	//2. 解析路径参数
	// 匹配 /user/geektutu
	//http://localhost:9999/user/geektutu
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	// 3. 获取Query参数
	// 匹配users?name=xxx&role=xxx，role可选
	//http://localhost:9999/users?name=Tom&role=student
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is an %s", name, role)
	})

	// 4. 获取POST参数
	//$ curl http://localhost:9999/form  -X POST -d 'username=geektutu&password=1234'
	//{"password":"1234","username":"geektutu"}
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// 5. GET 和 POST 混合
	//curl "http://localhost:9999/posts?id=9876&page=7"  -X POST -d 'username=geektutu&password=1234'
	//{"id":"9876","page":"7","password":"1234","username":"geektutu"}
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	//6. Map字典参数
	//curl -g "http://localhost:9999/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'
	//{"ids":{"Jack":"001","Tom":"002"},"names":{"a":"Sam","b":"David"}}
	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	//7. 重定向（Redirect）
	// curl -i http://localhost:9999/redirect
	// curl "http://localhost:9999/goindex"
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	r.Run(":9999") // listen and serve on 0.0.0.0:8080
```

### 分组路由
类似Controller层类上面加个统一的前缀 /cpc/ ,后面方法/addCpc、/deleteCpc等。
- 利用分组路由还可以更好地实现权限控制，例如将需要登录鉴权的路由放到同一分组中去，简化权限控制。
```go
defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}

	//group: v1
	v1 := r.Group("/v1")

	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}

	//group: v2
	v2 := r.Group("/v2")

	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}
```

### 上传文件
```go
    r := gin.Default()

	// 单个文件
	r.POST("/upload1", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		// ctx.SaveUploadedFile(file, dst)
		ctx.String(http.StatusOK, "%s uploaded!", file.Filename)
	})

	// 多个文件
	r.POST("/upload2", func(ctx *gin.Context) {
		// Muiltpart form
		form, _ := ctx.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Panicln(file.Filename)
			// ctx.SaveUploadedFile(file, dst)
		}
		ctx.String(http.StatusOK, "%s files up;oaded!", len(files))
	})

	r.Run(":9999") // listen and serve on 0.0.0.0:8080
```


**debug的方法**：dlv
使用：
1、dlv debug xxx.go 指定需要debug的文件
2、进入dlv交互式窗口后，b <filename>:<line> 指定断点
3、r arg 指定运行参数
4、n 执行一行
5、c 运行至断点或程序结束

### HTML模板（Template）
```go
r.LoadHTMLGlob("templates/*")

	stu1 := &student{Name: "Bob", Age: 12}
	stu2 := &student{Name: "Lili", Age: 13}
	r.GET("/arr", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
```
还需要搭配写个arr.tmpl文件：
```html
<!-- templates/arr.tmpl -->
<html>
<body>
    <p>hello, {{.title}}</p>
    {{range $index, $ele := .stuArr }}
    <p>{{ $index }}: {{ $ele.Name }} is {{ $ele.Age }} years old</p>
    {{ end }}
</body>
</html>
```

- Gin默认使用模板Go语言标准库的模板text/template和html/template，语法与标准库一致，支持各种复杂场景的渲染。
- 参考官方文档[text/template](https://golang.org/pkg/text/template/)，[html/template](https://golang.org/pkg/html/template/)


### 中间件


### 热加载调试 Hot Reload
fresh： go get -v -u github.com/pilu/fresh

**使用：** fresh run main.go

# 开发实践