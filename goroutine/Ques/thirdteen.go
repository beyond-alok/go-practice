package Ques

import (
	"fmt"
	"sync"
	"time"
)

func sum(s []int,ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for _,v := range s {
		sum = sum + v
	}
	ch <- sum
}

func Thirdteen() {

	var s = []int{1,2,3,4,5,6,7,8,9,10}
	ch := make(chan int)
	var wg sync.WaitGroup
	var sm int

	mid := len(s)/2

	wg.Add(2)
	go sum(s[:mid+1],ch,&wg)
	go sum(s[mid+1:],ch,&wg)

	go func () {
		wg.Wait()
		close(ch)
	}()

	for s := range ch {
		sm = sm + s
	}

	fmt.Println(sm)
	time.Sleep(5*time.Millisecond)

}