package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	"math/rand"
)

func GoChan(c chan bool) {
	go func() {
		c <- true
		fmt.Printf("Go Channel.")
		//切记 迭代的时候明确的关闭它 避免死锁
		close(c)
	}()
//	<-c
	for v := range c {
		fmt.Println(v)
	}
}

func FindNum(c chan bool, index int) {
	a := 1
	for i := 0; i<1000000 ; i++ {
		a+=i
	}

	fmt.Println(index,a)

	c <- true
}

func SyncNum(wg *sync.WaitGroup, index int) {
	a := 1
	for i:=0;i<100000;i++ {
		a+=i
	}
	fmt.Println(index,a)

	wg.Done()
}

func SelectChannel(c1 chan int,c2 chan string) {
	o := make(chan bool,2)
	go func() {
		for {
			select {
			case v,ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1",v)
			case v,ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2",v)
			}
		}
	}()
	c1 <- 1
	c2 <- "jojo"
	c1 <- 2
	c2 <- "pig"

	close(c1)

	for i := 0; i < 2; i++ {
		<- o
	}
}

func ReceviceChannel(c chan int) {
	//切记这里的顺序不要弄反了，goroutine总是先被执行
	go func(){
		for v := range c{
			fmt.Println(v)
		}
	}()

	num := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<10;i++ {
		select {
		case c <- num.Intn(1000):
		case c <- num.Intn(1000):
		}
	}
}


func main() {
	//有缓存 是异步的 没缓存是同步的
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := make(chan bool)
	GoChan(c)

	//第一种解决方案：设置channel的缓存，注意这里是异步，对每次都要有释放
	b := make(chan bool,10)
	for i:=0;i<10;i++ {
		go FindNum(b,i)
	}

	for i:=0;i<10;i++ {
		<-b
	}

	//第二种解决方案：采用sync的高级包 使用waitGroup
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i:=0;i<10;i++ {
		//注意这里是引入变量类型 &
		go SyncNum(&wg,i)
	}
	wg.Wait()

	//select channel
	c1,c2 := make(chan int),make(chan string)
	SelectChannel(c1,c2)

	d := make(chan int)
	ReceviceChannel(d)

}
