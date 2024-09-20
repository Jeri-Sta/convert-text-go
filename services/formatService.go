package services

import (
	"fmt"
	"strings"
)

func formatToJava() {
	fmt.Println("Lendo arquivo e formatando para Java...")

	var replacer = strings.NewReplacer("\\n", "\n", "\\t", "\t", "\\\"", "\"")

	input := ReadFile("input.txt")

	output := replacer.Replace(input)

	WriteFile("output.txt", output)
}

func formatToJson() {
	fmt.Println("Lendo arquivo e formatando para Java...")

	var replacer = strings.NewReplacer("\n", "\\n", "\t", "\\t", "\"", "\\\"")

	input := ReadFile("input.txt")

	output := replacer.Replace(input)

	WriteFile("output.txt", output)
}
