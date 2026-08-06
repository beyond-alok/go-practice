package Ques

import "fmt"

func Eight() {

	bf := make(chan string, 2)

	bf <- "foo"
	bf <- "bar"
	close(bf)

	fmt.Printf("length: %v, capacity : %v",len(bf), cap(bf))

	for s := range bf {
		fmt.Printf("\n%v",s)
	}

	
}
