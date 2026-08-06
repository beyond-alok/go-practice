package Ques

import (
	"fmt"
	"sync"
	"time"
)

func Fourteen() {
	integers := []int{1, 2, 3, 4, 5, 6}
	ch := make(chan int)
	var wg sync.WaitGroup


	wg.Add(2)

	go func(half []int, ch chan int,wg *sync.WaitGroup) {
		defer wg.Done()
		val := half[0]
		for i, v := range half {
			if i == 0 {
				continue
			}
			val = v * val
		}

		ch <- val

	}(integers[:3],ch,&wg)


	go func(half []int,ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		val := half[0]
		for i, v := range half {
			if i == 0 {
				continue
			}
			val = v * val
		}
		ch <- val
	}(integers[3:6],ch,&wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	result := 1
	for v := range ch {
		result = result * v
	}

	fmt.Println("multiplication result", result)
	time.Sleep(5 * time.Millisecond)
}
