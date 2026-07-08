package Ques 

import (
	"fmt"
	"time"
)

func One() {
	go func(){
		fmt.Println("Hello from goroutine")
	}()
	time.Sleep(2*time.Millisecond)
}
