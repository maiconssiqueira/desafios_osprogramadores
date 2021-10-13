package main

import (
	"fmt"
)

var listaPecas = map[int]string{
	1: "peao",
	2: "bispo",
	3: "cavalo",
	4: "torre",
	5: "rainha",
	6: "rei",
}

var tabuleiro = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
}

func main() {

	fmt.Println(listaPecas[3])

	for i := 0; i < 8; i++ {
		tabuleiro[0][0] = 4
		tabuleiro[0][1] = 3
		tabuleiro[0][2] = 2
		tabuleiro[0][3] = 5
		tabuleiro[0][4] = 6
		tabuleiro[0][5] = 2
		tabuleiro[0][6] = 3
		tabuleiro[0][7] = 4

		tabuleiro[1][i] = 1
		tabuleiro[6][i] = 1

		tabuleiro[7][0] = 4
		tabuleiro[7][1] = 3
		tabuleiro[7][2] = 2
		tabuleiro[7][3] = 5
		tabuleiro[7][4] = 6
		tabuleiro[7][5] = 2
		tabuleiro[7][6] = 3
		tabuleiro[7][7] = 4
	}

	for i, v := range tabuleiro {

		fmt.Println(i, v)
	}
}
