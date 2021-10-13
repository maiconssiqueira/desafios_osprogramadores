package main

import (
	"fmt"
)

func main() {

	rangeNumbers := 10000
	acc := 0

	for acc < rangeNumbers {
		acc += 1

		if acc%2 == 1 {
			fmt.Printf("O numero %v eh primo \n", acc)
		}
	}

}
