package services

import (
	"fmt"
	"log"
	"os"
)

func CreateFile(fileName string) {

	_, err := os.Open(fileName)

	if os.IsNotExist(err) {
		file, err := os.Create(fileName)

		if err != nil {
			log.Fatal("Não foi possível criar o arquivo: " + err.Error())
		}

		fmt.Println("Arquivos criados com sucesso!")

		defer file.Close()
	}
}

func ReadFile(fileName string) string {
	file, err := os.ReadFile(fileName)

	if os.IsNotExist(err) {
		log.Fatal("Não foi possível abrir o arquivo.")
	}

	return string(file)
}

func WriteFile(fileName string, content string) {
	err := os.WriteFile(fileName, []byte(content), 0664)

	if err != nil {
		log.Fatal("Erro ao escrever o arquivo")
	}
}
