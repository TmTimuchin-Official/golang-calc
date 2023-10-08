package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Выводимые ошибки
var operator string = "Ошибка: введён неверный математический оператор, используйте только +, -, * или /"
var systems string = "Ошибка: в выражении используются разные системы счисления"
var notMath string = "Ошибка: строка не является матиматической операцией"
var tooMuch string = "Ошибка: допустимо использовать максимум 2 операнда"
var diapason string = "Ошибка: допустимый диапазон значений операндов: 1-10"
var zero string = "Ошибка: римская система счисления не предполагает числа меньше единицы"

// Проверка введена ли цифра
func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// Перевод из римских чисел в арабские
func romanToInt(s string) int {
	romanToInt := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
	}
	res := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanToInt[string(s[i])] < romanToInt[string(s[i+1])] {
			res -= romanToInt[string(s[i])]
		} else {
			res += romanToInt[string(s[i])]
		}
	}
	return res
}

// Перевод из арабских чисел в римские
func intToRoman(s int) string {
	intToRoman := map[int]string{
		1:   "I",
		4:   "IV",
		5:   "V",
		9:   "IX",
		10:  "X",
		40:  "XL",
		50:  "L",
		90:  "XC",
		100: "C",
	}
	res := ""
	for s > 0 {
		for _, key := range []int{100, 90, 50, 40, 10, 9, 5, 4, 1} {
			if s >= key {
				s -= key
				res += intToRoman[key]
				break
			}
		}
	}
	return res
}

func main() {
	// Ввод выражения
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	arr := strings.Split(input, " ")
	// fmt.Printf("%q\n", arr)
	// fmt.Println(reflect.TypeOf(num1))
	// fmt.Println(num1)
	// os.Exit(0)

	// Проверка на количество операндов
	if len(arr) < 3 {
		fmt.Println(notMath)
		os.Exit(0)
	} else if len(arr) > 3 {
		fmt.Println(tooMuch)
		os.Exit(0)
	}

	// Преобразование строки в число
	var num1 int
	var num2 int
	var isRoman bool
	if isNumeric(arr[0]) == true { // Проверка является ли первый операнд арабским числом

		if isNumeric(arr[2]) == true { // Проверка является ли второй операнд арабским числом
			num1, _ = strconv.Atoi(arr[0])
			num2, _ = strconv.Atoi(arr[2])
			// fmt.Println(num1)
			isRoman = false
		} else if romanToInt(arr[2]) != 0 { // Проверка является ли второй операнд римским числом
			fmt.Println(systems)
			os.Exit(0)
		} else { // Второй операнд не является числом в принципе
			fmt.Println(notMath)
			os.Exit(0)
		}
	} else if romanToInt(arr[0]) != 0 { // Проверка является ли первый операнд римским числом
		if romanToInt(arr[2]) != 0 { // Проверка является ли второй операнд римским числом
			num1 = romanToInt(arr[0])
			num2 = romanToInt(arr[2])
			// fmt.Println(numR1)
			isRoman = true
		} else if isNumeric(arr[2]) == true { // Проверка является ли второй операнд римским числом
			fmt.Println(systems)
			os.Exit(0)
		} else { // Второй операнд не является числом в принципе
			fmt.Println(notMath)
			os.Exit(0)
		}
	} else { //Первый операнд не является числом в принципе
		fmt.Println(notMath)
		os.Exit(0)
	}

	// Проверка на диапазон 1-10
	if (num1 > 10) || (num1 < 1) || (num2 > 10) || (num2 < 1) {
		fmt.Println(diapason)
		os.Exit(0)
	}

	// Вычисление
	res := 0
	if arr[1] == "+" {
		res = num1 + num2
	} else if arr[1] == "-" {
		res = num1 - num2
	} else if arr[1] == "*" {
		res = num1 * num2
	} else if arr[1] == "/" {
		res = num1 / num2
	} else {
		fmt.Println(operator)
		os.Exit(0)
	}

	//Перевод в римские
	if isRoman {
		if res <= 0 {
			fmt.Println(zero)
			os.Exit(0)
		}
		fmt.Println(intToRoman(res))
	} else {
		fmt.Println(res)
	}
}
