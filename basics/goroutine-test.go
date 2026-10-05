package basic

import (
	"fmt"
)
	
func Test() {

	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
	)

	fmt.Println(Sunday,Monday,Tuesday)
}

