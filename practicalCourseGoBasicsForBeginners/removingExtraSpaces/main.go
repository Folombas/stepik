package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Имеет строку с лишними пробелами в начале и в конце
	text := "    Hello, Go!        "

	// 2. Создаёт новую строку, в которой все лишние пробелы удалены
	cleanedText := strings.TrimSpace(text)

	// 3. Выводит на экран исходную и чищенную строки
	fmt.Printf("Исходная строка: '%s'\n", text)
	fmt.Printf("Очищенная строка: '%s'\n", cleanedText)
}
