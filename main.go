package main

import (
	"fmt"
	"strconv"
)

func readInt(prompt string) (int, error) {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return strconv.Atoi(input)
}

func main() {
	var numbers = []int{}
	fmt.Println("===== Bem vindo(a) ao Gerenciador de Números 2000 ====")

	for {
		fmt.Println("\nMenu de escolha: ")
		fmt.Println(
			"1) Adicionar um número\n" +
				"2) Listar números\n" +
				"3) Remover um número por índice\n" +
				"4) Estatísticas\n" +
				"5) Divisão entre 2 números\n" +
				"6) Limpar lista\n" +
				"0) sair")

		option, err := readInt("Escolha uma das opções acima (entre 0 e 6): ")

		if err != nil {
			fmt.Println("\nDigite um valor válido!")
			continue
		}
		println()

		switch option {
		case 0:
			fmt.Println("\n\nTchau...        :(")
			return
		case 1:
			hanfleAddNumber(&numbers)
		case 2:
			fmt.Println(numbers)
		case 3:
			handleRemove(&numbers)
		case 4:
			handleStatistics(&numbers)
		case 5:
			handleDivision()
		case 6:
			numbers = []int{}
		default:
			fmt.Println("Opção invalida!")
		}
	}
}

func hanfleAddNumber(numbers *[]int) {
	number, err := readInt("Digite um numero inteiro: ")
	if err != nil {
		fmt.Println("Digite um valor válido!")
		return
	}
	addNumber(numbers, number)
}

func handleRemove(numbers *[]int) {
	var index int
	var removedNumber int
	index, err := readInt("Digite um índice para remover: ")
	if err != nil {
		fmt.Println("Valor inválido")
		return
	}
	removedNumber, err = removeNumber(numbers, index)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Número", removedNumber, "removido com sucesso!")
	}
}

func handleStatistics(numbers *[]int) {
	min, max, avg, err := getNumbersStatistics(numbers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Mínimo:", min, "\nMáximo:", max, "\nMédia:", avg)
}

func handleDivision() {
	var result float64
	fmt.Println("Digite os numeros para a divisão:")
	dividend, err := readFloat("Dividendo: ")
	if err != nil {
		fmt.Println("Valor inválido!")
		return
	}
	divider, err := readFloat("Divisor: ")
	if err != nil {
		fmt.Println("Valor inválido!")
		return
	}
	result, err = division(dividend, divider)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(dividend, "dividido por", divider, "é:", result)
}

func readFloat(prompt string) (float64, error) {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return strconv.ParseFloat(input, 64)
}

func addNumber(numbers *[]int, number int) {
	*numbers = append(*numbers, number)
}

func removeNumber(numbers *[]int, index int) (int, error) {
	if index < 0 || index >= len(*numbers) {
		return 0, fmt.Errorf("Índice inválido.")
	}
	s := *numbers
	removedNumber := s[index]
	*numbers = append(s[:index], s[index+1:]...)
	return removedNumber, nil
}

func getNumbersStatistics(numbers *[]int) (int, int, float64, error) {
	if len(*numbers) == 0 {
		return 0, 0, 0.0, fmt.Errorf("Insira números para Calcular as estatísticas.")
	}
	return getMin(*numbers), getMax(*numbers), getAverage(*numbers), nil
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

func division(dividend float64, divider float64) (float64, error) {
	if divider == 0 {
		return 0, fmt.Errorf("Divisor não pode ser 0")
	}
	return dividend / divider, nil
}
