package main

import (
	"io"
	"log"
	"net"
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
/*func producer(ch chan int, count int) {
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
	count := 5
	go producer(ch, count)
	consumer(ch, count)
}
*/

// goroutine 并行计算
/*func spinner(delay time.Duration) {
	for {
		for _, r := range "abcdefg" {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x - 1) + fib(x - 2)
}

func main() {
	go spinner(100 * time.Microsecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\nFibonacci(%d) = %d\n", n, fibN)
}*/

// 并发的clock服务
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
