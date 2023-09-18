package main

import (
	"fmt"
	"strconv"
)

var result int
var romanresult string

func calculate(num1, num2, operator string) int {
	var x, y int

	x, _ = strconv.Atoi(num1)
	y, _ = strconv.Atoi(num2)

	switch operator {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		result = x / y
	}

	return result
}

func RomanToArabic(arabic [10]string, roman [10]string, rome string) string {

	var ArabRome string

	for i := 0; i < len(roman); i++ {
		if rome == roman[i] {
			ArabRome = arabic[i]
		}
	}
	return ArabRome
}

func ArabicToRoman(Result int) string {
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

	for _, conversion := range conversions {
		for Result >= conversion.value {
			romanresult += conversion.digit
			Result -= conversion.value
		}
	}
	return romanresult
}

func checkingoperands(arabic [10]string, roman [10]string, num string) bool {

	for _, elem := range arabic {
		if elem == num {
			return true
		}
	}

	for _, elem := range roman {
		if elem == num {
			return true
		}
	}

	return false
}

func exit(num1, num2, operator, check string) {
	var num1rom, num2rom string
	var flag int

	//Arrays of valid values
	arabic := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	roman := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	//Function for checking the number of operands
	if check != "" {
		fmt.Print("Результат: Операция должна состоять из двух операндов\n")
		return
	}

	//The function of checking the correct operator input
	if operator == "+" || operator == "-" || operator == "*" || operator == "/" {
		flag = 1
	}
	if flag != 1 || num1 == "" || num2 == "" || operator == "" {
		fmt.Print("Результат: Строка не является математической операцией\n")
		return
	}

	//Error in the range of input values
	if !checkingoperands(arabic, roman, num1) || !checkingoperands(arabic, roman, num2) {
		fmt.Print("Результат: Операнд не от 1 до 10\n")
	}

	//The main function of the calculation
	for i := 0; i < len(arabic); i++ {
		for j := 0; j < len(arabic); j++ {
			if num1 == arabic[i] && num2 == arabic[j] {
				calculate(num1, num2, operator)
				fmt.Print("Результат: ", result)
			}
			if num1 == roman[i] && num2 == roman[j] {
				num1rom = RomanToArabic(arabic, roman, num1)
				num2rom = RomanToArabic(arabic, roman, num2)

				calculate(num1rom, num2rom, operator)

				if result <= 0 {
					fmt.Print("Результат: В римской системе счисления нет отрицательных чисел и нуля\n")
					break
				}

				ArabicToRoman(result)
				fmt.Print("Результат: ", romanresult)
			}
			if num1 == arabic[i] && num2 == roman[j] || num1 == roman[i] && num2 == arabic[j] {
				fmt.Print("Результат: Ошибка, системы счета не совпадают\n")
			}
		}
	}
}

func main() {
	var num1, num2, operator, check string

	fmt.Print("Введите операцию из 2 операндов через пробелы:\n")
	fmt.Scanln(&num1, &operator, &num2, &check)
	exit(num1, num2, operator, check)
}
