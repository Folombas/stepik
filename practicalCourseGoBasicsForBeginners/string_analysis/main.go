package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Исходная строка на русском языке
	str := "Завтра наступит новый день"

	// Подсчёт количества байтов
	byteCount := len(str)

	// Подсчёт количества символов (рун)
	runeCount := utf8.RuneCountInString(str)

	// Вывод результатов
	fmt.Print("Исходная строка: '%s'\n", str)
	fmt.Printf("Количество байтов: %d\n", byteCount)
	fmt.Printf("Количество символов: %d\n", runeCount)
}
