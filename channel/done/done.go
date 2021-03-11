package main

import (
	"fmt"
)
/***
因为我们每一次等待都是我们自己设定等待一毫秒，但是如何真正的判断，这个任务输出结束
是我们这里要做的事情
这里演示的是使用Channel等待任务的结束
go语言的作者给我们强调，通过通信来共享内存
 */
func doWorker(id int ,
	c chan int,done chan bool)  {
	for  {
		//在这里做判断，判断还有没有值传入
		n,ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %c\n",id,n)
		done <- true
	}
}
func createWoker(id int) woker {
	w := woker{
		in:make(chan int),
		done:make(chan bool),
	}
	go doWorker(id,w.in,w.done)
	return w
}

type woker struct {
	in chan int
	done chan bool
}

func chanDeamo()  {
	var wokers [10]woker
	for i :=0;i<10;i++ {
		//创建channel要使用make函数
		wokers[i] = createWoker(i)
	}
	//给channel发送数据
	/***
	需要注意的是，使用channel要使用一个goroutine来接收，不然会反锁
	*/
	for i := 0;i<10;i++{
		wokers[i].in <- 'a'+i
	}

	for _,worker := range wokers{
		<-worker.done
	}

}



func main() {

	chanDeamo()

}
