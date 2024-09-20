package main

import (
	"bufio"
	"fmt"
	"os"

	"br.com.jeriel/convert/services"
)

func main() {
	services.CreateFile("input.txt")
	services.CreateFile("output.txt")

	fmt.Println("Olá!")

	scanner := bufio.NewScanner(os.Stdin)

loop:
	for {
		fmt.Print(`Qual opção de formatação você deseja realizar?
				1 - Formatar para JSON
				2 - Formatar para código Java
				0 - Fechar programa
				`)

		scanner.Scan()
		var input = scanner.Text()

		if scanner.Err() != nil {
			println("Falha ao ler a entrada.")
			continue
		}

		switch input {
		case "1":
			fmt.Println("Você escolheu a opção 1 - Formatar para JSON")
			services.FormatToJson()
		case "2":
			fmt.Println("Você escolheu a opção 2 - Formatar para código Java")
			services.FormatToJava()
		case "0":
			fmt.Println("Você escolheu a opção 0 - Fechar programa")
			break loop
		default:
			fmt.Println("Opção inválida")
		}
	}

}
