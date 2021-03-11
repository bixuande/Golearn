package main

import (
	"fmt"
	"time"
)

func main() {
	for i:= 0;i<1000;i++{

		go func(ii int) {
			for {
				fmt.Printf("hello from " +
					"gouroutine %d\n",i)
				//这里打印输出还是抢占式的，是因为Printf里面是io
				//iO是会跳出
			}}(i)

	}
	//因为我们是并发执行的所以最里面的for还没有来得及输出，所有的goroutine
	//就被结束了
  time.Sleep(time.Minute)
}
