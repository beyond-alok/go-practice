package Ques

import (
	"fmt"
	"sync"
	"time"
)

func findMax(ss []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	var max int

	for _, v := range ss {
		if v > max {
			max = v
		}

		ch <- v
	}
}

func Fifteen() {

	arr := []int{2, 3, 5, 1, 3, 2, 3, 5, 3, 4, 4, 5, 4, 6, 1, 3, 1, 3, 4, 5, 3, 2, 2, 4}
	n := 4

	ch := make(chan int)
	var wg sync.WaitGroup
	start := 0

	for n <= len(arr) {
		wg.Add(1)
		if n > len(arr) {
			n = len(arr)
		}
		go findMax(arr[start:n], ch, &wg)
		start = n
		n += 4
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var max int
	for v := range ch {
		if v > max {
			max = v
		}
	}

	fmt.Println("max : ", max)

	time.Sleep(5 * time.Millisecond)
}
