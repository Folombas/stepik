package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num1, num2 int

	// Чтение входных данных
	fmt.Print("Введите два числа через пробел: ")
	_, err := fmt.Scan(&num1, &num2)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	// Проверка диапазона (0-10000)
	if num1 < 0 || num1 > 10000 || num2 < 0 || num2 > 10000 {
		fmt.Println("Числа должны быть в диапазоне от 0 до 10000")
		return
	}

	// Преобразование чисел в строки для работы с цифрами
	str1 := strconv.Itoa(num1)
	str2 := strconv.Itoa(num2)

	// Проверка на уникальность цифр в каждом числе
	if !hasUniqueDigits(str1) || !hasUniqueDigits(str2) {
		fmt.Println("Ошибка: цифры в числах не должны повторяться")
		return
	}

	// Поиск общих цифр
	commonDigits := findCommonDigits(str1, str2)

	// Вывод результата
	if len(commonDigits) == 0 {
		fmt.Println("Общих цифр не найдено")
	} else {
		fmt.Println("Общие цифры:", strings.Join(commonDigits, " "))
	}
}

// hasUniqueDigits проверяет, что все цифры в числе уникальны
func hasUniqueDigits(str string) bool {
	seen := make(map[rune]bool)
	for _, ch := range str {
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}

// findCommonDigits находит цифры, которые есть в обоих числах
// и возвращает их в порядке появления в первом числе
func findCommonDigits(str1, str2 string) []string {
	var result []string
	digitsInSecond := make(map[rune]bool)

	// Создаём множество цифр второго числа
	for _, ch := range str2 {
		digitsInSecond[ch] = true
	}

	// Проверяем цифры первого числа
	for _, ch := range str1 {
		if digitsInSecond[ch] {
			result = append(result, string(ch))
		}
	}

	return result
}
