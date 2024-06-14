package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToIntMap = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
	"XXI": 21, "XXII": 22, "XXIII": 23, "XXIV": 24, "XXV": 25, "XXVI": 26, "XXVII": 27, "XXVIII": 28, "XXIX": 29, "XXX": 30,
	"XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37, "XXXVIII": 38, "XXXIX": 39, "XL": 40,
	"XLI": 41, "XLII": 42, "XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50,
	"LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55, "LVI": 56, "LVII": 57, "LVIII": 58, "LIX": 59, "LX": 60,
	"LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65, "LXVI": 66, "LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70,
	"LXXI": 71, "LXXII": 72, "LXXIII": 73, "LXXIV": 74, "LXXV": 75, "LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79, "LXXX": 80,
	"LXXXI": 81, "LXXXII": 82, "LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85, "LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90,
	"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95, "XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100,
}

var intToRomanMap = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
	"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
	"XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
	"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L",
	"LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
	"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
	"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
	"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
	"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C",
}

func romanToInt(roman string) (int, error) {
	if val, ok := romanToIntMap[roman]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("Недопустимая римская цифра")
}

func intToRoman(num int) (string, error) {
	if num <= 0 || num >= len(intToRomanMap) {
		return "", fmt.Errorf("Полученная римская цифра выходит за пределы")
	}
	return intToRomanMap[num], nil
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("неподдерживаемый оператор")
	}
}

func isRoman(input string) bool {
	_, exists := romanToIntMap[input]
	return exists
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например, 3 + 4 или V * VI):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	tokens := strings.Split(input, " ")
	if len(tokens) != 3 {
		panic("Недопустимый формат выражения")
	}

	aStr, operator, bStr := tokens[0], tokens[1], tokens[2]

	var a, b int
	var err error
	var isRomanNumerals bool

	if isRoman(aStr) && isRoman(bStr) {
		a, err = romanToInt(aStr)
		if err != nil {
			panic("Недопустимая римская цифра: " + aStr)
		}
		b, err = romanToInt(bStr)
		if err != nil {
			panic("Недопустимая римская цифра: " + bStr)
		}
		isRomanNumerals = true
	} else if !isRoman(aStr) && !isRoman(bStr) {
		a, err = strconv.Atoi(aStr)
		if err != nil {
			panic("Недопустимое целое: " + aStr)
		}
		b, err = strconv.Atoi(bStr)
		if err != nil {
			panic("Недопустимое целое: " + bStr)
		}
		if a < 1 || a > 10 || b < 1 || b > 10 {
			panic("числа должны быть от 1 до 10")
		}
		isRomanNumerals = false
	} else {
		panic("смешанные системы числительных не допускаются")
	}

	result := calculate(a, b, operator)

	if isRomanNumerals {
		if result < 1 {
			panic("результирующая римская цифра меньше I не допускается")
		}
		romanResult, err := intToRoman(result)
		if err != nil {
			panic(err)
		}
		fmt.Println("Вывод:", romanResult)
	} else {
		fmt.Println("Вывод:", result)
	}
}
