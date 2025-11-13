package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Имеет предложение
	sentence := "Go - это отличный язык программирования"

	// 2. Подсчитывает количество слов с помощью strings.Fields
	words := strings.Fields(sentence)
	wordCount := len(words)

	// 3. Выводит итоговое количество слов на экран
	fmt.Printf("В предложении '%s' %d слов(a).\n", sentence, wordCount)
}
