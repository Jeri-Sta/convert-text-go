package formatService

import (
	"fmt"
	"strings"

	"br.com.jeriel/convert/fileService"
)

func FormatToJava() {
	fmt.Println("Lendo arquivo e formatando para Java...")

	var replacer = strings.NewReplacer("\\n", "\n", "\\t", "\t", "\\\"", "\"")

	input := fileService.ReadFile("input.txt")

	output := replacer.Replace(input)

	fileService.WriteFile("output.txt", output)
}

func FormatToJson() {
	fmt.Println("Lendo arquivo e formatando para Java...")

	var replacer = strings.NewReplacer("\n", "\\n", "\t", "\\t", "\"", "\\\"")

	input := fileService.ReadFile("input.txt")

	output := replacer.Replace(input)

	fileService.WriteFile("output.txt", output)
}
