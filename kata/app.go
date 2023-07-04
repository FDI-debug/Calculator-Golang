package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumeric = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func ClassicToRoman(num int) string {
	conversions := []struct {
		value int
		digit string
	}{
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for num >= conversion.value {
			roman.WriteString(conversion.digit)
			num -= conversion.value
		}
	}

	return roman.String()
}

func Contains(a [10]string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func IdentifyClassicArray() [10]string {
	return [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
}

func IdentifyRomanArray() [10]string {
	return [...]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
}

func MathOperation(elements []string, firstOperand int, secondOperand int) string {
	if elements[1] == "+" {
		return strconv.Itoa(firstOperand + secondOperand)
	} else if elements[1] == "-" {
		return strconv.Itoa(firstOperand - secondOperand)
	} else if elements[1] == "*" {
		return strconv.Itoa(firstOperand * secondOperand)
	} else if elements[1] == "/" {
		return strconv.Itoa(firstOperand / secondOperand)
	}
	return "Ошибка операции, получен неизвестный знак."
}

func Calculator(elements []string) string {
	if len(elements) != 3 {
		return "Ошибка операции. Строка не является математическим выражением, либо не удовлетворяет заданию."
	}
	classicNum := IdentifyClassicArray()
	romanNum := IdentifyRomanArray()
	if Contains(classicNum, elements[0]) && Contains(classicNum, elements[2]) {
		firstOperand, _ := strconv.Atoi(elements[0])
		secondOperand, _ := strconv.Atoi(elements[2])
		return MathOperation(elements, firstOperand, secondOperand)
	}
	if Contains(romanNum, elements[0]) && Contains(romanNum, elements[2]) {
		firstOperand := romanNumeric[elements[0]]
		secondOperand := romanNumeric[elements[2]]
		classicResult, _ := strconv.Atoi(MathOperation(elements, firstOperand, secondOperand))
		if classicResult <= 0 {
			return "Ошибка операции. В римской системе есть только положительные числа."
		}
		return ClassicToRoman(classicResult)
	} else {
		return "Ошибка операции. Используется одновременно разные системы счисления, либо строка не удовлетворяет заданию."
	}
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	fmt.Println(Calculator(strings.Fields(input)))
}
