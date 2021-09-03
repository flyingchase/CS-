package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Human struct {
	name  string
	age   int
	phone string
}
type Studnet struct {
	Human  //匿名字段
	school string
	loan   float64
}
type Employee struct {
	Human
	company string
	money   float64
}

func (h *Human) SayHi() {
	fmt.Println("hello " + h.name + h.phone)
}
func (h *Human) Sling(lyrics string) {
	fmt.Println("lalalal", lyrics)
}
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle", beerStein)
}
func (e *Employee) SayHi() {
	fmt.Println("hello"+e.name+e.phone, e.company)
}

type Men interface {
	SayHi()
	Sling(lyrics string)
	Guzzle(beerStein string)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total

}
func fibonacc(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}

	}

}
func intSeq() func() int{
	i:=0
	return func() int {
	    i++
	    return i
	}
}
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("start: ",id)
	time.Sleep(time.Second)
	fmt.Println("done: ",id)

}

func main() {

	var wg sync.WaitGroup
	for i:=1;i<=5;i++{
		wg.Add(1)
		go worker(i,&wg)

	}
	wg.Wait()
	//nextIntL:=intSeq()
	//fmt.Println(nextIntL())
	//fmt.Println(nextIntL())
	//fmt.Println(nextIntL())
	//newIntL:=intSeq()
	//fmt.Println(newIntL())
//timer1:=time.NewTicker(2*time.Second)
//<-timer1.C
//fmt.Println("time 1 fired")
//
//timer2:=time.NewTicker(time.Second)
//go func() {
//
//	<-timer2.C
//	fmt.Println("tim 2 fired")
//}()
//Stop2:=timer2.Stop()
//	if Stop2{
//		fmt.Println("timer 2 stopped")
//	}
//	time.Sleep(2*time.Second)

	//s:=make([]string,3)
	//fmt.Println("emp: ",s)
	//
	//s[0]="a"
	//s[1]="b"
	//s[2]="c"
	//fmt.Println("set: ",s)
	//fmt.Println("get: ",s[2])
	//println(len(s))
	//s = append(s, "d", "e", "f")
	//fmt.Println("append: ",s)
	//
	//c:=make([]string,len(s))
	//
	//copy(c,s)
	//
	//fmt.Println("cpy: ",c)
	//
	//l:=s[2:5]
	//fmt.Println(cap(l))
	//fmt.Println(len(l))
	//
	//
	//towD:=make([][]int,3)
	//
	//for i:=0;i<3;i++ {
	//	innerLen:=i+1
	//	towD[i]=make([]int,innerLen)
	//	for j := 0; j < innerLen; j++ {
	//		towD[i][j]=i+j;
	//	}
	//
	//}
	//
	//fmt.Println(towD)
	//whatAmI := func(i interface{}) {
	//	switch t := i.(type) {
	//	case bool:
	//		fmt.Println("bool")
	//	case int:
	//		fmt.Println("int")
	//	default:
	//		fmt.Println(reflect.TypeOf(t))
	//	}
	//}
	//whatAmI("1")
	//whatAmI(true)
	//whatAmI(1)
	//whatAmI(Studnet{})
	//c := make(chan int)
	//quit := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	//fibonacc(c, quit)
}
