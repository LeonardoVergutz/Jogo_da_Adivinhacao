package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var aleatorio int
	var guess int
	var jogar int
	numeroderodadas := 1
	var matriz [][]int
	fmt.Print("Bem-vindo ao jogo da adivinhação! \n")

inicio:
	quantidade := 0
	fmt.Print("\nTente adivinhar o número entre 1 e 100. \n")
	for {
		aleatorio = rand.Intn(100) + 1
		if aleatorio >= 1 && aleatorio <= 100 {
			break
		}
	}

chute:

	fmt.Print("\nSua tentativa:")
	fmt.Scan(&guess)

	if guess > 100 || guess < 1 {
		fmt.Println("Chute inválido.")
		goto chute
	} else if guess != aleatorio {
		if aleatorio > guess {
			fmt.Print("O número é maior que ", guess, ".")
			quantidade++
		} else {
			fmt.Print("O número é menor que ", guess, ".")
			quantidade++
		}
		goto chute
	} else {
		quantidade++
		matriz = append(matriz, []int{quantidade})
		fmt.Println("\nParabéns, você acertou o número em ", quantidade, "tentativa(s).")
	}

perguntar:
	fmt.Print("Deseja jogar novamente? Digite 1 para sim e 2 para não. ")
	fmt.Scan(&jogar)

	if jogar != 1 && jogar != 2 {
		fmt.Println("Resposta inválida.")
		goto perguntar
	} else if jogar == 1 {
		numeroderodadas++
		goto inicio
	}

	rodadas := 1
	for i := 0; i < len(matriz); i++ {
		fmt.Println("\nVocê utilizou", matriz[i], "tentativas na rodada", rodadas)
		rodadas++
	}
	fmt.Println("Obrigado por jogar!")
}
