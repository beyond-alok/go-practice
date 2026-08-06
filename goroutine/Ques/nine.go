package Ques

import (
	"fmt"
)

func Nine() {
	bf := make(chan int, 1)

	bf <- 5

	go func(bf chan int) {
		bf <- 10
		close(bf)
	}(bf)

	for i := range bf {
		fmt.Println(i)
	}

}
