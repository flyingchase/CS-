# Go cheatsheet

## Variables

- 短声明不可在函数外部
  - 短声明至少一个新变量
- 常量 const
  - 可忽略类型 也可指定类型

**关键字**

1. `break default      func    interface select`
2. `case     defer        go      map          struct`
3. `chan     else goto package switch`
4. `const    fallthrough  if      range        type`
5. `continue for import return var`

### iota

- 枚举使用  为 const 常量赋值

- 按照 const 内的行数从 0 递增  同一行的值相等

- 每次遇见 const 重置为 0







rune 是 unit32 别称

byte 是 unit8 别称







### make new

#### make

- 内建 map slices channel 的内存分配
- 返回有初始值的 T类型 非*T 指针



#### new

- 各种类型的内存分配
- 返回指针 指向新分配的类型的零值





## Basic types

### Strings

``` go
// 不会将换行符\n 转化 直接输出换行
str:=`Multiline
string`
```

- 字符串不可变
  - 改变需要转化为`[]byte `类型 
  - +和切片操作可以更改  `s=“c”+s[1:]`
- 再强转 string(ch)即可
- 可使用`+`链接
- 



### Numbers

```go
num1:=3
num2:=3.1
num3:=3+4i
num4:=byte('a') // byte 底层为alias unit8 

fmt.Printf("reflect.TypeOf(num1): %v\n", reflect.TypeOf(num1))
fmt.Printf("reflect.TypeOf(num2): %v\n", reflect.TypeOf(num2))
fmt.Printf("reflect.TypeOf(num3): %v\n", reflect.TypeOf(num3))
fmt.Printf("reflect.TypeOf(num4): %v\n", reflect.TypeOf(num4))

// reflect.TypeOf(num1): int
// reflect.TypeOf(num2): float64
// reflect.TypeOf(num3): complex128
// reflect.TypeOf(num4): uint8
```

### Array

```go
numbers:=[...]int{1,2,3,4,5} // 忽略数组长度 直接按照后续初始化计长
for _,num:=range numbers{
  println(num)
}
```



### Slices

- `var slice []byte `
  - Slice[ : ：] 第一个默认位置为`0` 第二个默认位置是 len 第三个位置是容量的最后一个对应数组的下标 则 cap=(最后一位-第一位)
    - 前闭后开
- 类似结构体 内含有 *指针（指向 数组中slice 的开始位置）、len（代表 slice长度）、cap（代表 slice 开始位置到对应数组的最后位置长度）*
  - append 追加——>指针 会影响slice 指向的数组  
    - 当 slices(cap-len==0) 时slices 分配新的空间 返回新的数组指针 原数组内容不变
  - len 长度 
  - cap 容量 
  - copy 复制到目标 dst 返回复制元素个数

### Pointers



### map

- 字典无序 长度不固定，是引用类型
- len 获取 key 的数量
- 不存在 k 对应的 v 返回为 bool
- *API*
  - delete 删除 delete(map,key)
  - v







## Flow control

- if内声明的变量作用域仅在代码块内
- 无 while
- Switch 中 case 默认带 break 需要 `fallthrough`

### Conditional

#### siwtch

- switch中 Expr 与 case 的各个 expr判断必须是同类型的 
- switch 无表达式则默认为 `true`
- 

### Goto

与 `Here:`标签搭配  必须跳转到当前函数内定义的标签 

- 标签名大小写敏感
- 









## Functions

- 函数可作为参数传递

  - 成为嵌套函数  在函数中调用其他函数

  ```go
  package main
  
  import "fmt"
  // 首先需要定义函数类型 作为其它函数的传入参数
  type testInt func(int) bool
  // 按照函数类型来编写函数
  func isOdd(integer int) bool {
     if integer%2 == 0 {
        return false
     }
     return true
  }
  func isEven(integer int) bool {
     if integer%2 != 0 {
        return false
     }
     return true
  }
  func filter(slice []int, f testInt) []int {
     var res []int
     for _, val := range slice {
        if f(val) {
           res = append(res, val)
        }
     }
     return res
  }
  func main() {
     slice := []int{1, 2, 3, 4, 5, 6, 7}
     fmt.Println(slice)
     odd := filter(slice, isOdd) // 将函数作为参数传入函数
     fmt.Println(odd)
     even := filter(slice, isEven)
     fmt.Println(even)
  }
  ```

  - 将函数作为参数传递要求
    - 定义函数类型 
    - 要作为参数传递的函数类型与定义的函数类型相同
  
  #### Variadic Functions 可变参数
  
  - ``` go
    func funcName (nums ...int) {  // 不指定参数个数
        fmpl(nums)
        for _,num:=rang nums{
            sum+=num;
        }
        fmpl(sum)
    }
    func main() {
        sum(1,2)
        sum(1,2,3)
        nums:=[]int{1,2,3,4}
        sum(nums...)
    }
    ```
  
  #### Closures 返回func 
  
  - ```go
    func intSeq() func() int{    // 函数返回值为函数
       i:=0
       return func() int {
           i++
           return i
       }
    }
    func main() {
       nextIntL:=intSeq()
       fmt.Println(nextIntL()) // 1
       fmt.Println(nextIntL()) //2
       fmt.Println(nextIntL())//3
       newIntL:=intSeq()
       fmt.Println(newIntL()) //1 
    }
    ```

#### Recursion



```go
func fact (n int) int{
	if n==0 {
		return 1
	}
	return n* fact(n-1)  
}

func main(){
	fmt.Println(fact(7))  

	var fib func(n int) int  // 将函数定义为变量 

	fib=func (n int) int {
		if n<2 {
			return n
		}
		return fib(n-1)+fib(n-2)
	}
	fmt.Println(fib(7))
}
```


### Lambda



`arg …int`表示函数接受不定数量的变参

Func myFunc (arg ..int) int{}







### defer

- 函数执行到最后return 后  defer 逆序执行 最后返回值退出
- 



## Packages

两个保留函数 main 和 init

- `main 函数`只能用于`package main` `init函数`可用于各个 package
- 两者无参数和返回值



### 点操作

**. “packhageName”：**

- 点操作含义是调用该包内函数时 可忽略前缀包名
  - `fmt.Println()`可简写为`Println()`

### 别名操作

- 将包名在本文件内作为自定义名称
- 防止冲突/易于记忆包名





### _操作 

- 导入该包但不直接使用包内函数 调用包内的 `init函数`









## Error control

panic 和 recover 机制





## Structs

- 匿名字段可包含对应 struct 的所有内容

- 内置类型 Int 等也可以作为匿名字段

- 内外层有重复的字段 优先访问最外层的字段



## Methods

- 附属在给定的 struct 上 在函数基础上加上 receiver
- 方法可与在任何自定义的类型、内置类型、struct 上
  - type typeName typeLiteral   称 typeName 为自定义类型 
  - struct 是自定义中的一种





- 继承 
  - 匿名字段的 method 可被继承 包含该匿名字段的 struct 可调用期内的 methods





- 重写 override
  - 外部 struct 可重写内部 struct 的 methods

## Interfaces

- interface 可被任意对象实现
- 一个对象可实现多个 interface
- 任意类型都实现了空 interface

- 是一抽象方法的集合



#### 变量存储类型

- 断言`Comma-ok`
  - `value,ok:=element(T)` ok 是 bool 类型 true 代表是该类型 多在 if-else 使用
  - `switch` 使用`element.(type)` 作为判断和 case

```go
type Element interface {
} // 可作用在任何类型上
type List[] Element // element 为空接口并将其自定义名称为 List[]
type Person struct {
   name string
   age int
}

func (p *Person) String() string {
   return p.name+strconv.Itoa(p.age)

}
func main() {
list:=make(List,3)
list[0] = 1
list[1] = "hello"
list[2]=Person{name : "w",age :100}
for index,element:=range list {
   if value, ok := element.(int); ok {
      fmt.Println(index, value)
   }else if value, ok := element.(Person); ok {
      fmt.Println(index,value)
   }else if value, ok := element.(string); ok {

      fmt.Println(index,value)
   }else {
      fmt.Println(index)
   }
}
```



函数赋值给变量 函数内是空接口

```go
whatAmI := func(i interface{}) {
  switch t := i.(type) {
  case bool:
     fmt.Println("bool")
  case int:
     fmt.Println("int")
  default:
     fmt.Println(reflect.TypeOf(t))
  }
}
whatAmI("1")
whatAmI(true)
whatAmI(1)
whatAmI(Studnet{})
```





## Concurrency

- 通过通信来共享而非通过共享来通信
- goroutine 运行在相同的地址空间

#### channel

- channel 必须使用 make 创建并定义发送到 channel 的类型

- cap 可以读取 channel 的缓存容量

  - ``` go
    ch:=make(chan int) // 定义同时规定发送到 chan 的类型
    v:=1
    ch<-v 	// 将 v 发送到 chan
    w:=<-ch	// 从 ch中读取并赋值给 w
    ```

- 默认无缓存chan 接受和发送是阻塞的  除非另外一端准备好  

  - `value:<-c` 读取会被阻塞 直到有数据接受

  - `ch<-5` 发送会被阻塞 直到有被读取

  - ```go
    func sum(a []int, c chan int)  {
       total:=0
       for _,v:=range a{
          total+=v
       }
       c<-total
    
    }
    func main() {
       a:=[...]int {1, 2, 3,4,5,6,7,8,9,0}
       c:=make(chan int)
       go sum(a[:len(a)/2],c)
       go sum(a[len(a)/2:],c)
       x,y:=<-c,<-c
       fmt.Println(x,y,x+y)
    }
    ```

- Buffer channel 带缓存的 channel

  - `ch：=make(chan type,cap)` 在 cap 内的读写是无阻塞的 超过 cap 时候需要等待其他 goroutine 从 channel 中读取元素 释放空间

  - `for i:=rang c` 可以不断读取 channel 直到 channel 被显示关闭 

  - `close`关闭 channel 无法发送数据 但可以在消费方通过`v,ok:=c;ok`*断言*来测试是否关闭 channel

    - 在生产者方关闭 channle  在消费者方容易产生 panic
    - 结束 range 循环/无数据需要发送时候关闭 channel 才使用 close显式关闭

  - ```go
    func fibonacc(n int, c chan int)  {
       x,y:=1,1
       for i:=0;i<n;i++ {
          c<-x
          x,y=y,x+y
       }
       close(c)	// 显示关闭 channel
    }
    func main() {
       c:=make(chan int, 10)
       go fibonacc(cap(c),c)
       for i:=range c {
          fmt.Println(i)
       }
    }
    ```



- 多 channel 下 *select*

  - select 关键字监听 channel 上的数据流动

  - 默认是阻塞的  只有监听的 channel 上数据流动才运行  多个 channel 准备好时select 随机选择执行

  - 类似 switch 有 `default`   可设置超时来退出循环

  - ```go
    func fibonacc(c, quit chan int) {
       x, y := 1, 1
       for {
          select {		// 类似 switch  监听 channel 上的数据流动
          case c <- x:
             x, y  = y, x+y
          case <-quit:
             fmt.Println("quit")
             return
          case <-time.After(5*time.Second) // 阻塞超过5 s情况
              fmpl("timeOut")
              break;
          }
       }
    }
    
    func main() {
       c := make(chan int)
       quit:=make(chan int)
       go func() {
          for i := 0; i < 10; i++ {
             fmt.Println(<-c)
          }
          quit<-0
       }()
       fibonacc(c,quit)
    }
    ```

- `Goexit` 
  - 退出当前 Goroutine 但defer 仍调用
- `Gosched`
  - 让出 goRoutine 权限
- `NumCPU` `NumGoroutine`
  - 返回 cpu数目 
  - 返回执行和排队的任务总数
- `GOMAXPROCS`
  - 并行计算的 cpu 核数最大值 返回之前的核数值

## References





