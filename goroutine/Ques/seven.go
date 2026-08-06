package Ques

import "fmt"

func Seven() {

	bf := make(chan int,3)

	bf <- 1
	bf <- 2
	bf <- 3

	for v := range bf {
		fmt.Println(v)
	}

}
