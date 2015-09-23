package main

import (
	"fmt"
	"time"
	"math/rand"
)

func Producer(p chan int) {

	defer close(p)
	for i:=0; i<10; i++ {
		p <- rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000)
	}
}

func Customer(c chan int) {

	timeoutSignal := make(chan bool)
	defer close(timeoutSignal)
	go func() {
		time.Sleep(time.Second * 5)
		timeoutSignal <- true
	}()

	for {
		if v,ok:= <-c;ok{
			fmt.Println(v)
		}
	}

	//超时有问题

}

func readChannel(c chan int) {
//	time.Sleep(time.Second)
	for v := range c {
		fmt.Println("ReadChannel: ", v)
	}
}

func writeChannel(i int,c chan int){

	fmt.Println("write:",i)
	c <- i
}

func main() {
//	p := make(chan int)
//	go Producer(p)
//	Customer(p)
	c := make(chan int)
	go readChannel(c)
	for i:=0;i<15;i++ {
		writeChannel(i,c)
//		time.Sleep(time.Second)
	}
}