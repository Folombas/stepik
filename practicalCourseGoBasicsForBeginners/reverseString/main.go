package main

import (
	"fmt"
)

func reverseString(s string) string {
	// Преобразуем строку в слайс рун (для корректной работы с Unicode)
	runes := []rune(s)

	// Переворачиваем слайс рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Преобразуем обратно в строку
	return string(runes)
}

func main() {
	original := "GoLang"
	reversed := reverseString(original)

	fmt.Printf("Исходная строка: %s\n", original)
	fmt.Printf("Обратная строка: %s\n", reversed)
}
