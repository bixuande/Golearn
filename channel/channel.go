package main

import (
	"fmt"
	"time"
)

func worker(id int ,c chan int)  {
	for  {
		//在这里做判断，判断还有没有值传入
		n,ok := <-c
		if !ok {
     break
		}
		fmt.Printf("Worker %d received %c\n",id,n)
	}
}
func createWoker(id int) chan<- int {
	c := make(chan int)
	go worker(id,c)
	return c
}
func chanDeamo()  {
	var channles [10]chan<- int
	for i :=0;i<10;i++ {
		//创建channel要使用make函数
		channles[i] = createWoker(i)

	}
	 //给channel发送数据
	 /***
	 需要注意的是，使用channel要使用一个goroutine来接收，不然会反锁
	  */
	 for i := 0;i<10;i++{
	 	channles[i] <- 'a'+i
	 }
	 time.Sleep(time.Millisecond)
}

func bufferedChannel()  {
	/***
	这里的3是缓存三个数据的意思
	 */
	c :=make (chan int,3)
go worker(0,c)
  c <- 'a'
	c <- 'd'
	c <- 's'
	time.Sleep(time.Millisecond)
}

func channelClose()  {
	/***
	这里的3是缓存三个数据的意思
	*/
	c :=make (chan int,3)
	go worker(0,c)
	c <- 'a'
	c <- 'd'
	c <- 's'
	/***
	这里的关闭操作，是发送方做的事情
	因为我们设置了一毫秒的时间，所以在这一毫秒中
	我们的程序，接收完了过后，继续在运行，所以接收了零值
	所以打印出了很多的空值
	所以我们进行了下一步的优化就是判断接收完毕后进行的关闭处理
	 */
	close(c)

	time.Sleep(time.Millisecond)
}

func main() {
fmt.Println("Channel as first-class citizen")
	chanDeamo()
fmt.Println("Buffered channel")
bufferedChannel()
fmt.Println("Channel  close and range")
channelClose()
}
