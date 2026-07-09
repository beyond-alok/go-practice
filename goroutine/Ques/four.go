package Ques

import (
	"fmt"
	"time"
)

func Four() {
	uf := make(chan int)

	go func(uf chan int) {
		uf <- 42
	}(uf)

	time.Sleep(2*time.Millisecond)

	x := <- uf
	fmt.Println(x)
}
