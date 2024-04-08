package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var partOne string
	var operator string
	var partTwo string
	var partOneInt int
	var partTwoInt int
	var romeOper bool
	var result int

	// карта допустимых значений на вход
	mm := map[string]data{
		"one":          {"I", 1},
		"two":          {"II", 2},
		"three":        {"III", 3},
		"four":         {"IV", 4},
		"five":         {"V", 5},
		"six":          {"VI", 6},
		"seven":        {"VII", 7},
		"eight":        {"VIII", 8},
		"nine":         {"IX", 9},
		"ten":          {"X", 10},
		"eleven":       {"XI", 11},
		"twelve":       {"XII", 12},
		"thirteen":     {"XIII", 13},
		"fourteen":     {"XIV", 14},
		"fifteen":      {"XV", 15},
		"sixteen":      {"XVI", 16},
		"seventeen":    {"XVII", 17},
		"eighteen":     {"XVIII", 18},
		"nineteen":     {"XIX", 19},
		"twenty":       {"XX", 20},
		"twentyone":    {"XXI", 21},
		"twentytwo":    {"XXII", 22},
		"twentythree":  {"XXIII", 23},
		"twentyfour":   {"XXIV", 24},
		"twentyfive":   {"XXV", 25},
		"twentysix":    {"XXVI", 26},
		"twentyseven":  {"XXVII", 27},
		"twentyeight":  {"XXVIII", 28},
		"twentynine":   {"XXIX", 29},
		"thirty":       {"XXX", 30},
		"thirtyone":    {"XXXI", 31},
		"thirtytwo":    {"XXXII", 32},
		"thirtythree":  {"XXXIII", 33},
		"thirtyfour":   {"XXXIV", 34},
		"thirtyfive":   {"XXXV", 35},
		"thirtysix":    {"XXXVI", 36},
		"thirtyseven":  {"XXXVII", 37},
		"thirtyeight":  {"XXXVIII", 38},
		"thirtynine":   {"XXXIX", 39},
		"forty":        {"XL", 40},
		"fortyone":     {"XLI", 41},
		"fortytwo":     {"XLII", 42},
		"fortythree":   {"XLIII", 43},
		"fortyfour":    {"XLIV", 44},
		"fortyfive":    {"XLV", 45},
		"fortysix":     {"XLVI", 46},
		"fortyseven":   {"XLVII", 47},
		"fortyeight":   {"XLVIII", 48},
		"fortynine":    {"XLIX", 49},
		"fifty":        {"L", 50},
		"fiftyone":     {"LI", 51},
		"fiftytwo":     {"LII", 52},
		"fiftythree":   {"LIII", 53},
		"fiftyfour":    {"LIV", 54},
		"fiftyfive":    {"LV", 55},
		"fiftysix":     {"LVI", 56},
		"fiftyseven":   {"LVII", 57},
		"fiftyeight":   {"LVIII", 58},
		"fiftynine":    {"LIX", 59},
		"sixty":        {"LX", 60},
		"sixtyone":     {"LXI", 61},
		"sixtytwo":     {"LXII", 62},
		"sixtythree":   {"LXIII", 63},
		"sixtyfour":    {"LXIV", 64},
		"sixtyfive":    {"LXV", 65},
		"sixtysix":     {"LXVI", 66},
		"sixtyseven":   {"LXVII", 67},
		"sixtyeight":   {"LXVIII", 68},
		"sixtynine":    {"LXIX", 69},
		"seventy":      {"LXX", 70},
		"seventyone":   {"LXXI", 71},
		"seventytwo":   {"LXXII", 72},
		"seventythree": {"LXXIII", 73},
		"seventyfour":  {"LXXIV", 74},
		"seventyfive":  {"LXXV", 75},
		"seventysix":   {"LXXVI", 76},
		"seventyseven": {"LXXVII", 77},
		"seventyeight": {"LXXVIII", 78},
		"seventynine":  {"LXXIX", 79},
		"eighty":       {"LXXX", 80},
		"eightyone":    {"LXXXI", 81},
		"eightytwo":    {"LXXXII", 82},
		"eightythree":  {"LXXXIII", 83},
		"eightyfour":   {"LXXXIV", 84},
		"eightyfive":   {"LXXXV", 85},
		"eightysix":    {"LXXXVI", 86},
		"eightyseven":  {"LXXXVII", 87},
		"eightyeight":  {"LXXXVIII", 88},
		"eightynine":   {"LXXXIX", 89},
		"ninety":       {"XC", 90},
		"ninetyone":    {"XCI", 91},
		"ninetytwo":    {"XCII", 92},
		"ninetythree":  {"XCIII", 93},
		"ninetyfour":   {"XCIV", 94},
		"ninetyfive":   {"XCV", 95},
		"ninetysix":    {"XCVI", 96},
		"ninetyseven":  {"XCVII", 97},
		"ninetyeight":  {"XCVIII", 98},
		"ninetynine":   {"XCIX", 99},
		"hundred":      {"C", 100},
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
			fmt.Println("Некорретное выражение. Укажите в формате a+b, либо другой оператор '-*/' ")
			return
		}

		// Извлекаем значения введенных переменных
		partOne = text[:opIndex]
		operator = text[opIndex : opIndex+1]
		partTwo = text[opIndex+1:]

		partOne = strings.TrimSpace(partOne)   // Очищает все пустоты (пробелы, табуляцию)
		operator = strings.TrimSpace(operator) // Очищает все пустоты (пробелы, табуляцию)
		partTwo = strings.TrimSpace(partTwo)   // Очищает все пустоты (пробелы, табуляцию)

		partOneInt, _ = strconv.Atoi(partOne) // Преобразование строки в число
		partTwoInt, _ = strconv.Atoi(partTwo) // Преобразование строки в число

		// Проверяем систему нумерации и выдаем результат
		if OnlyArabRome(partOneInt, partTwoInt) == "rome" {
			key, found := GetKeyByValueRome(mm, partOne)
			valueMm := mm[key]
			if found {
				partOneInt = valueMm.arabNum
			} else {
				panic("Выдача паники. Недопустимое значение, только 1-10, либо I - X")
			}
			key, found = GetKeyByValueRome(mm, partTwo)
			valueMm = mm[key]
			if found {
				partTwoInt = valueMm.arabNum
				romeOper = true
			} else {
				panic("Выдача паники, так как используются одновременно разные системы счисления.")
			}
			if partOneInt >= 11 || partTwoInt >= 11 {
				panic("Выдача паники. Недопустимое значение, только 1-10, либо I - X")
			} else {
				result = GetOperationResult(operator, partOneInt, partTwoInt)
			}
			if result <= 0 && romeOper == true {
				panic("В риме не было отрицательных и нуля, ну вот так и жили.")
			} else {
				key, _ := GetKeyByValueArab(mm, result)
				valueMm := mm[key]
				resultRome := valueMm.romeNum
				fmt.Printf("Output:\n%v\n", resultRome)
			}
		} else {
			result = GetOperationResult(operator, partOneInt, partTwoInt)

			fmt.Printf("Output:\n%v\n", result)
		}
	}
}

func FindOperatorIndex(input string) int { // Проверка на доступный оператор
	operators := "+-*/"

	// Ищем индекс знака операции
	for i := range input {
		if strings.Contains(operators, string(input[i])) {
			return i
		}
	}
	return -1
}

func OnlyArabRome(partOneInt, partTwoInt int) string { // Проверка на условие тоьлко араб. или рим. операнды
	if partOneInt >= 1 && partOneInt <= 10 && partTwoInt >= 1 && partTwoInt <= 10 {
		return "arab"
	} else if partOneInt >= 1 && partTwoInt == 0 {
		panic("Выдача паники, лишний оператор или разные системы счисления.")
	} else if partOneInt == 0 && partTwoInt >= 1 {
		panic("Выдача паники, лишний оператор или разные системы счисления.")
	} else if partOneInt >= 11 || partTwoInt >= 11 {
		panic("Выдача паники, два операнда должны быть только  1 - 10, либо  I - X.")
	} else {
		return "rome"
	}
}

type data struct {
	romeNum string
	arabNum int
}

func GetKeyByValueRome(mm map[string]data, value string) (string, bool) { // Достаем из мапы по рим. числу
	for key, v := range mm {
		if v.romeNum == value {
			return key, true
		}
	}
	return "", false
}

func GetKeyByValueArab(mm map[string]data, valueInt int) (string, bool) { // Достаем из мапы по араб. числу
	for key, b := range mm {
		if b.arabNum == valueInt {
			return key, true
		}
	}
	return "", false
}

func GetOperationResult(operator string, partOneInt int, partTwoInt int) int { // Выполнение операции

	switch operator {
	case "+":
		result := partOneInt + partTwoInt
		return result
	case "-":
		result := partOneInt - partTwoInt
		return result
	case "*":
		result := partOneInt * partTwoInt
		return result
	case "/":
		if partTwoInt != 0 {
			result := partOneInt / partTwoInt
			return result
		} else {
			panic("Деление на ноль недопустимо.")
		}
	default:
		panic("Неверная операция.")
	}
}
