package Ques

import (
	"fmt"
	"time"
)

func Five() {
	first := make(chan string)

	second := make(chan string)

	go func(first,second chan string){
		first <- "ping"
		fmt.Println(<-second)
	}(first,second) 

	time.Sleep(2*time.Millisecond)
	fmt.Println(<- first)

	second <- "pong"
	
}
