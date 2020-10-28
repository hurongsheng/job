package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Microsecond * 3)
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	fmt.Println("i: 1111")
	wg.Wait()
}

/**
i:  9 //todo 为什么第一个是9
i:  0
i:  1
i:  2
i:  3
i:  4
i:  5
i:  6
i:  7
i:  8
*/
