package calculator

import (
	"strconv"
)

type Operator int

const (
	Plus Operator = iota
	Minus
	Divide
	Multi
	Equal
	Clear
	Number
)

type Token struct {
	Operator
	Value int
}

func tokenize(inputs []string) ([]Token, error) {
	var tokens []Token

	for _, v := range inputs {
		token, err := parse(v)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}

func parse(input string) (Token, error) {
	switch input {
	case "+":
		return Token{
			Operator: Plus,
			Value:    -1,
		}, nil

	case "-":
		return Token{Operator: Minus, Value: -1}, nil

	case "/":
		return Token{Operator: Divide, Value: -1}, nil

	case "*":
		return Token{Operator: Multi, Value: -1}, nil

	case "=":
		return Token{Operator: Equal, Value: -1}, nil

	case "C":
		return Token{Operator: Clear, Value: -1}, nil

	default:
		v, err := strconv.Atoi(input)
		return Token{Operator: Number, Value: v}, err
	}
}

func calculate(base int, operator Operator, value int) (int, error) {
	switch operator {
	case Plus:
		return base + value, nil
	case Minus:
		return base - value, nil
	case Multi:
		return base * value, nil
	case Divide:
		return base / value, nil
	case Equal:
		return base, nil
	}

	return base, nil
}
