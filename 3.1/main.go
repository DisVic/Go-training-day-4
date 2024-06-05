package main

import (
	"fmt"
)

func main() {
	var number int32
	_, _ = fmt.Scan(&number)
	if number == 0 {
		fmt.Println(0)
	} else if number > 0 {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}
