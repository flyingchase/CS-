# Go Web



## web 概念

### Request

include post , get , cookie,  url







### Response





### Conn

用户的每次请求链接





### Handler

处理请求和生成返回信息的处理逻辑



## Http 包执行流程

![image-20210903113614403](https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/image-20210903113614403.png)

- 创建 Listen Socket 监听窗口
- Listen Socket 接收客户端请求——>client Socket 
- Client socket 读取 HTTP请求的协议头 交给对象的 handler 处理 再通过 client socket 写给客户端



#### 监听端口

- `ListenAndServe`
  初始化 serve 对象，调用`net.Listen("tcp",addr)` 底层使用 tcp 协议搭建服务来监听端口
- 







#### 接收客户端请求

- 调用`srv.Serve(net.Listener)`函数 

- 每次请求创建 Conn





#### 分配 Handler

- conn 解析 request `c.readRequest()` 获取 handler 即为调用 ListenAndServe 的第二个参数 
  - 为空 nil 则默认`handler=DefaultServeMax `  为路由器 匹配 url 跳转到对应的 handle 函数
  - 路由请求规则`“/”`： 
    - 跳转到 `http.HandleFunc("/",selfFunc)` 中的自定义函数 
    - DefaultServeMux 调用 ServeHTTP 方法 内部及调用 selfFunc







































































































# 实战

## 创建 web server

使用 http.ListenAndServer()

- 第一个参数网络地址	

  - “” 则为所有网络接口的 80 端口

- 第二个参数 handler

  - nil 则为 DefaultServeMux(multipledxer) 看做路由器

  

使用http.Server

- struct
  - Addr 字段表示网络地址
  - Handler 字段
    - nil defaulSerMux
  - ListenAndServe() 方法

``` go
// way One 
http.ListenAndServe("localhost:9090",nil)
// way Two
server:=http.Server{Addr: "localhost:8080",Handler: nil}
server.ListenAndServe()
```

以上只能执行 http 而非 https

需要分别加上：

- http.ListenAndServeTLS()
- Server.ListenAndServeTLS()



#### Handler

是一个 interface 接口

- 定义了一个方法ServeHTTP()

  - HTTPResponseWriter
  - 指向 Request（struct）的指针

  ```go
  type Handler interface {
     ServeHTTP(ResponseWriter, *Request)
  }
  ```

  







#### DefaultServeMux

Multiplexer 多路复用器（可被视为路由器） 是 ServerMux 的一个指针变量

- 也是一个 Handler
- 转发调用其他 handler
- 调用 http.Handle 函数实际上调用的是 DefaultServeMux 上的 Hanler 方法









不指定 server struct 中的 handler 字段值

- 可以使用 http.Handle 将某个 Handler 附加到 DefaultServeMux
  - http 包有一个 Handle 函数
  - ServerMux struct 也有一个 Handle 方法













#### http.Handle

func Handle(pattern string, handler Handler)

-  第二个参数是 handler（注意是*指针*）

  ``` go
  type Handler interface {
   ServeHTTP(ResponseWriter, *Request)
  // 实现 ServerHTTP 方法的类型均可视为 handler
   }
  ```

  

```go
server := http.Server{
   Addr: "localhost:1090",
   Handler: nil,  // use DefaultServeMux
}
http.Handle("/wo",&mh)
server.ListenAndServe()
```



#### http.HandleFunc

Handler函数行为与hanlder 类似 将 f 适配为 handler 使得handler 具有方法 f  类似*类型转换*

作用即为： Handler 函数转化为 Handler 内部还是调用 http.Handle 函数 

- Handler 函数的签名与 ServeHTTP 方法的签名一样，接收：
  - 一个 http.ResponseWriter
  - 一个 指向 http.Request 的指针

```go
// 第二个参数是 func 但是不要带()  带（）就直接执行了
http.HandleFunc("/home",welcome)

func welcome(w http.ResponseWriter, r *http.Request) {
   w.Write([]byte("Home!"))
}
```

HandleFunc 源码：

```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    // 调用 DefaultServeMux 的 HandleFunc
   DefaultServeMux.HandleFunc(pattern, handler) 
}

// DefaultServeMux 的 HandleFunc 第二个参数是 Handler函数（不同于 http.Handle 第二个参数是 Handler）
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	if handler == nil {
		panic("http: nil handler")
	}
    // 内部还是调用 http.Handle 函数
	mux.Handle(pattern, HandlerFunc(handler))
}
```





## 内置 Handlers

### NotFoundHandler

```go
// NotFound replies to the request with an HTTP 404 not found error.
func NotFound(w ResponseWriter, r *Request) { Error(w, "404 page not found", StatusNotFound) }

// NotFoundHandler returns a simple request handler
// that replies to each request with a ``404 page not found'' reply.
func NotFoundHandler() Handler { return HandlerFunc(NotFound) }
```

给每个请求的响应均为404

### RedirectHandler

```go
// Redirect to a fixed URL
type redirectHandler struct {
   url  string
   code int
}

// The provided code should be in the 3xx range and is usually
// StatusMovedPermanently, StatusFound or StatusSeeOther.
func RedirectHandler(url string, code int) Handler {
	return &redirectHandler{url, code}
}
```

将每个请求使用给定的状态码code——>指定的 url 跳转到提供的第一个参数



### StripPrefix

```go
func StripPrefix(prefix string, h Handler) Handler
```

从请求的 URL去掉指定的前缀prefix 再调用第二个参数 handler h

- 请求的 URL 与前缀 prefix 不符合 404
- h handler 将会在 请求 url 被去除 prefix 后调用 用于接收请求



### TimeoutHandler

```go
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
```

- time.Duration 表示一段时间 alias int64 的别名
- 返回 handler 在指定时间 dt duration 内运行传入的 h
  - h 将要被修饰的 handler
  - msg 超时则返回 msg 信息表示响应时间过长
  - dt 处理 h 的允许时间



### FileServer

```go
func FileServer(root FileSystem) Handler
```

返回 handler 基于 root 文件系统来响应请求  root 是字符串作为根目录

```go
// filestystem 可以自定义  通常委托给 
type FileSystem interface {
   Open(name string) (File, error)
}
```



eg 使用内置 filehandler 实现 handleFunc 

```go
http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
      http.ServeFile(w,r,"wwwroot"+r.URL.Path)
   })
   
http.ListenAndServe(":9090",nil)
http.ListenAndServe("8080",http.FileServer(http.Dir("wwwwroot")))
```













## 请求 Request

### HTTP 请求











### Request







### URL





### Header







### Body





























































































































