
# Go RPC & TLS 鉴权简明教程
Go 语言远程过程调用(Remote Procedure Call, RPC)的使用方式，示例基于 Golang 标准库 net/rpc，同时介绍了如何基于 TLS/SSL 实现服务器端和客户端的单向鉴权、双向鉴权。

## RPC简介
>远程过程调用（英语：Remote Procedure Call，缩写为 RPC）是一个计算机通信协议。
- 该协议允许运行于一台计算机的程序调用另一个地址空间（通常为一个开放网络的一台计算机）的子程序，而**程序员就像调用本地程序一样，无需额外地为这个交互作用编程**（无需关注细节）。
- RPC是一种服务器-客户端（Client/Server）模式，经典实现是一个通过发送请求-接受回应进行信息交互的系统。
- RPC 协议假定某种传输协议(TCP, UDP)存在，为通信程序之间携带信息数据。


- 相比于调用本地的接口，RPC 还需要知道的是服务器端的地址信息。


## 一个简单的Demo
计算二次方的程序：
- Cal 结构体，提供了 Square 方法，用于计算传入参数 num 的 二次方。


**本地调用版本**：


```go
// main.go
package main

import "log"

type Result struct {
	Num, Ans int
}

type Cal int

func (cal *Cal) Square(num int) *Result {
	return &Result{
		Num: num,
		Ans: num * num,
	}
}

func main() {
	cal := new(Cal)
	result := cal.Square(12)
	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
```


**RPC需要满足什么条件？**
- 有一定的约束和规范，Golang标准库中的`net/rpc`的方法需要长下面这样：

```go
func (t *T) MethodName(argType T1, replyType *T2) error
```

即需要满足以下 5 个条件：
1. 方法类型（T）是导出的（首字母大写）
2. 方法名（MethodName）是导出的
3. 方法有2个参数(argType T1, replyType *T2)，均为导出/内置类型
   1. net/rpc 对参数个数的限制比较严格，仅能有2个，
   2. 第一个参数是调用者提供的请求参数，
   3. 第二个参数是返回给调用者的响应参数
4. 方法的第2个参数一个指针(replyType *T2)
5. 方法的返回值类型是 error


**满足RPC条件的版本**


```go
func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func main() {
	cal := new(Cal)
	var result Result
	cal.Square(15, &result)
	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
```
* Cal 和 Square 均为导出类型，满足条件 1) 和 2)
* 2 个参数，num int 为内置类型，result *Result 为导出类型，满足条件 3)
* 第2个参数 result *Result 是一个指针，满足条件 4)
>方法 Cal.Square 满足了 RPC 调用的5个条件。可以用rpc进行改造了。


## RPC服务与调用


### 基于HTTP，启动RPC服务
RPC是典型的CS架构：需要将 Cal.Square 方法放在服务端。
- 服务端需要提供一个套接字服务，处理客户端发送的请求。通常可以基于 HTTP 协议，监听一个端口，等待 HTTP 请求。


**服务端Server**
```go
type Result struct {
	Num, Ans int
}

type Cal int

func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}
func main() {
	// 发布 Cal 中满足 RPC 注册条件的方法（Cal.Square）
	rpc.Register(new(Cal))
	// 注册用于处理 RPC 消息的 HTTP Handler
	rpc.HandleHTTP()

	// 监听 1234 端口，等待 RPC 请求。
	log.Printf("Serving RPC server on port %d", 1234)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}
```
- RPC 服务启动，等待客户端的调用。



**客户端Client**


```go
type Result struct {
	Num, Ans int
}

func main() {
	// 创建了 HTTP 客户端 client，并且创建了与 localhost:1234 的链接，1234 恰好是 RPC 服务监听的端口
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var result Result
	// 调用远程方法，第1个参数是方法名 Cal.Square，后两个参数与 Cal.Square 的定义的参数相对应。
	if err := client.Call("Cal.Square", 12, &result); err != nil {
		log.Fatal("Failed to call Cal.Square.", err)
	}
	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
```
- rpc.Call 调用远程方法，**第1个参数是方法名** Cal.Square，后两个参数与 Cal.Square 的定义的参数相对应。


## 异步调用
- `client.Call` 是同步调用的方式，会阻塞当前的程序，直到结果返回。
- `client.Go`: 是异步调用,客户端调用后，程序可以继续往下走。
```go
	// 创建了 HTTP 客户端 client，并且创建了与 localhost:1234 的链接，1234 恰好是 RPC 服务监听的端口
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var result Result
	// 异步调用
	asyncCall := client.Go("Cal.Square", 12, &result, nil)

	log.Printf("%d^2 = %d", result.Num, result.Ans) // 2020/01/13 21:34:26 0^2 = 0

	// 阻塞当前程序直到 RPC 调用结束
	<-asyncCall.Done
	log.Printf("%d^2 = %d", result.Num, result.Ans) // 2020/01/13 21:34:26 12^2 = 144
}
```


## 证书鉴权（TLS/SSL）

### 客户端对服务器端鉴权
- HTTP 协议默认是不加密的，我们可以使用证书来保证通信过程的安全。

- 生成私钥和自签名的证书，并将 server.key 权限设置为只读，保证私钥的安全。


**生成私钥和自签名的证书**
```s
# 生成私钥
openssl genrsa -out server.key 2048
# 生成证书
openssl req -new -x509 -key server.key -out server.crt -days 3650
# 只读权限
chmod 400 server.key
```


**服务端可以使用证书启动TLS端口监听**
```go
    // 发布 Cal 中满足 RPC 注册条件的方法（Cal.Square）
	rpc.Register(new(Cal))
	cert, _ := tls.LoadX509KeyPair("server.pem", "server.key")
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	// 服务器端可以使用生成的 server.crt 和 server.key 文件启动 TLS 的端口监听。
	listener, _ := tls.Listen("tcp", ":1234", config)
	log.Printf("Serving RPC server on port %d", 1234)

	/*
		listener.Accept() 阻塞等待客户端与服务端建立连接，建立连接后交给 rpc.ServeConn 异步处理。
		因为可能有多个客户端建立连接，所以需要无限循环，
		每建立一个链接，就异步处理，然后继续等待下一个连接建立。
	*/
	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
```


**客户端使用tsl.Dial，并将服务端的证书添加到信任证书池中**
```go
    certPool := x509.NewCertPool()
	certBytes, err := ioutil.ReadFile("/Users/bobo-mac/Documents/code/study/go/src/helloworld/rpc/tls/server/server.pem")
	if err != nil {
		log.Fatal("Failed to read server.pem")
	}
	certPool.AppendCertsFromPEM(certBytes)

	config := &tls.Config{
		// 将服务端的证书添加到信任证书池中
		RootCAs: certPool,
	}

	conn, _ := tls.Dial("tcp", "localhost:1234", config)
	defer conn.Close()
	client := rpc.NewClient(conn)

	var result Result
	if err := client.Call("Cal.Square", 12, &result); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)
	}

	log.Printf("%d^2 = %d", result.Num, result.Ans)
```


### 服务端对客户端的鉴权
与上面的类似，重点是tls.Config的配置：
- 把对方的证书添加到自己的信任证书池 RootCAs(客户端配置)，ClientCAs(服务器端配置) 中。
- 创建链接时，配置自己的证书 Certificates


**客户端**
```go
/ client/main.go

cert, _ := tls.LoadX509KeyPair("client.crt", "client.key")
certPool := x509.NewCertPool()
certBytes, _ := ioutil.ReadFile("../server/server.crt")
certPool.AppendCertsFromPEM(certBytes)
config := &tls.Config{
	Certificates: []tls.Certificate{cert},
	RootCAs: certPool,
}
```


**服务端**
```go
// server/main.go

cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
certPool := x509.NewCertPool()
certBytes, _ := ioutil.ReadFile("../client/client.crt")
certPool.AppendCertsFromPEM(certBytes)
config := &tls.Config{
	Certificates: []tls.Certificate{cert},
	ClientAuth:   tls.RequireAndVerifyClientCert,
	ClientCAs:    certPool,
}

```


# 参考或者说照搬
- 极客兔兔：https://geektutu.com/post/quick-go-rpc.html