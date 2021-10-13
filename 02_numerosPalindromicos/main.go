package main

import (
	"fmt"
	"strconv"
)

func reverter(revert string) string {
	result := []rune(revert)
	var res []rune

	for i := len(result) - 1; i >= 0; i-- {
		res = append(res, result[i])
	}

	return fmt.Sprintf("%v", string(res))
}

func main() {

	rangeNumbers := 10000

	for acc := 0; acc < rangeNumbers; acc++ {

		toRevert := strconv.Itoa(acc)

		if toRevert == reverter(toRevert) {
			fmt.Printf("O numero %s eh Palindromico \n", toRevert)
		}
	}
}
