package main

import("fmt")



func main(){
	var number int
	var numbers = []int{}
	fmt.Println("===== Bem vindo(a) ao Gerenciador de Números 2000 ====")
	
	for{
		var option int
		fmt.Println("\nMenu de escolha: ")
		fmt.Println(
			"1) Adicionar um números\n" +
			"2) Listar números\n" +
			"3) Remover um número\n" +
			"4) Estatísticas\n" +
			"5) Divisão entre 2 números\n" +
			"6) Limpar lista\n" +
			"0) sair")
		fmt.Print("Escolha uma das opções acima (entre 0 e 6): ")
		fmt.Scanln(&option)
		
		switch option{
		case 0:
			fmt.Println("\n\nTchau :(")
			return
		case 1:
			fmt.Print("Digite um numero inteiro: ")
			fmt.Scanln(&number)
			numbers = addNumber(numbers, number)
		case 2:
			fmt.Println(numbers)
		case 3:
			fmt.Println("desenvolvimento")
		case 4:
			fmt.Println("desenvolvimento")
		case 5:
			fmt.Println("desenvolvimento")
		case 6:
			numbers = []int{}
		default:
			fmt.Println("Opção invalida!")
		}
	}
}

func addNumber(numbers []int, number int) []int {
		return append(numbers, number)
}