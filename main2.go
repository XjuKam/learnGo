package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
1. принимаю на вход значение
2. убираю пробелы
4. разделяю значения на 3 переменные
5. проверяю 1 и 3 переменные на возможность изменения типа на int (это можно, так как при конвертации стринг (римские) будет 0, а на вход нельзя 0)
    1.  если инт = арабские - идем дальше
    2. если float - выдаем панику
    3. если стринг
        1. дополнительная проверка на соответствие значений в мапе римских цифер
        2. конвертируем их в арабские
    4. если 1 и 3 переменны имеют разные типы - выдаем панику
    5. проверяем переменные на ограничения от 1 до 10
6. используем свитч по операцияю +*-/
1. если ок - выводим результат
1. если были римские - идем в мапу для получения соответсвующего ключа значению результата
2. если нет - выдаем панику

*/

func main() {

	var v1 string
	var v2 string
	var v3 string
	var v1int int
	var v3int int
	var romeOper bool
	var result int

	// карта допустимых значений на вход
	mm := map[string]data{
		"one":       {"I", 1},
		"two":       {"II", 2},
		"three":     {"III", 3},
		"four":      {"IV", 4},
		"five":      {"V", 5},
		"six":       {"VI", 6},
		"seven":     {"VII", 7},
		"eight":     {"VIII", 8},
		"nine":      {"IX", 9},
		"ten":       {"X", 10},
		"eleven":    {"XI", 11},
		"twelve":    {"XII", 12},
		"thirteen":  {"XIII", 13},
		"fourteen":  {"XIV", 14},
		"fifteen":   {"XV", 15},
		"sixteen":   {"XVI", 16},
		"seventeen": {"XVII", 17},
		"eighteen":  {"XVIII", 18},
		"nineteen":  {"XIX", 19},
		"twenty":    {"XX", 20},
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Input")
		text, _ := reader.ReadString('\n') // Ждет ввода данных в формате строки
		text = strings.TrimSpace(text)     // Очищает все пустоты (пробелы, табуляцию)

		// Ищем индекс знака операции
		opIndex := FindOperatorIndex(text)

		// Проверяем, что знак операции найден
		if opIndex == -1 {
			fmt.Println("Знак операции не найден в выражении")
			return
		}

		// Извлекаем значения введенных переменных
		v1 = text[:opIndex]
		v2 = text[opIndex : opIndex+1]
		v3 = text[opIndex+1:]

		v1 = strings.TrimSpace(v1) // Очищает все пустоты (пробелы, табуляцию)
		v2 = strings.TrimSpace(v2) // Очищает все пустоты (пробелы, табуляцию)
		v3 = strings.TrimSpace(v3) // Очищает все пустоты (пробелы, табуляцию)

		v1int, _ = strconv.Atoi(v1) // Преобразование строки в число
		v3int, _ = strconv.Atoi(v3) // Преобразование строки в число

		// Проверяем систему нумерации
		if OnlyArabRome(v1int, v3int) == "rome" {
			key, found := GetKeyByValueRome(mm, v1)
			valueMm := mm[key]

			if found {
				v1int = valueMm.arabNum

			} else {
				panic("Выдача паники. Недопустимое значение, только 1-10, либо I - X")
			}
			key, found = GetKeyByValueRome(mm, v3)
			valueMm = mm[key]
			if found {
				v3int = valueMm.arabNum

				romeOper = true
			} else {
				panic("Выдача паники, так как используются одновременно разные системы счисления.")
			}

			// Выполнение операции

			switch v2 {
			case "+":
				result = v1int + v3int
			case "-":
				result = v1int - v3int
			case "*":
				result = v1int * v3int
			case "/":
				if v3int != 0 {
					result = v1int / v3int
				} else {
					fmt.Println("Деление на ноль недопустимо.")
					return
				}
			default:
				fmt.Println("Неверная операция.")
				return
			}

			if result <= -1 && OnlyArabRome(v1int, v3int) == "rome" {
				panic("В риме не было отрицательных, ну вот так и жили.")
			} else if romeOper == true {
				key, _ := GetKeyByValueArab(mm, result)
				valueMm := mm[key]
				resultRome := valueMm.romeNum
				fmt.Printf("Output:\n%v\n", resultRome)
			} else {
				fmt.Printf("Output:\n%v\n", result)
			}
		} else {

			// Выполнение операции
			switch v2 {
			case "+":
				result = v1int + v3int
			case "-":
				result = v1int - v3int
			case "*":
				result = v1int * v3int
			case "/":
				if v3int != 0 {
					result = v1int / v3int
				} else {
					fmt.Println("Деление на ноль недопустимо.")
					return
				}
			default:
				fmt.Println("Неверная операция.")
				return
			}

			if result <= -1 && OnlyArabRome(v1int, v3int) == "rome" {
				panic("В риме не было отрицательных, ну вот так и жили.")
			} else if romeOper == true {
				key, _ := GetKeyByValueRome(mm, strconv.Itoa(result))
				valueMm := mm[key]
				resultRome := valueMm.romeNum
				fmt.Printf("Output:\n%v", resultRome)
			} else {
				fmt.Printf("Output:\n%v\n", result)
			}
		}
	}
}
func FindOperatorIndex(input string) int {
	operators := "+-*/"

	// Ищем индекс знака операции
	for i := range input {
		if strings.Contains(operators, string(input[i])) {
			return i
		}
	}
	return -1
}

func OnlyArabRome(v1int, v3int int) string {
	if v1int == 0 || v3int == 0 {
		//	panic("Выдача паники, два операнда должны быть только  1 - 10, либо  I - X.")
		return "rome"
	} else if v1int >= 1 && v1int <= 10 && v3int >= 1 && v3int <= 10 {
		return "arab"
	} else if v1int >= 1 && v3int == 0 {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	} else if v1int == 0 && v3int >= 1 {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	} else if v1int >= 11 || v3int >= 11 {
		panic("Выдача паники, два операнда должны быть только  1 - 10, либо  I - X.")
	} else {
		return "rome"
	}
}

type data struct {
	romeNum string
	arabNum int
}

func GetKeyByValueRome(mm map[string]data, value string) (string, bool) {
	for key, v := range mm {
		if v.romeNum == value {
			return key, true
		}
	}
	return "", false
}

func GetKeyByValueArab(mm map[string]data, valueInt int) (string, bool) {
	for key, b := range mm {
		if b.arabNum == valueInt {
			return key, true
		}
	}
	return "", false
}
