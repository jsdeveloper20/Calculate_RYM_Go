/*
package main

import (
	"bufio"
	"fmt"
	_ "io"
	"os"
	"strconv"
	"strings"
)

func conv() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}



func main() {

	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}

}
*/

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
}

var intToRomanMap = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
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
