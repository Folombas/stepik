package main

import (
	"fmt"
	"strings"
)

func main() {
	// Исходное предложение
	original := "Я изучаю язык Java"

	// Замена слова "Java" на "Go"
	changed := strings.ReplaceAll(original, "Java", "Go")

	// Вывод результатов
	fmt.Printf("Исходное предложение: '%s'\n", original)
	fmt.Printf("Изменённое предложение: '%s'\n", changed)
}
