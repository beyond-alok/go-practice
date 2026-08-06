package Ques

import (
	"fmt"
	"sync"
)

var index []int
var mtx sync.Mutex

func appendIndex(i int) {
	mtx.Lock()
	index = append(index, i)
	mtx.Unlock()
}

func Twelve() {

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Go(func() {appendIndex(i)})
	}

	wg.Wait()

	fmt.Println(index)

}