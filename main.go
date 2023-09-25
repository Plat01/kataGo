package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type State int

const (
	Arabic State = iota
	Roman
)

var state State = Arabic

type RomanPair struct {
	Value  int
	Symbol string
}

var symbols = []RomanPair{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func romanToInt(roman string) (int, error) {
	romanToValue := map[string]int{}
	for _, pair := range symbols {
		romanToValue[pair.Symbol] = pair.Value
	}

	total := 0
	i := 0

	prevVal := 1001

	for i < len(roman) {
		if i+1 < len(roman) {
			symbol := roman[i : i+2]
			if value, found := romanToValue[symbol]; found {
				if value <= prevVal {
					total += value
					i += 2
					prevVal = value
					continue
				} else {
					return 0, errors.New(roman + " number have wrong order ")
				}
			}
		}

		symbol := string(roman[i])
		if value, found := romanToValue[symbol]; found {
			if value <= prevVal {
				total += value
				i++
			} else {
				return 0, errors.New(roman + " number have wrong order ")
			}
		} else {
			return 0, errors.New(roman + " is invalid Roman numeral")
		}
	}

	return total, nil
}

func intToRoman(num int) (string, error) {

	if num <= 0 || num > 4000 {
		return "", errors.New("Roman number can't be less than 1 and bigger than 4000 ")
	}

	roman := ""

	for num > 0 {
		for _, v := range symbols {
			for v.Value <= num {
				roman += v.Symbol
				num -= v.Value
			}
		}
	}

	return roman, nil
}

func splitString(str string) []string {
	extract := make([]string, 0)

	for _, s := range strings.Split(str, " ") {
		trimmed := strings.TrimSpace(s)
		if len(trimmed) > 0 {
			extract = append(extract, trimmed)
		}
	}

	return extract
}

func strToInt(s string) (int, error) {

	val, err := strconv.Atoi(s)
	if err != nil {
		val, err := romanToInt(s)
		state = Roman
		if err != nil {
			return 0, err
		}
		return val, err
	}

	return val, err
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("%e", err)
	}
	line = strings.ToUpper(line)

	//romNum, err := romanToInt("LXXXV")
	//if err != nil {
	//	fmt.Printf("%e", err)
	//}
	//fmt.Println(romNum)
	//
	//arNum, err := intToRoman(905)
	//if err != nil {
	//	fmt.Printf("%e", err)
	//}
	//fmt.Printf("%s", arNum)

	exp := splitString(line)
	res := 0

	if len(exp) < 2 {
		fmt.Println("Expression should consist at least 2 arguments and operation sign!")
	} else {

		val, err := strToInt(exp[0])
		if err != nil {
			fmt.Printf("%e", err)
			return
		}
		res += val

		for i := 1; i < len(exp); i += 2 {
			if len(exp) >= i+1 {
				val, err := strToInt(exp[i+1])
				if err != nil {
					fmt.Printf("%e", err)
					return
				}
				switch exp[i] {
				case "+":
					res += val
				case "-":
					res -= val
				case "*":
					res *= val
				case "/":
					res /= val
				default:
					fmt.Println("Unexpected operand " + exp[i])
					return
				}

			} else {
				fmt.Println("Expression doesn't consist last operand")
				return
			}

		}

		if state == Arabic {
			fmt.Println(res)
		} else {
			res, err := intToRoman(res)
			if err != nil {
				fmt.Printf("%e", err)
				return
			}
			fmt.Println(res)
		}
		return
	}
}
