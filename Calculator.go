package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func main() {

	result := calculator("X + V")
	fmt.Println(result)
}

func calculator(str string) string {
	result := ""
	res := 0
	arabicNum := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	romanNum := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	numArr := strings.Split(str, " ")
	if len(numArr) != 3 {
		err := errors.New("Формат математической операции не удовлетворяет заданию — два операнда и один оператор, через пробел")
		log.Fatal(err)
	}

	num1, znak, num2 := numArr[0], numArr[1], numArr[2]

	if slices.Contains(arabicNum, num1) && slices.Contains(romanNum, num2) || slices.Contains(romanNum, num1) && slices.Contains(arabicNum, num2) {
		err := errors.New("numbers must be the same format")
		log.Fatal(err)
	}

	if (!slices.Contains(arabicNum, num1) || !slices.Contains(arabicNum, num2)) && (!slices.Contains(romanNum, num1) || !slices.Contains(romanNum, num2)) {
		err := errors.New("wrong numbers")
		log.Fatal(err)
	}

	// arabic numbers calculating
	if slices.Contains(arabicNum, num1) && slices.Contains(arabicNum, num2) {
		oper1, _ := strconv.Atoi(num1)

		oper2, _ := strconv.Atoi(num2)

		switch znak {
		case "+":
			res = oper1 + oper2
		case "-":
			res = oper1 - oper2
		case "/":
			res = oper1 / oper2
		case "*":
			res = oper1 * oper2
		default:
			err := errors.New("invalid arithmetic sign")
			log.Fatal(err)
		}
		result = strconv.Itoa(res)

		// roman numbers calculating
	} else if slices.Contains(romanNum, num1) && slices.Contains(romanNum, num2) {
		var romOper1 int
		var romOper2 int

		switch num1 {
		case "I":
			romOper1 = 1
		case "II":
			romOper1 = 2
		case "III":
			romOper1 = 3
		case "IV":
			romOper1 = 4
		case "V":
			romOper1 = 5
		case "VI":
			romOper1 = 6
		case "VII":
			romOper1 = 7
		case "VIII":
			romOper1 = 8
		case "IX":
			romOper1 = 9
		case "X":
			romOper1 = 10
		default:
			err := errors.New("invalid arithmetic sign")
			log.Fatal(err)
		}

		switch num2 {
		case "I":
			romOper2 = 1
		case "II":
			romOper2 = 2
		case "III":
			romOper2 = 3
		case "IV":
			romOper2 = 4
		case "V":
			romOper2 = 5
		case "VI":
			romOper2 = 6
		case "VII":
			romOper2 = 7
		case "VIII":
			romOper2 = 8
		case "IX":
			romOper2 = 9
		case "X":
			romOper2 = 10
		default:
			err := errors.New("invalid arithmetic sign")
			log.Fatal(err)
		}

		switch znak {
		case "+":
			res = romOper1 + romOper2
		case "-":
			res = romOper1 - romOper2
		case "/":
			res = romOper1 / romOper2
		case "*":
			res = romOper1 * romOper2
		default:
			err := errors.New("invalid arithmetic sign")
			log.Fatal(err)
		}
		if res <= 0 {
			err := errors.New("there are no zero or negative numbers in Roman numerals")
			log.Fatal(err)
		}

		result = integerToRoman(res)

	}

	return result
}

// arabic to roman numbers conversion
func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}
