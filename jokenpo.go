package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	jogador1, jogador2 := cadastrarJogadores()
	jogadaJogador1, jogadaJogador2 := cadastrarJogadas(jogador1, jogador2)
	nomeGanhador := defineVencedor(jogadaJogador1, jogadaJogador2, jogador1, jogador2)
	exibeVencedor(nomeGanhador)
}

func exibeVencedor(nomeGanhador string) {
	fmt.Println("======== Resultado ========")
	if nomeGanhador != "Empate" {
		fmt.Println("Vencedor da partida:", nomeGanhador, "!")
	} else {
		fmt.Println("O jogo empatou!")
	}
}

func defineVencedor(jogadaJogador1 string, jogadaJogador2 string, jogador1 string, jogador2 string) string {
	if jogadaJogador1 == "PEDRA" && jogadaJogador2 == "TESOURA" {
		return jogador1
	} else if jogadaJogador1 == "PEDRA" && jogadaJogador2 == "PAPEL" {
		return jogador2
	} else if jogadaJogador1 == "PAPEL" && jogadaJogador2 == "PEDRA" {
		return jogador1
	} else if jogadaJogador1 == "PAPEL" && jogadaJogador2 == "TESOURA" {
		return jogador2
	} else if jogadaJogador1 == "TESOURA" && jogadaJogador2 == "PAPEL" {
		return jogador1
	} else if jogadaJogador1 == "TESOURA" && jogadaJogador2 == "PEDRA" {
		return jogador2
	} else if jogadaJogador2 == "PEDRA" && jogadaJogador1 == "TESOURA" {
		return jogador2
	} else if jogadaJogador2 == "PAPEL" && jogadaJogador1 == "PEDRA" {
		return jogador2
	} else if jogadaJogador2 == "TESOURA" && jogadaJogador1 == "PAPEL" {
		return jogador2
	} else {
		return "Empate"
	}
}

func cadastrarJogadas(jogador1 string, jogador2 string) (string, string) {
	jogadaJogador1 := criaJogada(jogador1)
	jogadaJogador2 := criaJogada(jogador2)

	return jogadaJogador1, jogadaJogador2
}

func criaJogada(jogador string) string {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("======== Cria jogada ======== ")
		fmt.Println("Defina sua jogada,", jogador, ":")
		fmt.Println("Pedra")
		fmt.Println("Papel")
		fmt.Println("Tesoura")
		scanner.Scan()
		jogada := scanner.Text()

		jogada = strings.ToUpper(jogada)
		jogada = strings.TrimSpace(jogada)
		if jogada == "" || (jogada != "PEDRA" && jogada != "PAPEL" && jogada != "TESOURA") {
			fmt.Println("Opcao incorreta, tente nomente")
			continue
		} else {
			return jogada
		}
	}
}

func cadastrarJogadores() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite o nome do jogador 1: ")
	scanner.Scan()
	jogador1 := scanner.Text()
	fmt.Println("Digite o nome do jogador 2: ")
	scanner.Scan()
	jogador2 := scanner.Text()

	return jogador1, jogador2
}
