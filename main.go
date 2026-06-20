package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Digite seu código")
	var code string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		code = scanner.Text()
	}

	command, tokens := Lexer(code)

	success, message := Parser(command, tokens)

	if success {
		fmt.Println(message)
	} else {
		fmt.Println(message)
	}
}

func Lexer(code string) (command string, values []string) {
	tokens := strings.Split(code, " ")

	return tokens[0], tokens[1:]
}

func Parser(command string, tokens []string) (success bool, message string) {
	if command == "print" {
		text := strings.Join(tokens, " ")
		text = strings.ReplaceAll(text, `"`, "")

		fmt.Println(text)
	}

	if command == "sum" {
		if len(tokens) >= 3 {
			return false, "Não consigo somar mais que dois numeros"
		} else if len(tokens) <= 1 {
			return false, "não é possivel somar menos de um numero"
		}
		var sum int
		for _, token := range tokens {
			num, err := strconv.Atoi(token)

			if err == nil {
				sum += num
			}
		}
		fmt.Println(tokens[0], "+", tokens[1], "=", sum)
	}

	return true, "Sucesso na compilação"
}
