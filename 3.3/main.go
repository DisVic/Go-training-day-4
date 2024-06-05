package main

import (
	"fmt"
)

func main() {
	var month uint8
	_, _ = fmt.Scan(&month)
	switch month {
	case 1:
		fmt.Println("Зима")
	case 2:
		fmt.Println("Зима")
	case 3:
		fmt.Println("Весна")
	case 4:
		fmt.Println("Весна")
	case 5:
		fmt.Println("Весна")
	case 6:
		fmt.Println("Лето")
	case 7:
		fmt.Println("Лето")
	case 8:
		fmt.Println("Лето")
	case 9:
		fmt.Println("Осень")
	case 10:
		fmt.Println("Осень")
	case 11:
		fmt.Println("Осень")
	case 12:
		fmt.Println("Зима")
	}
}
