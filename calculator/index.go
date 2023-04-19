package calculator

import (
	"errors"
	"fmt"
	"math"
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

type Calclator struct {
	Buffer       string
	Accumulator  int
	PrevOperator Operator
	NeedClear    bool
}

type Token struct {
	Operator
	Value string
}

func NewCalculator() Calclator {
	c := Calclator{Buffer: "", Accumulator: 0, PrevOperator: Clear, NeedClear: false}

	return c
}

func (c *Calclator) AddStr(str string) (string, error) {
	token, err := parse(str)

	if err != nil {
		return "", err
	}

	switch token.Operator {
	case Number:
		// 直前に演算系のボタンが押されていたらバッファを数字で置き換え
		if c.NeedClear {
			c.Buffer = token.Value
			c.NeedClear = false
			break
		}
		if c.Buffer == "0" {
			if token.Value != "0" {
				c.Buffer = token.Value
			}
			break
		}
		c.Buffer += token.Value

	case Clear:
		c.Accumulator = 0
		c.Buffer = ""
		c.PrevOperator = Clear
		c.NeedClear = false

	case Equal:
		// TODO: = を連打したときの挙動を再現したい
		value, err := strconv.Atoi(c.Buffer)
		if err != nil {
			return "", err
		}
		result, _ := calculate(c.Accumulator, c.PrevOperator, value)
		fmt.Printf("eq result:%v prevO:%v acc:%v value:%v\n", result, c.PrevOperator, c.Accumulator, value)
		c.Accumulator = result
		c.Buffer = strconv.Itoa(result)
		c.NeedClear = true
	default:
		value, err := strconv.Atoi(c.Buffer)
		if err != nil {
			return "", err
		}

		if c.PrevOperator == Clear {
			c.Accumulator = value
		} else {
			result, _ := calculate(c.Accumulator, token.Operator, value)
			fmt.Printf("result:%v op:%v acc:%v value:%v\n", result, token.Operator, c.Accumulator, value)
			c.Accumulator = result
		}
		c.PrevOperator = token.Operator
		c.NeedClear = true
	}

	return c.Buffer, nil
}

func parse(input string) (*Token, error) {
	switch input {
	case "+":
		return &Token{
			Operator: Plus,
			Value:    "",
		}, nil

	case "-":
		return &Token{Operator: Minus, Value: ""}, nil

	case "/":
		return &Token{Operator: Divide, Value: ""}, nil

	case "*":
		return &Token{Operator: Multi, Value: ""}, nil

	case "=":
		return &Token{Operator: Equal, Value: ""}, nil

	case "C":
		return &Token{Operator: Clear, Value: ""}, nil

	default:
		r := int([]rune(input)[0])
		// code point を見て "0" から "9" かを見る
		if r >= 48 && r <= 57 {
			return &Token{Operator: Number, Value: input}, nil
		}

		return nil, errors.New("Unknown operator:" + input)
	}
}

func calculate(acc int, operator Operator, value int) (int, error) {
	switch operator {
	case Plus:
		return acc + value, nil
	case Minus:
		return acc - value, nil
	case Multi:
		return acc * value, nil
	case Divide:
		// TODO: 小数点未対応
		if value == 0 {
			return int(math.Inf(0)), errors.New("Division by zero")
		}
		return acc / value, nil
	case Equal:
		return acc, nil
	case Clear:
		return value, nil
	}

	return acc, nil
}
