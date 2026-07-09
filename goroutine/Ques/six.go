package Ques

import (
	"fmt"
	"time"
)

func double(in,out chan int) {
	x := <- in
	x = 2*x
	out <- x
}

func Six() {
	in := make(chan int)
	out := make(chan int)

	go double(in,out)

	time.Sleep(2*time.Millisecond)

	in <- 2

	x := <- out

	fmt.Println(x)
}
