package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"strings"
)

type Args struct {
	Number1, Number2 float64
}

func main() {
	for {
		var err error
		var exit string
		var number1Str string
		var operation string
		var number2Str string

		fmt.Print("Deseja sair? (s/n): ")
		exit, err = readFromKeyboard()

		if exit == "s" {
			os.Exit(0)
		}

		fmt.Print("Digite o número 1: ")
		number1Str, err = readFromKeyboard()

		if err != nil {
			fmt.Println("Entrada para o número 1 inválida")
			fmt.Println(err)
			continue
		}

		fmt.Print("Digite a operação (+, -, *, / ou ^): ")
		operation, err = readFromKeyboard()

		if err != nil {
			fmt.Println("Entrada para a operação inválida")
			fmt.Println(err)
			continue
		}

		fmt.Print("Digite o número 2: ")
		number2Str, err = readFromKeyboard()

		if err != nil {
			fmt.Println("Entrada para o número 2 inválida")
			fmt.Println(err)
			continue
		}

		args := new(Args)

		args.Number1, err = strconv.ParseFloat(number1Str, 64)

		if err != nil {
			fmt.Println("Conversão do número 1 teve erro")
			fmt.Println(err)
			continue
		}

		args.Number2, err = strconv.ParseFloat(number2Str, 64)

		if err != nil {
			fmt.Println("Conversão do número 2 teve erro")
			fmt.Println(err)
			continue
		}

		server, err := rpc.Dial("tcp", "localhost:3000")

		if err != nil {
			log.Fatal(err)
		}

		var function string
		var result float64

		switch operation {
		case "+":
			function = "Calculator.Sum"
		case "-":
			function = "Calculator.Sub"
		case "*":
			function = "Calculator.Mult"
		case "/":
			function = "Calculator.Div"
		case "^":
			function = "Calculator.Exp"
		default:
			fmt.Println("Operação não é suportada, apenas essas: +, -, *, / ou ^")
			continue
		}

		err = server.Call(function, args, &result)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v %v %v = %v\n", args.Number1, operation, args.Number2, result)

		fmt.Print("\n\n\n")
	}
}

func readFromKeyboard() (string, error) {
	var err error
	var input string

	reader := bufio.NewReader(os.Stdin)

	input, err = reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	input = strings.Replace(input, "\r\n", "", -1)
	input = strings.Replace(input, "\n", "", -1)

	return input, nil
}
