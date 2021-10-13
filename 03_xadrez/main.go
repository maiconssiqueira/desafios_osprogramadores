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

var (
	peao   int
	bispo  int
	cavalo int
	torre  int
	rainha int
	rei    int
)

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

	fmt.Println("Tabuleiro de Xadrez")
	for _, v := range tabuleiro {

		for _, v := range v {
			if v == 1 {
				peao++
			} else if v == 2 {
				bispo++
			} else if v == 3 {
				cavalo++
			} else if v == 4 {
				torre++
			} else if v == 5 {
				rainha++
			} else if v == 6 {
				rei++
			}
		}
		fmt.Printf("%v \n", v)
	}
	fmt.Printf("\nO total de peoes eh:\t %v \n", peao)
	fmt.Printf("O total de bispos eh:\t %v \n", bispo)
	fmt.Printf("O total de cavalos eh:\t %v \n", cavalo)
	fmt.Printf("O total de torres eh:\t %v \n", torre)
	fmt.Printf("O total de rainhas eh:\t %v \n", rainha)
	fmt.Printf("O total de reis eh:\t %v \n", rei)
}
