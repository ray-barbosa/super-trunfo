package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type Carta struct {
	estado                    rune
	codigo                    string
	nomeCidade                string
	populacao                 int
	area                      float64
	pib                       float64
	quantidadePontosTuristicos int
}

func cadastrarCarta(cardDetails Carta) Carta {
	return cardDetails
}

func boasVindas() {
	fmt.Println("Bem-vindo ao sistema de cadastro de cartas!")
	fmt.Println("------------------------------------------")
}

func exibirCarta(card Carta) {
	fmt.Printf("Estado: %c\n", card.estado)
	fmt.Printf("Código: %s\n", card.codigo)
	fmt.Printf("Nome da Cidade: %s\n", card.nomeCidade)
	fmt.Printf("População: %d\n", card.populacao)
	fmt.Printf("Área: %.2f km²\n", card.area)
	fmt.Printf("PIB: R$ %.2f\n", card.pib)
	fmt.Printf("Número de Pontos Turísticos: %d\n", card.quantidadePontosTuristicos)
	fmt.Println("------------------------------------------")
}

func main() {
	boasVindas()

	var cartas []Carta
	var continuar string

	reader := bufio.NewReader(os.Stdin)

	for {
		var card Carta

		fmt.Println("Digite o código do estado (uma letra de 'A' a 'H'):")

		estadoInput, _ := reader.ReadString('\n')
        card.estado = rune(estadoInput[0])

		fmt.Println("Digite o código da carta (de 01 a 04):")
		
		card.codigo, _ = reader.ReadString('\n')
		card.codigo = strings.TrimSpace(card.codigo)

		fmt.Println("Digite o nome da cidade:")
		card.nomeCidade, _ = reader.ReadString('\n')
		card.nomeCidade = strings.TrimSpace(card.nomeCidade)

		fmt.Println("Digite o número de habitantes:")
		fmt.Scanf("%d\n", &card.populacao)


		fmt.Println("Digite a área (em km²):")
		fmt.Scanf("%f\n", &card.area)

		fmt.Println("Digite o PIB (em R$):")
		fmt.Scanf("%f\n", &card.pib)

		fmt.Println("Digite a quantidade de pontos turísticos:")
		fmt.Scanf("%d\n", &card.quantidadePontosTuristicos)

		cartas = append(cartas, cadastrarCarta(card))
		exibirCarta(card)

		fmt.Print("Deseja cadastrar outra carta? (s/n): ")
		fmt.Scanln(&continuar)
		if strings.ToLower(continuar) != "s" {
			break
		}
		fmt.Println()
		fmt.Printf("Carta %d:\n", len(cartas))
		exibirCarta(cartas[len(cartas)-1])
	}

		fmt.Printf("\nTotal de cartas cadastradas: %d\n", len(cartas))
		fmt.Printf("Cartas cadastradas:\n")
		for i, carta := range cartas {
			fmt.Printf("Carta %d:\n", i+1)
			exibirCarta(carta)
		}
	}


