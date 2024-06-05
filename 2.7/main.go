package main

import (
	"fmt"
	"math"
)

func main() {
	var catheret1, catheret2 float64
	_, _ = fmt.Scan(&catheret1, &catheret2)
	hypotenuse := math.Sqrt(math.Pow(catheret1, 2) + math.Pow(catheret2, 2))
	fmt.Println(hypotenuse)
}
