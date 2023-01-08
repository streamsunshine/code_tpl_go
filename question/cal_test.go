// Package question
package question

import (
	"fmt"
	"strconv"
	"testing"
)

//  (1-2+ 12 * 14 * )
func Cal(str string) int {
	num := []int{}
	symbol := []byte{}

	priorityMap := map[byte]int{
		'(': 0,
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	var calStack func() bool
	calStack = func() bool {
		fmt.Println(num)
		fmt.Println(symbol)

		symbolLen := len(symbol)
		if symbolLen == 0 {
			return false
		}
		oper := symbol[symbolLen-1]
		if oper == '(' {
			symbol = symbol[:(symbolLen - 1)]
			return false
		}

		numLen := len(num)
		right := num[numLen-1]
		left := num[numLen-2]

		switch oper {
		case '+':
			left = left + right
		case '-':
			left = left - right
		case '*':
			left = left * right
		case '/':
			left = left / right
		}

		num[numLen-2] = left
		num = num[:(numLen - 1)]
		symbol = symbol[:(symbolLen - 1)]
		return true
	}

	strBytes := []byte(str)
	for i := 0; i < len(strBytes); i++ {
		v := strBytes[i]
		switch v {
		case '(':
			symbol = append(symbol, v)
		case ')':
			for calStack() {
			}
		case '+', '-', '*', '/':
			symLen := len(symbol)
			if symLen == 0 {
				symbol = append(symbol, v)
			} else if priorityMap[symbol[symLen-1]] >= priorityMap[v] {

				for j := symLen - 1; priorityMap[symbol[j]] >= priorityMap[v]; j-- {
					if !calStack() {
						break
					}
				}
				symbol = append(symbol, v)
			} else {
				symbol = append(symbol, v)
			}
		default:
			if v >= '0' && v <= '9' {
				tmpNum := 0
				for ; i < len(strBytes) && strBytes[i] >= '0' && strBytes[i] <= '9'; i++ {
					tmp, _ := strconv.Atoi(string(strBytes[i]))
					tmpNum = tmpNum*10 + tmp
				}
				i--
				num = append(num, tmpNum)
			}
		}
	}
	for calStack() {
	}
	return num[0]
}

func TestCal(t *testing.T) {
	rs := Cal("(1 + 4) * ((4 -2) + 10)")
	fmt.Println(rs)
}
