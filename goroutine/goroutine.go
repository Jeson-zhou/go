package main

import (
	"fmt"
	"time"
)

// go并发示例
/*func hello() {
	fmt.Println("Hello, Goroutine")
}

func main() {
	go hello()
	fmt.Println("Main, Goroutine")
}*/

// goroutine的同步
/*func main() {
	var v int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		v = 1
		wg.Done()
	}()
	go func() {
		fmt.Println(v)
		wg.Done()
	}()
	wg.Wait()
}*/

// goroutine为什么需要同步？
/*var shareCnt int
func incrShareCnt() {
	for i := 0; i < 100000; i++ {
		shareCnt++
	}
}

func main() {
	for i := 0; i < 2; i++ {
		go incrShareCnt()
	}
	time.Sleep(time.Second * 10)
	fmt.Println(shareCnt)
}*/

// goroutine 实现同步
func producer(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("大妈做的第%d个面包\n", i)
		ch <- i
		time.Sleep(time.Second * time.Duration(1))
	}
}

func consumer(ch chan int, count int) {
	for v := range ch {
		fmt.Printf("大叔吃了第%d个面包\n", v)
		count--
		if count == 0 {
			fmt.Println("没面包了，吃饱了")
			close(ch)
		}
	}
}

func main() {
	ch := make(chan int)
	count := 8
	go producer(ch, count)
	consumer(ch, count)
}
