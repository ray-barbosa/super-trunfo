    package main

    import "fmt"

	func cadastroJogador(nome string) string {
		return "Jogador " + nome + " cadastrado com sucesso!"
	}

	func somar(num1 int, num2 int) int {
		return num1 + num2
	}

	func subtrair(num1 int, num2 int) int {
		return num1 - num2
	}

    func main() {
        fmt.Println(cadastroJogador("Alice"))
		fmt.Println(cadastroJogador("Bob"))
		fmt.Println(somar(5, 10))
		fmt.Println(subtrair(10, 5))
	   }






