package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Имеет строку, содержащую русские буквы
	text := "Я изучаю программирование на Go"

	// 2. Подсчитывает количество гласных и согласных букв
	vowels := "аеёиоуыэюяАЕЁИОУЫЭЮЯ"
	consonants := "бвгджзйклмнпрстфхцчшщъьБВГДЖЗЙКЛМНПРСТФХЦЧШЩЪЬ"

	vowelCount := 0
	consonantCount := 0

	// Перебираем каждый символ в строке
	for _, char := range text {
		if strings.ContainsRune(vowels, char) {
			vowelCount++
		} else if strings.ContainsRune(consonants, char) {
			consonantCount++
		}
		// Пробелы и другие символы игнорируются
	}

	// 3. Выводит итоговое количество гласных и согласных на экран
	fmt.Printf("Гласных: %d\n", vowelCount)
	fmt.Printf("Согласных: %d\n", consonantCount)
}
