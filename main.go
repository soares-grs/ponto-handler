package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Permissão de usuário padrão do SO
const defaultPermission = 0666

func main() {
	showMenu()

	command := listenCommand()

	switch command {
	case 1:
		checkIn("Entrada Expediente")
	case 2:
		checkIn("Saída Expediente")
	case 3:
		checkIn("Entrada Intervalo")
	case 4:
		checkIn("Saída Intervalo")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Este comando não foi reconhecido :/ ")
		os.Exit(-1)
	}
}

func showMenu() {
	fmt.Println("1 - Entrada Expediente")
	fmt.Println("2 - Saída Expediente")
	fmt.Println("3 - Entrada Intervalo")
	fmt.Println("4 - Saída Intervalo")
	fmt.Println("0 - Sair do Programa")

}

func listenCommand() int {
	var listenedCommand int
	fmt.Scan(&listenedCommand)

	return listenedCommand
}
func verifyCheckInExists(checkInFile string, checkInWords ...string) bool {
	file, err := os.Open(checkInFile)
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if containsAllWords(line, checkInWords...) {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ocorreu um erro ao ler o arquivo:", err)
	}

	return false
}

func containsAllWords(line string, words ...string) bool {
	for _, word := range words {
		if !strings.Contains(line, word) {
			return false
		}
	}
	return true
}
func checkIn(checkInType string) {

	var filePath string = "checkin-" + time.Now().Format("02-01-2006") + ".txt"
	var checkInToRegist string = time.Now().Format("02/01/2006 15:04") + " - " + checkInType + "\n"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, defaultPermission)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if verifyCheckInExists(filePath, checkInType) == true {
		fmt.Println("O ponto já foi registrado")
	} else {
		file.WriteString(checkInToRegist)
		fmt.Println("Ponto registrado com sucesso!")
	}

	file.Close()
}
