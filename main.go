package main

import (
	"fmt"
	"strings"
)

func main() {
	nome := "Conferência Go"
	nomeDoUsuario := ""
	sobrenome := ""
	email := ""
	ticketsReservados := 0
	ticketsSobrando := 100
	const ticketsConferencia = 100
	pessoasComReserva := []string{}

	apresentacao(nome)

	for {
		fmt.Printf("\nTemos um total de %v tickets e %v tickets disponíveis para compra. Garanta logo sua vaga!", ticketsConferencia, ticketsSobrando)

		// Pega nome do usuário e guarda na variável nomeDoUsuario
		fmt.Println("\nDigite seu nome: ")
		fmt.Scan(&nomeDoUsuario)

		// sobrenome
		fmt.Println("Digite seu sobrenome: ")
		fmt.Scan(&sobrenome)

		// email
		fmt.Println("Digite seu email: ")
		fmt.Scan(&email)

		// numero de tickets que a pessoa deseja reservar
		fmt.Println("Digite quantos tickets deseja reservar: ")
		fmt.Scan(&ticketsReservados)

		if ticketsReservados > ticketsSobrando {
			fmt.Printf("\nSó há %v tickets disponíveis, dessa maneira você não pode reservar %v tickets.\n", ticketsSobrando, ticketsReservados)
			continue
		}

		ticketsSobrando = ticketsSobrando - ticketsReservados

		pessoasComReserva = append(pessoasComReserva, nomeDoUsuario+" "+sobrenome)

		nomes := pegaNomes(pessoasComReserva)

		fmt.Printf("Lista de pessoas com reservas já garantidas %v\n", nomes)

		if ticketsSobrando == 0 {
			fmt.Println("As entradas da nossa super conferência Go se esgotaram")
			break
		}
	}
}

func apresentacao(nome string) {
	fmt.Println("Bem vindo à", nome, "a melhor conferência do mundo sobre Go")
}

func pegaNomes(pessoasComReserva []string) []string {
	nomes := []string{}
	for _, pessoaComReserva := range pessoasComReserva {
		var nomesCompletos = strings.Fields(pessoaComReserva)
		var nome = nomesCompletos[0]

		nomes = append(nomes, nome)
	}

	return nomes
}
