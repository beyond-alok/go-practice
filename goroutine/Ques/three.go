package Ques

import (
	"fmt"
	"time"
)

func greet(name string) {
	fmt.Println("Hello,",name)
}


func Three() {
	go greet("Alok")

	time.Sleep(2*time.Millisecond)

}
