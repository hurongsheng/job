package main

import (
	"fmt"
	"sync"
	"time"
)

func main2() {
	chanConcurrent()
}
func chanConcurrent() {
	ch := make(chan int, 1)
	go produceChan(ch, "a")
	go produceChan(ch, "b")
	consumerChan(ch)
}
func consumerChan(ch chan int) {
	i := 0
	for {
		select {
		case <-time.After(time.Millisecond * 100):
			ch <- i
			i++
			ch <- i
			i++
		}
	}
}
func produceChan(ch chan int, name string) {
	for {
		select {
		case c := <-ch:
			fmt.Printf("%v get %v\n", name, c)
		}
	}
}

//下面的迭代会有什么问题？
type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	// ch := make(chan interface{}) // 解除注释看看！
	ch := make(chan interface{},1)
	go func() {
		set.RLock()

		for elem, value := range set.s {
			ch <- elem
			fmt.Println("Iter:", elem, value)
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func main() {

	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	ch := th.Iter()
	time.Sleep(time.Second)
	fmt.Printf("%s%v\n", "ch", <-ch)
	fmt.Printf("%s%v\n", "ch", <-ch)
	fmt.Printf("%s%v\n", "ch", <-ch)
}
