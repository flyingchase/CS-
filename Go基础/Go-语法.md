

## 1 并发

主要应用在网络服务器编程



### 1.1 Goroutines

并发执行单元 

主函数在单独goroutine执行 成为 `main goroutunes` 通过关键字 `go`实现

终止程序/主函数退出 即可中断goroutines执行 

```go
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 2; i++ {
    runtime.Gosched()	//GoSched()表示cpu把时间片让出
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

//world和hello交替输出两次

```



```go
package main

import (
	"fmt"
	"time"
)
//func spinner和fib在两个函数中但是同时执行 
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
//主函数返回后均中断goruntines 输出fib(45)
func main() {
	go spinner(100 * time.Microsecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d)=%d\n", n, fibN)	

}

```





#### 1.1.1并发clock服务



```go
// clock1.go文件 
package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	go spinner(100 * time.Microsecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d)=%d\n", n, fibN)

}

```



```go
//tcp client only-read netcat1.go文件 实现类似iterm中的nc localhost 8000功能

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	coon, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer coon.Close()
	mustCopy(os.Stdout, coon)
}

```



终端效果

<img src="/Users/qlzhou/Library/Application Support/typora-user-images/image-20201226125802166.png" alt="image-20201226125802166" style="zoom:70%;" />

### 1.2channels

goroutines运行在相同的地址空间 goroutines之间的数据通信机制 channel

- 定义channel时必须定义发送到channel的值的类型

- 必须使用make创建channel

```go
ci:=make(chan int)
cs:=make(chan string)
cf:=make(chan interface{})
//channel用于参数传递时为引用

ch:=make(chan int ,1)//带缓存的channe 只要后面容量大于0
```

- 使用操作符 `<-`发送和接受数据

```go
ch <- v	//发送v到chaanel ch 
v:=<-ch	//接受数据赋值给v
<-ch //无接受结果的接受操作合法
```

- 

#### 1.2.1 同步channels 无缓存channels

- 默认中channel接受/发送是阻塞的 直到另外一端准备 针对无缓存的channels

  读取 `value:=<-ch` 接受赋值会被堵塞直到有数据

  发送 `ch<-5`会被阻塞 直到被读取 

  基于无缓存的channel发送操作导致发送者gorotines阻塞 知道另外一个gorotines在相同channels执行接受操作 

  基于无缓存的channels的发送和接受操作会导致两个goroutines做一次同步 

```go
package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total 
  
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

//输出 -5 17 12
```

- 同类型channel可以 `==`运算符比较 结果为boolean 或者与零值 `nil`比较
- `close()`关闭channel 再对其发送操作导致panic异常 但是接受操作可以正常



#### 1.2.2 带缓存的channel    buffered channels

`cap(ch)`返回ch的容量 `len(ch)`返回有效元素个数 返回值为int类型

向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素。如果内部缓存队列是满的，那么发送操作将阻塞直到因另一个goroutine执行接收操作而释放了新的队列空间。相反，如果channel是空的，接收操作将阻塞直到有另一个goroutine执行发送操作而向队列插入元素。

```go
ch := make(chan type,value)
ch:=make(chan int,4)
```

value>1 channe是有缓冲、非阻塞 

```go
func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
```

- 使用range进行类似slice和map操作带缓存的channel 

```go
package main

import "fmt"

func fib(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fib(cap(c), c)
  // for i:=range c 一直读取channel中数据直到被显式关闭
	for i := range c {	
		fmt.Println(i)
	}
  v,ok:=<-c
  fmt.Println(ok)	//ok为fales则表明chan被关闭
  
}


```



- 多个channele操作 **select**

关键字select监听channel数据流 默认阻塞 只有被监听的channel发生操作时运行

```go
package main

import "fmt"

func fib(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y := y, x+y
			_, _ = x, y //避免定义未使用的报错
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	// fmt.Println("first c is", <-c)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fib(c, quit)
}

```



#### 1.2.3 串联channels ——>pipeline

channels输出作为下一个channels的输入 

利用close()关闭channles 利用多个channels的接受结果判断是否关闭

```go
package main

import "fmt"
// 三个goroutines 利用channels在之间传递
func main() {
	squares, naturals := make(chan int), make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
  //close()之后读取是正常的 无法传入
	for x := range squares {
		fmt.Println(x)
	}
}

```



