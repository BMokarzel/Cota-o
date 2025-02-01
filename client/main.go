package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Escolha uma moeda para consultar a cotação:")
	fmt.Println("1. USD")
	fmt.Println("2. EUR")
	fmt.Println("3. BTC")

	// Leitura da opção do usuário
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite o número correspondente à moeda: ")
	scanner.Scan()
	option := scanner.Text()

	var currency string
	switch option {
	case "1":
		currency = "USD"
	case "2":
		currency = "EUR"
	case "3":
		currency = "BTC"
	default:
		fmt.Println("Opção inválida!")
		return
	}

	url := "http://localhost:8080/" + currency

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		scanner := bufio.NewScanner(response.Body)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	} else {
		fmt.Printf("Erro: código HTTP %d\n", response.StatusCode)
	}
}
