package main

import (
	"fmt"
	"math/rand"
	"time"
)

/***
制作一个消息的生成器
 */

func msgGen(name string) chan string{
	c := make(chan string)
      go func() {
      	i:=0
		  for  {
		  	time.Sleep(time.Duration(rand.Intn(2000))*time.Millisecond)
			  c <- fmt.Sprintf("serveice %s message %d",name,i)
			  i++
		  }
	  }()
return c

}

//优化的时候，直接做一个循环。就是可以接收不确定个的线程，当不知道有多少个
//并发任务的时候，我们使用这个前去解决
func fanIn(c1, c2 chan string)chan string {
	c:= make(chan string)
	go func(){
		for  {
			c <- <-c1

		}
	}()
	go func(){
		for  {
			c <- <-c2

		}
	}()
	return c
}

//知道有多少个任务的时候使用这个
func fanInbySelect(c1,c2 chan string)chan string{
	c:=make(chan string)
	go func() {
		select {
		case m:= <-c1:
			c<- m
		case m:= <-c2:
			c<- m
		}
	}()
	return c
}
/***
下面两个都是超时等待，非阻塞等待
 */
func nonBlockingWait(c chan string) (string,bool)  {
	select{
	case m := <- c:
	return m ,true
	default:
		return "",false
	}
}

func timeoutWait(c chan string,timeout time.Duration)(string ,bool)  {

	select {
	case m:=<-c:
		return m,true
	case  <-time.After(timeout):
		return "",false
	}
	
}
/***
这里是同时等待两个服务，也就是两个服务同时进行
 */
func main() {
 m1 := msgGen("ce1")
	m2 := msgGen("ce2")

	for  {
		fmt.Println(<-m1)
		if m,ok := nonBlockingWait(m2);ok{
			fmt.Println(m)
		}else {
			fmt.Println("no message from service2")
		}
	}
}
