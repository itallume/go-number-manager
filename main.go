package main

import("fmt")



func main(){
	var number int
	var numbers = []int{}
	fmt.Println("===== Bem vindo(a) ao Gerenciador de Números 2000 ====")
	
	for{
		var option int
		fmt.Println("Menu de escolha: ")
		fmt.Println(
			"1) Adicionar um números" +
			"2) Listar números" +
			"3) Remover um número" +
			"4) Estatísticas " +
			"5) Divisão entre 2 números" +
			"6) Limpar lista" +
			"0) sair")
		fmt.Print("Escolha uma das opções acima (entre 0 e 6): ")
		fmt.Scan(&option)
		
		switch option{
		case 1:


		}
		fmt.Print("Digite um numero inteiro: ")
		fmt.Scanln(&number)
		numbers = append(numbers, number)

	}

	func addNumber(number int, numbers *[]int)  {
		
	}
	
	fmt.Print(numbers)
}