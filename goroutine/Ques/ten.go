package Ques

import (
	"fmt"
	"sync"
)

func Ten() {
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fmt.Println("goroutine",idx)
		}(i)
	}
	wg.Wait()
}
