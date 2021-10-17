package main


import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Product struct {
	// 生成该产品的生产者序列号
	prodNum int
	value   int
}

/*
采用 channel 实现产品的双向传递，缓冲区设置为 10
即实现了 empty 和 full 这两个资源信号量的功能
任意时刻只有一个 goroutine 可以对 channel 进行访问：
即实现 mutex 作用
*/
var ch = make(chan Product, 10)

// 终止生成的信号
var stop = false

func Producer(wg *sync.WaitGroup, proNum int) {
	for !stop {
		p := Product{
			prodNum: proNum,
			value:   rand.Int(),
		}
		ch <- p
		fmt.Printf("Producer %v produce a product : %#v\n", proNum, p)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
	wg.Done()
}
func Consumer(wg *sync.WaitGroup, conNum int) {
	for p := range ch {
		fmt.Printf("consumer %v consumes a product : %#v\n", conNum, p)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
	wg.Done()
}
func main() {
	var wgp sync.WaitGroup
	var wgc sync.WaitGroup
	wgp.Add(5)
	wgc.Add(5)
	// producer and consumer are both 5
	for i := 0; i < 5; i++ {
		go Producer(&wgp, i)
		go Consumer(&wgc, i)
	}
	time.Sleep(time.Duration(1) * time.Second)
	stop = true
	wgp.Wait()
	close(ch)
	wgc.Wait()
}
