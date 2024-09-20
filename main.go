package main

import (
	"fmt"

	"br.com.jeriel/convert/fileService"
	"br.com.jeriel/convert/formatService"
)

func main() {
	fileService.CreateFile("input.txt")
	fileService.CreateFile("output.txt")

	fmt.Println("Olá!")

loop:
	for {
		fmt.Print(`Qual opção de formatação você deseja realizar?
				1 - Formatar para JSON
				2 - Formatar para código Java
				0 - Fechar programa
				`)

		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			fmt.Println("Você escolheu a opção 1 - Formatar para JSON")
			formatService.FormatToJson()
		case 2:
			fmt.Println("Você escolheu a opção 2 - Formatar para código Java")
			formatService.FormatToJava()
		case 0:
			fmt.Println("Você escolheu a opção 0 - Fechar programa")
			break loop
		default:
			fmt.Println("Opção inválida")
		}
	}

}
