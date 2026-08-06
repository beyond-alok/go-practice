package Ques

import (
	"fmt"
	"sync"
)

func squareIndex(i int) {
		fmt.Printf("square of index %v is %v\n",i,i*i)
	}

func Eleven() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Go(func() {squareIndex(i)})
	}
	wg.Wait()
}