package main

import (
	"fmt"
)

func main() {
	var number int
	var numbers = []int{}
	fmt.Println("===== Bem vindo(a) ao Gerenciador de Números 2000 ====")

	for {
		var option int
		var err error
		fmt.Println("\nMenu de escolha: ")
		fmt.Println(
			"1) Adicionar um número\n" +
				"2) Listar números\n" +
				"3) Remover um número por índice\n" +
				"4) Estatísticas\n" +
				"5) Divisão entre 2 números\n" +
				"6) Limpar lista\n" +
				"0) sair")
		fmt.Print("Escolha uma das opções acima (entre 0 e 6): ")
		fmt.Scanln(&option)
		println()

		switch option {
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
			var index int
			var removedNumber int
			fmt.Print("Digite um índice para remover: ")
			fmt.Scanln(&index)
			numbers, removedNumber, err = removeNumber(numbers, index)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Número", removedNumber, "removido com sucesso!")
			}
		case 4:
			var min int
			var max int
			var avg float64
			min, max, avg, err = getNumbersStatistics(numbers)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("Mínimo:", min, "\nMáximo:", max, "\nMédia:", avg)
		case 5:
			var divider float64
			var dividend float64
			var result float64
			fmt.Println("Digite os numeros para a divisão:")
			fmt.Print("Divisor: ")
			fmt.Scanln(&divider)
			fmt.Print("Dividendo: ")
			fmt.Scanln(&dividend)
			result, err = division(divider, dividend)
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

func removeNumber(numbers []int, index int) ([]int, int, error) {
	if index < 0 || index >= len(numbers) {
		return numbers, 0, fmt.Errorf("Índice inválido.")
	}
	removedNumber := numbers[index]
	return append(numbers[:index], numbers[index+1:]...), removedNumber, nil
}

func getNumbersStatistics(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0.0, fmt.Errorf("Insira números para Calcular as estatísticas.")
	}
	return getMin(numbers), getMax(numbers), getAverage(numbers), nil
}

func getMin(numbers []int) int {
	min := numbers[0]
	for _, v := range numbers {
		if v <= min {
			min = v
		}
	}
	return min
}

func getMax(numbers []int) int {
	max := numbers[0]
	for _, v := range numbers {
		if v >= max {
			max = v
		}
	}
	return max
}

func getAverage(numbers []int) float64 {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return float64(sum) / float64(len(numbers))
}

func division(divider float64, dividend float64) (float64, error) {

	if divider == 0 {
		return 0.0, fmt.Errorf("Divisor não pode ser 0")
	}
	return divider / dividend, nil
}
