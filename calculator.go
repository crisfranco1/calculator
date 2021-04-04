package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calculator struct{}

func ReadFromConsole() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func convertStringToInt(stringNumber string) int {
	number, _ := strconv.Atoi(stringNumber)
	return number
}

func (calculator) operate(input string, operator string) int {
	operands := strings.Split(input, operator)
	operand1 := convertStringToInt(operands[0])
	operand2 := convertStringToInt(operands[1])
	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2
	default:
		fmt.Println(operator, "not supported.")
		return 0
	}
}
