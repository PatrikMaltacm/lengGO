package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Token struct {
	Type  string
	Value string
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Arquivo não encontrado!")
		return
	}

	fileName := os.Args[1]

	isValidExtension := strings.HasSuffix(fileName, ".lg")

	if !isValidExtension {
		fmt.Println("Extensão inválida para o arquivo!")
		return
	}

	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Falha ao ler arquivo!")
		return
	}

	var code string
	code = string(data)

	tokens := Lexer(code)

	fmt.Println(tokens)

	success, message := Parser(tokens)

	if success {
		fmt.Println(message)
	} else {
		fmt.Println(message)
	}
}

func Lexer(code string) []Token {
	var tokens []Token
	var currentWord string
	inString := false
	currentType := "IDENTIFIER"

	for _, characther := range code {
		if characther == '\n' || characther == '\r' {
			continue
		}
		if characther == '"' {
			inString = !inString
			currentType = "STRING"
			continue
		}
		if characther >= '0' && characther <= '9' && !inString {
			currentType = "NUMBER"
		}
		if characther == ';' {
			if currentWord != "" {
				newToken := Token{
					Type:  currentType,
					Value: currentWord,
				}
				tokens = append(tokens, newToken)
			}
			currentWord = ""
			currentType = "END_STATEMENT"
			newToken := Token{
				Type:  currentType,
				Value: currentWord,
			}
			tokens = append(tokens, newToken)
			currentType = "IDENTIFIER"
			continue
		}
		if characther == ' ' && !inString {
			if currentWord != "" {
				newToken := Token{
					Type:  currentType,
					Value: currentWord,
				}
				tokens = append(tokens, newToken)
				currentType = "IDENTIFIER"
				currentWord = ""
			}
		} else {
			currentWord += string(characther)
		}
	}

	if currentWord != "" {
		newToken := Token{
			Type:  currentType,
			Value: currentWord,
		}
		tokens = append(tokens, newToken)
	}

	fmt.Println(len(tokens))

	return tokens
}

func Parser(tokens []Token) (success bool, message string) {
	var currentStatement []Token

	for _, token := range tokens {
		if token.Type == "END_STATEMENT" {
			// 1. O primeiro item da currentStatement é o comando (ex: "print" ou "sum")
			// 2. Os itens do índice 1 em diante são os valores.
			if currentStatement[0].Value == "print" {
				for _, args := range currentStatement[1:] {
					fmt.Println(args.Value)
				}
			}

			if currentStatement[0].Value == "sum" {
				if len(currentStatement[1:]) >= 3 {
					return false, "Não consigo somar mais que dois numeros"
				} else if len(currentStatement[1:]) <= 1 {
					return false, "não é possivel somar menos de um numero"
				}
				var sum int
				for _, args := range currentStatement[1:] {
					num, err := strconv.Atoi(args.Value)

					if err == nil {
						sum += num
					}
				}
				fmt.Println(currentStatement[1].Value, "+", currentStatement[2].Value, "=", sum)
			}
			// Depois de executar, limpa para o próximo comando:
			currentStatement = []Token{}
			continue
		}
		currentStatement = append(currentStatement, token)
	}

	return true, "Sucesso na compilação"
}
