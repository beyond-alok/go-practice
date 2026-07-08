package Ques

import (
	"fmt"
	"time"
)


func Two() {
	 
	go func() {
		fmt.Println("goroutine A")
	}()

	go func() {
		fmt.Println("goroutine B")
	}()

	go func() {
		fmt.Println("goroutine C")
	}()

	time.Sleep(2*time.Millisecond)
}


