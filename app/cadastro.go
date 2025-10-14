package main

import "fmt"

// Função para cadastrar um novo jogador

func cadastroJogador(nome string) string {
	//no futuro iremos salvar o jogador em um banco de dados
	//por enquanto, apenas retornamos uma mensagem de sucesso
	return fmt.Sprintf("Jogador %s cadastrado com sucesso!", nome)
}





